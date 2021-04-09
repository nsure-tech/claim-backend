package service

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"nsure/vote/common"
	"nsure/vote/models"
	"nsure/vote/models/mysql"
	"nsure/vote/utils"
	"time"
)

func ApplyVote(claimId int64, status common.ClaimStatus, arbiterId, coverHash, signHash string) (string, error) {
	tx, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return "", err
	}
	defer func() { _ = tx.Rollback() }()

	vote, err := tx.GetVoteForUpdate(claimId, arbiterId)
	if err != nil {
		return "", err
	}
	if vote == nil {
		return "", errors.New("vote is nil errors")
	}

	if vote.Status != common.ClaimStatusArbiter {
		return "", errors.New("already claim pass or deny")
	}

	if vote.CoverHash != coverHash {
		return "", fmt.Errorf("cover_hash %v!=%v", vote.CoverHash, coverHash)
	}

	vote.Status = status
	vote.SignHash = signHash

	err = tx.UpdateVote(vote)
	if err != nil {
		return "", err
	}

	apply, err := tx.GetApplyForUpdate(claimId, arbiterId)
	if err != nil {
		return "", err
	}
	if vote.Status == common.ClaimStatusPass {
		apply.Status = common.ApplyStatusPass
	} else if vote.Status == common.ClaimStatusDeny {
		apply.Status = common.ApplyStatusDeny
	}
	if err = tx.UpdateApply(apply); err != nil {
		return "", err
	}

	return string(apply.Status), tx.CommitTx()
}

func ApplyVoteAdmin(claimId int64, status common.ClaimStatus, arbiterId, coverHash, signHash string) (string, error) {
	tx, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return "", err
	}
	defer func() { _ = tx.Rollback() }()

	if votes, err := GetAlreadyVoteByClaimId(claimId); err != nil {
		return "", err
	} else if len(votes) >= common.ArbiterMaxNum {
		return "", fmt.Errorf("already vote:%v", len(votes))
	}

	if vote, err := tx.GetVoteForUpdate(claimId, arbiterId); err != nil {
		return "", err
	} else if vote != nil {
		return "", fmt.Errorf("vote errors isn't nil")
	}

	if !AdminAddress(arbiterId) {
		return "", fmt.Errorf("%v isn't admin address", arbiterId)
	}
	claim, err := tx.GetClaimById(claimId)
	if err != nil {
		return "", err
	}
	if claim.CoverHash != coverHash {
		return "", fmt.Errorf("cover_hash %v!=%v", claim.CoverHash, coverHash)
	}
	if claim.ArbiterAt.Add(time.Duration(common.VoteMinute) * time.Minute).After(time.Now()) {
		return "", fmt.Errorf("vote time error %v", claim.VoteAt)
	}
	if claim.Settled {
		return "", fmt.Errorf("claimId:%v already close", claimId)
	}

	vote := &models.Vote{
		ClaimId:      claimId,
		ArbiterId:    arbiterId,
		UserId:       claim.UserId,
		Product:      claim.Product,
		CoverId:      claim.CoverId,
		CoverHash:    claim.CoverHash,
		Currency:     claim.Currency,
		Amount:       claim.Amount,
		Reward:       claim.Reward,
		ArbiterAt:    claim.ArbiterAt,
		SubmitAt:     claim.SubmitAt,
		CoverBeginAt: claim.CoverBeginAt,
		SignHash:     signHash,
		Status:       status,
	}
	if err = tx.AddVote(vote); err != nil {
		return "", err
	}

	return string(status), tx.CommitTx()
}

func GetAlreadyVoteByClaimId(claimId int64) ([]*models.Vote, error) {
	statuses := []common.ClaimStatus{common.ClaimStatusPass, common.ClaimStatusDeny}
	return mysql.SharedStore().GetVoteByClaimIdStatus(claimId, statuses)
}

func GetVoteByVoteNum(voteNum uint8) ([]*common.ClaimId, error) {
	statuses := []common.ClaimStatus{common.ClaimStatusPass, common.ClaimStatusDeny}
	return mysql.SharedStore().GetVoteByVoteNum(statuses, voteNum)
}

func ExecuteVote(claimId int64) error {
	tx, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	claim, err := tx.GetClaimForUpdate(claimId)
	if err != nil {
		return err
	}
	if claim == nil {
		return errors.New("claim error")
	}

	votes, err := tx.GetVoteByClaimId(claimId)
	if err != nil {
		return err
	}
	pass := 0
	deny := 0
	for _, vote := range votes {
		if vote.Status == common.ClaimStatusPass {
			pass++
		} else if vote.Status == common.ClaimStatusDeny {
			deny++
		}
	}
	rewardNum := 0
	claim.VoteAt = time.Now()
	claim.VoteNum = pass + deny

	if pass >= common.VoteMaxNum {
		claim.Status = common.ClaimStatusPass
		claim.PaymentStatus = common.PaymentStatusPass
		rewardNum = pass
	} else if deny >= common.VoteMaxNum {
		claim.Status = common.ClaimStatusDeny
		claim.PaymentStatus = common.PaymentStatusFail
		rewardNum = deny
	} else {
		claim.Status = common.ClaimStatusArbiterFail
		claim.PaymentStatus = common.PaymentStatusFail
		claim.Settled = true
	}

	err = tx.UpdateClaim(claim)
	if err != nil {
		return err
	}
	if claim.Settled {
		return tx.CommitTx()
	}

	for _, vote := range votes {
		voteFill := &models.VoteFill{
			VoteId:        vote.Id,
			ClaimId:       vote.ClaimId,
			CoverId:       vote.CoverId,
			CoverHash:     claim.CoverHash,
			Currency:      claim.Currency,
			Amount:        claim.Amount,
			Reward:        claim.Reward,
			ArbiterId:     vote.ArbiterId,
			VoteAt:        claim.VoteAt,
			ClaimStatus:   claim.Status,
			PaymentStatus: claim.PaymentStatus,
			VoteStatus:    vote.Status,
		}
		if vote.Status == common.ClaimStatusArbiter {
			voteFill.Status = common.FillStatusAbstain
		} else if claim.Status == vote.Status {
			voteFill.Status = common.FillStatusEqual
			voteFill.RewardNum = rewardNum
		} else if claim.Status == common.ClaimStatusArbiterFail {
			voteFill.Status = common.FillStatusEmpty
		} else {
			voteFill.Status = common.FillStatusDifferent
		}
		err = tx.AddVoteFill(voteFill)
		if err != nil {
			return err
		}

		vote.Settled = true
		err = tx.UpdateVote(vote)
		if err != nil {
			return err
		}
	}

	return tx.CommitTx()
}

