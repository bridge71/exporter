package database

import "exporter/internal/models"

func (s *service) FzzFindCustByEmail(custs *[]models.Cust, email string) {
	s.gormDB.Where("email LIKE ?", "%"+email+"%").Find(custs)
}

func (s *service) FindBankAccountByMercId(bankAccounts *[]models.BankAccount, MercId uint) {
	s.gormDB.Where("mercId = ?", MercId).Find(bankAccounts)
}

func (s *service) FzzFindCustByWechat(custs *[]models.Cust, wechat string) {
	s.gormDB.Where("wechat LIKE ?", "%"+wechat+"%").Find(custs)
}

func (s *service) FzzFindCustByQQ(custs *[]models.Cust, qq string) {
	s.gormDB.Where("qq LIKE ?", "%"+qq+"%").Find(custs)
}

func (s *service) FzzFindCustByPhoneNum(custs *[]models.Cust, phoneNum string) {
	s.gormDB.Where("phoneNum LIKE ?", "%"+phoneNum+"%").Find(custs)
}

func (s *service) FindMerchantByMercCode(merchants *[]models.Merchant, mercCode string) {
	s.gormDB.Where("mercCode = ?", mercCode).Find(merchants)
}

func (s *service) FindMerchantByMercAbbr(merchants *[]models.Merchant, mercAbbr string) {
	s.gormDB.Where("mercAbbr = ?", mercAbbr).Find(merchants)
}

func (s *service) FindMerchantByShortMerc(merchants *[]models.Merchant, shortMerc string) {
	s.gormDB.Where("shortMerc = ?", shortMerc).Find(merchants)
}

func (s *service) FindMerchantByMerc(merchants *[]models.Merchant, merc string) {
	s.gormDB.Where("merc = ?", merc).Find(merchants)
}

func (s *service) FindMerchantByEngName(merchants *[]models.Merchant, engName string) {
	s.gormDB.Where("engName = ?", engName).Find(merchants)
}

func (s *service) FindMerchantByAddress(merchants *[]models.Merchant, address string) {
	s.gormDB.Where("address = ?", address).Find(merchants)
}

func (s *service) FindMerchantByNation(merchants *[]models.Merchant, nation string) {
	s.gormDB.Where("nation = ?", nation).Find(merchants)
}

func (s *service) FindMerchantByPhoneNum(merchants *[]models.Merchant, phoneNum string) {
	s.gormDB.Where("phoneNum = ?", phoneNum).Find(merchants)
}

func (s *service) FindMerchantByEmail(merchants *[]models.Merchant, email string) {
	s.gormDB.Where("email = ?", email).Find(merchants)
}

func (s *service) FindMerchantByFax(merchants *[]models.Merchant, fax string) {
	s.gormDB.Where("fax = ?", fax).Find(merchants)
}

func (s *service) FindMerchantByWebsite(merchants *[]models.Merchant, website string) {
	s.gormDB.Where("website = ?", website).Find(merchants)
}

func (s *service) FindMerchantByTaxType(merchants *[]models.Merchant, taxType string) {
	s.gormDB.Where("taxType = ?", taxType).Find(merchants)
}

func (s *service) FindMerchantByTaxCode(merchants *[]models.Merchant, taxCode string) {
	s.gormDB.Where("taxCode = ?", taxCode).Find(merchants)
}

func (s *service) FindMerchantByMercType(merchants *[]models.Merchant, mercType string) {
	s.gormDB.Where("mercType = ?", mercType).Find(merchants)
}

func (s *service) FindMerchantByRegCap(merchants *[]models.Merchant, regCap string) {
	s.gormDB.Where("regCap = ?", regCap).Find(merchants)
}

func (s *service) FindMerchantByNotes(merchants *[]models.Merchant, notes string) {
	s.gormDB.Where("notes = ?", notes).Find(merchants)
}

func (s *service) FindCustByName(custs *[]models.Cust, name string) {
	s.gormDB.Where("name = ?", name).Find(custs)
}

