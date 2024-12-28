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

func (s *Server) FindAcctHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	s.db.FindAcct(accts)
	c.JSON(http.StatusOK, models.Message{
		Acct: *accts,
	})
}

func (s *Server) FindAcctByAcctNameHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	acctName := c.PostForm("acctName")
	s.db.FindAcctByAcctName(accts, acctName)
	c.JSON(http.StatusOK, models.Message{
		Acct: *accts,
	})
}

func (s *Server) FindAcctBankByAccNameHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	accName := c.PostForm("accName")
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
	accNum := c.PostForm("accNum")
	s.db.FindAcctBankByAccNum(acctBanks, accNum)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

// FindAcctBankByCurrencyHandler 根据货币查找账户银行处理函数
func (s *Server) FindAcctBankByCurrencyHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	currency := c.PostForm("currency")
	s.db.FindAcctBankByCurrency(acctBanks, currency)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

// FindAcctBankByBankNameHandler 根据银行名称查找账户银行处理函数
func (s *Server) FindAcctBankByBankNameHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	bankName := c.PostForm("bankName")
	s.db.FindAcctBankByBankName(acctBanks, bankName)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

// FindAcctBankByIdHandler 根据ID查找账户银行处理函数
func (s *Server) FindAcctBankByIdHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	acctId := c.PostForm("acctId")
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
	acctAddr := c.PostForm("acctAddr")
	s.db.FindAcctByAcctAddr(accts, acctAddr)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FzzFindAcctByAcctAddrHandler 根据账户地址模糊查找账户处理函数
func (s *Server) FzzFindAcctByAcctAddrHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	acctAddr := c.PostForm("acctAddr")
	s.db.FzzFindAcctByAcctAddr(accts, acctAddr)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FindAcctByNationHandler 根据国家查找账户处理函数
func (s *Server) FindAcctByNationHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	nation := c.PostForm("nation")
	s.db.FindAcctByNation(accts, nation)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FindAcctByTaxTypeHandler 根据税务类型查找账户处理函数
func (s *Server) FindAcctByTaxTypeHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	taxType := c.PostForm("taxType")
	s.db.FindAcctByTaxType(accts, taxType)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FindAcctByTaxCodeHandler 根据税务代码查找账户处理函数
func (s *Server) FindAcctByTaxCodeHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	taxCode := c.PostForm("taxCode")
	s.db.FindAcctByTaxCode(accts, taxCode)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FindAcctByPhoneNumHandler 根据电话号码查找账户处理函数
func (s *Server) FindAcctByPhoneNumHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	phoneNum := c.PostForm("phoneNum")
	s.db.FindAcctByPhoneNum(accts, phoneNum)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FzzFindAcctByPhoneNumHandler 根据电话号码模糊查找账户处理函数
func (s *Server) FzzFindAcctByPhoneNumHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	phoneNum := c.PostForm("phoneNum")
	s.db.FzzFindAcctByPhoneNum(accts, phoneNum)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FindAcctByEmailHandler 根据电子邮件查找账户处理函数
func (s *Server) FindAcctByEmailHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	email := c.PostForm("email")
	s.db.FindAcctByEmail(accts, email)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FzzFindAcctByEmailHandler 根据电子邮件模糊查找账户处理函数
func (s *Server) FzzFindAcctByEmailHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	email := c.PostForm("email")
	s.db.FzzFindAcctByEmail(accts, email)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FindAcctByWebsiteHandler 根据网站查找账户处理函数
func (s *Server) FindAcctByWebsiteHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	website := c.PostForm("website")
	s.db.FindAcctByWebsite(accts, website)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FzzFindAcctByWebsiteHandler 根据网站模糊查找账户处理函数
func (s *Server) FzzFindAcctByWebsiteHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	website := c.PostForm("website")
	s.db.FzzFindAcctByWebsite(accts, website)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FindAcctByRegDateHandler 根据注册日期查找账户处理函数
func (s *Server) FindAcctByRegDateHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	regDate := c.PostForm("regDate")
	s.db.FindAcctByRegDate(accts, regDate)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FzzFindAcctByRegDateHandler 根据注册日期模糊查找账户处理函数
func (s *Server) FzzFindAcctByRegDateHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	regDate := c.PostForm("regDate")
	s.db.FzzFindAcctByRegDate(accts, regDate)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FindAcctByNotesHandler 根据备注查找账户处理函数
func (s *Server) FindAcctByNotesHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	notes := c.PostForm("notes")
	s.db.FindAcctByNotes(accts, notes)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FzzFindAcctByNotesHandler 根据备注模糊查找账户处理函数
func (s *Server) FzzFindAcctByNotesHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	notes := c.PostForm("notes")
	s.db.FzzFindAcctByNotes(accts, notes)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FzzFindAcctBankByBankNameHandler 根据银行名称模糊查找账户银行处理函数
func (s *Server) FzzFindAcctBankByBankNameHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	bankName := c.PostForm("bankName")
	s.db.FzzFindAcctBankByBankName(acctBanks, bankName)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

