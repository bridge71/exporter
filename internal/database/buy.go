package database

import (
	"exporter/internal/models"

	"gorm.io/gorm"
)

func (s *service) FindBuys(Buy *[]models.Buy) {
	s.gormDB.
		Preload("DocReq").        // 加载 DocReq
		Preload("PrdtInfos").     // 加载 PrdtInfos
		Preload("Purrecs").       // 加载 Purrecs
		Preload("ShouldOuts").    // 加载 ShouldOuts
		Preload("Outs").          // 加载 Outs
		Preload("Acct").          // 加载 Acct
		Preload("Merchant").      // 加载 Merchant
		Preload("PayMentMethod"). // 加载 PayMentMethod
		Preload("PackSpec").      // 加载 PackSpec
		Preload("AcctBank").      // 加载 AcctBank
		Preload("BankAccount").   // 加载 BankAccount
		Order("id desc").         // 按 ID 降序排序
		Find(Buy)                 // 查询结果存储到 buys 中
}

func (s *service) FindBuyPrdtInfo(Buy *models.Buy) {
	s.gormDB.Preload("PrdtInfos").Find(Buy)
}

func (s *service) FindBuyPurrec(Buy *models.Buy) {
	s.gormDB.Preload("Purrecs").Find(Buy)
}

func (s *service) DeleteBuyOuts(Buy *models.Buy, Outs *models.Out) error {
	return s.gormDB.Model(Buy).Association("Outs").Delete(Outs)
}

func (s *service) DeleteBuyShouldOuts(Buy *models.Buy, ShouldOuts *models.ShouldOut) error {
	return s.gormDB.Model(Buy).Association("ShouldOuts").Delete(ShouldOuts)
}

func (s *service) FindBuyOuts(Buy *models.Buy) {
	s.gormDB.Preload("Outs").Find(Buy)
}

func (s *service) FindBuyShouldOuts(Buy *models.Buy) {
	s.gormDB.Preload("ShouldOuts").Find(Buy)
}

func (s *service) SaveBuy(Buy *models.Buy) error {
	return s.gormDB.Omit("Purrecs", "PrdtInfos", "DocReq", "Outs", "ShouldOuts").Save(Buy).Error
}

func (s *service) DeleteBuyDocReq(Buy *models.Buy, DocReq *[]models.DocReq) error {
	return s.gormDB.Model(Buy).Association("DocReq").Delete(DocReq)
}

func (s *service) DeleteBuyPrdtInfo(Buy *models.Buy, prdtInfo *models.PrdtInfo) error {
	return s.gormDB.Model(Buy).Association("PrdtInfos").Delete(prdtInfo)
}

func (s *service) DeleteBuyPurrec(Buy *models.Buy, Purrec *models.Purrec) error {
	return s.gormDB.Model(Buy).Association("Purrecs").Delete(Purrec)
}

func (s *service) TransBuy(Buy *models.Buy, DocReq *[]models.DocReq) error {
	return s.gormDB.Transaction(func(tx *gorm.DB) error {
		if len(*DocReq) != 0 {
			if err := s.gormDB.Model(Buy).Association("DocReq").Delete(DocReq); err != nil {
				return err
			}
		}
		if err := s.gormDB.Omit("Purrecs", "PrdtInfos").Save(Buy).Error; err != nil {
			return err
		}
		return nil
	})
}
