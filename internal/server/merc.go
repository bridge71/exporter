package server

import (
	"exporter/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 根据 BankAccName 查找 BankAccount
func (s *Server) FindBankAccountByBankAccNameHandler(c *gin.Context) {
	bankAccounts := &[]models.BankAccount{}
	bankAccName := c.Query("BankAccName")
	s.db.FindBankAccountByBankAccName(bankAccounts, bankAccName)
	c.JSON(http.StatusOK, models.Message{
		BankAccount: *bankAccounts,
	})
}

// 根据 CompName 查找 BankAccount
func (s *Server) FindBankAccountByCompNameHandler(c *gin.Context) {
	bankAccounts := &[]models.BankAccount{}
	compName := c.Query("CompName")
	s.db.FindBankAccountByCompName(bankAccounts, compName)
	c.JSON(http.StatusOK, models.Message{
		BankAccount: *bankAccounts,
	})
}

// 根据 Currency 查找 BankAccount
func (s *Server) FindBankAccountByCurrencyHandler(c *gin.Context) {
	bankAccounts := &[]models.BankAccount{}
	currency := c.Query("Currency")
	s.db.FindBankAccountByCurrency(bankAccounts, currency)
	c.JSON(http.StatusOK, models.Message{
		BankAccount: *bankAccounts,
	})
}

// 根据 BankName 查找 BankAccount
func (s *Server) FindBankAccountByBankNameHandler(c *gin.Context) {
	bankAccounts := &[]models.BankAccount{}
	bankName := c.Query("BankName")
	s.db.FindBankAccountByBankName(bankAccounts, bankName)
	c.JSON(http.StatusOK, models.Message{
		BankAccount: *bankAccounts,
	})
}

// 根据 BankEng 查找 BankAccount
func (s *Server) FindBankAccountByBankEngHandler(c *gin.Context) {
	bankAccounts := &[]models.BankAccount{}
	bankEng := c.Query("BankEng")
	s.db.FindBankAccountByBankEng(bankAccounts, bankEng)
	c.JSON(http.StatusOK, models.Message{
		BankAccount: *bankAccounts,
	})
}

// 根据 BankAddr 查找 BankAccount
func (s *Server) FindBankAccountByBankAddrHandler(c *gin.Context) {
	bankAccounts := &[]models.BankAccount{}
	bankAddr := c.Query("BankAddr")
	s.db.FindBankAccountByBankAddr(bankAccounts, bankAddr)
	c.JSON(http.StatusOK, models.Message{
		BankAccount: *bankAccounts,
	})
}

// 根据 Notes 查找 BankAccount
func (s *Server) FindBankAccountByNotesHandler(c *gin.Context) {
	bankAccounts := &[]models.BankAccount{}
	notes := c.Query("Notes")
	s.db.FindBankAccountByNotes(bankAccounts, notes)
	c.JSON(http.StatusOK, models.Message{
		BankAccount: *bankAccounts,
	})
}

// 根据 Addr 模糊查找 Cust
func (s *Server) FzzFindCustByAddrHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	addr := c.Query("Addr")
	s.db.FzzFindCustByAddr(custs, addr)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 Email 模糊查找 Cust
func (s *Server) FzzFindCustByEmailHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	email := c.Query("Email")
	s.db.FzzFindCustByEmail(custs, email)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 PhoneNum 模糊查找 Cust
func (s *Server) FzzFindCustByPhoneNumHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	phoneNum := c.Query("PhoneNum")
	s.db.FzzFindCustByPhoneNum(custs, phoneNum)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 QQ 模糊查找 Cust
func (s *Server) FzzFindCustByQQHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	qq := c.Query("QQ")
	s.db.FzzFindCustByQQ(custs, qq)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 Wechat 模糊查找 Cust
func (s *Server) FzzFindCustByWechatHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	wechat := c.Query("Wechat")
	s.db.FzzFindCustByWechat(custs, wechat)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 Merc 模糊查找 Cust
func (s *Server) FzzFindCustByMercHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	merc := c.Query("Merc")
	s.db.FzzFindCustByMerc(custs, merc)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 Post 模糊查找 Cust
func (s *Server) FzzFindCustByPostHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	post := c.Query("Post")
	s.db.FzzFindCustByPost(custs, post)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 Notes 模糊查找 Cust
func (s *Server) FzzFindCustByNotesHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	notes := c.Query("Notes")
	s.db.FzzFindCustByNotes(custs, notes)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 Name 精确查找 Cust
func (s *Server) FindCustByNameHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	name := c.Query("Name")
	s.db.FindCustByName(custs, name)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 Gender 精确查找 Cust
func (s *Server) FindCustByGenderHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	gender := c.Query("Gender")
	s.db.FindCustByGender(custs, gender)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 Nation 精确查找 Cust
func (s *Server) FindCustByNationHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	nation := c.Query("Nation")
	s.db.FindCustByNation(custs, nation)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 Addr 精确查找 Cust
func (s *Server) FindCustByAddrHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	addr := c.Query("Addr")
	s.db.FindCustByAddr(custs, addr)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 Email 精确查找 Cust
func (s *Server) FindCustByEmailHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	email := c.Query("Email")
	s.db.FindCustByEmail(custs, email)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 PhoneNum 精确查找 Cust
func (s *Server) FindCustByPhoneNumHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	phoneNum := c.Query("PhoneNum")
	s.db.FindCustByPhoneNum(custs, phoneNum)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 QQ 精确查找 Cust
func (s *Server) FindCustByQQHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	qq := c.Query("QQ")
	s.db.FindCustByQQ(custs, qq)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 Wechat 精确查找 Cust
func (s *Server) FindCustByWechatHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	wechat := c.Query("Wechat")
	s.db.FindCustByWechat(custs, wechat)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 Merc 精确查找 Cust
func (s *Server) FindCustByMercHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	merc := c.Query("Merc")
	s.db.FindCustByMerc(custs, merc)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 Post 精确查找 Cust
func (s *Server) FindCustByPostHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	post := c.Query("Post")
	s.db.FindCustByPost(custs, post)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 Notes 精确查找 Cust
func (s *Server) FindCustByNotesHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	notes := c.Query("Notes")
	s.db.FindCustByNotes(custs, notes)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 MercCode 精确查找 Merchant
func (s *Server) FindMerchantByMercCodeHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	mercCode := c.Query("MercCode")
	s.db.FindMerchantByMercCode(merchants, mercCode)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 MercAbbr 精确查找 Merchant
func (s *Server) FindMerchantByMercAbbrHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	mercAbbr := c.Query("MercAbbr")
	s.db.FindMerchantByMercAbbr(merchants, mercAbbr)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 ShortMerc 精确查找 Merchant
func (s *Server) FindMerchantByShortMercHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	shortMerc := c.Query("ShortMerc")
	s.db.FindMerchantByShortMerc(merchants, shortMerc)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 Merc 精确查找 Merchant
func (s *Server) FindMerchantByMercHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	merc := c.Query("Merc")
	s.db.FindMerchantByMerc(merchants, merc)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 EngName 精确查找 Merchant
func (s *Server) FindMerchantByEngNameHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	engName := c.Query("EngName")
	s.db.FindMerchantByEngName(merchants, engName)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 Address 精确查找 Merchant
func (s *Server) FindMerchantByAddressHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	address := c.Query("Address")
	s.db.FindMerchantByAddress(merchants, address)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 Nation 精确查找 Merchant
func (s *Server) FindMerchantByNationHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	nation := c.Query("Nation")
	s.db.FindMerchantByNation(merchants, nation)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 PhoneNum 精确查找 Merchant
func (s *Server) FindMerchantByPhoneNumHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	phoneNum := c.Query("PhoneNum")
	s.db.FindMerchantByPhoneNum(merchants, phoneNum)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 Email 精确查找 Merchant
func (s *Server) FindMerchantByEmailHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	email := c.Query("Email")
	s.db.FindMerchantByEmail(merchants, email)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 Fax 精确查找 Merchant
func (s *Server) FindMerchantByFaxHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	fax := c.Query("Fax")
	s.db.FindMerchantByFax(merchants, fax)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 Website 精确查找 Merchant
func (s *Server) FindMerchantByWebsiteHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	website := c.Query("Website")
	s.db.FindMerchantByWebsite(merchants, website)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 TaxType 精确查找 Merchant
func (s *Server) FindMerchantByTaxTypeHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	taxType := c.Query("TaxType")
	s.db.FindMerchantByTaxType(merchants, taxType)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 TaxCode 精确查找 Merchant
func (s *Server) FindMerchantByTaxCodeHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	taxCode := c.Query("TaxCode")
	s.db.FindMerchantByTaxCode(merchants, taxCode)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 MercType 精确查找 Merchant
func (s *Server) FindMerchantByMercTypeHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	mercType := c.Query("MercType")
	s.db.FindMerchantByMercType(merchants, mercType)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 RegCap 精确查找 Merchant
func (s *Server) FindMerchantByRegCapHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	regCap := c.Query("RegCap")
	s.db.FindMerchantByRegCap(merchants, regCap)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 Notes 精确查找 Merchant
func (s *Server) FindMerchantByNotesHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	notes := c.Query("Notes")
	s.db.FindMerchantByNotes(merchants, notes)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

func (s *Server) SaveMerchantHandler(c *gin.Context) {
	merchant := &models.Merchant{}
	err := c.ShouldBind(merchant)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			// RetMessage: "error in bind of merchant",
			RetMessage: err.Error(),
		})
		return
	}
	err, merchant.FileId, merchant.FileName = s.SaveFile(c)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error when saving file ",
		})
		return
	}
	err = s.db.SaveMerchant(merchant)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error when saving merchant ",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved merchant successfully",
		})
	}
}