func (s *service) FindCustByGender(custs *[]models.Cust, gender string) {
	s.gormDB.Where("gender = ?", gender).Find(custs)
}

func (s *service) FindCustByNation(custs *[]models.Cust, nation string) {
	s.gormDB.Where("nation = ?", nation).Find(custs)
}

func (s *service) FindCustByAddr(custs *[]models.Cust, addr string) {
	s.gormDB.Where("addr = ?", addr).Find(custs)
}

func (s *service) FindCustByEmail(custs *[]models.Cust, email string) {
	s.gormDB.Where("email = ?", email).Find(custs)
}

func (s *service) FindCustByPhoneNum(custs *[]models.Cust, phoneNum string) {
	s.gormDB.Where("phoneNum = ?", phoneNum).Find(custs)
}

func (s *service) FindCustByQQ(custs *[]models.Cust, qq string) {
	s.gormDB.Where("qq = ?", qq).Find(custs)
}

func (s *service) FindCustByWechat(custs *[]models.Cust, wechat string) {
	s.gormDB.Where("wechat = ?", wechat).Find(custs)
}

func (s *service) FindCustByMerc(custs *[]models.Cust, merc string) {
	s.gormDB.Where("merc = ?", merc).Find(custs)
}

func (s *service) FindCustByPost(custs *[]models.Cust, post string) {
	s.gormDB.Where("post = ?", post).Find(custs)
}

func (s *service) FindCustByNotes(custs *[]models.Cust, notes string) {
	s.gormDB.Where("notes = ?", notes).Find(custs)
}

func (s *service) FindBankAccountByBankAccName(bankAccounts *[]models.BankAccount, bankAccName string) {
	s.gormDB.Where("bankAccName = ?", bankAccName).Find(bankAccounts)
}

func (s *service) FindBankAccountByCompName(bankAccounts *[]models.BankAccount, compName string) {
	s.gormDB.Where("compName = ?", compName).Find(bankAccounts)
}

func (s *service) FindBankAccountByAcctNum(bankAccounts *[]models.BankAccount, acctNum string) {
	s.gormDB.Where("acctNum = ?", acctNum).Find(bankAccounts)
}

func (s *service) FindBankAccountByCurrency(bankAccounts *[]models.BankAccount, currency string) {
	s.gormDB.Where("currency = ?", currency).Find(bankAccounts)
}

func (s *service) FindBankAccountByBankName(bankAccounts *[]models.BankAccount, bankName string) {
	s.gormDB.Where("bankName = ?", bankName).Find(bankAccounts)
}

func (s *service) FindBankAccountByBankEng(bankAccounts *[]models.BankAccount, bankEng string) {
	s.gormDB.Where("bankEng = ?", bankEng).Find(bankAccounts)
}

func (s *service) FindBankAccountByBankNum(bankAccounts *[]models.BankAccount, bankNum string) {
	s.gormDB.Where("bankNum = ?", bankNum).Find(bankAccounts)
}

func (s *service) FindBankAccountBySwiftCode(bankAccounts *[]models.BankAccount, swiftCode string) {
	s.gormDB.Where("swiftCode = ?", swiftCode).Find(bankAccounts)
}

func (s *service) FindBankAccountByBankAddr(bankAccounts *[]models.BankAccount, bankAddr string) {
	s.gormDB.Where("bankAddr = ?", bankAddr).Find(bankAccounts)
}

func (s *service) FindBankAccountByNotes(bankAccounts *[]models.BankAccount, notes string) {
	s.gormDB.Where("notes = ?", notes).Find(bankAccounts)
}

// FindCustByMercId 根据 MercId 查找关联的 Cust 数据
func (s *service) FindCustByMercId(custs *[]models.Cust, MercId uint) {
	s.gormDB.Where("mercId = ?", MercId).Find(custs)
}

func (s *service) DeleteMerchant(merchant *models.Merchant) error {
	return s.gormDB.Delete(merchant).Error
}

