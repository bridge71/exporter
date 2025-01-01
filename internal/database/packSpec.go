package database

import "exporter/internal/models"

func (s *service) DeletePackSpec(PackSpec *models.PackSpec) error {
	return s.gormDB.Delete(PackSpec).Error
}

func (s *service) SavePackSpec(PackSpec *models.PackSpec) error {
	return s.gormDB.Save(PackSpec).Error
}

func (s *service) FindPackSpec(PackSpecs *[]models.PackSpec) {
	s.gormDB.Find(PackSpecs)
}
