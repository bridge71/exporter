package database

import "exporter/internal/models"

func (s *service) DeleteAcct(acct *models.Acct) error {
	return s.gormDB.Delete(acct).Error
}

func (s *service) SaveAcct(acct *models.Acct) error {
	return s.gormDB.Omit("AcctBanks", "Sales", "Sends").Save(acct).Error
}

func (s *service) FindAcct(accts *[]models.Acct) {
	s.gormDB.Preload("AcctBanks").Find(accts)
}

func (s *service) FirstAcct(id uint, Acct *[]models.Acct) {
	s.gormDB.First(Acct, id)
}
