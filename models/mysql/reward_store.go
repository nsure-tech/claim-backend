package mysql

import (
	"github.com/jinzhu/gorm"
	"nsure/vote/common"
	"nsure/vote/models"
	"time"
)

func (s *Store) AddReward(reward *models.Reward) error {
	reward.CreatedAt = time.Now()
	return s.db.Create(reward).Error
}

func (s *Store) UpdateReward(reward *models.Reward) error {
	reward.UpdatedAt = time.Now()
	return s.db.Save(reward).Error
}

func (s *Store) GetRewardByEnd(endTime uint) ([]*models.Reward, error) {
	db := s.db.Where("settled=0 AND ADDDATE(created_at,interval ? minute) < now()", endTime)

	var rewards []*models.Reward
	err := db.Find(&rewards).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return rewards, err
}
func (s *Store) GetRewardForUpdate(id int64) (*models.Reward, error) {
	var reward models.Reward
	err := s.db.Raw("SELECT * FROM v_reward WHERE id=?  AND settled=0 FOR UPDATE", id).Scan(&reward).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &reward, err
}

func (s *Store) GetRewardTotal() (*common.Reward, error) {
	var reward common.Reward
	err := s.db.Table("v_reward").Select("SUM(amount) as reward").Scan(&reward).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &reward, err
}

func (s *Store) GetRewardByArbiterId(arbiterId string) (*common.Reward, error) {
	var reward common.Reward
	err := s.db.Table("v_reward").Select("SUM(amount) as reward").Where("settled=1 AND arbiter_id = ?", arbiterId).Scan(&reward).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &reward, err
}

func (s *Store) GetUnsettledRewards(count int) ([]*models.Reward, error) {
	db := s.db.Where("settled =?", 0).Order("id ASC").Limit(count)

	var rewards []*models.Reward
	err := db.Find(&rewards).Error
	return rewards, err
}

func (s *Store) AddRewardCha(rewardCha *models.RewardCha) error {
	rewardCha.CreatedAt = time.Now()
	return s.db.Create(rewardCha).Error
}

func (s *Store) UpdateRewardCha(rewardCha *models.RewardCha) error {
	rewardCha.UpdatedAt = time.Now()
	return s.db.Save(rewardCha).Error
}

func (s *Store) GetRewardChaByEnd(endTime uint) ([]*models.RewardCha, error) {
	db := s.db.Where("settled=0 AND ADDDATE(created_at,interval ? minute) < now()", endTime)

	var rewardChas []*models.RewardCha
	err := db.Find(&rewardChas).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return rewardChas, err
}
func (s *Store) GetRewardChaForUpdate(id int64) (*models.RewardCha, error) {
	var rewardCha models.RewardCha
	err := s.db.Raw("SELECT * FROM v_reward_cha WHERE id=? AND settled=0 FOR UPDATE", id).Scan(&rewardCha).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &rewardCha, err
}

func (s *Store) AddPunishCha(punishCha *models.PunishCha) error {
	punishCha.CreatedAt = time.Now()
	return s.db.Create(punishCha).Error
}

func (s *Store) UpdatePunishCha(punishCha *models.PunishCha) error {
	punishCha.UpdatedAt = time.Now()
	return s.db.Save(punishCha).Error
}
func (s *Store) GetPunishChaByEnd(endTime uint) ([]*models.PunishCha, error) {
	db := s.db.Where("settled=0 AND ADDDATE(created_at,interval ? minute) < now()", endTime)

	var punishChas []*models.PunishCha
	err := db.Find(&punishChas).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return punishChas, err
}
func (s *Store) GetPunishChaForUpdate(id int64) (*models.PunishCha, error) {
	var punishCha models.PunishCha
	err := s.db.Raw("SELECT * FROM v_punish_cha WHERE id=? AND settled=0 FOR UPDATE", id).Scan(&punishCha).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &punishCha, err
}
