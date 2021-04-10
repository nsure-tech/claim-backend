package mysql

import (
	"github.com/jinzhu/gorm"
	"nsure/vote/models"
	"time"
)

func (s *Store) AddPayment(payment *models.Payment) error {
	payment.CreatedAt = time.Now()
	return s.db.Create(payment).Error
}

func (s *Store) UpdatePayment(payment *models.Payment) error {
	payment.UpdatedAt = time.Now()
	return s.db.Save(payment).Error
}

func (s *Store) GetPaymentByEnd(endTime uint) ([]*models.Payment, error) {
	db := s.db.Where("settled=0 AND ADDDATE(created_at,interval ? minute) < now()", endTime)

	var payments []*models.Payment
	err := db.Find(&payments).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return payments, err
}
func (s *Store) GetPaymentForUpdate(id int64) (*models.Payment, error) {
	var payment models.Payment
	err := s.db.Raw("SELECT * FROM v_payment WHERE id=? AND settled=0 FOR UPDATE", id).Scan(&payment).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &payment, err
}

func (s *Store) GetPaymentByClaimForUpdate(claimId int64) (*models.Payment, error) {
	var payment models.Payment
	err := s.db.Raw("SELECT * FROM v_payment WHERE claim_id=? AND settled=0 FOR UPDATE", claimId).Scan(&payment).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &payment, err
}

func (s *Store) GetUnsettledPayments(count int) ([]*models.Payment, error) {
	db := s.db.Where("settled =?", 0).Order("id ASC").Limit(count)

	var payments []*models.Payment
	err := db.Find(&payments).Error
	return payments, err
}