func (s *service) SaveMerchant(merchant *models.Merchant) error {
	return s.gormDB.Save(merchant).Error
}

func (s *service) FindMerchant(merchant *[]models.Merchant) {
	s.gormDB.Find(merchant)
}

// DeleteBankAccount 删除银行账户
func (s *service) DeleteBankAccount(bankAccount *models.BankAccount) error {
	return s.gormDB.Delete(bankAccount).Error
}

// SaveBankAccount 保存或更新银行账户
func (s *service) SaveBankAccount(bankAccount *models.BankAccount) error {
	return s.gormDB.Save(bankAccount).Error
}

// FindBankAccount 查询所有银行账户
func (s *service) FindBankAccount(bankAccount *[]models.BankAccount) {
	s.gormDB.Find(bankAccount)
}

// DeleteCust 删除客户
func (s *service) DeleteCust(cust *models.Cust) error {
	return s.gormDB.Delete(cust).Error
}

// SaveCust 保存或更新客户
func (s *service) SaveCust(cust *models.Cust) error {
	return s.gormDB.Save(cust).Error
}

// FindCust 查询所有客户
func (s *service) FindCust(cust *[]models.Cust) {
	s.gormDB.Find(cust)
}

func (s *service) FzzFindMerchantByEngName(merchant *[]models.Merchant, EngName string) {
	s.gormDB.Where("EngName like ?", "%"+EngName+"%").Find(merchant)
}

// 根据 MercId 查找 Merchant
func (s *service) FindMerchantByMercId(merchant *[]models.Merchant, MercId uint) {
	s.gormDB.Where("mercId = ?", MercId).Find(merchant)
}

// 根据 MercAbbr 模糊查找 Merchant
func (s *service) FzzFindMerchantByMercAbbr(merchant *[]models.Merchant, MercAbbr string) {
	s.gormDB.Where("mercAbbr LIKE ?", "%"+MercAbbr+"%").Find(merchant)
}

// 根据 ShortMerc 模糊查找 Merchant
func (s *service) FzzFindMerchantByShortMerc(merchant *[]models.Merchant, ShortMerc string) {
	s.gormDB.Where("shortMerc LIKE ?", "%"+ShortMerc+"%").Find(merchant)
}

// 根据 Merc 模糊查找 Merchant
func (s *service) FzzFindMerchantByMerc(merchant *[]models.Merchant, Merc string) {
	s.gormDB.Where("merc LIKE ?", "%"+Merc+"%").Find(merchant)
}

// 根据 Address 模糊查找 Merchant
func (s *service) FzzFindMerchantByAddress(merchant *[]models.Merchant, Address string) {
	s.gormDB.Where("address LIKE ?", "%"+Address+"%").Find(merchant)
}

// 根据 Nation 模糊查找 Merchant
func (s *service) FzzFindMerchantByNation(merchant *[]models.Merchant, Nation string) {
	s.gormDB.Where("nation LIKE ?", "%"+Nation+"%").Find(merchant)
}

// 根据 Website 模糊查找 Merchant
func (s *service) FzzFindMerchantByWebsite(merchant *[]models.Merchant, Website string) {
	s.gormDB.Where("website LIKE ?", "%"+Website+"%").Find(merchant)
}

// 根据 TaxType 模糊查找 Merchant
func (s *service) FzzFindMerchantByTaxType(merchant *[]models.Merchant, TaxType string) {
	s.gormDB.Where("taxType LIKE ?", "%"+TaxType+"%").Find(merchant)
}

// 根据 MercType 模糊查找 Merchant
func (s *service) FzzFindMerchantByMercType(merchant *[]models.Merchant, MercType string) {
	s.gormDB.Where("mercType LIKE ?", "%"+MercType+"%").Find(merchant)
}

// 根据 Notes 模糊查找 Merchant
func (s *service) FzzFindMerchantByNotes(merchant *[]models.Merchant, Notes string) {
	s.gormDB.Where("notes LIKE ?", "%"+Notes+"%").Find(merchant)
}

