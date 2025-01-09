package database

import "exporter/internal/models"

// DeleteCat 删除一个 Cat 记录
func (s *service) DeleteCat(cat *models.Cat) error {
	return s.gormDB.Delete(cat).Error
}

// SaveCat 保存或更新一个 Cat 记录
func (s *service) SaveCat(cat *models.Cat) error {
	return s.gormDB.Save(cat).Error
}

// FindCat 查询所有 Cat 记录
func (s *service) FindCat(cats *[]models.Cat) {
	s.gormDB.Find(cats)
}

// FindCatByID 根据 CatID 查询 Cat 记录
func (s *service) FindCatByID(cat *models.Cat, catID uint) error {
	return s.gormDB.First(cat, catID).Error
}
