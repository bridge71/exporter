package database

import "exporter/internal/models"

func (s *service) DeleteAcct(acct *models.Acct) error {
	return s.gormDB.Delete(acct).Error
}

func (s *service) SaveAcct(acct *models.Acct) error {
	return s.gormDB.Omit("AcctBanks").Omit("Sales").Save(acct).Error
}

func (s *service) FindAcct(accts *[]models.Acct) {
	s.gormDB.Find(accts)
}

func (s *service) FirstAcct(id uint, Acct *[]models.Acct) {
	s.gormDB.First(Acct, id)
}
