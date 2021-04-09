package mysql

import (
	"github.com/jinzhu/gorm"
	"nsure/vote/models"
	"time"
)

func (s *Store) AddQualification(qualifications *models.Qualification) error {
	qualifications.CreatedAt = time.Now()
	return s.db.Create(qualifications).Error
}

func (s *Store) GetQualificationByArbiterId(arbiterId string) (*models.Qualification, error) {
	db := s.db.Where("arbiter_id=?", arbiterId)
	var qualifications models.Qualification
	err := db.Find(&qualifications).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &qualifications, err
}

func (s *Store) GetQualificationForUpdate(arbiterId string) (*models.Qualification, error) {
	var qualifications models.Qualification
	err := s.db.Raw("SELECT * FROM v_qualification WHERE arbiter_id=? FOR UPDATE", arbiterId).Scan(&qualifications).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &qualifications, err
}
func (s *Store) UpdateQualification(qualifications *models.Qualification) error {
	qualifications.UpdatedAt = time.Now()
	return s.db.Save(qualifications).Error
}

func (s *Store) AddPending(pending *models.Pending) error {
	pending.CreatedAt = time.Now()
	return s.db.Create(pending).Error
}
func (s *Store) GetPendingByArbiterId(arbiterId string) (*models.Pending, error) {
	db := s.db.Where("arbiter_id=?", arbiterId)
	var pending models.Pending
	err := db.Find(&pending).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &pending, err
}

func (s *Store) GetPendingForUpdate(arbiterId string) (*models.Pending, error) {
	var pending models.Pending
	err := s.db.Raw("SELECT * FROM v_pending WHERE arbiter_id=? FOR UPDATE", arbiterId).Scan(&pending).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &pending, err
}

func (s *Store) UpdatePending(pending *models.Pending) error {
	pending.UpdatedAt = time.Now()
	return s.db.Save(pending).Error
}

func (s *Store) GetUnsettledPends(pendingTime uint, count int) ([]*models.Pending, error) {
	db := s.db.Where("settled =?", 0)
	db = db.Where("ADDDATE(submit_at,interval ? minute) < now()", pendingTime)
	db = db.Order("id ASC").Limit(count)

	var pends []*models.Pending
	err := db.Find(&pends).Error
	return pends, err
}
