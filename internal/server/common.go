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
	MercType1 := models.MercType{MercType: "国外终端"}
	_ = s.db.SaveMercType(&MercType1)

	MercType2 := models.MercType{MercType: "国外贸易商"}
	_ = s.db.SaveMercType(&MercType2)

	MercType3 := models.MercType{MercType: "国内贸易商"}
	_ = s.db.SaveMercType(&MercType3)

	SuprType1 := models.SuprType{SuprType: "国内生产工厂"}
	_ = s.db.SaveSuprType(&SuprType1)

	SuprType2 := models.SuprType{SuprType: "国内生产工厂销售公司"}
	_ = s.db.SaveSuprType(&SuprType2)

	SuprType3 := models.SuprType{SuprType: "国内贸易商"}
	_ = s.db.SaveSuprType(&SuprType3)

	PrdtType1 := models.PrdtType{PrdtType: "食品添加剂"}
	_ = s.db.SavePrdtType(&PrdtType1)

	PrdtType2 := models.PrdtType{PrdtType: "饲料添加剂"}
	_ = s.db.SavePrdtType(&PrdtType2)

	PrdtType3 := models.PrdtType{PrdtType: "油田化学品"}
	_ = s.db.SavePrdtType(&PrdtType3)

	FoodAddType1 := models.FoodAddType{FoodAddType: "甜味剂"}
	_ = s.db.SaveFoodAddType(&FoodAddType1)

	FoodAddType2 := models.FoodAddType{FoodAddType: "酸味剂"}
	_ = s.db.SaveFoodAddType(&FoodAddType2)

	FoodAddType3 := models.FoodAddType{FoodAddType: "防腐剂"}
	_ = s.db.SaveFoodAddType(&FoodAddType3)

	FeedAddType1 := models.FeedAddType{FeedAddType: "氨基酸"}
	_ = s.db.SaveFeedAddType(&FeedAddType1)

	FeedAddType2 := models.FeedAddType{FeedAddType: "维生素"}
	_ = s.db.SaveFeedAddType(&FeedAddType2)

	FeedAddType3 := models.FeedAddType{FeedAddType: "兽药类 "}
	_ = s.db.SaveFeedAddType(&FeedAddType3)

	UnitMeas1 := models.UnitMeas{UnitMeas: "KGS"}
	_ = s.db.SaveUnitMeas(&UnitMeas1)

	UnitMeas2 := models.UnitMeas{UnitMeas: "MT"}
	_ = s.db.SaveUnitMeas(&UnitMeas2)

	UnitMeas3 := models.UnitMeas{UnitMeas: "GRAMS"}
	_ = s.db.SaveUnitMeas(&UnitMeas3)

	PackType1 := models.PackType{PackType: "塑编袋"}
	_ = s.db.SavePackType(&PackType1)

	PackType2 := models.PackType{PackType: "牛皮纸袋"}
	_ = s.db.SavePackType(&PackType2)

	PackType3 := models.PackType{PackType: "铝箔袋"}
	_ = s.db.SavePackType(&PackType3)

	ConType1 := models.ConType{ConType: "20GP-20尺普柜"}
	_ = s.db.SaveConType(&ConType1)

	ConType2 := models.ConType{ConType: "40HQ-40尺高柜"}
	_ = s.db.SaveConType(&ConType2)

	ConType3 := models.ConType{ConType: "40GP-40尺普柜"}
	_ = s.db.SaveConType(&ConType3)

	Currency1 := models.Currency{Currency: "USD"}
	_ = s.db.SaveCurrency(&Currency1)

	Currency2 := models.Currency{Currency: "CNY"}
	_ = s.db.SaveCurrency(&Currency2)

	Currency3 := models.Currency{Currency: "EURO"}
	_ = s.db.SaveCurrency(&Currency3)

	TradeTerm1 := models.TradeTerm{TradeTerm: "FOB"}
	_ = s.db.SaveTradeTerm(&TradeTerm1)

	TradeTerm2 := models.TradeTerm{TradeTerm: "CIF"}
	_ = s.db.SaveTradeTerm(&TradeTerm2)

	TradeTerm3 := models.TradeTerm{TradeTerm: "CNF"}
	_ = s.db.SaveTradeTerm(&TradeTerm3)

	Nation1 := models.Nation{Nation: "CHINA-中国"}
	_ = s.db.SaveNation(&Nation1)

	Nation2 := models.Nation{Nation: "SOUTH KOREA-韩国"}
	_ = s.db.SaveNation(&Nation2)

	Nation3 := models.Nation{Nation: "VIETNAM-越南"}
	_ = s.db.SaveNation(&Nation3)

	Port1 := models.Port{Port: "DALIAN, CHINA"}
	_ = s.db.SavePort(&Port1)

	Port2 := models.Port{Port: "LIANYUNGANG, CHINA"}
	_ = s.db.SavePort(&Port2)

	Port3 := models.Port{Port: "QINZHOU, CHINA"}
	_ = s.db.SavePort(&Port3)

	TaxType1 := models.TaxType{TaxType: "USCI"}
	_ = s.db.SaveTaxType(&TaxType1)

	TaxType2 := models.TaxType{TaxType: "ORGANIZATION CODE"}
	_ = s.db.SaveTaxType(&TaxType2)

	TaxType3 := models.TaxType{TaxType: "TRADE REGISTER NUMBER"}
	_ = s.db.SaveTaxType(&TaxType3)

	BrandType1 := models.BrandType{BrandType: "境内自主品牌"}
	_ = s.db.SaveBrandType(&BrandType1)

	BrandType2 := models.BrandType{BrandType: "境外品牌"}
	_ = s.db.SaveBrandType(&BrandType2)

	BrandType3 := models.BrandType{BrandType: "无品牌"}
	_ = s.db.SaveBrandType(&BrandType3)

	EduLevel1 := models.EduLevel{EduLevel: "中专/高中"}
	_ = s.db.SaveEduLevel(&EduLevel1)

	EduLevel2 := models.EduLevel{EduLevel: "专科"}
	_ = s.db.SaveEduLevel(&EduLevel2)

	EduLevel3 := models.EduLevel{EduLevel: "本科"}
	_ = s.db.SaveEduLevel(&EduLevel3)

	Dept1 := models.Dept{Dept: "销售部"}
	_ = s.db.SaveDept(&Dept1)

	Dept2 := models.Dept{Dept: "采购部"}
	_ = s.db.SaveDept(&Dept2)

	Dept3 := models.Dept{Dept: "运营部"}
	_ = s.db.SaveDept(&Dept3)

	Position1 := models.Position{Position: "总经理"}
	_ = s.db.SavePosition(&Position1)

	Position2 := models.Position{Position: "销售经理"}
	_ = s.db.SavePosition(&Position2)

	Position3 := models.Position{Position: "销售总监"}
	_ = s.db.SavePosition(&Position3)

	QualStd1 := models.QualStd{QualStd: "国标 GB Standard"}
	_ = s.db.SaveQualStd(&QualStd1)

	QualStd2 := models.QualStd{QualStd: "企标 Corporate Standards"}
	_ = s.db.SaveQualStd(&QualStd2)

	QualStd3 := models.QualStd{QualStd: "无 None"}
	_ = s.db.SaveQualStd(&QualStd3)

	InvLoc1 := models.InvLoc{InvLoc: "FACTORY-工厂"}
	_ = s.db.SaveInvLoc(&InvLoc1)

	InvLoc2 := models.InvLoc{InvLoc: "QINGDAO, CHINA-中国青岛港"}
	_ = s.db.SaveInvLoc(&InvLoc2)

	InvLoc3 := models.InvLoc{InvLoc: "DALIAN, CHINA-中国大连港"}
	_ = s.db.SaveInvLoc(&InvLoc3)

	DocReq1 := models.DocReq{DocReq: "Bill of Lading"}
	_ = s.db.SaveDocReq(&DocReq1)

	DocReq2 := models.DocReq{DocReq: "Commercial Invoice"}
	_ = s.db.SaveDocReq(&DocReq2)

	DocReq3 := models.DocReq{DocReq: "Packing List"}
	_ = s.db.SaveDocReq(&DocReq3)

	PayMth1 := models.PayMth{PayMth: "T/T"}
	_ = s.db.SavePayMth(&PayMth1)

	PayMth2 := models.PayMth{PayMth: "D/P"}
	_ = s.db.SavePayMth(&PayMth2)

	PayMth3 := models.PayMth{PayMth: "L/C"}
	_ = s.db.SavePayMth(&PayMth3)

	PayLimit1 := models.PayLimit{PayLimit: "7 Days Upon B/L Date"}
	_ = s.db.SavePayLimit(&PayLimit1)

	PayLimit2 := models.PayLimit{PayLimit: "10 Days Upon B/L Date"}
	_ = s.db.SavePayLimit(&PayLimit2)

	PayLimit3 := models.PayLimit{PayLimit: "20 Days Upon B/L Date"}
	_ = s.db.SavePayLimit(&PayLimit3)

	FinaDocStatus1 := models.FinaDocStatus{FinaDocStatus: "未支付"}
	_ = s.db.SaveFinaDocStatus(&FinaDocStatus1)

	FinaDocStatus2 := models.FinaDocStatus{FinaDocStatus: "已部分支付"}
	_ = s.db.SaveFinaDocStatus(&FinaDocStatus2)

	FinaDocStatus3 := models.FinaDocStatus{FinaDocStatus: "已全部支付"}
	_ = s.db.SaveFinaDocStatus(&FinaDocStatus3)

	FinaDocType1 := models.FinaDocType{FinaDocType: "应收货款"}
	_ = s.db.SaveFinaDocType(&FinaDocType1)

	FinaDocType2 := models.FinaDocType{FinaDocType: "应收佣金"}
	_ = s.db.SaveFinaDocType(&FinaDocType2)

	FinaDocType3 := models.FinaDocType{FinaDocType: "应收费用"}
	_ = s.db.SaveFinaDocType(&FinaDocType3)

	ExpType1 := models.ExpType{ExpType: "海运费"}
	_ = s.db.SaveExpType(&ExpType1)

	ExpType2 := models.ExpType{ExpType: "ENS/ICS2费用"}
	_ = s.db.SaveExpType(&ExpType2)

	ExpType3 := models.ExpType{ExpType: "CargoX费用"}
	_ = s.db.SaveExpType(&ExpType3)

	Rates1 := models.Rates{Rates: "/20GP-per 20GP"}
	_ = s.db.SaveRates(&Rates1)

	Rates2 := models.Rates{Rates: "/40HC-per 40HC"}
	_ = s.db.SaveRates(&Rates2)

	Rates3 := models.Rates{Rates: "/票-per Bill"}
	_ = s.db.SaveRates(&Rates3)

	BussOrderSta1 := models.BussOrderSta{BussOrderSta: "已确认"}
	_ = s.db.SaveBussOrderSta(&BussOrderSta1)

	BussOrderSta2 := models.BussOrderSta{BussOrderSta: "已收款未发货"}
	_ = s.db.SaveBussOrderSta(&BussOrderSta2)

	BussOrderSta3 := models.BussOrderSta{BussOrderSta: "已发货未收款"}
	_ = s.db.SaveBussOrderSta(&BussOrderSta3)

	Acct := models.Acct{AcctCode: "ASDFG", AcctAbbr: "AS", EtyAbbr: "简称1", AcctName: "实体1"}
	_ = s.db.SaveAcct(&Acct)

	Acct1 := models.Acct{AcctCode: "DFG", AcctAbbr: "DG", EtyAbbr: "简称2", AcctName: "实体2"}
	_ = s.db.SaveAcct(&Acct1)

	Acct2 := models.Acct{AcctCode: "FG", AcctAbbr: "G", EtyAbbr: "简称3", AcctName: "实体3"}
	_ = s.db.SaveAcct(&Acct2)

	AcctBank := models.AcctBank{AcctID: 1, AccName: "11", AccNum: "3424"}
	_ = s.db.Save(&AcctBank)

	AcctBank1 := models.AcctBank{AcctID: 2, AccName: "22", AccNum: "343424"}
	_ = s.db.Save(&AcctBank1)

	AcctBank2 := models.AcctBank{AcctID: 3, AccName: "33", AccNum: "354424"}
	_ = s.db.Save(&AcctBank2)

	AcctBank3 := models.AcctBank{AcctID: 1, AccName: "44", AccNum: "3490824"}
	_ = s.db.Save(&AcctBank3)

	AcctBank4 := models.AcctBank{AcctID: 1, AccName: "55", AccNum: "3321332"}
	_ = s.db.Save(&AcctBank4)

	AcctBank5 := models.AcctBank{AcctID: 3, AccName: "66", AccNum: "7983424"}
	_ = s.db.Save(&AcctBank5)

	Merchant := models.Merchant{MercCode: "ldaf", ShortMerc: "erer"}
	_ = s.db.SaveMerchant(&Merchant)

	Merchant1 := models.Merchant{MercCode: "dfsde", ShortMerc: "dfs"}
	_ = s.db.SaveMerchant(&Merchant1)

	BankAccount := models.BankAccount{BankAccName: "yourAccount", CompName: "sof", AcctNum: "7896", BankName: "sf", MerchantID: 1}
	_ = s.db.Save(&BankAccount)

	BankAccount1 := models.BankAccount{BankAccName: "myAccount", CompName: "ersf", AcctNum: "j77896", BankName: "s9ujf", MerchantID: 1}
	_ = s.db.Save(&BankAccount1)

	BankAccount2 := models.BankAccount{BankAccName: "ourAccount", CompName: "hhersf", AcctNum: "92386", BankName: "jdjd", MerchantID: 2}
	_ = s.db.Save(&BankAccount2)

	Spot := models.Spot{InvLocName: "QINGDAO"}
	_ = s.db.Save(&Spot)
	Spot1 := models.Spot{InvLocName: "DALIAN"}
	_ = s.db.Save(&Spot1)
	Spot2 := models.Spot{InvLocName: "SHENZHEN"}
	_ = s.db.Save(&Spot2)

	PayMent := models.PayMentMethod{PayMtdName: "10T7O"}
	_ = s.db.Save(&PayMent)

	PayMent1 := models.PayMentMethod{PayMtdName: "30S5O"}
	_ = s.db.Save(&PayMent1)

	PackSpec := models.PackSpec{SpecName: "nokoa"}
	_ = s.db.Save(&PackSpec)

	PackSpec1 := models.PackSpec{SpecName: "tfska"}
	_ = s.db.Save(&PackSpec1)
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

func (s *Server) SaveFile(c *gin.Context, name string) (error, uint, string) {
	file, err := c.FormFile(name)
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
