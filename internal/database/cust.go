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
	s.gormDB.Find(custs)
}

// FindCustById 根据 MercId 查询 Cust 记录
func (s *service) FindCustById(custs *[]models.Cust, mercId uint) {
	s.gormDB.Where("mercId = ?", mercId).Find(custs)
}
