package database

import (
	"exporter/internal/models"
)

func (s *service) FindShouldOuts(ShouldOut *[]models.ShouldOut) {
	s.gormDB.Order("id desc").Find(ShouldOut)
}

func (s *service) FindShouldOutPurrec(ShouldOut *models.ShouldOut) {
	s.gormDB.Preload("Purrecs").Find(ShouldOut)
}

func (s *service) FindShouldOutOut(ShouldOut *models.ShouldOut) {
	s.gormDB.Preload("Outs").Find(ShouldOut)
}

func (s *service) FindShouldOutBuy(ShouldOut *models.ShouldOut) {
	s.gormDB.Preload("Buys").Find(ShouldOut)
}

func (s *service) SaveShouldOut(ShouldOut *models.ShouldOut) error {
	return s.gormDB.Omit("Purrecs", "Buys", "Outs").Save(ShouldOut).Error
}

func (s *service) DeleteShouldOutPurrec(ShouldOut *models.ShouldOut, purrec *models.Purrec) error {
	return s.gormDB.Model(ShouldOut).Association("Purrecs").Delete(purrec)
}

func (s *service) DeleteShouldOutBuy(ShouldOut *models.ShouldOut, buy *models.Buy) error {
	return s.gormDB.Model(ShouldOut).Association("Buys").Delete(buy)
}

func (s *service) DeleteShouldOutOut(ShouldOut *models.ShouldOut, out *models.Out) error {
	return s.gormDB.Model(ShouldOut).Association("Outs").Delete(out)
}
