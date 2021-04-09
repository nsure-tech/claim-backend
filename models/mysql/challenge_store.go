package mysql

import (
	"nsure/vote/common"
	"nsure/vote/models"
	"time"
)

func (s *Store) AddChallenge(challenge *models.Challenge) error {
	challenge.CreatedAt = time.Now()
	return s.db.Create(challenge).Error
}
func (s *Store) UpdateChallenge(challenge *models.Challenge) error {
	challenge.UpdatedAt = time.Now()
	return s.db.Save(challenge).Error
}

func (s *Store) GetChallengeForUpdate(claimId int64) (*models.Challenge, error) {
	var challenge models.Challenge
	err := s.db.Raw("SELECT * FROM v_challenge WHERE claim_id=? AND settled=0 FOR UPDATE", claimId).Scan(&challenge).Error
	return &challenge, err
}

func (s *Store) GetChallengeByStatus(status common.ChallengeStatus, count int) ([]*models.Challenge, error) {
	db := s.db.Where("settled=0 and status = ?", status).Order("id ASC").Limit(count)

	var challenges []*models.Challenge
	err := db.Find(&challenges).Error
	return challenges, err
}

func (s *Store) AddChallengeFill(challengeFill *models.ChallengeFill) error {
	challengeFill.CreatedAt = time.Now()
	return s.db.Create(challengeFill).Error
}
func (s *Store) UpdateChallengeFill(challengeFill *models.ChallengeFill) error {
	challengeFill.UpdatedAt = time.Now()
	return s.db.Save(challengeFill).Error
}

func (s *Store) GetChallengeFillForUpdate(claimId int64) (*models.ChallengeFill, error) {
	var challengeFill models.ChallengeFill
	err := s.db.Raw("SELECT * FROM v_challenge_fill WHERE claim_id=? AND settled=0 FOR UPDATE", claimId).Scan(&challengeFill).Error
	return &challengeFill, err
}

func (s *Store) GetChallengeFillByStatus(status common.ChallengeStatus, count int) ([]*models.ChallengeFill, error) {
	db := s.db.Where("settled=0 and status = ?", status).Order("id ASC").Limit(count)

	var challengeFills []*models.ChallengeFill
	err := db.Find(&challengeFills).Error
	return challengeFills, err
}
