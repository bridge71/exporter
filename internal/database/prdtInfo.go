package database

import "exporter/internal/models"

func (s *service) FindPrdtInfo(PrdtInfo *[]models.PrdtInfo) {
	s.gormDB.Preload("Spot").Preload("PackSpec").Preload("Brand").Preload("Cat").Find(PrdtInfo)
}
