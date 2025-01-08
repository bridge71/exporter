package database

import (
	"exporter/internal/models"
)

func (s *service) FindIns(In *[]models.In) {
	s.gormDB.Order("id desc").Find(In)
}

func (s *service) FindInSend(In *models.In) {
	s.gormDB.Preload("Sends").Find(In)
}

func (s *service) FindInShouldIn(In *models.In) {
	s.gormDB.Preload("ShouldIns").Find(In)
}

func (s *service) FindInSale(In *models.In) {
	s.gormDB.Preload("Sales").Find(In)
}

func (s *service) SaveIn(In *models.In) error {
	return s.gormDB.Omit("Sends", "Sales", "ShouldIns").Save(In).Error
}

func (s *service) DeleteInSend(In *models.In, send *models.Send) error {
	return s.gormDB.Model(In).Association("Sends").Delete(send)
}

func (s *service) DeleteInSale(In *models.In, sale *models.Sale) error {
	return s.gormDB.Model(In).Association("Sales").Delete(sale)
}

func (s *service) DeleteInShouldIn(In *models.In, shouldIns *models.ShouldIn) error {
	return s.gormDB.Model(In).Association("ShouldIns").Delete(shouldIns)
}
