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

func (s *service) FindBrandByID(brand *models.Brand, brandID uint) error {
	return s.gormDB.First(brand, brandID).Error
}
