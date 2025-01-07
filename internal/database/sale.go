package database

import "exporter/internal/models"

func (s *service) FindSales(sale *[]models.Sale) {
	s.gormDB.Preload("DocReq").Order("id desc").Find(sale)
}

func (s *service) FindSalePrdtInfo(sale *models.Sale) {
	s.gormDB.Preload("PrdtInfos").Find(sale)
}

func (s *service) FindSaleSend(sale *models.Sale) {
	s.gormDB.Preload("Sends").Find(sale)
}

func (s *service) SaveSale(sale *models.Sale) error {
	return s.gormDB.Save(sale).Error
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
