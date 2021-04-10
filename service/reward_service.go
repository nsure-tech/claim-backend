package service

import (
	"fmt"
	"github.com/shopspring/decimal"
	"nsure/vote/common"
	"nsure/vote/models"
	"nsure/vote/models/mysql"
)

func GetRewardByEnd(endTime uint) ([]*models.Reward, error) {
	return mysql.SharedStore().GetRewardByEnd(endTime)
}
func GetPunishChaByEnd(endTime uint) ([]*models.PunishCha, error) {
	return mysql.SharedStore().GetPunishChaByEnd(endTime)
}
func GetRewardChaByEnd(endTime uint) ([]*models.RewardCha, error) {
	return mysql.SharedStore().GetRewardChaByEnd(endTime)
}

func ExecuteReward(id int64) error {
	tx, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	reward, err := tx.GetRewardForUpdate(id)
	if err != nil {
		return err
	}
	if reward == nil {
		return fmt.Errorf("reward is nil")
	}
	if err := AddBill(tx, common.AccountNSure, reward.Currency, reward.Amount.Neg(), decimal.Zero,
		common.BillTypeReward, ""); err != nil {
		return err
	}
	if err := AddBill(tx, reward.ArbiterId, reward.Currency, reward.Amount, decimal.Zero,
		common.BillTypeReward, ""); err != nil {
		return err
	}
	reward.Settled = true
	if err := tx.UpdateReward(reward); err != nil {
		return err
	}
	return tx.CommitTx()
}

func ExecuteRewardCha(id int64) error {
	tx, err := mysql.SharedStore().BeginTx()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	rewardCha, err := tx.GetRewardChaForUpdate(id)
	if err != nil {
		return err
	}
	if rewardCha == nil {
		return fmt.Errorf("rewardCha is nil")
	}
	if err := AddBill(tx, common.AccountChallenge, rewardCha.Currency, rewardCha.Amount.Neg(), decimal.Zero,
		common.BillTypeReward, ""); err != nil {
		return err
	}
	if err := AddBill(tx, rewardCha.ChallengeId, rewardCha.Currency, rewardCha.Amount, decimal.Zero,
		common.BillTypeReward, ""); err != nil {
		return err
	}
	rewardCha.Settled = true
	if err := tx.UpdateRewardCha(rewardCha); err != nil {
		return err
	}
	return tx.CommitTx()
}
