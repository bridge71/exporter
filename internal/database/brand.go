package database

import "exporter/internal/models"

func (s *service) DeleteBrand(brand *models.Brand) error {
	return s.gormDB.Delete(brand).Error
}

func (s *service) SaveBrand(brand *models.Brand) error {
	return s.gormDB.Save(brand).Error
}

func (s *service) FindBrand(brands *[]models.Brand) {
	s.gormDB.Find(brands)
}

func (s *service) FindBrandById(brand *models.Brand, brandId uint) error {
	return s.gormDB.First(brand, brandId).Error
}
