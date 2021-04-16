package service

import (
	"fmt"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"nsure/vote/common"
	"nsure/vote/contract"
	"nsure/vote/log"
	"nsure/vote/models"
	"nsure/vote/models/mysql"
	"nsure/vote/utils"
	"time"
)

func GetWithdrawNonceByUserId(userId string, status common.WithdrawStatus) (*common.Nonce, error) {
	return mysql.SharedStore().GetWithdrawNonceByUserId(userId, status)
}

func JudgeNonce(userId string, nonce uint64) bool {
	if chainNonce, err := contract.GetNonceByAddress(userId); err != nil {
		return false
	} else if chainNonce != nonce {
		return false
	}
	if seq, err := GetWithdrawNonceByUserId(userId, common.WithdrawStatusSuccess); err == nil {
		if seq != nil && seq.Nonce != nil {
			if nonce == *seq.Nonce+1 {
				return true
			}
		} else {
			if nonce == 0 {
				return true
			}
		}
	}
	return false
}

func AddWithdraw(userId, currencyContr string, amount decimal.Decimal, nonce uint64) (*common.WithdrawResult, error) {
	withdrawResult := &common.WithdrawResult{}
	if oldWithdraw, err := mysql.SharedStore().GetWithdrawByUserNonce(userId, nonce); err == nil && oldWithdraw != nil {
		return withdrawResult, fmt.Errorf("please wait")
	}
	if !JudgeNonce(userId, nonce) {
		return withdrawResult, fmt.Errorf("userId %v nonce %v error", userId, nonce)
	}
	currency, err := GetCurrency(currencyContr)
	if err != nil {
		return nil, fmt.Errorf("currency %v", currencyContr)
	}
	endAt := time.Now().Add(time.Duration(common.DurationClaim) * time.Minute)
	deadline := utils.I64ToA(endAt.Unix())

	if sigHash, _, err := contract.SignClaim(userId, currencyContr, amount.String(), utils.U64ToA(nonce), deadline); err != nil {
		log.GetLog().Error("contract SignClaim", zap.Error(err))
		return nil, err
	} else {
		if len(sigHash) != 65 {
			return nil, fmt.Errorf("sign error")
		}
		if sigHash[64] < 27 {
			sigHash[64] += 27
		}
		//withdrawResult.SigHash = sigHash.String()
		withdrawResult.R = sigHash[0:32].String()
		withdrawResult.S = sigHash[32:64].String()
		withdrawResult.V = sigHash[64:].String()
		withdrawResult.DeadLine = deadline
	}

	db, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return withdrawResult, err
	}
	defer func() { _ = db.Rollback() }()

	enough, err := HasEnoughBalance(userId, currency, amount)
	if err != nil {
		return nil, err
	}
	if !enough {
		return nil, fmt.Errorf("no enough %v : request=%v", currency, amount)
	}

	account, err := db.GetAccountForUpdate(userId, currency)
	if err != nil {
		return nil, err
	}

	account.Available = account.Available.Sub(amount)
	account.Hold = account.Hold.Add(amount)

	bill := &models.Bill{
		UserId:    userId,
		Currency:  currency,
		Available: amount.Neg(),
		Hold:      amount,
		Type:      common.BillTypeWithdraw,
		Settled:   true,
		Notes:     "",
	}
	if err = db.AddBill(bill); err != nil {
		return nil, err
	}

	if err = db.UpdateAccount(account); err != nil {
		return nil, err
	}

	withDraw := &models.Withdraw{
		UserId:   userId,
		Currency: currency,
		Amount:   amount,
		Nonce:    nonce,
		Status:   common.WithdrawStatusApply,
		EndAt:    endAt,
	}
	if err = db.AddWithdraw(withDraw); err != nil {
		log.GetLog().Error("Add Withdraw", zap.Error(err))
		return nil, err
	}
	if err = db.CommitTx(); err != nil {
		return nil, err
	}
	utils.AlarmMessage(userId, currency, amount, nonce)
	return withdrawResult, nil
}

