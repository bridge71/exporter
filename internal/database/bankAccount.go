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
	s.gormDB.Find(bankAccounts)
}

// FindBankAccountById 根据 MercId 查询 BankAccount 记录
func (s *service) FindBankAccountById(bankAccounts *[]models.BankAccount, mercId uint) {
	s.gormDB.Where("mercId = ?", mercId).Find(bankAccounts)
}
