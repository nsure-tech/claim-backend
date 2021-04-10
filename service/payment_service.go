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

func GetPaymentUnsettled() ([]*models.Payment, error) {
	return mysql.SharedStore().GetUnsettledPayments(common.PaymentCount)
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
		common.BillTypePayment, ""); err != nil {
		return err
	}
	if err := AddDelayBill(tx, payment.UserId, payment.Currency, payment.Amount, decimal.Zero,
		common.BillTypePayment, ""); err != nil {
		return err
	}
	payment.Settled = true
	if err := tx.UpdatePayment(payment); err != nil {
		return err
	}
	return tx.CommitTx()
}

func PaymentByAdmin(adminId string, claimId int64, pay decimal.Decimal) (bool, error) {
	// todo if !ChallengeAddress(adminId){return false, fmt.Errorf("%v isn't payment admin address", adminId)}
	tx, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return false, err
	}
	defer func() { _ = tx.Rollback() }()

	payment, err := tx.GetPaymentByClaimForUpdate(claimId)
	if err != nil {
		return false, err
	}
	if payment == nil {
		return false, fmt.Errorf("payment is already finish")
	}
	if err = AddDelayBill(tx, common.AccountPayment, payment.Currency, pay.Neg(), decimal.Zero,
		common.BillTypePayment, ""); err != nil {
		return false, err
	}
	if err = AddDelayBill(tx, payment.UserId, payment.Currency, pay, decimal.Zero,
		common.BillTypePayment, ""); err != nil {
		return false, err
	}
	payment.Pay = pay
	payment.AdminId = adminId
	payment.Settled = true
	if err = tx.UpdatePayment(payment); err != nil {
		return false, err
	}

	return true, tx.CommitTx()
}
