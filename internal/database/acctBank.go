package database

import "exporter/internal/models"

func (s *service) DeleteAcctBank(acctBank *models.AcctBank) error {
	return s.gormDB.Delete(acctBank).Error
}

func (s *service) SaveAcctBank(acctBank *models.AcctBank) error {
	return s.gormDB.Omit("Sales").Save(acctBank).Error
}

func (s *service) FindAcctBank(acctBanks *[]models.AcctBank) {
	s.gormDB.Preload("Acct").Find(acctBanks)
}

func (s *service) FindAcctBankByID(acctBanks *[]models.AcctBank, acctID uint) {
	s.gormDB.Where("acctID = ?", acctID).Find(acctBanks)
}
