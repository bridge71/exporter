package database

import "exporter/internal/models"

func (s *service) SavePurrec(Purrec *models.Purrec) error {
	return s.gormDB.Omit("Buys", "LoadingInfos", "PrdtInfos").Save(Purrec).Error
}

func (s *service) FindPurrecs(Purrec *[]models.Purrec) {
	s.gormDB.Order("id desc").Find(Purrec)
}

func (s *service) FindPurrecPrdtInfo(Purrec *models.Purrec) {
	s.gormDB.Preload("PrdtInfos").Find(Purrec)
}

func (s *service) FindPurrecLoadingInfo(Purrec *models.Purrec) {
	s.gormDB.Preload("LoadingInfos").Find(Purrec)
}

func (s *service) FindPurrecBuy(Purrec *models.Purrec) {
	s.gormDB.Preload("Buys").Find(Purrec)
}

func (s *service) DeletePurrecPrdtInfo(Purrec *models.Purrec, prdtInfo *models.PrdtInfo) error {
	return s.gormDB.Model(Purrec).Association("PrdtInfos").Delete(prdtInfo)
}

func (s *service) DeletePurrecBuy(Purrec *models.Purrec, buy *models.Buy) error {
	return s.gormDB.Model(Purrec).Association("Buys").Delete(buy)
}

func (s *service) FindPurrecOuts(Purrec *models.Purrec) {
	s.gormDB.Preload("Outs").Find(Purrec)
}

func (s *service) FindPurrecShouldOuts(Purrec *models.Purrec) {
	s.gormDB.Preload("ShouldOuts").Find(Purrec)
}

func (s *service) DeletePurrecOuts(Purrec *models.Purrec, out *models.Out) error {
	return s.gormDB.Model(Purrec).Association("Outs").Delete(out)
}

func (s *service) DeletePurrecShouldOuts(Purrec *models.Purrec, shouldOut *models.ShouldOut) error {
	return s.gormDB.Model(Purrec).Association("ShouldOuts").Delete(shouldOut)
}

func (s *service) DeletePurrecLoadingInfo(Purrec *models.Purrec, LoadingInfo *models.LoadingInfo) error {
	return s.gormDB.Model(Purrec).Association("LoadingInfos").Delete(LoadingInfo)
}
