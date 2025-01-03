package database

import "exporter/internal/models"

func (s *service) DeleteAcct(acct *models.Acct) error {
	return s.gormDB.Delete(acct).Error
}

func (s *service) SaveAcct(acct *models.Acct) error {
	return s.gormDB.Omit("CreatedAt").Save(acct).Error
}

func (s *service) FindAcct(accts *[]models.Acct) {
	s.gormDB.Find(accts)
}
