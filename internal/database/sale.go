package database

import (
	"exporter/internal/models"

	"gorm.io/gorm"
)

func (s *service) FindSales(sale *[]models.Sale) {
	s.gormDB.
		Preload("DocReq").        // 加载 DocReq
		Preload("PrdtInfos").     // 加载 PrdtInfos
		Preload("Sends").         // 加载 Sends
		Preload("ShouldIns").     // 加载 ShouldIns
		Preload("Ins").           // 加载 Ins
		Preload("Acct").          // 加载 Acct
		Preload("Merchant").      // 加载 Merchant
		Preload("PayMentMethod"). // 加载 PayMentMethod
		Preload("PackSpec").      // 加载 PackSpec
		Preload("AcctBank").      // 加载 AcctBank
		Preload("BankAccount").   // 加载 BankAccount
		Order("id desc").         // 按 ID 降序排序
		Find(sale)                // 查询结果存储到 sales 中
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

func (s *service) TransSale(sale *models.Sale, DocReq *[]models.DocReq) error {
	return s.gormDB.Transaction(func(tx *gorm.DB) error {
		if len(*DocReq) != 0 {
			if err := s.gormDB.Model(sale).Association("DocReq").Delete(DocReq); err != nil {
				return err
			}
		}
		if err := s.gormDB.Save(sale).Error; err != nil {
			return err
		}
		return nil
	})
}
