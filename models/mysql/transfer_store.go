package mysql

import (
	"nsure/vote/models"
	"time"
)

func (s *Store) AddTransfer(transfer *models.Transfer) error {
	transfer.CreatedAt = time.Now()
	return s.db.Create(transfer).Error
}

func (s *Store) UpdateTransfer(transfer *models.Transfer) error {
	transfer.UpdatedAt = time.Now()
	return s.db.Save(transfer).Error
}
