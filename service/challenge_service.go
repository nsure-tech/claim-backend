package service

import (
	"errors"
	"fmt"
	"nsure/vote/common"
	"nsure/vote/models"
	"nsure/vote/models/mysql"
	"nsure/vote/utils"
	"time"
)

func GetChallengeByStatus(status common.ChallengeStatus, count int) ([]*models.Challenge, error) {
	return mysql.SharedStore().GetChallengeByStatus(status, count)
}

func ApplyChallenge(challengeId, coverHash string, claimId int64) (bool, error) {
	db, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return false, err
	}
	defer func() { _ = db.Rollback() }()

	claim, err := db.GetClaimForUpdate(claimId)
	if err != nil {
		return false, err
	}

	if claim.CoverHash != coverHash {
		return false, fmt.Errorf("cover_hash %v!=%v", claim.CoverHash, coverHash)
	}

	holdSize, err := utils.ChallengeNSure(claim.Amount)
	if err != nil {
		return false, err
	}
	if _, err = HoldBalance(db, challengeId, common.CurrencyNSure, holdSize, common.BillTypeChallenge); err != nil {
		return false, err
	}

	if claim.Status == common.ClaimStatusPass {
		claim.Status = common.ClaimStatusPassCha
	} else if claim.Status == common.ClaimStatusDeny {
		claim.Status = common.ClaimStatusDenyCha
	} else {
		return false, errors.New("claim status errors")
	}
	claim.PaymentStatus = common.PaymentStatusCha
	claim.ChallengeAt = time.Now()
	claim.Challenged = true

	challenge := &models.Challenge{
		ChallengeId: challengeId,
		ChallengeAt: claim.ChallengeAt,
		ClaimId:     claim.Id,
		CoverId:     claim.CoverId,
		CoverHash:   claim.CoverHash,
		Currency:    claim.Currency,
		Amount:      claim.Amount,
		Reward:      claim.Reward,
		Hold:        holdSize,
		ClaimStatus: claim.Status,
		Status:      common.ChallengeStatusApply,
	}

	if err = db.AddChallenge(challenge); err != nil {
		return false, err
	}
	if err = db.UpdateClaim(claim); err != nil {
		return false, err
	}
	return true, db.CommitTx()
}

func ChallengeVoteByAdmin(adminId string, claimId int64, status common.ChallengeStatus) (bool, error) {
	if !ChallengeAddress(adminId) {
		return false, fmt.Errorf("%v isn't challenge admin address", adminId)
	}

	db, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return false, err
	}
	defer func() { _ = db.Rollback() }()

	challenge, err := db.GetChallengeForUpdate(claimId)
	if err != nil {
		return false, err
	}
	if challenge == nil {
		return false, fmt.Errorf("challenge is null")
	}
	if challenge.Status == common.ChallengeStatusApply {
		challenge.Status = status
		if err = db.UpdateChallenge(challenge); err != nil {
			return false, err
		}

		return true, db.CommitTx()
	} else {
		return false, fmt.Errorf("challenge already success or fail")
	}

}

func ApplyChallengeOld(challengeId string, claimId int64) (bool, error) {
	db, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return false, err
	}
	defer func() { _ = db.Rollback() }()

	claim, err := db.GetClaimForUpdate(claimId)
	if err != nil {
		return false, err
	}

	holdSize, err := utils.ChallengeNSure(claim.Amount)
	if err != nil {
		return false, err
	}
	if _, err = HoldBalance(db, challengeId, common.CurrencyNSure, holdSize, common.BillTypeChallenge); err != nil {
		return false, err
	}

	if claim.Status == common.ClaimStatusPass {
		claim.Status = common.ClaimStatusPassCha
	} else if claim.Status == common.ClaimStatusDeny {
		claim.Status = common.ClaimStatusDenyCha
	} else {
		return false, errors.New("claim status errors")
	}
	claim.PaymentStatus = common.PaymentStatusCha
	claim.ChallengeAt = time.Now()
	claim.Challenged = true

	challenge := &models.Challenge{
		ChallengeId: challengeId,
		ChallengeAt: claim.ChallengeAt,
		ClaimId:     claim.Id,
		CoverId:     claim.CoverId,
		CoverHash:   claim.CoverHash,
		Currency:    claim.Currency,
		Amount:      claim.Amount,
		Reward:      claim.Reward,
		Hold:        holdSize,
		ClaimStatus: claim.Status,
		Status:      common.ChallengeStatusApply,
	}

	if err = db.AddChallenge(challenge); err != nil {
		return false, err
	}
	if err = db.UpdateClaim(claim); err != nil {
		return false, err
	}
	return true, db.CommitTx()
}

