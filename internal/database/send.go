package database

import "exporter/internal/models"

func (s *service) SaveSend(Send *models.Send) error {
	return s.gormDB.Omit("Sales", "LoadingInfos", "PrdtInfos").Save(Send).Error
}

func (s *service) FindSends(Send *[]models.Send) {
	s.gormDB.Order("id desc").Find(Send)
}

func (s *service) FindSendPrdtInfo(Send *models.Send) {
	s.gormDB.Preload("PrdtInfos").Find(Send)
}

func (s *service) FindSendLoadingInfo(Send *models.Send) {
	s.gormDB.Preload("LoadingInfos").Find(Send)
}

func (s *service) FindSendSale(Send *models.Send) {
	s.gormDB.Preload("Sales").Find(Send)
}

func (s *service) DeleteSendPrdtInfo(Send *models.Send, prdtInfo *models.PrdtInfo) error {
	return s.gormDB.Model(Send).Association("PrdtInfos").Delete(prdtInfo)
}

func (s *service) DeleteSendSale(Send *models.Send, sale *models.Sale) error {
	return s.gormDB.Model(Send).Association("Sales").Delete(sale)
}

func (s *service) FindSendIns(Send *models.Send) {
	s.gormDB.Preload("Ins").Find(Send)
}

func (s *service) FindSendShouldIns(Send *models.Send) {
	s.gormDB.Preload("ShouldIns").Find(Send)
}

func (s *service) DeleteSendIns(Send *models.Send, ins *models.In) error {
	return s.gormDB.Model(Send).Association("Ins").Delete(ins)
}

func (s *service) DeleteSendShouldIns(Send *models.Send, shouldIns *models.ShouldIn) error {
	return s.gormDB.Model(Send).Association("ShouldIns").Delete(shouldIns)
}

func (s *service) DeleteSendLoadingInfo(Send *models.Send, LoadingInfo *models.LoadingInfo) error {
	return s.gormDB.Model(Send).Association("LoadingInfos").Delete(LoadingInfo)
}