func (s *Server) SaveCustHandler(c *gin.Context) {
	cust := &models.Cust{}
	err := c.ShouldBind(cust)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of cust",
		})
		return
	}

	err = s.db.SaveCust(cust)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "error when saving cust",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved cust successfully",
		})
	}
}

func (s *Server) SaveBankAccountHandler(c *gin.Context) {
	bankAccount := &models.BankAccount{}
	err := c.ShouldBind(bankAccount)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of bank account",
		})
		return
	}

	err, bankAccount.FileId, bankAccount.FileName = s.SaveFile(c)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error when saving file",
		})
		return
	}

	err = s.db.SaveBankAccount(bankAccount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "error when saving bank account",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved bank account successfully",
		})
	}
}

func (s *Server) DeleteMerchantHandler(c *gin.Context) {
	merchant := &models.Merchant{}
	err := c.ShouldBind(merchant)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
		})
		return
	}
	log.Printf("%v\n", merchant)

	err = s.db.DeleteMerchant(merchant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete merchant",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete merchant successfully",
		})
	}
}

func (s *Server) DeleteCustHandler(c *gin.Context) {
	cust := &models.Cust{}
	err := c.ShouldBind(cust)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of cust",
		})
		return
	}
	log.Printf("%v\n", cust)

	err = s.db.DeleteCust(cust)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete cust",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete cust successfully",
		})
	}
}

