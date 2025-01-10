package database

import (
	"exporter/internal/models"
)

func (s *service) FindShouldOuts(ShouldOut *[]models.ShouldOut) {
	s.gormDB.
		Preload("Merchant").    // 加载 Merchant
		Preload("Acct").        // 加载 Acct
		Preload("BankAccount"). // 加载 BankAccount
		Preload("AcctBank").    // 加载 AcctBank
		Preload("Purrecs").     // 加载 Purrecs
		Preload("Buys").        // 加载 Buys
		Preload("Outs").        // 加载 Outs
		Order("id desc").       // 按 ID 降序排序
		Find(ShouldOut)         // 查询结果存储到 shouldOuts 中
}

func (s *service) FindShouldOutPurrec(ShouldOut *models.ShouldOut) {
	s.gormDB.Preload("Purrecs").Find(ShouldOut)
}

func (s *service) DeleteShouldOutCostInfo(ShouldOut *models.ShouldOut, costInfo *models.CostInfo) error {
	return s.gormDB.Model(ShouldOut).Association("CostInfos").Delete(costInfo)
}

func (s *service) FindShouldOutCostInfo(ShouldOut *models.ShouldOut) {
	s.gormDB.Preload("CostInfos").Find(ShouldOut)
}

func (s *service) FindShouldOutOut(ShouldOut *models.ShouldOut) {
	s.gormDB.Preload("Outs").Find(ShouldOut)
}

func (s *service) FindShouldOutBuy(ShouldOut *models.ShouldOut) {
	s.gormDB.Preload("Buys").Find(ShouldOut)
}

func (s *service) SaveShouldOut(ShouldOut *models.ShouldOut) error {
	return s.gormDB.Save(ShouldOut).Error
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
