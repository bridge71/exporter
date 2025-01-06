package database

import "exporter/internal/models"

func (s *service) FindSalePrdtInfo(sale *[]models.Sale) {
	s.gormDB.Preload("PrdtInfos").Preload("DocReq").Find(sale)
}

func (s *service) SaveSale(sale *models.Sale) error {
	return s.gormDB.Save(sale).Error
}