func (s *Server) DeleteBankAccountHandler(c *gin.Context) {
	bankAccount := &models.BankAccount{}
	err := c.ShouldBind(bankAccount)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of bank account",
		})
		return
	}
	log.Printf("%v\n", bankAccount)

	err = s.db.DeleteBankAccount(bankAccount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete bank account",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete bank account successfully",
		})
	}
}

func (s *Server) FindMerchantHandler(c *gin.Context) {
	merchants := []models.Merchant{}
	s.db.FindMerchant(&merchants)
	for i := range merchants {
		// 查找关联的 Cust 数据
		custs := []models.Cust{}
		s.db.FindCustByMercId(&custs, merchants[i].MercId)
		merchants[i].Custs = append(merchants[i].Custs, custs...)

		// 查找关联的 BankAccount 数据
		bankAccounts := []models.BankAccount{}
		s.db.FindBankAccountByMercId(&bankAccounts, merchants[i].MercId)
		merchants[i].BankAccounts = append(merchants[i].BankAccounts, bankAccounts...)
	}
	c.JSON(http.StatusOK, models.Message{
		Merchant: merchants,
	})
}

func (s *Server) FindCustHandler(c *gin.Context) {
	custs := []models.Cust{}
	s.db.FindCust(&custs)
	c.JSON(http.StatusOK, models.Message{
		Cust: custs,
	})
}

