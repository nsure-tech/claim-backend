package service

import (
	"github.com/shopspring/decimal"
	"nsure/vote/common"
	"nsure/vote/models"
	"nsure/vote/models/mysql"
	"time"
)

func ExecuteBill(bill *models.Bill) error {
	tx, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	account, err := GetAccountForUpdate(tx, bill.UserId, bill.Currency)
	if err != nil {
		return err
	}

	account.Available = account.Available.Add(bill.Available)
	account.Hold = account.Hold.Add(bill.Hold)
	bill.Settled = true

	err = tx.UpdateBill(bill)
	if err != nil {
		return err
	}

	err = tx.UpdateAccount(account)
	if err != nil {
		return err
	}

	return tx.CommitTx()
}

func AddBill(store models.Store, userId string, currency string, available, hold decimal.Decimal, billType common.BillType, notes string) error {
	bill := &models.Bill{
		UserId:    userId,
		Currency:  currency,
		Available: available,
		Hold:      hold,
		Type:      billType,
		Settled:   false,
		Notes:     notes,
	}
	return store.AddBill(bill)
}
func GetUnsettledBills(count int) ([]*models.Bill, error) {
	return mysql.SharedStore().GetUnsettledBills(count)
}

func GetBillsCountByUserId(userId string, statuses []common.BillType) (int, error) {
	return mysql.SharedStore().GetBillsCountByUserId(userId, statuses)
}
func GetBillsByUserId(userId string, statuses []common.BillType, offset, limit int) ([]*models.Bill, error) {
	return mysql.SharedStore().GetBillsByUserId(userId, statuses, offset, limit)
}

func AddWaitBill(store models.Store, userId string, currency string, available, hold decimal.Decimal, billType common.BillType, notes string) error {
	bill := &models.WaitBill{
		UserId:    userId,
		Currency:  currency,
		Available: available,
		Hold:      hold,
		Type:      billType,
		Settled:   false,
		EndAt:     time.Now().Add(time.Duration(common.BillMinute) * time.Minute),
		Notes:     notes,
	}
	return store.AddWaitBill(bill)
}
func GetUnsettledWaitBills(count int) ([]*models.WaitBill, error) {
	return mysql.SharedStore().GetUnsettledWaitBills(count)
}
func ExecuteWaitBill(waitBill *models.WaitBill) error {
	tx, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	bill := &models.Bill{
		UserId:    waitBill.UserId,
		Currency:  waitBill.Currency,
		Available: waitBill.Available,
		Hold:      waitBill.Hold,
		Type:      waitBill.Type,
		Settled:   false,
		Notes:     waitBill.Notes,
	}
	if err := tx.AddBill(bill); err != nil {
		return err
	}

	waitBill.Settled = true
	err = tx.UpdateWaitBill(waitBill)
	if err != nil {
		return err
	}

	return tx.CommitTx()
}
