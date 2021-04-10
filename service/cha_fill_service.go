package service

import (
	"github.com/shopspring/decimal"
	"nsure/vote/common"
	"nsure/vote/models"
	"nsure/vote/models/mysql"
)

func GetChallengeFillByStatus(status common.ChallengeStatus, count int) ([]*models.ChallengeFill, error) {
	return mysql.SharedStore().GetChallengeFillByStatus(status, count)
}

func AddChallengeFill(store models.Store, challenge *models.Challenge, userId, product string) (*models.ChallengeFill, error) {
	challengeFill := &models.ChallengeFill{
		ChallengeId: challenge.ChallengeId,
		ChallengeAt: challenge.ChallengeAt,
		ClaimId:     challenge.ClaimId,
		CoverId:     challenge.CoverId,
		CoverHash:   challenge.CoverHash,
		UserId:      userId,
		Product:     product,
		Currency:    challenge.Currency,
		Amount:      challenge.Amount,
		Reward:      challenge.Reward,
		Hold:        challenge.Hold,
		ClaimStatus: challenge.ClaimStatus,
		Status:      challenge.Status,
	}
	if err := store.AddChallengeFill(challengeFill); err != nil {
		return nil, err
	}
	return challengeFill, nil
}
func ExecuteChallengeFillSuccess(claimId int64) error {
	db, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return err
	}
	defer func() { _ = db.Rollback() }()

	challengeFill, err := db.GetChallengeFillForUpdate(claimId)
	if err != nil {
		return err
	}

	voteFills, err := db.GetVoteFillsByClaimId(claimId)
	if err != nil {
		return err
	}
	challengeFunds := decimal.Zero
	rewardNum := 0
	paymentStatus := common.PaymentStatusFail
	if common.ClaimStatusDenyChaPass == challengeFill.ClaimStatus {
		paymentStatus = common.PaymentStatusPass
	}
	for _, voteFill := range voteFills {
		voteFill.ChallengeStatus = common.ChallengeStatusSuccess
		voteFill.ClaimStatus = challengeFill.ClaimStatus
		voteFill.PaymentStatus = paymentStatus
		if voteFill.Status == common.FillStatusDifferent {
			voteFill.Status = common.FillStatusChaEqual
			rewardNum++
		} else if voteFill.Status == common.FillStatusEqual {
			voteFill.Status = common.FillStatusChaDifferent
			voteFill.RewardNum = 0
		}
	}

	for _, voteFill := range voteFills {
		if voteFill.Status == common.FillStatusChaEqual {
			voteFill.RewardNum = rewardNum
		}
		if punishFunds, err := UpdateVoteFill(db, voteFill); err != nil {
			return err
		} else {
			challengeFunds = challengeFunds.Add(punishFunds)
		}
	}

	if err := AddWaitBill(db, common.AccountNSure, common.CurrencyNSure, challengeFunds.Neg(), decimal.Zero,
		common.BillTypeChallengeSuccess, ""); err != nil {
		return err
	}
	if err := AddWaitBill(db, challengeFill.ChallengeId, common.CurrencyNSure, challengeFunds, decimal.Zero,
		common.BillTypeChallengeSuccess, ""); err != nil {
		return err
	}
	if err := AddWaitBill(db, challengeFill.ChallengeId, common.CurrencyNSure, challengeFill.Hold, challengeFill.Hold.Neg(),
		common.BillTypeChallengeSuccess, ""); err != nil {
		return err
	}

	if common.ClaimStatusDenyChaPass == challengeFill.ClaimStatus {
		payment := newPayment(challengeFill)
		if err = db.AddPayment(payment); err != nil {
			return err
		}
	}

	challengeFill.Settled = true
	if err = db.UpdateChallengeFill(challengeFill); err != nil {
		return err
	}

	return db.CommitTx()

}
func ExecuteChallengeFillFail(claimId int64) error {
	db, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return err
	}
	defer func() { _ = db.Rollback() }()

	chaFill, err := db.GetChallengeFillForUpdate(claimId)
	if err != nil {
		return err
	}

	voteFills, err := db.GetVoteFillsByClaimId(claimId)
	if err != nil {
		return err
	}
	for _, voteFill := range voteFills {
		voteFill.ChallengeStatus = common.ChallengeStatusFail
		voteFill.ClaimStatus = chaFill.ClaimStatus
		if _, err = UpdateVoteFill(db, voteFill); err != nil {
			return err
		}
	}

	if err := AddWaitBill(db, common.AccountChallenge, common.CurrencyNSure, chaFill.Hold, decimal.Zero,
		common.BillTypeChallengeFail, ""); err != nil {
		return err
	}

	if err := AddWaitBill(db, chaFill.ChallengeId, common.CurrencyNSure, decimal.Zero, chaFill.Hold.Neg(),
		common.BillTypeChallengeFail, ""); err != nil {
		return err
	}

	punishCha := &models.PunishCha{
		ClaimId:         chaFill.ClaimId,
		CoverId:         chaFill.CoverId,
		CoverHash:       chaFill.CoverHash,
		ChallengeId:     chaFill.ChallengeId,
		ClaimStatus:     chaFill.ClaimStatus,
		ChallengeStatus: chaFill.Status,
		Currency:        common.CurrencyNSure,
		Amount:          chaFill.Hold,
	}
	if err := db.AddPunishCha(punishCha); err != nil {
		return err
	}

	if common.ClaimStatusPassChaDeny == chaFill.ClaimStatus {
		payment := newPayment(chaFill)
		if err = db.AddPayment(payment); err != nil {
			return err
		}
	}

	chaFill.Settled = true
	if err = db.UpdateChallengeFill(chaFill); err != nil {
		return err
	}

	return db.CommitTx()

}

func newPayment(chaFill *models.ChallengeFill) *models.Payment {
	return &models.Payment{
		UserId:      chaFill.UserId,
		Product:     chaFill.Product,
		CoverId:     chaFill.CoverId,
		CoverHash:   chaFill.CoverHash,
		Currency:    chaFill.Currency,
		Amount:      chaFill.Amount,
		ClaimId:     chaFill.ClaimId,
		ClaimStatus: chaFill.ClaimStatus,
	}
}
