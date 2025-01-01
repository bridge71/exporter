package database

import "exporter/internal/models"

func (s *service) DeleteEmpl(Empl *models.Empl) error {
	return s.gormDB.Delete(Empl).Error
}

func (s *service) SaveEmpl(Empl *models.Empl) error {
	return s.gormDB.Save(Empl).Error
}

func (s *service) FindEmpl(Empls *[]models.Empl) {
	s.gormDB.Find(Empls)
}
