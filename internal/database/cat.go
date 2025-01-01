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

// FindCatById 根据 CatId 查询 Cat 记录
func (s *service) FindCatById(cat *models.Cat, catId uint) error {
	return s.gormDB.First(cat, catId).Error
}
