package database

import "exporter/internal/models"

// DeletePayMentMethod 删除 PayMentMethod 记录
func (s *service) DeletePayMentMethod(payMentMethod *models.PayMentMethod) error {
	return s.gormDB.Delete(payMentMethod).Error
}

// SavePayMentMethod 保存或更新 PayMentMethod 记录
func (s *service) SavePayMentMethod(payMentMethod *models.PayMentMethod) error {
	return s.gormDB.Save(payMentMethod).Error
}

// FindPayMentMethods 查询所有 PayMentMethod 记录
func (s *service) FindPayMentMethods(payMentMethods *[]models.PayMentMethod) {
	s.gormDB.Find(payMentMethods)
}

// FindPayMentMethodByID 根据 ID 查询 PayMentMethod 记录
func (s *service) FindPayMentMethodByID(payMentMethod *models.PayMentMethod, id uint) error {
	return s.gormDB.First(payMentMethod, id).Error
}