func ChallengeVote(claimId int64, status bool) (bool, error) {
	db, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return false, err
	}
	defer func() { _ = db.Rollback() }()

	challenge, err := db.GetChallengeForUpdate(claimId)
	if err != nil {
		return false, err
	}
	if status {
		challenge.Status = common.ChallengeStatusSuccess
	} else {
		challenge.Status = common.ChallengeStatusFail
	}
	if err = db.UpdateChallenge(challenge); err != nil {
		return false, err
	}

	return true, db.CommitTx()
}

func ExecuteChallengeSuccess(claimId int64) error {
	db, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return err
	}
	defer func() { _ = db.Rollback() }()
	challenge, err := db.GetChallengeForUpdate(claimId)
	if err != nil {
		return err
	}
	claim, err := db.GetClaimForUpdate(claimId)
	if err != nil {
		return err
	}
	if claim.Status == common.ClaimStatusPassCha {
		claim.Status = common.ClaimStatusPassChaPass
		claim.PaymentStatus = common.PaymentStatusFail
	} else if claim.Status == common.ClaimStatusDenyCha {
		claim.Status = common.ClaimStatusDenyChaPass
		claim.PaymentStatus = common.PaymentStatusPass
	} else {
		return errors.New("claim status error")
	}

	//voteFills, err := db.GetVoteFillsByClaimId(claimId)
	//if err != nil {
	//	return err
	//}
	//for _, voteFill := range voteFills {
	//	voteFill.ChallengeStatus = common.ChallengeStatusSuccess
	//	db.UpdateVoteFill(voteFill)
	//}

	claim.Settled = true
	if err = db.UpdateClaim(claim); err != nil {
		return err
	}

	challenge.ClaimStatus = claim.Status
	if _, err = AddChallengeFill(db, challenge, claim.UserId, claim.Product); err != nil {
		return err
	}
	challenge.Settled = true
	if err = db.UpdateChallenge(challenge); err != nil {
		return err
	}
	if err = db.CommitTx(); err != nil {
		return err
	}
	common.ClaimActive--
	common.ClaimClosed++
	return nil
}

func ExecuteChallengeFail(claimId int64) error {
	db, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return err
	}
	defer func() { _ = db.Rollback() }()

	challenge, err := db.GetChallengeForUpdate(claimId)
	if err != nil {
		return err
	}

	claim, err := db.GetClaimForUpdate(claimId)
	if err != nil {
		return err
	}
	if claim.Status == common.ClaimStatusPassCha {
		claim.Status = common.ClaimStatusPassChaDeny
		claim.PaymentStatus = common.PaymentStatusPass
	} else if claim.Status == common.ClaimStatusDenyCha {
		claim.Status = common.ClaimStatusDenyChaDeny
		claim.PaymentStatus = common.PaymentStatusFail
	} else {
		return errors.New("claim status error")
	}

	//voteFills, err := db.GetVoteFillsByClaimId(claimId)
	//if err != nil {
	//	return err
	//}
	//for _, voteFill := range voteFills {
	//	voteFill.ChallengeStatus = common.ChallengeStatusFail
	//	db.UpdateVoteFill(voteFill)
	//}

	claim.Settled = true
	if err = db.UpdateClaim(claim); err != nil {
		return err
	}

	challenge.ClaimStatus = claim.Status
	if _, err = AddChallengeFill(db, challenge, claim.UserId, claim.Product); err != nil {
		return err
	}

	challenge.Settled = true
	if err = db.UpdateChallenge(challenge); err != nil {
		return err
	}
	if err = db.CommitTx(); err != nil {
		return err
	}

	common.ClaimActive--
	common.ClaimClosed++
	return nil
}
