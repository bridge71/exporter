package database

import (
	"exporter/internal/models"
)

func (s *service) FindOuts(Out *[]models.Out) {
	s.gormDB.
		Preload("Merchant").    // 加载 Merchant
		Preload("Acct").        // 加载 Acct
		Preload("BankAccount"). // 加载 BankAccount
		Preload("AcctBank").    // 加载 AcctBank
		Preload("Purrecs").     // 加载 Purrecs
		Preload("Buys").        // 加载 Buys
		Preload("ShouldOuts").  // 加载 ShouldOuts
		Order("id desc").       // 按 ID 降序排序
		Find(Out)               // 查询结果存储到 outs 中
}

func (s *service) FindOutPurrec(Out *models.Out) {
	s.gormDB.Preload("Purrecs").Find(Out)
}

func (s *service) FindOutShouldOut(Out *models.Out) {
	s.gormDB.Preload("ShouldOuts").Find(Out)
}

func (s *service) FindOutBuy(Out *models.Out) {
	s.gormDB.Preload("Buys").Find(Out)
}

func (s *service) SaveOut(Out *models.Out) error {
	return s.gormDB.Omit("Purrecs", "Buys", "ShouldOuts").Save(Out).Error
}

func (s *service) DeleteOutPurrec(Out *models.Out, purrec *models.Purrec) error {
	return s.gormDB.Model(Out).Association("Purrecs").Delete(purrec)
}

func (s *service) DeleteOutBuy(Out *models.Out, buy *models.Buy) error {
	return s.gormDB.Model(Out).Association("Buys").Delete(buy)
}

func (s *service) DeleteOutShouldOut(Out *models.Out, shouldOut *models.ShouldOut) error {
	return s.gormDB.Model(Out).Association("ShouldOuts").Delete(shouldOut)
}
