package database

import "exporter/internal/models"

func (s *service) FindSalePrdtInfo(sale *[]models.Sale) {
	s.gormDB.Preload("PrdtInfos").Find(sale)
}

func (s *service) SaveSale(sale *models.Sale) error {
	return s.gormDB.Omit("PrdtInfos").Save(sale).Error
}