func GetUnsettledFills(count int) ([]*models.VoteFill, error) {
	return mysql.SharedStore().GetUnsettledVoteFills(count)
}

func ExecuteVoteFill(claimId int64) error {
	db, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return err
	}
	defer func() { _ = db.Rollback() }()
	claim, err := db.GetClaimForUpdate(claimId)
	if err != nil {
		return err
	}
	if claim.Status == common.ClaimStatusPass {
		claim.Status = common.ClaimStatusPassEnd
		claim.PaymentStatus = common.PaymentStatusPass
		payment := &models.Payment{
			UserId:      claim.UserId,
			Product:     claim.Product,
			CoverId:     claim.CoverId,
			CoverHash:   claim.CoverHash,
			Currency:    claim.Currency,
			Amount:      claim.Amount,
			ClaimId:     claim.Id,
			ClaimStatus: claim.Status,
		}

		if err = db.AddPayment(payment); err != nil {
			return err
		}
	} else if claim.Status == common.ClaimStatusDeny {
		claim.Status = common.ClaimStatusDenyEnd
		claim.PaymentStatus = common.PaymentStatusFail
	} else {
		return errors.New("claim status error")
	}
	voteFills, err := db.GetVoteFillsByClaimId(claimId)
	if err != nil {
		return err
	}
	for _, voteFill := range voteFills {
		voteFill.ClaimStatus = claim.Status
		if _, err = UpdateVoteFill(db, voteFill); err != nil {
			return err
		}
	}
	claim.Settled = true
	if err = db.UpdateClaim(claim); err != nil {
		return err
	}
	common.ClaimActive--
	common.ClaimClosed++

	return db.CommitTx()
}

func UpdateVoteFill(tx models.Store, voteFill *models.VoteFill) (decimal.Decimal, error) {
	punishFunds := decimal.Zero
	if common.ChallengeStatusSuccess == voteFill.ChallengeStatus &&
		(voteFill.Status == common.FillStatusChaDifferent || voteFill.Status == common.FillStatusAbstain) {
		punishFunds = utils.ArbiterNSure()
		if err := AddDelayBill(tx, common.AccountNSure, common.CurrencyNSure, punishFunds, decimal.Zero,
			common.BillTypeVotePunish, ""); err != nil {
			return decimal.Zero, err
		}
		if err := AddDelayBill(tx, voteFill.ArbiterId, common.CurrencyNSure, decimal.Zero, punishFunds.Neg(),
			common.BillTypeVotePunish, ""); err != nil {
			return decimal.Zero, err
		}
		if _, err := SubUsedAddClosed(tx, voteFill.ArbiterId, 1); err != nil {
			return decimal.Zero, err
		}
	} else if voteFill.Status == common.FillStatusAbstain {
		if _, err := ReturnAvailableNum(tx, voteFill.ArbiterId, 1); err != nil {
			return decimal.Zero, err
		}
	} else {
		if _, err := ReturnAvailableAddClosed(tx, voteFill.ArbiterId, 1); err != nil {
			return decimal.Zero, err
		}
	}

	if voteFill.Status == common.FillStatusEqual || voteFill.Status == common.FillStatusChaEqual {
		if err := AddReward(tx, voteFill); err != nil {
			return decimal.Zero, err
		}
	}

	voteFill.Settled = true
	if err := tx.UpdateVoteFill(voteFill); err != nil {
		return decimal.Zero, err
	}
	return punishFunds, nil
}

func AddReward(store models.Store, voteFill *models.VoteFill) error {
	if voteFill.RewardNum <= 0 {
		return errors.New("reward num is zero error")
	}
	reward := &models.Reward{
		VoteId:      voteFill.VoteId,
		ClaimId:     voteFill.ClaimId,
		CoverId:     voteFill.CoverId,
		CoverHash:   voteFill.CoverHash,
		ArbiterId:   voteFill.ArbiterId,
		ClaimStatus: voteFill.ClaimStatus,
		VoteStatus:  voteFill.VoteStatus,
		Currency:    voteFill.Currency,
		Amount:      voteFill.Reward.Div(decimal.NewFromInt(int64(voteFill.RewardNum))).Floor(),
	}
	return store.AddReward(reward)
}
