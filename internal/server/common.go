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

func (s *Server) AddDicData(c *gin.Context) {
	MercType := models.MercType{MercType: "6"}
	_ = s.db.SaveMercType(&MercType)

	SuprType := models.SuprType{SuprType: "SupplierType1"}
	_ = s.db.SaveSuprType(&SuprType)

	PrdtType := models.PrdtType{PrdtType: "ProductType1"}
	_ = s.db.SavePrdtType(&PrdtType)

	FoodAddType := models.FoodAddType{FoodAddType: "FoodAdditiveType1"}
	_ = s.db.SaveFoodAddType(&FoodAddType)

	FeedAddType := models.FeedAddType{FeedAddType: "FeedAdditiveType1"}
	_ = s.db.SaveFeedAddType(&FeedAddType)

	UnitMeas := models.UnitMeas{UnitMeas: "Unit1"}
	_ = s.db.SaveUnitMeas(&UnitMeas)

	PackType := models.PackType{PackType: "PackType1"}
	_ = s.db.SavePackType(&PackType)

	ConType := models.ConType{ConType: "ContainerType1"}
	_ = s.db.SaveConType(&ConType)

	Currency := models.Currency{Currency: "USD"}
	_ = s.db.SaveCurrency(&Currency)

	TradeTerm := models.TradeTerm{TradeTerm: "FOB"}
	_ = s.db.SaveTradeTerm(&TradeTerm)

	Nation := models.Nation{Nation: "China"}
	_ = s.db.SaveNation(&Nation)

	Port := models.Port{Port: "Port1"}
	_ = s.db.SavePort(&Port)

	TaxType := models.TaxType{TaxType: "TaxType1"}
	_ = s.db.SaveTaxType(&TaxType)

	BrandType := models.BrandType{BrandType: "BrandType1"}
	_ = s.db.SaveBrandType(&BrandType)

	EduLevel := models.EduLevel{EduLevel: "Bachelor"}
	_ = s.db.SaveEduLevel(&EduLevel)

	Dept := models.Dept{Dept: "Department1"}
	_ = s.db.SaveDept(&Dept)

	Position := models.Position{Position: "Position1"}
	_ = s.db.SavePosition(&Position)

	QualStd := models.QualStd{QualStd: "Standard1"}
	_ = s.db.SaveQualStd(&QualStd)

	InvLoc := models.InvLoc{InvLoc: "Location1"}
	_ = s.db.SaveInvLoc(&InvLoc)

	DocReq := models.DocReq{DocReq: "Document1"}
	_ = s.db.SaveDocReq(&DocReq)

	PayMth := models.PayMth{PayMth: "Method1"}
	_ = s.db.SavePayMth(&PayMth)

	PayLimit := models.PayLimit{PayLimit: "Limit1"}
	_ = s.db.SavePayLimit(&PayLimit)

	FinaDocStatus := models.FinaDocStatus{FinaDocStatus: "Status1"}
	_ = s.db.SaveFinaDocStatus(&FinaDocStatus)

	FinaDocType := models.FinaDocType{FinaDocType: "Type1"}
	_ = s.db.SaveFinaDocType(&FinaDocType)

	ExpType := models.ExpType{ExpType: "ExpenseType1"}
	_ = s.db.SaveExpType(&ExpType)

	Rates := models.Rates{Rates: "Rate1"}
	_ = s.db.SaveRates(&Rates)

	BussOrderSta := models.BussOrderSta{BussOrderSta: "OrderStatus1"}
	_ = s.db.SaveBussOrderSta(&BussOrderSta)

	Acct := models.Acct{AcctCode: "nihao", AcctAbbr: "hao", EtyAbbr: "room", AcctName: "hh"}
	_ = s.db.SaveAcct(&Acct)

	AcctBank := models.AcctBank{AcctId: 1, AccName: "sss", AccNum: "3424"}
	_ = s.db.Save(&AcctBank)

	Merchant := models.Merchant{MercCode: "ldaf", MercAbbr: "weq", ShortMerc: "erer"}
	_ = s.db.SaveMerchant(&Merchant)

	BankAccount := models.BankAccount{BankAccName: "sss", CompName: "dfds", AcctNum: "dfa", BankName: "fff", MerchantId: 1}
	_ = s.db.Save(&BankAccount)

	PayMent := models.PayMentMethod{PayMtdName: "dfads"}
	_ = s.db.Save(&PayMent)

	PackSpec := models.PackSpec{SpecName: "fsd"}
	_ = s.db.Save(&PackSpec)
}

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
		return nil, File.ID, FileName
	}
	return nil, 0, ""
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
