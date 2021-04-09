package service

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"nsure/vote/common"
	"nsure/vote/log"
	"nsure/vote/models"
	"nsure/vote/models/mysql"
	"time"
)

func AddClaim(userId, product, coverId, coverHash, currency string, amount, cost, reward decimal.Decimal, submitAt, beginAt, endAt time.Time, desc, cred, loss string) (int64, error) {
	claim := &models.Claim{
		UserId:       userId,
		Product:      product,
		Currency:     currency,
		CoverId:      coverId,
		CoverHash:    coverHash,
		Amount:       amount,
		Cost:         cost,
		Reward:       reward,
		SubmitAt:     submitAt,
		ArbiterAt:    time.Now().AddDate(10, 0, 0),
		VoteAt:       time.Now().AddDate(10, 1, 0),
		ChallengeAt:  time.Now().AddDate(10, 2, 0),
		CoverBeginAt: beginAt,
		CoverEndAt:   endAt,
		Description:  desc,
		Credential:   cred,
		Loss:         loss,
		Status:       common.ClaimStatusNew,
	}
	oldClaim, err := mysql.SharedStore().GetClaimByHash(coverHash)
	if err != nil {
		return 0, err
	}
	if oldClaim != nil {
		return oldClaim.Id, fmt.Errorf("already claim error")
	}

	db, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return 0, err
	}
	defer func() { _ = db.Rollback() }()

	err = db.AddClaim(claim)
	if err != nil {
		log.GetLog().Debug("Add Claim", zap.Error(err))
		return 0, err
	}

	if err = db.CommitTx(); err != nil {
		return 0, err
	}
	common.ClaimActive++

	return claim.Id, nil
}

func GetClaimByApply() ([]*models.Claim, error) {
	return mysql.SharedStore().GetClaimByApply(common.ClaimStatusNew, common.ApplyMinute, common.ApplyMaxNum)
}
func GetClaimByEndApply() ([]*models.Claim, error) {
	return mysql.SharedStore().GetClaimByEndApply(common.ClaimStatusNew, common.ApplyMinute)
}
func GetClaimByEndVote() ([]*models.Claim, error) {
	return mysql.SharedStore().GetClaimByEndVote(common.ClaimStatusArbiter, common.VoteMinute)
}

func GetClaimClose(statuses []common.ClaimStatus) ([]*models.Claim, error) {
	return mysql.SharedStore().GetClaimClose(statuses, common.CloseMinute)
}

func GetClaimTotal() (int, error) {
	return mysql.SharedStore().GetClaimTotal()
}
func GetClaim(userId string, product string, status common.ClaimStatus, beforeId, afterId int64, limit int) ([]*models.Claim, error) {
	return mysql.SharedStore().GetClaim(userId, product, status, beforeId, afterId, limit)
}
func GetClaimList(userId string, product string, status common.ClaimStatus, offset, limit int) ([]*models.Claim, error) {
	return mysql.SharedStore().GetClaimList(userId, product, status, offset, limit)
}
func GetClaimCount(settled bool) (int, error) {
	return mysql.SharedStore().GetClaimCount(settled)
}
func GetClaimByUserId(userId string) ([]*models.Claim, error) {
	return mysql.SharedStore().GetClaimByUserId(userId)
}
func GetClaimById(claimId int64) (*models.Claim, error) {
	return mysql.SharedStore().GetClaimById(claimId)
}

//func GetClaimVote(claimId int64) (*models.Claim, []*models.ClaimDesc, []*models.ClaimCred, error) {
//	claim, err := mysql.SharedStore().GetClaimById(claimId)
//	if err != nil {
//		return nil, nil, nil, err
//	}
//	desc, err := mysql.SharedStore().GetClaimDescByClaimId(claimId)
//	if err != nil {
//		return nil, nil, nil, err
//	}
//	cred, err := mysql.SharedStore().GetClaimCredByClaimId(claimId)
//	if err != nil {
//		return nil, nil, nil, err
//	}
//	return claim, desc, cred, nil
//}

