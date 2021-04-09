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

func GetQualificationByArbiterId(arbiterId string) (*models.Qualification, error) {
	return mysql.SharedStore().GetQualificationByArbiterId(arbiterId)
}

func GetQualificationAccount(arbiterId string) (*models.Qualification, *models.Account, error) {
	qualifications, err := GetQualificationByArbiterId(arbiterId)
	if err != nil {
		return nil, nil, err
	}
	account, err := GetBalanceByUserId(arbiterId)
	if err != nil {
		return nil, nil, err
	}
	return qualifications, account, nil
}

func ApplyQualification(arbiterId string, num int) (*models.Qualification, *models.Account, error) {
	holdSize := decimal.NewFromInt(int64(num)).Mul(utils.ArbiterNSure())
	db, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return nil, nil, err
	}
	defer func() { _ = db.Rollback() }()

	account, err := HoldBalance(db, arbiterId, common.CurrencyNSure, holdSize, common.BillTypeArbiter)
	if err != nil {
		return nil, nil, errors.New("insufficient Nsure balance\n\nTo add 1role, you need 5000 Nsure deposited.\n")
	}

	qualification, err := AddAvailableNum(db, arbiterId, num)
	if err != nil {
		return nil, nil, err
	}

	return qualification, account, db.CommitTx()
}

func PendingQualifications(arbiterId string, num int) (*models.Qualification, *models.Account, error) {
	db, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return nil, nil, err
	}
	defer func() { _ = db.Rollback() }()

	account, err := GetAccount(arbiterId, common.CurrencyNSure)
	if err != nil {
		return nil, nil, err
	}

	qualification, err := SubAvailableAddPending(db, arbiterId, num)
	if err != nil {
		return nil, nil, err
	}

	return qualification, account, db.CommitTx()
}

func AddAvailableNum(db models.Store, userId string, num int) (*models.Qualification, error) {
	qualification, err := db.GetQualificationForUpdate(userId)
	if err != nil {
		return nil, err
	}
	if qualification == nil {
		qualification = &models.Qualification{
			ArbiterId: userId,
			Available: num,
		}
		db.AddQualification(qualification)
	} else {
		qualification.Available += num
		db.UpdateQualification(qualification)
	}
	return qualification, nil
}

func SubAvailableNum(db models.Store, userId string, num int) (*models.Qualification, error) {
	enough, err := HasEnoughAvailable(userId, num)
	if err != nil {
		return nil, err
	}
	if !enough {
		return nil, errors.New(fmt.Sprintf("no enough userId=%v request=%v", userId, num))
	}

	qualification, err := db.GetQualificationForUpdate(userId)
	if err != nil {
		return nil, err
	}
	if qualification == nil {
		return nil, errors.New("no enough")
	}
	qualification.Available -= num
	qualification.Used += num
	db.UpdateQualification(qualification)
	return qualification, nil
}

func ReturnAvailableNum(db models.Store, arbiterId string, num int) (*models.Qualification, error) {
	qualification, err := db.GetQualificationForUpdate(arbiterId)
	if err != nil {
		return nil, err
	}
	if qualification == nil {
		return nil, errors.New("no enough")
	}
	qualification.Available += num
	qualification.Used -= num
	db.UpdateQualification(qualification)
	return qualification, nil
}

func ReturnAvailableAddClosed(db models.Store, arbiterId string, num int) (*models.Qualification, error) {
	qualification, err := db.GetQualificationForUpdate(arbiterId)
	if err != nil {
		return nil, err
	}
	if qualification == nil {
		return nil, errors.New("no enough")
	}
	qualification.Available += num
	qualification.Used -= num
	qualification.Closed += num
	db.UpdateQualification(qualification)
	return qualification, nil
}

func SubUsedAddClosed(db models.Store, arbiterId string, num int) (*models.Qualification, error) {
	qualification, err := db.GetQualificationForUpdate(arbiterId)
	if err != nil {
		return nil, err
	}
	if qualification == nil {
		return nil, errors.New("no enough")
	}
	qualification.Used -= num
	qualification.Closed += num
	db.UpdateQualification(qualification)
	return qualification, nil
}

func SubAvailableAddPending(db models.Store, userId string, num int) (*models.Qualification, error) {
	enough, err := HasEnoughAvailable(userId, num)
	if err != nil {
		return nil, err
	}
	if !enough {
		return nil, errors.New(fmt.Sprintf("no available roles to deduct"))
	}

	pending := &models.Pending{
		ArbiterId: userId,
		Pending:   num,
		SubmitAt:  time.Now(),
	}

	qualification, err := db.GetQualificationForUpdate(userId)
	if err != nil {
		return nil, err
	}
	if qualification == nil {
		return nil, errors.New("no enough")
	}
	qualification.Available -= num
	qualification.Pending += num

	db.AddPending(pending)
	db.UpdateQualification(qualification)

	return qualification, nil
}

func HasEnoughAvailable(userId string, num int) (bool, error) {
	qualification, err := GetQualificationByArbiterId(userId)
	if err != nil {
		return false, err
	}
	if qualification == nil {
		return false, nil
	}
	return qualification.Available >= num, nil
}

func GetUnsettledPends(count int) ([]*models.Pending, error) {
	return mysql.SharedStore().GetUnsettledPends(common.PendingMinute, count)
}

func ExecutePending(pending *models.Pending) error {
	tx, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	remainingFunds := decimal.NewFromInt(int64(pending.Pending)).Mul(utils.ArbiterNSure())
	err = AddDelayBill(tx, pending.ArbiterId, common.CurrencyNSure, remainingFunds, remainingFunds.Neg(),
		common.BillTypePending, "")
	if err != nil {
		return err
	}

	qualification, err := tx.GetQualificationForUpdate(pending.ArbiterId)
	if err != nil {
		return err
	}
	if qualification == nil {
		return errors.New("qualifications error")
	}
	qualification.Pending -= pending.Pending
	tx.UpdateQualification(qualification)

	pending.Settled = true
	err = tx.UpdatePending(pending)
	if err != nil {
		return err
	}

	return tx.CommitTx()
}
