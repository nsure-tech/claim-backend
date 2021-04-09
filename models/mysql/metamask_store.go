package mysql

import (
	"nsure/vote/models"
	"time"
)

func (s *Store) AddMetamask(metamask *models.Metamask) error {
	metamask.CreatedAt = time.Now()
	return s.db.Create(metamask).Error
}

func (s *Store) AddConfig(config *models.Config) error {
	config.CreatedAt = time.Now()
	return s.db.Create(config).Error
}

func (s *Store) GetConfig(keyWord string) (*models.Config, error) {
	var config models.Config
	err := s.db.Raw("SELECT * FROM v_config WHERE key_word=?", keyWord).Scan(&config).Error
	return &config, err
}

func (s *Store) GetConfigForUpdate(keyWord string) (*models.Config, error) {
	var config models.Config
	err := s.db.Raw("SELECT * FROM v_config WHERE key_word=? FOR UPDATE", keyWord).Scan(&config).Error
	return &config, err
}

func (s *Store) UpdateConfig(config *models.Config) error {
	config.UpdatedAt = time.Now()
	return s.db.Save(config).Error
}