// FindAcctBankBySwiftCodeHandler 根据SWIFT代码查找账户银行处理函数
func (s *Server) FindAcctBankBySwiftCodeHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	swiftCode := c.PostForm("swiftCode")
	s.db.FindAcctBankBySwiftCode(acctBanks, swiftCode)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

// FindAcctBankByBankAddrHandler 根据银行地址查找账户银行处理函数
func (s *Server) FindAcctBankByBankAddrHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	bankAddr := c.PostForm("bankAddr")
	s.db.FindAcctBankByBankAddr(acctBanks, bankAddr)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

// FzzFindAcctBankByBankAddrHandler 根据银行地址模糊查找账户银行处理函数
func (s *Server) FzzFindAcctBankByBankAddrHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	bankAddr := c.PostForm("bankAddr")
	s.db.FzzFindAcctBankByBankAddr(acctBanks, bankAddr)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

// FzzFindAcctBankByNotesHandler 根据备注模糊查找账户银行处理函数
func (s *Server) FzzFindAcctBankByNotesHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	notes := c.PostForm("notes")
	s.db.FzzFindAcctBankByNotes(acctBanks, notes)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

// FindAcctBankByNotesHandler 根据备注查找账户银行处理函数
func (s *Server) FindAcctBankByNotesHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	notes := c.PostForm("notes")
	s.db.FindAcctBankByNotes(acctBanks, notes)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

// FindAcctByAcctCodeHandler 根据账户代码查找账户处理函数
func (s *Server) FindAcctByAcctCodeHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	acctCode := c.PostForm("acctCode")
	s.db.FindAcctByAcctCode(accts, acctCode)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

// FindAcctByAcctAbbrHandler 根据账户缩写查找账户处理函数
func (s *Server) FindAcctByAcctAbbrHandler(c *gin.Context) {
	accts := &[]models.Acct{}
	acctAbbr := c.PostForm("acctAbbr")
	s.db.FindAcctByAcctAbbr(accts, acctAbbr)
	c.JSON(http.StatusOK, models.Message{Acct: *accts})
}

func (s *Server) CreateAcctHandler(c *gin.Context) {
	acct := &models.Acct{}
	err := c.ShouldBind(acct)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of acct",
		})
		return
	}

	bankIndex := 0
	for {
		accNum := c.PostForm("AccBanks[" + strconv.Itoa(bankIndex) + "].AccNum")
		accName := c.PostForm("AccBanks[" + strconv.Itoa(bankIndex) + "].AccName")
		isUpload := c.PostForm("AccBanks[" + strconv.Itoa(bankIndex) + "].IsUpload")
		currency := c.PostForm("AccBanks[" + strconv.Itoa(bankIndex) + "].Currency")
		bankName := c.PostForm("AccBanks[" + strconv.Itoa(bankIndex) + "].BankName")
		bankNum := c.PostForm("AccBanks[" + strconv.Itoa(bankIndex) + "].BankNum")
		swiftCode := c.PostForm("AccBanks[" + strconv.Itoa(bankIndex) + "].SwiftCode")
		bankAddr := c.PostForm("AccBanks[" + strconv.Itoa(bankIndex) + "].BankAddr")
		notes := c.PostForm("AccBanks[" + strconv.Itoa(bankIndex) + "].Notes")

		flag := false
		if isUpload == "true" {
			flag = true
		}

		if accNum == "" || accName == "" || isUpload == "" {
			break
		}
		acct.AcctBanks = append(acct.AcctBanks, models.AcctBank{
			AccName:   accName,
			AccNum:    accNum,
			IsUpload:  flag,
			Currency:  currency,
			BankName:  bankName,
			BankNum:   bankNum,
			SwiftCode: swiftCode,
			BankAddr:  bankAddr,
			Notes:     notes,
		})
		bankIndex++
	}

	now := 0
	allFiles, err := c.MultipartForm()
	attaches := allFiles.File["attach[]"]
	if err == nil {
		for index, attach := range attaches {
			log.Printf("%d index\n", index)
			src, err := attach.Open()
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.Message{
					RetMessage: "error when opening file",
				})
				return
			}
			defer src.Close()

			hash := md5.New()
			if _, err := io.Copy(hash, src); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":   "文件读取失败",
					"details": err.Error(),
				})
				return
			}
			MD5 := hex.EncodeToString(hash.Sum(nil))
			FileName := attach.Filename
			Suffix := filepath.Ext(attach.Filename)
			File := &models.File{
				Name:   FileName,
				MD5:    MD5,
				Suffix: Suffix,
			}
			s.db.CreateFile(File)
			for ; now < bankIndex; now++ {
				if acct.AcctBanks[now].IsUpload {
					acct.AcctBanks[now].FileId = File.FileId
					now++
					break
				}
			}
			if now >= bankIndex && acct.IsUpload {
				acct.FileId = File.FileId
			}

			c.SaveUploadedFile(attach, "../files/"+MD5+Suffix) // 以main程序所在目录为基准

			log.Printf(attach.Filename + "\n")
		}
	}

	log.Printf("%v", acct)
	err = s.db.CreateAcct(acct)
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
