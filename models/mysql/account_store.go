package mysql

import (
	"github.com/jinzhu/gorm"
	"nsure/vote/models"
	"time"
)

func (s *Store) GetAccount(userId string, currency string) (*models.Account, error) {
	var account models.Account
	err := s.db.Raw("SELECT * FROM v_account WHERE user_id=? AND currency=?", userId,
		currency).Scan(&account).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &account, err
}

func (s *Store) GetAccountsByArbiterId(userId string) ([]*models.Account, error) {
	db := s.db.Where("user_id=?", userId)

	var accounts []*models.Account
	err := db.Find(&accounts).Error
	return accounts, err
}

func (s *Store) GetAccountForUpdate(userId string, currency string) (*models.Account, error) {
	var account models.Account
	err := s.db.Raw("SELECT * FROM v_account WHERE user_id=? AND currency=? FOR UPDATE", userId, currency).Scan(&account).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &account, err
}

func (s *Store) AddAccount(account *models.Account) error {
	account.CreatedAt = time.Now()
	return s.db.Create(account).Error
}

func (s *Store) UpdateAccount(account *models.Account) error {
	account.UpdatedAt = time.Now()
	return s.db.Save(account).Error
}
