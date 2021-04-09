package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"nsure/vote/common"
	"nsure/vote/models"
	"strings"
	"time"
)

func (s *Store) AddBills(bills []*models.Bill) error {
	if len(bills) == 0 {
		return nil
	}
	var valueStrings []string
	for _, bill := range bills {
		valueString := fmt.Sprintf("(NOW(),'%v', '%v', %v, %v, '%v', %v, '%v')",
			bill.UserId, bill.Currency, bill.Available, bill.Hold, bill.Type, bill.Settled, bill.Notes)
		valueStrings = append(valueStrings, valueString)
	}
	sql := fmt.Sprintf("INSERT INTO v_bill (created_at, user_id,currency,available,hold, type,settled,notes) VALUES %s", strings.Join(valueStrings, ","))
	return s.db.Exec(sql).Error
}

func (s *Store) AddBill(bill *models.Bill) error {
	bill.CreatedAt = time.Now()
	return s.db.Create(bill).Error
}

func (s *Store) GetUnsettledBills(count int) ([]*models.Bill, error) {
	db := s.db.Where("settled =?", 0).Order("id ASC").Limit(count)

	var bills []*models.Bill
	err := db.Find(&bills).Error
	return bills, err
}

func (s *Store) GetBillsCountByUserId(userId string, statuses []common.BillType) (int, error) {
	var count int
	db := s.db.Model(&models.Bill{}).Where("settled = ?", 1)
	db = db.Where("user_id =?", userId)
	db = db.Where("available !=0")

	if len(statuses) != 0 {
		db = db.Where("type IN (?)", statuses)
	}
	err := db.Count(&count).Error
	if err == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return count, err
}

func (s *Store) GetBillsByUserId(userId string, statuses []common.BillType, offset, limit int) ([]*models.Bill, error) {
	db := s.db.Where("settled =?", 1)
	db = db.Where("user_id =?", userId)
	db = db.Where("available !=0")

	if len(statuses) != 0 {
		db = db.Where("type IN (?)", statuses)
	}
	if limit <= 0 {
		limit = 100
	}
	db = db.Order("id DESC")
	if offset > 0 {
		db = db.Offset(offset).Limit(limit)
	} else {
		db = db.Limit(limit)
	}

	var bills []*models.Bill
	err := db.Find(&bills).Error
	return bills, err
}

func (s *Store) UpdateBill(bill *models.Bill) error {
	bill.UpdatedAt = time.Now()
	return s.db.Save(bill).Error
}