func (s *Server) FindBankAccountHandler(c *gin.Context) {
	bankAccounts := []models.BankAccount{}
	s.db.FindBankAccount(&bankAccounts)
	c.JSON(http.StatusOK, models.Message{
		BankAccount: bankAccounts,
	})
}

// 根据 MercAbbr 模糊查找 Merchant
func (s *Server) FzzFindMerchantByMercAbbrHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	mercAbbr := c.PostForm("MercAbbr")
	s.db.FzzFindMerchantByMercAbbr(merchants, mercAbbr)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 ShortMerc 模糊查找 Merchant
func (s *Server) FzzFindMerchantByShortMercHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	shortMerc := c.PostForm("ShortMerc")
	s.db.FzzFindMerchantByShortMerc(merchants, shortMerc)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 Merc 模糊查找 Merchant
func (s *Server) FzzFindMerchantByMercHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	merc := c.PostForm("Merc")
	s.db.FzzFindMerchantByMerc(merchants, merc)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 EngName 模糊查找 Merchant
func (s *Server) FzzFindMerchantByEngNameHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	engName := c.PostForm("EngName")
	s.db.FzzFindMerchantByEngName(merchants, engName)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 Address 模糊查找 Merchant
func (s *Server) FzzFindMerchantByAddressHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	address := c.PostForm("Address")
	s.db.FzzFindMerchantByAddress(merchants, address)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 Nation 模糊查找 Merchant
func (s *Server) FzzFindMerchantByNationHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	nation := c.PostForm("Nation")
	s.db.FzzFindMerchantByNation(merchants, nation)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 Website 模糊查找 Merchant
func (s *Server) FzzFindMerchantByWebsiteHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	website := c.PostForm("Website")
	s.db.FzzFindMerchantByWebsite(merchants, website)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 TaxType 模糊查找 Merchant
func (s *Server) FzzFindMerchantByTaxTypeHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	taxType := c.PostForm("TaxType")
	s.db.FzzFindMerchantByTaxType(merchants, taxType)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 MercType 模糊查找 Merchant
func (s *Server) FzzFindMerchantByMercTypeHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	mercType := c.PostForm("MercType")
	s.db.FzzFindMerchantByMercType(merchants, mercType)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 Notes 模糊查找 Merchant
func (s *Server) FzzFindMerchantByNotesHandler(c *gin.Context) {
	merchants := &[]models.Merchant{}
	notes := c.PostForm("Notes")
	s.db.FzzFindMerchantByNotes(merchants, notes)
	c.JSON(http.StatusOK, models.Message{
		Merchant: *merchants,
	})
}

// 根据 Name 模糊查找 Cust
func (s *Server) FzzFindCustByNameHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	name := c.PostForm("Name")
	s.db.FzzFindCustByName(custs, name)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 Nation 模糊查找 Cust
func (s *Server) FzzFindCustByNationHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	nation := c.PostForm("Nation")
	s.db.FzzFindCustByNation(custs, nation)
	c.JSON(http.StatusOK, models.Message{
		Cust: *custs,
	})
}

// 根据 BankAccName 模糊查找 BankAccount
func (s *Server) FzzFindBankAccountByBankAccNameHandler(c *gin.Context) {
	bankAccounts := &[]models.BankAccount{}
	bankAccName := c.PostForm("BankAccName")
	s.db.FzzFindBankAccountByBankAccName(bankAccounts, bankAccName)
	c.JSON(http.StatusOK, models.Message{
		BankAccount: *bankAccounts,
	})
}

