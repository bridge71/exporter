package database

import "exporter/internal/models"

func (s *service) FindFile(file *models.File) {
	s.gormDB.Where("ID = ?", file.ID).Find(file)
}

func (s *service) CreateFile(file *models.File) error {
	return s.gormDB.Create(file).Error
}
