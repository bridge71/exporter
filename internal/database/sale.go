package database

import (
	"exporter/internal/models"

	"gorm.io/gorm"
)

func (s *service) FindSales(sale *[]models.Sale) {
	s.gormDB.Preload("DocReq").Order("id desc").Find(sale)
}

func (s *service) FindSalePrdtInfo(sale *models.Sale) {
	s.gormDB.Preload("PrdtInfos").Find(sale)
}

func (s *service) FindSaleSend(sale *models.Sale) {
	s.gormDB.Preload("Sends").Find(sale)
}

func (s *service) DeleteSaleIns(sale *models.Sale, ins *models.In) error {
	return s.gormDB.Model(sale).Association("Ins").Delete(ins)
}

func (s *service) DeleteSaleShouldIns(sale *models.Sale, shouldins *models.ShouldIn) error {
	return s.gormDB.Model(sale).Association("ShouldIns").Delete(shouldins)
}

func (s *service) FindSaleIns(sale *models.Sale) {
	s.gormDB.Preload("Ins").Find(sale)
}

func (s *service) FindSaleShouldIns(sale *models.Sale) {
	s.gormDB.Preload("ShouldIns").Find(sale)
}

func (s *service) SaveSale(sale *models.Sale) error {
	return s.gormDB.Omit("Sends", "PrdtInfos", "DocReq", "Ins", "ShouldIns").Save(sale).Error
}

func (s *service) DeleteSaleDocReq(sale *models.Sale, DocReq *[]models.DocReq) error {
	return s.gormDB.Model(sale).Association("DocReq").Delete(DocReq)
}

func (s *service) DeleteSalePrdtInfo(sale *models.Sale, prdtInfo *models.PrdtInfo) error {
	return s.gormDB.Model(sale).Association("PrdtInfos").Delete(prdtInfo)
}

func (s *service) DeleteSaleSend(sale *models.Sale, send *models.Send) error {
	return s.gormDB.Model(sale).Association("Sends").Delete(send)
}

func (s *service) TransSale(sale *models.Sale, DocReq *[]models.DocReq) error {
	return s.gormDB.Transaction(func(tx *gorm.DB) error {
		if err := s.gormDB.Model(sale).Association("DocReq").Delete(DocReq); err != nil {
			return err
		}
		if err := s.gormDB.Omit("Sends", "PrdtInfos", "DocReq").Save(sale).Error; err != nil {
			return err
		}
		return nil
	})
}
