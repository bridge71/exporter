package database

import "exporter/internal/models"

func (s *service) FindLoadingInfo(LoadingInfo *[]models.LoadingInfo) {
	s.gormDB.Preload("PackSpec").Preload("Brand").Preload("Cat").Find(LoadingInfo)
}
