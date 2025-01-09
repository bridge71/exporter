package database

import "exporter/internal/models"

// DeleteBankAccount 删除一个 BankAccount 记录
func (s *service) DeleteBankAccount(bankAccount *models.BankAccount) error {
	return s.gormDB.Delete(bankAccount).Error
}

// SaveBankAccount 保存或更新一个 BankAccount 记录
func (s *service) SaveBankAccount(bankAccount *models.BankAccount) error {
	return s.gormDB.Save(bankAccount).Error
}

// FindBankAccount 查询所有 BankAccount 记录
func (s *service) FindBankAccount(bankAccounts *[]models.BankAccount) {
	s.gormDB.Preload("Merchant").Find(bankAccounts)
}

// FindBankAccountByID 根据 MercID 查询 BankAccount 记录
func (s *service) FindBankAccountByID(bankAccounts *[]models.BankAccount, mercID uint) {
	s.gormDB.Where("mercID = ?", mercID).Find(bankAccounts)
}