// 根据 CompName 模糊查找 BankAccount
func (s *Server) FzzFindBankAccountByCompNameHandler(c *gin.Context) {
	bankAccounts := &[]models.BankAccount{}
	compName := c.PostForm("CompName")
	s.db.FzzFindBankAccountByCompName(bankAccounts, compName)
	c.JSON(http.StatusOK, models.Message{
		BankAccount: *bankAccounts,
	})
}

// 根据 AcctNum 查找 BankAccount
func (s *Server) FindBankAccountByAcctNumHandler(c *gin.Context) {
	bankAccounts := &[]models.BankAccount{}
	acctNum := c.PostForm("AcctNum")
	s.db.FindBankAccountByAcctNum(bankAccounts, acctNum)
	c.JSON(http.StatusOK, models.Message{
		BankAccount: *bankAccounts,
	})
}

// 根据 Currency 模糊查找 BankAccount
func (s *Server) FzzFindBankAccountByCurrencyHandler(c *gin.Context) {
	bankAccounts := &[]models.BankAccount{}
	currency := c.PostForm("Currency")
	s.db.FzzFindBankAccountByCurrency(bankAccounts, currency)
	c.JSON(http.StatusOK, models.Message{
		BankAccount: *bankAccounts,
	})
}

// 根据 BankName 模糊查找 BankAccount
func (s *Server) FzzFindBankAccountByBankNameHandler(c *gin.Context) {
	bankAccounts := &[]models.BankAccount{}
	bankName := c.PostForm("BankName")
	s.db.FzzFindBankAccountByBankName(bankAccounts, bankName)
	c.JSON(http.StatusOK, models.Message{
		BankAccount: *bankAccounts,
	})
}

// 根据 BankEng 模糊查找 BankAccount
func (s *Server) FzzFindBankAccountByBankEngHandler(c *gin.Context) {
	bankAccounts := &[]models.BankAccount{}
	bankEng := c.PostForm("BankEng")
	s.db.FzzFindBankAccountByBankEng(bankAccounts, bankEng)
	c.JSON(http.StatusOK, models.Message{
		BankAccount: *bankAccounts,
	})
}

// 根据 BankNum 查找 BankAccount
func (s *Server) FindBankAccountByBankNumHandler(c *gin.Context) {
	bankAccounts := &[]models.BankAccount{}
	bankNum := c.PostForm("BankNum")
	s.db.FindBankAccountByBankNum(bankAccounts, bankNum)
	c.JSON(http.StatusOK, models.Message{
		BankAccount: *bankAccounts,
	})
}

// 根据 SwiftCode 查找 BankAccount
func (s *Server) FindBankAccountBySwiftCodeHandler(c *gin.Context) {
	bankAccounts := &[]models.BankAccount{}
	swiftCode := c.PostForm("SwiftCode")
	s.db.FindBankAccountBySwiftCode(bankAccounts, swiftCode)
	c.JSON(http.StatusOK, models.Message{
		BankAccount: *bankAccounts,
	})
}

// 根据 BankAddr 模糊查找 BankAccount
func (s *Server) FzzFindBankAccountByBankAddrHandler(c *gin.Context) {
	bankAccounts := &[]models.BankAccount{}
	bankAddr := c.PostForm("BankAddr")
	s.db.FzzFindBankAccountByBankAddr(bankAccounts, bankAddr)
	c.JSON(http.StatusOK, models.Message{
		BankAccount: *bankAccounts,
	})
}

// 根据 Notes 模糊查找 BankAccount
func (s *Server) FzzFindBankAccountByNotesHandler(c *gin.Context) {
	bankAccounts := &[]models.BankAccount{}
	notes := c.PostForm("Notes")
	s.db.FzzFindBankAccountByNotes(bankAccounts, notes)
	c.JSON(http.StatusOK, models.Message{
		BankAccount: *bankAccounts,
	})
}