// 根据 CustId 查找 Cust
func (s *service) FindCustByCustId(cust *[]models.Cust, CustId uint) {
	s.gormDB.Where("custId = ?", CustId).Find(cust)
}

// 根据 Name 模糊查找 Cust
func (s *service) FzzFindCustByName(cust *[]models.Cust, Name string) {
	s.gormDB.Where("name LIKE ?", "%"+Name+"%").Find(cust)
}

// 根据 Nation 模糊查找 Cust
func (s *service) FzzFindCustByNation(cust *[]models.Cust, Nation string) {
	s.gormDB.Where("nation LIKE ?", "%"+Nation+"%").Find(cust)
}

// 根据 Addr 模糊查找 Cust
func (s *service) FzzFindCustByAddr(cust *[]models.Cust, Addr string) {
	s.gormDB.Where("addr LIKE ?", "%"+Addr+"%").Find(cust)
}

// 根据 Merc 模糊查找 Cust
func (s *service) FzzFindCustByMerc(cust *[]models.Cust, Merc string) {
	s.gormDB.Where("merc LIKE ?", "%"+Merc+"%").Find(cust)
}

// 根据 Post 模糊查找 Cust
func (s *service) FzzFindCustByPost(cust *[]models.Cust, Post string) {
	s.gormDB.Where("post LIKE ?", "%"+Post+"%").Find(cust)
}

// 根据 Notes 模糊查找 Cust
func (s *service) FzzFindCustByNotes(cust *[]models.Cust, Notes string) {
	s.gormDB.Where("notes LIKE ?", "%"+Notes+"%").Find(cust)
}

// 根据 BankAccId 查找 BankAccount
func (s *service) FindBankAccountByBankAccId(bankAccount *[]models.BankAccount, BankAccId uint) {
	s.gormDB.Where("bankAccId = ?", BankAccId).Find(bankAccount)
}

// 根据 BankAccName 模糊查找 BankAccount
func (s *service) FzzFindBankAccountByBankAccName(bankAccount *[]models.BankAccount, BankAccName string) {
	s.gormDB.Where("bankAccName LIKE ?", "%"+BankAccName+"%").Find(bankAccount)
}

// 根据 CompName 模糊查找 BankAccount
func (s *service) FzzFindBankAccountByCompName(bankAccount *[]models.BankAccount, CompName string) {
	s.gormDB.Where("compName LIKE ?", "%"+CompName+"%").Find(bankAccount)
}

// 根据 Currency 模糊查找 BankAccount
func (s *service) FzzFindBankAccountByCurrency(bankAccount *[]models.BankAccount, Currency string) {
	s.gormDB.Where("currency LIKE ?", "%"+Currency+"%").Find(bankAccount)
}

// 根据 BankName 模糊查找 BankAccount
func (s *service) FzzFindBankAccountByBankName(bankAccount *[]models.BankAccount, BankName string) {
	s.gormDB.Where("bankName LIKE ?", "%"+BankName+"%").Find(bankAccount)
}

// 根据 BankEng 模糊查找 BankAccount
func (s *service) FzzFindBankAccountByBankEng(bankAccount *[]models.BankAccount, BankEng string) {
	s.gormDB.Where("bankEng LIKE ?", "%"+BankEng+"%").Find(bankAccount)
}

// 根据 BankAddr 模糊查找 BankAccount
func (s *service) FzzFindBankAccountByBankAddr(bankAccount *[]models.BankAccount, BankAddr string) {
	s.gormDB.Where("bankAddr LIKE ?", "%"+BankAddr+"%").Find(bankAccount)
}

// 根据 Notes 模糊查找 BankAccount
func (s *service) FzzFindBankAccountByNotes(bankAccount *[]models.BankAccount, Notes string) {
	s.gormDB.Where("notes LIKE ?", "%"+Notes+"%").Find(bankAccount)
}
