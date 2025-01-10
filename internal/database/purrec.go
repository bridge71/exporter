package database

import "exporter/internal/models"

func (s *service) SavePurrec(Purrec *models.Purrec) error {
	return s.gormDB.Save(Purrec).Error
}

func (s *service) FindPurrecs(Purrec *[]models.Purrec) {
	s.gormDB.
		Preload("Buys").          // 加载 Buys
		Preload("PrdtInfos").     // 加载 PrdtInfos
		Preload("LoadingInfos").  // 加载 LoadingInfos
		Preload("ShouldOuts").    // 加载 ShouldOuts
		Preload("Outs").          // 加载 Outs
		Preload("Merchant").      // 加载 Merchant
		Preload("PackSpec").      // 加载 PackSpec
		Preload("PayMentMethod"). // 加载 PayMentMethod
		Preload("AcctBank").      // 加载 AcctBank
		Order("id desc").         // 按 ID 降序排序
		Find(Purrec)              // 查询结果存储到 purrecs 中
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
