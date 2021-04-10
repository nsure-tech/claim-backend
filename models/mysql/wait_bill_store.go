package mysql

import (
	"nsure/vote/models"
	"time"
)

func (s *Store) AddWaitBill(bill *models.WaitBill) error {
	bill.CreatedAt = time.Now()
	return s.db.Create(bill).Error
}

func (s *Store) GetUnsettledWaitBills(count int) ([]*models.WaitBill, error) {
	db := s.db.Where("settled =? AND end_at<now()", 0).Order("id ASC").Limit(count)
	var bills []*models.WaitBill
	err := db.Find(&bills).Error
	return bills, err
}

func (s *Store) UpdateWaitBill(bill *models.WaitBill) error {
	bill.UpdatedAt = time.Now()
	return s.db.Save(bill).Error
}
