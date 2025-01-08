package database

import (
	"exporter/internal/models"
)

func (s *service) FindShouldIns(ShouldIn *[]models.ShouldIn) {
	s.gormDB.Order("id desc").Find(ShouldIn)
}

func (s *service) FindShouldInSend(ShouldIn *models.ShouldIn) {
	s.gormDB.Preload("Sends").Find(ShouldIn)
}

func (s *service) FindShouldInIn(ShouldIn *models.ShouldIn) {
	s.gormDB.Preload("Ins").Find(ShouldIn)
}

func (s *service) FindShouldInSale(ShouldIn *models.ShouldIn) {
	s.gormDB.Preload("Sales").Find(ShouldIn)
}

func (s *service) SaveShouldIn(ShouldIn *models.ShouldIn) error {
	return s.gormDB.Omit("Sends", "Sales", "Ins").Save(ShouldIn).Error
}

func (s *service) DeleteShouldInSend(ShouldIn *models.ShouldIn, send *models.Send) error {
	return s.gormDB.Model(ShouldIn).Association("Sends").Delete(send)
}

func (s *service) DeleteShouldInSale(ShouldIn *models.ShouldIn, sale *models.Sale) error {
	return s.gormDB.Model(ShouldIn).Association("Sales").Delete(sale)
}

func (s *service) DeleteShouldInIn(ShouldIn *models.ShouldIn, in *models.In) error {
	return s.gormDB.Model(ShouldIn).Association("Ins").Delete(in)
}
