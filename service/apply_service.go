package service

import (
	"errors"
	"fmt"
	"math/rand"
	"nsure/vote/common"
	"nsure/vote/models"
	"nsure/vote/models/mysql"
	"time"
)

func GetApplyByArbiterId(arbiterId string, status *common.ApplyStatus) ([]*models.Apply, error) {
	return mysql.SharedStore().GetApplyByArbiterId(arbiterId, status)
}

func GetApplyCount(arbiterId string, status common.ApplyStatus) (int, error) {
	return mysql.SharedStore().GetApplyCountByArbiterId(arbiterId, status)
}
func GetApplyList(arbiterId string, status common.ApplyStatus, begin, limit int) ([]*models.Apply, error) {
	return mysql.SharedStore().GetApplyListByArbiterId(arbiterId, status, begin, limit)
}
func GetApplyByClaimId(claimId int64, status *common.ApplyStatus) ([]*models.Apply, error) {
	return mysql.SharedStore().GetApplyByClaimId(claimId, status)
}
func GetApplyByApplyNum(applyNum uint8, status *common.ApplyStatus) ([]*models.Apply, error) {
	return mysql.SharedStore().GetApplyByApplyNum(applyNum, status)
}

func ExecuteApply(claimId int64) error {
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

	applies, err := GetApplyByClaimId(claimId, nil)
	if err != nil {
		return err
	}

	claim.Status = common.ClaimStatusArbiter
	claim.ArbiterAt = time.Now()
	if claim.ApplyNum != len(applies) {
		return fmt.Errorf("execute apply error")
	}

	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(applies), func(i int, j int) {
		applies[i], applies[j] = applies[j], applies[i]
	})
	for i, apply := range applies {
		apply.Settled = true
		if i < common.ArbiterMaxNum {
			apply.ArbiterAt = claim.SubmitAt
			apply.Status = common.ApplyStatusSuccess
			vote := &models.Vote{
				ClaimId:      claimId,
				ArbiterId:    apply.ArbiterId,
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
				Status:       common.ClaimStatusArbiter,
			}
			err = tx.AddVote(vote)
			if err != nil {
				return err
			}
		} else {
			apply.ArbiterAt = claim.SubmitAt
			apply.Status = common.ApplyStatusFail
			ReturnAvailableNum(tx, apply.ArbiterId, 1)
		}
		err = tx.UpdateApply(apply)
		if err != nil {
			return err
		}
	}

	err = tx.UpdateClaim(claim)
	if err != nil {
		return err
	}

	return tx.CommitTx()
}
func ReturnApply(claimId int64) error {
	tx, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	applies, err := GetApplyByClaimId(claimId, nil)
	if err != nil {
		return err
	}
	for _, apply := range applies {
		apply.Settled = true
		apply.Status = common.ApplyStatusFail
		ReturnAvailableNum(tx, apply.ArbiterId, 1)
		err = tx.UpdateApply(apply)
		if err != nil {
			return err
		}
	}

	claim, err := tx.GetClaimForUpdate(claimId)
	if err != nil {
		return err
	}
	if claim == nil {
		return errors.New("claim error")
	}
	claim.Status = common.ClaimStatusApplyFail
	claim.Settled = true
	if claim.ApplyNum != len(applies) {
		return fmt.Errorf("execute apply error")
	}
	err = tx.UpdateClaim(claim)
	if err != nil {
		return err
	}

	return tx.CommitTx()
}
