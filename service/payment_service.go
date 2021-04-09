package service

import (
	"fmt"
	"github.com/shopspring/decimal"
	"nsure/vote/common"
	"nsure/vote/models"
	"nsure/vote/models/mysql"
)

func GetPaymentByEnd(endTime uint) ([]*models.Payment, error) {
	return mysql.SharedStore().GetPaymentByEnd(endTime)
}

func ExecutePayment(id int64) error {
	tx, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	payment, err := tx.GetPaymentForUpdate(id)
	if err != nil {
		return err
	}
	if payment == nil {
		return fmt.Errorf("payment is nil")
	}
	if err := AddDelayBill(tx, common.AccountPayment, payment.Currency, payment.Amount.Neg(), decimal.Zero,
		common.BillTypeReward, ""); err != nil {
		return err
	}
	if err := AddDelayBill(tx, payment.UserId, payment.Currency, payment.Amount, decimal.Zero,
		common.BillTypeReward, ""); err != nil {
		return err
	}
	payment.Settled = true
	if err := tx.UpdatePayment(payment); err != nil {
		return err
	}
	return tx.CommitTx()
}
