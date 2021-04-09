package mysql

import (
	"github.com/jinzhu/gorm"
	"nsure/vote/common"
	"nsure/vote/models"
	"time"
)

func (s *Store) GetClaimByApply(status common.ClaimStatus, applyTime uint, applyNum uint8) ([]*models.Claim, error) {
	db := s.db.Where("status=?", status)
	db = db.Where("apply_num<?", applyNum)
	db = db.Where("settled=0 AND ADDDATE(submit_at,interval ? minute) > now()", applyTime)

	var claims []*models.Claim
	err := db.Find(&claims).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return claims, err
}

func (s *Store) GetClaimTotal() (int, error) {
	var count int
	err := s.db.Model(&models.Claim{}).Count(&count).Error
	if err == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return count, err
}

func (s *Store) GetClaimCount(settled bool) (int, error) {
	var count int
	err := s.db.Model(&models.Claim{}).Where("settled = ?", settled).Count(&count).Error
	if err == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return count, err
}

func (s *Store) GetClaimByEndApply(status common.ClaimStatus, applyTime uint) ([]*models.Claim, error) {
	db := s.db.Where("status=?", status)
	db = db.Where("settled=0 AND ADDDATE(submit_at,interval ? minute) < now()", applyTime)

	var claims []*models.Claim
	err := db.Find(&claims).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return claims, err
}

func (s *Store) GetClaimByEndVote(status common.ClaimStatus, voteTime uint) ([]*models.Claim, error) {
	db := s.db.Where("status=?", status)
	db = db.Where("settled=0 AND ADDDATE(arbiter_at, interval ? minute) < now()", voteTime)

	var claims []*models.Claim
	err := db.Find(&claims).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return claims, err
}

func (s *Store) GetClaimClose(statuses []common.ClaimStatus, paymentTime uint) ([]*models.Claim, error) {
	db := s.db.Where("settled=0 AND ADDDATE(vote_at, interval ? minute) < now()", paymentTime)
	if len(statuses) != 0 {
		db = db.Where("status IN (?)", statuses)
	}

	var claims []*models.Claim
	err := db.Find(&claims).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return claims, err
}

func (s *Store) GetClaimForUpdate(claimId int64) (*models.Claim, error) {
	var claim models.Claim
	err := s.db.Raw("SELECT * FROM v_claim WHERE id=? AND settled=0 FOR UPDATE", claimId).Scan(&claim).Error
	return &claim, err
}

func (s *Store) GetClaimByHash(coverHash string) (*models.Claim, error) {
	var claim models.Claim
	err := s.db.Raw("SELECT * FROM v_claim WHERE cover_hash=? ", coverHash).Scan(&claim).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &claim, err
}

func (s *Store) GetClaimById(claimId int64) (*models.Claim, error) {
	var claim models.Claim
	err := s.db.Raw("SELECT * FROM v_claim WHERE id=?", claimId).Scan(&claim).Error
	return &claim, err
}

func (s *Store) GetClaimByUserId(userId string) ([]*models.Claim, error) {
	db := s.db.Where("user_id=?", userId)

	var claims []*models.Claim
	err := db.Find(&claims).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return claims, err
}

func (s *Store) GetClaimList(userId string, product string, status common.ClaimStatus, offset, limit int) ([]*models.Claim, error) {
	db := s.db

	if len(userId) != 0 {
		db = db.Where("user_id=?", userId)
	}
	if len(product) != 0 {
		db = db.Where("product=?", product)
	}
	if len(status) != 0 {
		db = db.Where("status=?", status)
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

	var claims []*models.Claim
	err := db.Find(&claims).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return claims, err
}

func (s *Store) GetClaim(userId string, product string, status common.ClaimStatus, beforeId, afterId int64, limit int) ([]*models.Claim, error) {
	db := s.db

	if len(userId) != 0 {
		db = db.Where("user_id=?", userId)
	}
	if len(product) != 0 {
		db = db.Where("product=?", product)
	}
	if len(status) != 0 {
		db = db.Where("status=?", status)
	}
	if beforeId > 0 {
		db = db.Where("id>=?", beforeId)
	}

	if afterId > 0 {
		db = db.Where("id<?", afterId)
	}

	if limit <= 0 {
		limit = 100
	}

	db = db.Order("id DESC").Limit(limit)

	var claims []*models.Claim
	err := db.Find(&claims).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return claims, err
}

func (s *Store) UpdateClaim(claim *models.Claim) error {
	claim.UpdatedAt = time.Now()
	return s.db.Save(claim).Error
}
func (s *Store) AddClaim(claim *models.Claim) error {
	claim.CreatedAt = time.Now()
	return s.db.Create(claim).Error
}

func (s *Store) AddClaimDesc(claimDesc *models.ClaimDesc) error {
	claimDesc.CreatedAt = time.Now()
	return s.db.Create(claimDesc).Error
}
func (s *Store) GetClaimDescByClaimId(claimId int64) ([]*models.ClaimDesc, error) {
	db := s.db.Where("claim_id=?", claimId)

	var claimDesc []*models.ClaimDesc
	err := db.Find(&claimDesc).Error
	//if err == gorm.ErrRecordNotFound {
	//	return nil, nil
	//}
	return claimDesc, err
}

func (s *Store) AddClaimCred(claimCred *models.ClaimCred) error {
	claimCred.CreatedAt = time.Now()
	return s.db.Create(claimCred).Error
}
func (s *Store) GetClaimCredByClaimId(claimId int64) ([]*models.ClaimCred, error) {
	db := s.db.Where("claim_id=?", claimId)

	var claimCred []*models.ClaimCred
	err := db.Find(&claimCred).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return claimCred, err
}
