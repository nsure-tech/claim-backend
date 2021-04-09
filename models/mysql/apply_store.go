package mysql

import (
	"github.com/jinzhu/gorm"
	"nsure/vote/common"
	"nsure/vote/models"
	"time"
)

func (s *Store) AddApply(apply *models.Apply) error {
	apply.CreatedAt = time.Now()
	return s.db.Create(apply).Error
}

func (s *Store) GetApplyCountByArbiterId(arbiterId string, status common.ApplyStatus) (int, error) {
	var count int
	db := s.db.Model(&models.Apply{}).Where("arbiter_id=?", arbiterId)
	if len(status) != 0 {
		db = db.Where("status=?", status)
	}
	err := db.Count(&count).Error
	if err == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return count, err
}

func (s *Store) GetApplyListByArbiterId(arbiterId string, status common.ApplyStatus, begin, limit int) ([]*models.Apply, error) {
	db := s.db.Where("arbiter_id=?", arbiterId)
	if len(status) != 0 {
		db = db.Where("status=?", status)
	}
	if limit <= 0 {
		limit = 100
	}

	db = db.Order("id DESC")
	if begin > 0 {
		db.Offset(begin).Limit(limit)
	} else {
		db.Limit(limit)
	}

	var applies []*models.Apply
	err := db.Find(&applies).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return applies, err
}

func (s *Store) GetApplyByArbiterId(arbiterId string, status *common.ApplyStatus) ([]*models.Apply, error) {
	db := s.db.Where("arbiter_id=?", arbiterId)
	if status != nil {
		db = db.Where("status=?", status)
	}

	var applies []*models.Apply
	err := db.Find(&applies).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return applies, err
}

func (s *Store) GetApplyByClaimId(claimId int64, status *common.ApplyStatus) ([]*models.Apply, error) {
	db := s.db.Where("claim_id=?", claimId)
	if status != nil {
		db = db.Where("status=?", status)
	}

	var applies []*models.Apply
	err := db.Find(&applies).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return applies, err
}

func (s *Store) GetApplyByApplyNum(applyNum uint8, status *common.ApplyStatus) ([]*models.Apply, error) {
	db := s.db.Where("settled=0 AND apply_num=?", applyNum)
	if status != nil {
		db = db.Where("status=?", status)
	}

	var applies []*models.Apply
	err := db.Find(&applies).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return applies, err
}

func (s *Store) GetApplyForUpdate(claimId int64, arbiterId string) (*models.Apply, error) {
	var apply models.Apply
	err := s.db.Raw("SELECT * FROM v_apply WHERE claim_id=? and arbiter_id=?  FOR UPDATE", claimId, arbiterId).Scan(&apply).Error
	return &apply, err
}
func (s *Store) UpdateApply(apply *models.Apply) error {
	apply.UpdatedAt = time.Now()
	return s.db.Save(apply).Error
}
