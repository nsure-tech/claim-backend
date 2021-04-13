package mysql

import (
	"github.com/jinzhu/gorm"
	"nsure/vote/common"
	"nsure/vote/models"
	"time"
)

func (s *Store) GetWithdrawNonceByUserId(userId string, status common.WithdrawStatus) (*common.Nonce, error) {
	var nonce common.Nonce
	err := s.db.Table("v_withdraw").Select("max(nonce) as nonce").Where("user_id=? AND status=?", userId, status).Scan(&nonce).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &nonce, err
}
func (s *Store) GetWithdrawsByEndAt(status common.WithdrawStatus, endTime uint) ([]*models.Withdraw, error) {
	db := s.db.Where("settled=0 AND status=?", status)
	db = db.Where("ADDDATE(end_at,interval ? minute) < now()", endTime)

	var withdraws []*models.Withdraw
	err := db.Find(&withdraws).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return withdraws, err
}

func (s *Store) GetWithdrawsByUserId(userId string) ([]*models.Withdraw, error) {
	db := s.db.Where("user_id=?", userId)

	var withdraws []*models.Withdraw
	err := db.Find(&withdraws).Error
	return withdraws, err
}

func (s *Store) GetWithdrawByUserNonce(userId string, nonce uint64) (*models.Withdraw, error) {
	var withdraw models.Withdraw
	err := s.db.Raw("SELECT * FROM v_withdraw WHERE user_id=? AND nonce=? AND settled=0 ", userId, nonce).Scan(&withdraw).Error
	return &withdraw, err
}

func (s *Store) GetWithdrawForUpdate(userId string, nonce uint64) (*models.Withdraw, error) {
	var withdraw models.Withdraw
	err := s.db.Raw("SELECT * FROM v_withdraw WHERE user_id=? AND nonce=? AND settled=0 FOR UPDATE", userId, nonce).Scan(&withdraw).Error
	return &withdraw, err
}

func (s *Store) AddWithdraw(withdraw *models.Withdraw) error {
	withdraw.CreatedAt = time.Now()
	return s.db.Create(withdraw).Error
}

func (s *Store) UpdateWithdraw(withdraw *models.Withdraw) error {
	withdraw.UpdatedAt = time.Now()
	return s.db.Save(withdraw).Error
}

func (s *Store) AddChainClaim(chainClaim *models.ChainClaim) error {
	chainClaim.CreatedAt = time.Now()
	return s.db.Create(chainClaim).Error
}

func (s *Store) UpdateChainClaim(chainClaim *models.ChainClaim) error {
	chainClaim.UpdatedAt = time.Now()
	return s.db.Save(chainClaim).Error
}
