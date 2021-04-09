package mysql

import (
	"github.com/jinzhu/gorm"
	"nsure/vote/common"
	"nsure/vote/models"
	"time"
)

func (s *Store) AddVote(vote *models.Vote) error {
	vote.CreatedAt = time.Now()
	return s.db.Create(vote).Error
}
func (s *Store) UpdateVote(vote *models.Vote) error {
	vote.UpdatedAt = time.Now()
	return s.db.Save(vote).Error
}

func (s *Store) GetVoteByVoteNum(statuses []common.ClaimStatus, voteNum uint8) ([]*common.ClaimId, error) {
	var claimIds []*common.ClaimId
	db := s.db.Table("v_vote").Select("claim_id").Where("settled=0")
	if len(statuses) != 0 {
		db = db.Where("status IN (?)", statuses)
	}
	err := db.Group("claim_id").Having("count(*) >= ?", voteNum).Scan(&claimIds).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return claimIds, err
}

//func (s *Store) GetVoteByVoteNum(voteNum uint8) ([]*common.ClaimId, error) {
//	var claimIds []*common.ClaimId
//	db := s.db.Table("v_vote").Select("claim_id").Where("settled=0 AND status != ?", common.ClaimStatusArbiter)
//	err := db.Group("claim_id").Having("count(*) >= ?", voteNum).Scan(&claimIds).Error
//	if err == gorm.ErrRecordNotFound {
//		return nil, nil
//	}
//	return claimIds, err
//}

func (s *Store) GetVoteForUpdate(claimId int64, arbiterId string) (*models.Vote, error) {
	var vote models.Vote
	err := s.db.Raw("SELECT * FROM v_vote WHERE claim_id=? AND arbiter_Id=? AND settled=0 FOR UPDATE", claimId, arbiterId).Scan(&vote).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &vote, err
}
func (s *Store) GetVoteByClaimIdStatus(claimId int64, statuses []common.ClaimStatus) ([]*models.Vote, error) {
	db := s.db.Where("claim_id=? ", claimId)
	if len(statuses) != 0 {
		db = db.Where("status IN (?)", statuses)
	}
	var votes []*models.Vote
	err := db.Find(&votes).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return votes, err

}

func (s *Store) GetVoteByClaimId(claimId int64) ([]*models.Vote, error) {
	db := s.db.Where("claim_id=? ", claimId)

	var votes []*models.Vote
	err := db.Find(&votes).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return votes, err

}

func (s *Store) GetVoteByEnd(voteTime uint) ([]*models.Vote, error) {
	db := s.db.Where("settled=0 AND ADDDATE(arbiter_at,interval ? minute) < now() ", voteTime)

	var votes []*models.Vote
	err := db.Find(&votes).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return votes, err
}

func (s *Store) GetVote(arbiterId string, product string, status common.ClaimStatus, beforeId, afterId int64, limit int) ([]*models.Vote, error) {
	db := s.db

	if len(arbiterId) != 0 {
		db = db.Where("arbiter_id=?", arbiterId)
	}
	if len(product) != 0 {
		db = db.Where("product=?", product)
	}
	if len(status) != 0 {
		db = db.Where("status =", status)
	}
	if beforeId > 0 {
		db = db.Where("id>?", beforeId)
	}

	if afterId > 0 {
		db = db.Where("id<?", afterId)
	}

	if limit <= 0 {
		limit = 100
	}

	db = db.Order("id DESC").Limit(limit)

	var votes []*models.Vote
	err := db.Find(&votes).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return votes, err
}

func (s *Store) AddVoteFill(voteFill *models.VoteFill) error {
	voteFill.CreatedAt = time.Now()
	return s.db.Create(voteFill).Error
}

func (s *Store) UpdateVoteFill(voteFill *models.VoteFill) error {
	voteFill.UpdatedAt = time.Now()
	return s.db.Save(voteFill).Error
}

func (s *Store) GetVoteFillsByClaimId(claimId int64) ([]*models.VoteFill, error) {
	db := s.db.Where("claim_id=?", claimId)

	var voteFills []*models.VoteFill
	err := db.Find(&voteFills).Error
	return voteFills, err
}

func (s *Store) GetUnsettledVoteFills(count int) ([]*models.VoteFill, error) {
	db := s.db.Where("settled =?", 0).Order("id ASC").Limit(count)

	var voteFills []*models.VoteFill
	err := db.Find(&voteFills).Error
	return voteFills, err
}
