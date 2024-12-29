package database

import "exporter/internal/models"

func (s *service) SaveAcct(acct *models.Acct) error {
	return s.gormDB.Save(acct).Error
}

func (s *service) SaveAcctBank(acctBank *models.AcctBank) error {
	return s.gormDB.Save(acctBank).Error
}

func (s *service) FindFile(file *models.File) {
	s.gormDB.Where("fileId = ?", file.FileId).Find(file)
}

func (s *service) FindAcct(accts *[]models.Acct) {
	s.gormDB.Find(accts)
}

func (s *service) FindAcctBank(acctBanks *[]models.AcctBank) {
	s.gormDB.Find(acctBanks)
}

func (s *service) FindAcctBankById(acctBanks *[]models.AcctBank, acctId uint) {
	s.gormDB.Where("acctId = ?", acctId).Find(acctBanks)
}

func (s *service) FindAcctBankByAccName(acctBanks *[]models.AcctBank, accName string) {
	s.gormDB.Where("accName = ?", accName).Find(acctBanks)
}

func (s *service) FindAcctBankByAccNum(acctBanks *[]models.AcctBank, accNum string) {
	s.gormDB.Where("accNum = ?", accNum).Find(acctBanks)
}

func (s *service) FindAcctBankByCurrency(acctBanks *[]models.AcctBank, currency string) {
	s.gormDB.Where("currency = ?", currency).Find(acctBanks)
}

func (s *service) FindAcctBankByBankName(acctBanks *[]models.AcctBank, bankName string) {
	s.gormDB.Where("bankName = ?", bankName).Find(acctBanks)
}

func (s *service) FzzFindAcctBankByBankName(acctBanks *[]models.AcctBank, bankName string) {
	s.gormDB.Where("bankName like ?", "%"+bankName+"%").Find(acctBanks)
}

func (s *service) FindAcctBankBySwiftCode(acctBanks *[]models.AcctBank, swiftCode string) {
	s.gormDB.Where("swiftCode = ?", swiftCode).Find(acctBanks)
}

func (s *service) FindAcctBankByBankAddr(acctBanks *[]models.AcctBank, bankAddr string) {
	s.gormDB.Where("bankAddr = ?", bankAddr).Find(acctBanks)
}

func (s *service) FzzFindAcctBankByBankAddr(acctBanks *[]models.AcctBank, bankAddr string) {
	s.gormDB.Where("bankAddr like ?", "%"+bankAddr+"%").Find(acctBanks)
}

func (s *service) FzzFindAcctBankByNotes(acctBanks *[]models.AcctBank, notes string) {
	s.gormDB.Where("notes like ?", "%"+notes+"%").Find(acctBanks)
}

func (s *service) FindAcctBankByNotes(acctBanks *[]models.AcctBank, notes string) {
	s.gormDB.Where("notes = ?", notes).Find(acctBanks)
}

func (s *service) FindAcctByAcctCode(accts *[]models.Acct, acctCode string) {
	s.gormDB.Where("acctCode = ?", acctCode).Find(accts)
}

func (s *service) FindAcctByAcctAbbr(accts *[]models.Acct, acctAbbr string) {
	s.gormDB.Where("acctAbbr = ?", acctAbbr).Find(accts)
}

func (s *service) FindAcctByEtyAbbr(accts *[]models.Acct, EtyAbbr string) {
	s.gormDB.Where("EtyAbbr = ?", EtyAbbr).Find(accts)
}

func (s *service) FindAcctByAcctName(accts *[]models.Acct, acctName string) {
	s.gormDB.Where("acctName = ?", acctName).Find(accts)
}

func (s *service) FindAcctByAcctAddr(accts *[]models.Acct, acctAddr string) {
	s.gormDB.Where("acctAddr = ?", acctAddr).Find(accts)
}

func (s *service) FzzFindAcctByAcctAddr(accts *[]models.Acct, acctAddr string) {
	s.gormDB.Where("acctAbbr like ?", "%"+acctAddr+"%").Find(accts)
}

func (s *service) FindAcctByNation(accts *[]models.Acct, nation string) {
	s.gormDB.Where("nation = ?", nation).Find(accts)
}

func (s *service) FindAcctByTaxType(accts *[]models.Acct, taxType string) {
	s.gormDB.Where("taxType = ?", taxType).Find(accts)
}

func (s *service) FindAcctByTaxCode(accts *[]models.Acct, taxCode string) {
	s.gormDB.Where("taxCode = ?", taxCode).Find(accts)
}

// FindAcctByPhoneNum 查找具有特定电话号码的账户
func (s *service) FindAcctByPhoneNum(accts *[]models.Acct, phoneNum string) {
	s.gormDB.Where("phoneNum = ?", phoneNum).Find(accts)
}

// FzzFindAcctByPhoneNum 查找电话号码包含特定字符串的账户
func (s *service) FzzFindAcctByPhoneNum(accts *[]models.Acct, phoneNum string) {
	s.gormDB.Where("phoneNum like ?", "%"+phoneNum+"%").Find(accts)
}

// FindAcctByEmail 查找具有特定电子邮件的账户
func (s *service) FindAcctByEmail(accts *[]models.Acct, email string) {
	s.gormDB.Where("email = ?", email).Find(accts)
}

// FzzFindAcctByEmail 查找电子邮件包含特定字符串的账户
func (s *service) FzzFindAcctByEmail(accts *[]models.Acct, email string) {
	s.gormDB.Where("email like ?", "%"+email+"%").Find(accts)
}

// FindAcctByWebsite 查找具有特定网站的账户
func (s *service) FindAcctByWebsite(accts *[]models.Acct, website string) {
	s.gormDB.Where("website = ?", website).Find(accts)
}

// FzzFindAcctByWebsite 查找网站包含特定字符串的账户
func (s *service) FzzFindAcctByWebsite(accts *[]models.Acct, website string) {
	s.gormDB.Where("website like ?", "%"+website+"%").Find(accts)
}

// FindAcctByRegDate 查找注册日期为特定日期的账户
func (s *service) FindAcctByRegDate(accts *[]models.Acct, regDate string) {
	s.gormDB.Where("regDate = ?", regDate).Find(accts)
}

// FzzFindAcctByRegDate 查找注册日期包含特定字符串的账户
func (s *service) FzzFindAcctByRegDate(accts *[]models.Acct, regDate string) {
	s.gormDB.Where("regDate like ?", "%"+regDate+"%").Find(accts)
}

// FindAcctByNotes 查找具有特定备注的账户
func (s *service) FindAcctByNotes(accts *[]models.Acct, notes string) {
	s.gormDB.Where("notes = ?", notes).Find(accts)
}

// FzzFindAcctByNotes 查找备注包含特定字符串的账户
func (s *service) FzzFindAcctByNotes(accts *[]models.Acct, notes string) {
	s.gormDB.Where("notes like ?", "%"+notes+"%").Find(accts)
}
