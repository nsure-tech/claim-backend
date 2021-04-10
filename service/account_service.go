package service

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"nsure/vote/common"
	"nsure/vote/models"
	"nsure/vote/models/mysql"
)

func AccountDeposit(userId, currency string, amount decimal.Decimal) (*models.Account, error) {
	tx, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()

	account, err := GetAccountForUpdate(tx, userId, currency)
	if err != nil {
		return nil, err
	}

	account.Available = account.Available.Add(amount)

	bill := &models.Bill{
		UserId:    userId,
		Currency:  currency,
		Available: amount,
		Hold:      decimal.NewFromInt(0),
		Type:      common.BillTypeDeposit,
		Settled:   true,
		Notes:     "",
	}
	err = tx.AddBill(bill)
	if err != nil {
		return nil, err
	}

	err = tx.UpdateAccount(account)
	if err != nil {
		return nil, err
	}
	return account, tx.CommitTx()
}

func AccountWithdraw(userId, currency string, amount decimal.Decimal) (*models.Account, error) {
	tx, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()

	enough, err := HasEnoughBalance(userId, currency, amount)
	if err != nil {
		return nil, err
	}
	if !enough {
		return nil, errors.New(fmt.Sprintf("no enough %v : request=%v", currency, amount))
	}

	account, err := tx.GetAccountForUpdate(userId, currency)
	if err != nil {
		return nil, err
	}

	account.Available = account.Available.Sub(amount)

	bill := &models.Bill{
		UserId:    userId,
		Currency:  currency,
		Available: amount.Neg(),
		Hold:      decimal.NewFromInt(0),
		Type:      common.BillTypeWithdraw,
		Settled:   true,
		Notes:     "",
	}
	err = tx.AddBill(bill)
	if err != nil {
		return nil, err
	}

	err = tx.UpdateAccount(account)
	if err != nil {
		return nil, err
	}
	return account, tx.CommitTx()
}

func GetBalanceByUserId(userId string) (*models.Account, error) {
	return GetAccount(userId, common.CurrencyNSure)
}

func HoldBalance(db models.Store, userId string, currency string, amount decimal.Decimal, billType common.BillType) (*models.Account, error) {
	if amount.LessThanOrEqual(decimal.Zero) {
		return nil, errors.New("size less than 0")
	}

	enough, err := HasEnoughBalance(userId, currency, amount)
	if err != nil {
		return nil, err
	}
	if !enough {
		return nil, errors.New(fmt.Sprintf("no enough %v : request=%v", currency, amount))
	}

	account, err := db.GetAccountForUpdate(userId, currency)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, errors.New("no enough")
	}

	account.Available = account.Available.Sub(amount)
	account.Hold = account.Hold.Add(amount)

	bill := &models.Bill{
		UserId:    userId,
		Currency:  currency,
		Available: amount.Neg(),
		Hold:      amount,
		Type:      billType,
		Settled:   true,
		Notes:     "",
	}
	err = db.AddBill(bill)
	if err != nil {
		return nil, err
	}

	err = db.UpdateAccount(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
func HasEnoughBalance(userId string, currency string, amount decimal.Decimal) (bool, error) {
	account, err := GetAccount(userId, currency)
	if err != nil {
		return false, err
	}
	if account == nil {
		return false, nil
	}
	return account.Available.GreaterThanOrEqual(amount), nil
}

func GetAccount(userId string, currency string) (*models.Account, error) {
	return mysql.SharedStore().GetAccount(userId, currency)
}

func GetAccountForUpdate(store models.Store, userId string, currency string) (*models.Account, error) {
	account, err := store.GetAccountForUpdate(userId, currency)
	if err != nil {
		return nil, err
	}

	if account == nil {
		err = store.AddAccount(&models.Account{
			UserId:   userId,
			Currency: currency,
		})
		if err != nil {
			return nil, err
		}
		account, err = store.GetAccountForUpdate(userId, currency)
		if err != nil {
			return nil, err
		}
	}
	return account, nil
}
