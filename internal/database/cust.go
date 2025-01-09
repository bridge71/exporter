package database

import "exporter/internal/models"

// DeleteCust 删除一个 Cust 记录
func (s *service) DeleteCust(cust *models.Cust) error {
	return s.gormDB.Delete(cust).Error
}

// SaveCust 保存或更新一个 Cust 记录
func (s *service) SaveCust(cust *models.Cust) error {
	return s.gormDB.Save(cust).Error
}

// FindCust 查询所有 Cust 记录
func (s *service) FindCust(custs *[]models.Cust) {
	s.gormDB.Preload("Merchant").Find(custs)
}

// FindCustByID 根据 MercID 查询 Cust 记录
func (s *service) FindCustByID(custs *[]models.Cust, mercID uint) {
	s.gormDB.Where("mercID = ?", mercID).Find(custs)
}
