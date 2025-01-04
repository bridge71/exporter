package database

import "exporter/internal/models"

// DeleteMerchant 删除一个 Merchant 记录
func (s *service) DeleteMerchant(merchant *models.Merchant) error {
	return s.gormDB.Delete(merchant).Error
}

// SaveMerchant 保存或更新一个 Merchant 记录
func (s *service) SaveMerchant(merchant *models.Merchant) error {
	return s.gormDB.Omit("Sales", "Custs", "BankAccounts").Save(merchant).Error
}

// FindMerchant 查询所有 Merchant 记录
func (s *service) FindMerchant(merchants *[]models.Merchant) {
	s.gormDB.Find(merchants)
}