func AddChainClaim(userId, currencyContr string, amount decimal.Decimal, nonce uint64, raw string) error {
	var claimErr error = nil
	currency, err := GetCurrency(currencyContr)
	if err != nil {
		claimErr = fmt.Errorf("%v", err)
	} else {
		if err = addClaimUpdateWithdraw(userId, currency, amount, nonce, raw); err != nil {
			claimErr = fmt.Errorf("%v %v", claimErr, err)
		}
	}

	if claimErr != nil {
		clainClaim := &models.ChainClaim{
			UserId:   userId,
			Currency: currency,
			Amount:   amount,
			Nonce:    nonce,
			Raw:      raw,
			Status:   common.WithdrawStatusFail,
			Settled:  true,
		}
		if err = mysql.SharedStore().AddChainClaim(clainClaim); err != nil {
			claimErr = fmt.Errorf("%v %v", claimErr, err)
		}
	}
	return claimErr

}

func addClaimUpdateWithdraw(userId, currency string, amount decimal.Decimal, nonce uint64, raw string) error {
	clainClaim := &models.ChainClaim{
		UserId:   userId,
		Currency: currency,
		Amount:   amount,
		Nonce:    nonce,
		Raw:      raw,
		Status:   common.WithdrawStatusSuccess,
		Settled:  true,
	}

	db, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return err
	}
	defer func() { _ = db.Rollback() }()

	withdraw, err := db.GetWithdrawByUserNonce(userId, nonce)
	if err != nil {
		return err
	}

	if !withdraw.Amount.Equal(amount) {
		return fmt.Errorf("amount %v!=%v", withdraw.Amount, amount)
	}

	if err = db.AddChainClaim(clainClaim); err != nil {
		return err
	}

	withdraw.Status = common.WithdrawStatusSuccess
	withdraw.Settled = true
	if err = db.UpdateWithdraw(withdraw); err != nil {
		return err
	}
	bill := &models.Bill{
		UserId:    userId,
		Currency:  currency,
		Available: decimal.Zero,
		Hold:      amount.Neg(),
		Type:      common.BillTypeWithdrawLock,
		Settled:   false,
		Notes:     "",
	}
	if err = db.AddBill(bill); err != nil {
		return err
	}

	return db.CommitTx()
}

func GetWithdrawsByEndAt(endTime uint) ([]*models.Withdraw, error) {
	return mysql.SharedStore().GetWithdrawsByEndAt(common.WithdrawStatusApply, endTime)
}

func ExecuteWithdraw(withdraw *models.Withdraw) error {
	db, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return err
	}
	defer func() { _ = db.Rollback() }()

	bill := &models.Bill{
		UserId:    withdraw.UserId,
		Currency:  withdraw.Currency,
		Available: withdraw.Amount,
		Hold:      withdraw.Amount.Neg(),
		Type:      common.BillTypeWithdrawBack,
		Settled:   false,
		Notes:     "",
	}
	if err = db.AddBill(bill); err != nil {
		return err
	}
	withdraw.Status = common.WithdrawStatusFail
	withdraw.Settled = true
	if err = db.UpdateWithdraw(withdraw); err != nil {
		return err
	}

	return db.CommitTx()
}

func GetRewardByArbiterId(arbiterId string) decimal.Decimal {
	if reward, err := mysql.SharedStore().GetRewardByArbiterId(arbiterId); err != nil {
		return decimal.Zero
	} else {
		return reward.Reward
	}
}

func GetRewardTotal() decimal.Decimal {
	if reward, err := mysql.SharedStore().GetRewardTotal(); err != nil {
		log.GetLog().Error("mysql GetRewardTotal()", zap.Error(err))
		return decimal.Zero
	} else {
		log.GetLog().Debug("mysql GetRewardTotal()", zap.String("reward", reward.Reward.String()))
		return reward.Reward
	}
}
