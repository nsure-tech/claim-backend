package service

import (
	"github.com/shopspring/decimal"
	"nsure/vote/common"
	"nsure/vote/models"
	"nsure/vote/models/mysql"
)

func DepositTreasury(currency, fromAddress, toAddress, raw string, status common.TransferStatus, amount decimal.Decimal) error {
	transfer := &models.Transfer{
		Currency:    currency,
		FromAddress: fromAddress,
		ToAddress:   toAddress,
		Amount:      amount,
		Raw:         raw,
		Status:      status,
	}
	mysql.SharedStore().AddTransfer(transfer)

	tx, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	account, err := GetAccountForUpdate(tx, fromAddress, currency)
	if err != nil {
		return err
	}

	account.Available = account.Available.Add(amount)

	bill := &models.Bill{
		UserId:    fromAddress,
		Currency:  currency,
		Available: amount,
		Hold:      decimal.NewFromInt(0),
		Type:      common.BillTypeDeposit,
		Settled:   true,
		Notes:     "",
	}
	if err = tx.AddBill(bill); err != nil {
		return err
	}

	if err = tx.UpdateAccount(account); err != nil {
		return err
	}

	transfer.Settled = true
	if err = tx.UpdateTransfer(transfer); err != nil {
		return err
	}
	return tx.CommitTx()
}
