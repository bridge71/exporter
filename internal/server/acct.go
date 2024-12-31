package server

import (
	"crypto/md5"
	"encoding/hex"
	"exporter/internal/models"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) Uploader(c *gin.Context) {
	file := &models.File{}
	err := c.ShouldBind(file)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of file",
		})
		return
	}
	s.db.FindFile(file)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename*=utf-8''"+file.Name+file.Suffix)
	c.File("../files/" + file.MD5 + file.Suffix)

	c.JSON(http.StatusOK, models.Message{
		RetMessage: file.Name,
	})
}

func (s *Server) FindAcctHandler(c *gin.Context) {
	accts := []models.Acct{}
	s.db.FindAcct(&accts)
	for i := range accts {
		acctBank := []models.AcctBank{}
		s.db.FindAcctBankById(&acctBank, accts[i].AcctId)
		accts[i].AcctBanks = append(accts[i].AcctBanks, acctBank...)
	}
	c.JSON(http.StatusOK, models.Message{
		Acct: accts,
	})
}

func (s *Server) FindAcctByAcctNameHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	acctName := c.PostForm("AcctName")
	s.db.FindAcctByAcctName(accts, acctName)
	c.JSON(http.StatusOK, models.Message{
		Acct: *accts,
	})
}

func (s *Server) FindAcctBankByAccNameHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	accName := c.PostForm("AccName")
	s.db.FindAcctBankByAccName(acctBanks, accName)
	c.JSON(http.StatusOK, models.Message{
		AcctBank: *acctBanks,
	})
}

// FindAcctBankHandler 查找账户银行处理函数
func (s *Server) FindAcctBankHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	s.db.FindAcctBank(acctBanks)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

// FindAcctBankByAccNumHandler 根据账户编号查找账户银行处理函数
func (s *Server) FindAcctBankByAccNumHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	accNum := c.PostForm("AccNum")
	s.db.FindAcctBankByAccNum(acctBanks, accNum)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

// FindAcctBankByCurrencyHandler 根据货币查找账户银行处理函数
func (s *Server) FindAcctBankByCurrencyHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	currency := c.PostForm("Currency")
	s.db.FindAcctBankByCurrency(acctBanks, currency)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

// FindAcctBankByBankNameHandler 根据银行名称查找账户银行处理函数
func (s *Server) FindAcctBankByBankNameHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	bankName := c.PostForm("BankName")
	s.db.FindAcctBankByBankName(acctBanks, bankName)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

// FindAcctBankByIdHandler 根据ID查找账户银行处理函数
func (s *Server) FindAcctBankByIdHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	acctId := c.PostForm("AcctId")
	idStr, _ := strconv.Atoi(acctId)
	id := uint(idStr)
	s.db.FindAcctBankById(acctBanks, id)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

// FindAcctByEtyAbbrHandler 根据实体缩写查找账户处理函数
func (s *Server) FindAcctByEtyAbbrHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	EtyAbbr := c.PostForm("EtyAbbr")
	s.db.FindAcctByEtyAbbr(accts, EtyAbbr)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FindAcctByAcctAddrHandler 根据账户地址查找账户处理函数
func (s *Server) FindAcctByAcctAddrHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	acctAddr := c.PostForm("AcctAddr")
	s.db.FindAcctByAcctAddr(accts, acctAddr)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FzzFindAcctByAcctAddrHandler 根据账户地址模糊查找账户处理函数
func (s *Server) FzzFindAcctByAcctAddrHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	acctAddr := c.PostForm("AcctAddr")
	s.db.FzzFindAcctByAcctAddr(accts, acctAddr)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FindAcctByNationHandler 根据国家查找账户处理函数
func (s *Server) FindAcctByNationHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	nation := c.PostForm("Nation")
	s.db.FindAcctByNation(accts, nation)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FindAcctByTaxTypeHandler 根据税务类型查找账户处理函数
func (s *Server) FindAcctByTaxTypeHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	taxType := c.PostForm("TaxType")
	s.db.FindAcctByTaxType(accts, taxType)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FindAcctByTaxCodeHandler 根据税务代码查找账户处理函数
func (s *Server) FindAcctByTaxCodeHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	taxCode := c.PostForm("TaxCode")
	s.db.FindAcctByTaxCode(accts, taxCode)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FindAcctByPhoneNumHandler 根据电话号码查找账户处理函数
func (s *Server) FindAcctByPhoneNumHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	phoneNum := c.PostForm("PhoneNum")
	s.db.FindAcctByPhoneNum(accts, phoneNum)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FzzFindAcctByPhoneNumHandler 根据电话号码模糊查找账户处理函数
func (s *Server) FzzFindAcctByPhoneNumHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	phoneNum := c.PostForm("PhoneNum")
	s.db.FzzFindAcctByPhoneNum(accts, phoneNum)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FindAcctByEmailHandler 根据电子邮件查找账户处理函数
func (s *Server) FindAcctByEmailHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	email := c.PostForm("Email")
	s.db.FindAcctByEmail(accts, email)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FzzFindAcctByEmailHandler 根据电子邮件模糊查找账户处理函数
func (s *Server) FzzFindAcctByEmailHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	email := c.PostForm("Email")
	s.db.FzzFindAcctByEmail(accts, email)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FindAcctByWebsiteHandler 根据网站查找账户处理函数
func (s *Server) FindAcctByWebsiteHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	website := c.PostForm("Website")
	s.db.FindAcctByWebsite(accts, website)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FzzFindAcctByWebsiteHandler 根据网站模糊查找账户处理函数
func (s *Server) FzzFindAcctByWebsiteHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	website := c.PostForm("Website")
	s.db.FzzFindAcctByWebsite(accts, website)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FindAcctByRegDateHandler 根据注册日期查找账户处理函数
func (s *Server) FindAcctByRegDateHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	regDate := c.PostForm("RegDate")
	s.db.FindAcctByRegDate(accts, regDate)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FzzFindAcctByRegDateHandler 根据注册日期模糊查找账户处理函数
func (s *Server) FzzFindAcctByRegDateHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	regDate := c.PostForm("RegDate")
	s.db.FzzFindAcctByRegDate(accts, regDate)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FindAcctByNotesHandler 根据备注查找账户处理函数
func (s *Server) FindAcctByNotesHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	notes := c.PostForm("Notes")
	s.db.FindAcctByNotes(accts, notes)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FzzFindAcctByNotesHandler 根据备注模糊查找账户处理函数
func (s *Server) FzzFindAcctByNotesHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	notes := c.PostForm("Notes")
	s.db.FzzFindAcctByNotes(accts, notes)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FzzFindAcctBankByBankNameHandler 根据银行名称模糊查找账户银行处理函数
func (s *Server) FzzFindAcctBankByBankNameHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	bankName := c.PostForm("BankName")
	s.db.FzzFindAcctBankByBankName(acctBanks, bankName)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

// FindAcctBankBySwiftCodeHandler 根据SWIFT代码查找账户银行处理函数
func (s *Server) FindAcctBankBySwiftCodeHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	swiftCode := c.PostForm("SwiftCode")
	s.db.FindAcctBankBySwiftCode(acctBanks, swiftCode)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

// FindAcctBankByBankAddrHandler 根据银行地址查找账户银行处理函数
func (s *Server) FindAcctBankByBankAddrHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	bankAddr := c.PostForm("BankAddr")
	s.db.FindAcctBankByBankAddr(acctBanks, bankAddr)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

// FzzFindAcctBankByBankAddrHandler 根据银行地址模糊查找账户银行处理函数
func (s *Server) FzzFindAcctBankByBankAddrHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	bankAddr := c.PostForm("BankAddr")
	s.db.FzzFindAcctBankByBankAddr(acctBanks, bankAddr)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