func GetClaimResult(claimId int64) (*models.Claim, []*models.Vote, error) {
	claim, err := GetClaimById(claimId)
	if err != nil {
		return nil, nil, err
	}
	votes, err := GetVoteByClaimId(claimId)
	if err != nil {
		return nil, nil, err
	}

	return claim, votes, nil
}

func GetClaimByArbiter(arbiterId string, product string, status common.ClaimStatus, beforeId, afterId int64, limit int) ([]*models.Vote, *models.Qualification, error) {
	votes, err := mysql.SharedStore().GetVote(arbiterId, product, status, beforeId, afterId, limit)
	if err != nil {
		return nil, nil, err
	}

	qualification, err := GetQualificationByArbiterId(arbiterId)
	if err != nil {
		return nil, nil, err
	}

	return votes, qualification, nil
}

func ClaimApply(claimId int64, arbiterId string) (string, error) {
	// status := common.ApplyStatusApply
	applies, err := GetApplyByClaimId(claimId, nil)
	if err != nil {
		return "", err
	}
	for _, apply := range applies {
		if apply.ArbiterId == arbiterId {
			return "", errors.New("already apply")
		}
	}
	applyNum := len(applies)
	if applyNum >= common.ApplyMaxNum {
		return "", errors.New("already max apply num")
	}

	tx, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return "", err
	}
	defer func() { _ = tx.Rollback() }()

	qualifications, err := SubAvailableNum(tx, arbiterId, 1)
	if err != nil {
		return "", err
	}
	if qualifications == nil {
		return "", errors.New("qualifications error")
	}

	claim, err := tx.GetClaimForUpdate(claimId)
	if err != nil {
		return "", err
	}
	if claim.ApplyNum != applyNum {
		return "", errors.New("apply error")
	}
	claim.ApplyNum++
	applyTime := time.Now()
	apply := &models.Apply{
		ClaimId:   claimId,
		SubmitAt:  claim.SubmitAt,
		CoverId:   claim.CoverId,
		CoverHash: claim.CoverHash,
		UserId:    claim.UserId,
		Product:   claim.Product,
		ApplyAt:   applyTime,
		ArbiterAt: applyTime,
		ApplyNum:  claim.ApplyNum,
		ArbiterId: arbiterId,
		Status:    common.ApplyStatusApply,
	}

	if err = tx.AddApply(apply); err != nil {
		return "", err
	}

	if err = tx.UpdateClaim(claim); err != nil {
		return "", err
	}

	return string(apply.Status), tx.CommitTx()
}

func GetClaimByArbiterId(arbiterId string) ([]*models.Claim, error) {
	qualification, err := GetQualificationByArbiterId(arbiterId)
	if err != nil {
		return nil, errors.New("no available qualifications")
	}
	if qualification.Available <= 0 {
		return nil, errors.New("no available qualifications")
	}
	claims, err := GetClaimByApply()
	if err != nil {
		return nil, err
	}
	// status := common.ApplyStatusApply
	applies, err := GetApplyByArbiterId(arbiterId, nil)
	if err != nil {
		return nil, err
	}
	applied := make(map[int64]struct{})
	for _, apply := range applies {
		applied[apply.ClaimId] = struct{}{}
	}

	var available []*models.Claim
	i := 0
	for _, claim := range claims {
		if _, found := applied[claim.Id]; found {
			continue
		}
		available = append(available, claim)
		i++
		if i >= qualification.Available {
			break
		}
	}
	return available, nil
}

func ExecuteClaim(claimId int64) error {
	applies, err := GetApplyByClaimId(claimId, nil)
	if err != nil {
		return err
	}
	if len(applies) < common.ArbiterMaxNum {
		//return ReturnApply(claimId)
		log.GetLog().Warn("claimId apply num", zap.Int64("claimId", claimId), zap.Int("apply num", len(applies)))
		return fmt.Errorf("claimId(%v) apply num(%v)", claimId, len(applies))
	} else {
		return ExecuteApply(claimId)
	}
}