// FzzFindAcctBankByNotesHandler 根据备注模糊查找账户银行处理函数
func (s *Server) FzzFindAcctBankByNotesHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	notes := c.PostForm("Notes")
	s.db.FzzFindAcctBankByNotes(acctBanks, notes)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

// FindAcctBankByNotesHandler 根据备注查找账户银行处理函数
func (s *Server) FindAcctBankByNotesHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	notes := c.PostForm("Notes")
	s.db.FindAcctBankByNotes(acctBanks, notes)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

// FindAcctByAcctCodeHandler 根据账户代码查找账户处理函数
func (s *Server) FindAcctByAcctCodeHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	acctCode := c.PostForm("AcctCode")
	s.db.FindAcctByAcctCode(accts, acctCode)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FindAcctByAcctAbbrHandler 根据账户缩写查找账户处理函数
func (s *Server) FindAcctByAcctAbbrHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	acctAbbr := c.PostForm("AcctAbbr")
	s.db.FindAcctByAcctAbbr(accts, acctAbbr)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

func (s *Server) DeleteAcctBankHandler(c *gin.Context) {
	acctBank := &models.AcctBank{}
	err := c.ShouldBind(acctBank)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of acctBank",
		})
		return
	}
	log.Printf("%v\n", acctBank)

	err = s.db.DeleteAcctBank(acctBank)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to delete acctBank",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete acctBank successfully",
		})
	}
}

func (s *Server) DeleteAcctHandler(c *gin.Context) {
	acct := &models.Acct{}
	err := c.ShouldBind(acct)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of acct",
		})
		return
	}
	log.Printf("%v\n", acct)

	err = s.db.DeleteAcct(acct)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to delete acct",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete acct successfully",
		})
	}
}

func (s *Server) SaveFile(c *gin.Context) (error, uint, string) {
	file, err := c.FormFile("file")
	if err == nil {
		src, err := file.Open()
		if err != nil {
			return err, 0, ""
		}
		defer src.Close()

		hash := md5.New()
		if _, err := io.Copy(hash, src); err != nil {
			return err, 0, ""
		}
		MD5 := hex.EncodeToString(hash.Sum(nil))
		FileName := file.Filename
		Suffix := filepath.Ext(file.Filename)
		File := &models.File{
			Name:   FileName,
			MD5:    MD5,
			Suffix: Suffix,
		}
		s.db.CreateFile(File)

		c.SaveUploadedFile(file, "../files/"+MD5+Suffix) // 以main文件所在目录为基准

		log.Printf(file.Filename + "\n")
		return nil, File.FileId, FileName
	}
	return nil, 0, ""
}

func (s *Server) SaveAcctBankHandler(c *gin.Context) {
	acctBank := &models.AcctBank{}
	err := c.ShouldBind(acctBank)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of acctBank",
		})
		return
	}
	log.Printf("%v\n", acctBank)

	err, acctBank.FileId, acctBank.FileName = s.SaveFile(c)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to save file",
		})
	}
	err = s.db.SaveAcctBank(acctBank)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to save acctBank",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved acctBank successfully",
		})
	}
}

func (s *Server) Str2Uint(str string) uint {
	// 使用 strconv.ParseUint 将字符串解析为 uint64
	// 参数 10 表示十进制，参数 0 表示自动推断位数
	value, err := strconv.ParseUint(str, 10, 0)
	if err != nil {
		// 如果解析失败，返回错误
		return 0
	}
	// 将 uint64 转换为 uint 并返回
	return uint(value)
}

func (s *Server) SaveAcctHandler(c *gin.Context) {
	acct := &models.Acct{}
	err := c.ShouldBind(acct)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of acct",
		})
		return
	}

	err, acct.FileId, acct.FileName = s.SaveFile(c)
	log.Printf("%v", acct)
	err = s.db.SaveAcct(acct)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error when insert accounting information",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "insert accounting information successfully",
		})
	}
}
