package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许所有来源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // 启用 cookies/auth
	}))
	r.GET("/", s.HelloWorldHandler)

	r.GET("/health", s.healthHandler)

	r.POST("/file", s.Uploader)
	r.POST("/save/acctBank", s.SaveAcctBankHandler)
	r.POST("/save/acct", s.SaveAcctHandler)
	r.POST("/delete/acctBank", s.DeleteAcctBankHandler)
	r.POST("/delete/acct", s.DeleteAcctHandler)
	r.GET("/find/acct", s.FindAcctHandler)
	r.GET("/find/acctBank", s.FindAcctBankHandler)

	r.POST("/create/user", s.CreateUserHandler)
	r.POST("/auth", s.LoginHandler)
	r.POST("/modify/user/password", s.ModifyUserPassHandler)

	// Merchant 相关路由
	r.POST("/save/merchant", s.SaveMerchantHandler)
	r.POST("/delete/merchant", s.DeleteMerchantHandler)
	r.GET("/find/merchant", s.FindMerchantHandler)

	// Cust 相关路由
	r.POST("/save/cust", s.SaveCustHandler)
	r.POST("/delete/cust", s.DeleteCustHandler)
	r.GET("/find/cust", s.FindCustHandler)

	// BankAccount 相关路由
	r.POST("/save/bankAccount", s.SaveBankAccountHandler)
	r.POST("/delete/bankAccount", s.DeleteBankAccountHandler)
	r.GET("/find/bankAccount", s.FindBankAccountHandler)

	deleteGroup := r.Group("/delete")
	{
		deleteGroup.POST("/mercType", s.DeleteMercTypeHandler)
		deleteGroup.POST("/suprType", s.DeleteSuprTypeHandler)
		deleteGroup.POST("/prdtType", s.DeletePrdtTypeHandler)
		deleteGroup.POST("/foodAddType", s.DeleteFoodAddTypeHandler)
		deleteGroup.POST("/feedAddType", s.DeleteFeedAddTypeHandler)
		deleteGroup.POST("/unitMeas", s.DeleteUnitMeasHandler)
		deleteGroup.POST("/packType", s.DeletePackTypeHandler)
		deleteGroup.POST("/conType", s.DeleteConTypeHandler)
		deleteGroup.POST("/currency", s.DeleteCurrencyHandler)
		deleteGroup.POST("/tradeTerm", s.DeleteTradeTermHandler)
		deleteGroup.POST("/nation", s.DeleteNationHandler)
		deleteGroup.POST("/port", s.DeletePortHandler)
		deleteGroup.POST("/taxType", s.DeleteTaxTypeHandler)
		deleteGroup.POST("/brandType", s.DeleteBrandTypeHandler)
		deleteGroup.POST("/degree", s.DeleteDegreeHandler)
		deleteGroup.POST("/dept", s.DeleteDeptHandler)
		deleteGroup.POST("/post", s.DeletePostHandler)
		deleteGroup.POST("/qualStd", s.DeleteQualStdHandler)
		deleteGroup.POST("/invLoc", s.DeleteInvLocHandler)
		deleteGroup.POST("/docReq", s.DeleteDocReqHandler)
		deleteGroup.POST("/payMth", s.DeletePayMthHandler)
		deleteGroup.POST("/payLimit", s.DeletePayLimitHandler)
		deleteGroup.POST("/finaDocStatus", s.DeleteFinaDocStatusHandler)
		deleteGroup.POST("/finaDocType", s.DeleteFinaDocTypeHandler)
		deleteGroup.POST("/expType", s.DeleteExpTypeHandler)
		deleteGroup.POST("/rates", s.DeleteRatesHandler)
		deleteGroup.POST("/bussOrderSta", s.DeleteBussOrderStaHandler)
	}
	findGroup := r.Group("/find")
	{
		findGroup.GET("/mercType", s.FindMercTypeHandler)
		findGroup.GET("/suprType", s.FindSuprTypeHandler)
		findGroup.GET("/prdtType", s.FindPrdtTypeHandler)
		findGroup.GET("/foodAddType", s.FindFoodAddTypeHandler)
		findGroup.GET("/feedAddType", s.FindFeedAddTypeHandler)
		findGroup.GET("/unitMeas", s.FindUnitMeasHandler)
		findGroup.GET("/packType", s.FindPackTypeHandler)
		findGroup.GET("/conType", s.FindConTypeHandler)
		findGroup.GET("/currency", s.FindCurrencyHandler)
		findGroup.GET("/tradeTerm", s.FindTradeTermHandler)
		findGroup.GET("/nation", s.FindNationHandler)
		findGroup.GET("/port", s.FindPortHandler)
		findGroup.GET("/taxType", s.FindTaxTypeHandler)
		findGroup.GET("/brandType", s.FindBrandTypeHandler)
		findGroup.GET("/degree", s.FindDegreeHandler)
		findGroup.GET("/dept", s.FindDeptHandler)
		findGroup.GET("/post", s.FindPostHandler)
		findGroup.GET("/qualStd", s.FindQualStdHandler)
		findGroup.GET("/invLoc", s.FindInvLocHandler)
		findGroup.GET("/docReq", s.FindDocReqHandler)
		findGroup.GET("/payMth", s.FindPayMthHandler)
		findGroup.GET("/payLimit", s.FindPayLimitHandler)
		findGroup.GET("/finaDocStatus", s.FindFinaDocStatusHandler)
		findGroup.GET("/finaDocType", s.FindFinaDocTypeHandler)
		findGroup.GET("/expType", s.FindExpTypeHandler)
		findGroup.GET("/rates", s.FindRatesHandler)
		findGroup.GET("/bussOrderSta", s.FindBussOrderStaHandler)

		// Merchant 查找路由
		merchantGroup := findGroup.Group("/merchant")
		{
			merchantGroup.GET("/mercCode", s.FindMerchantByMercCodeHandler)   // 参数: MercCode
			merchantGroup.GET("/mercAbbr", s.FindMerchantByMercAbbrHandler)   // 参数: MercAbbr
			merchantGroup.GET("/shortMerc", s.FindMerchantByShortMercHandler) // 参数: ShortMerc
			merchantGroup.GET("/merc", s.FindMerchantByMercHandler)           // 参数: Merc
			merchantGroup.GET("/engName", s.FindMerchantByEngNameHandler)     // 参数: EngName
			merchantGroup.GET("/address", s.FindMerchantByAddressHandler)     // 参数: Address
			merchantGroup.GET("/nation", s.FindMerchantByNationHandler)       // 参数: Nation
			merchantGroup.GET("/phoneNum", s.FindMerchantByPhoneNumHandler)   // 参数: PhoneNum
			merchantGroup.GET("/email", s.FindMerchantByEmailHandler)         // 参数: Email
			merchantGroup.GET("/fax", s.FindMerchantByFaxHandler)             // 参数: Fax
			merchantGroup.GET("/website", s.FindMerchantByWebsiteHandler)     // 参数: Website
			merchantGroup.GET("/taxType", s.FindMerchantByTaxTypeHandler)     // 参数: TaxType
			merchantGroup.GET("/taxCode", s.FindMerchantByTaxCodeHandler)     // 参数: TaxCode
			merchantGroup.GET("/mercType", s.FindMerchantByMercTypeHandler)   // 参数: MercType
			merchantGroup.GET("/regCap", s.FindMerchantByRegCapHandler)       // 参数: RegCap
			merchantGroup.GET("/notes", s.FindMerchantByNotesHandler)         // 参数: Notes
		}

		// Cust 查找路由
		custGroup := findGroup.Group("/cust")
		{
			custGroup.GET("/name", s.FindCustByNameHandler)         // 参数: Name
			custGroup.GET("/gender", s.FindCustByGenderHandler)     // 参数: Gender
			custGroup.GET("/nation", s.FindCustByNationHandler)     // 参数: Nation
			custGroup.GET("/addr", s.FindCustByAddrHandler)         // 参数: Addr
			custGroup.GET("/email", s.FindCustByEmailHandler)       // 参数: Email
			custGroup.GET("/phoneNum", s.FindCustByPhoneNumHandler) // 参数: PhoneNum
			custGroup.GET("/qq", s.FindCustByQQHandler)             // 参数: QQ
			custGroup.GET("/wechat", s.FindCustByWechatHandler)     // 参数: Wechat
			custGroup.GET("/merc", s.FindCustByMercHandler)         // 参数: Merc
			custGroup.GET("/post", s.FindCustByPostHandler)         // 参数: Post
			custGroup.GET("/notes", s.FindCustByNotesHandler)       // 参数: Notes
		}

		// BankAccount 查找路由
		bankAccountGroup := findGroup.Group("/bankAccount")
		{
			bankAccountGroup.GET("/bankAccName", s.FindBankAccountByBankAccNameHandler) // 参数: BankAccName
			bankAccountGroup.GET("/compName", s.FindBankAccountByCompNameHandler)       // 参数: CompName
			bankAccountGroup.GET("/currency", s.FindBankAccountByCurrencyHandler)       // 参数: Currency
			bankAccountGroup.GET("/bankName", s.FindBankAccountByBankNameHandler)       // 参数: BankName
			bankAccountGroup.GET("/bankEng", s.FindBankAccountByBankEngHandler)         // 参数: BankEng
			bankAccountGroup.GET("/bankAddr", s.FindBankAccountByBankAddrHandler)       // 参数: BankAddr
			bankAccountGroup.GET("/notes", s.FindBankAccountByNotesHandler)             // 参数: Notes
			bankAccountGroup.GET("/acctNum", s.FindBankAccountByAcctNumHandler)         // 参数: AcctNum
			bankAccountGroup.GET("/bankNum", s.FindBankAccountByBankNumHandler)         // 参数: BankNum
			bankAccountGroup.GET("/swiftCode", s.FindBankAccountBySwiftCodeHandler)     // 参数: SwiftCode
		}
	}

	// 创建“acct”子分组，用于账户相关的查找
	acctGroup := findGroup.Group("/acct")
	{
		// 精确查找接口
		acctGroup.GET("/acctCode", s.FindAcctByAcctCodeHandler) // 参数: AcctCode (账户代码)
		acctGroup.GET("/acctAbbr", s.FindAcctByAcctAbbrHandler) // 参数: AcctAbbr (账户缩写)
		acctGroup.GET("/EtyAbbr", s.FindAcctByEtyAbbrHandler)   // 参数: EtyAbbr (实体缩写)
		acctGroup.GET("/acctName", s.FindAcctByAcctNameHandler) // 参数: AcctName (账户名称)
		acctGroup.GET("/acctAddr", s.FindAcctByAcctAddrHandler) // 参数: AcctAddr (账户地址)
		acctGroup.GET("/nation", s.FindAcctByNationHandler)     // 参数: Nation (国家)
		acctGroup.GET("/taxType", s.FindAcctByTaxTypeHandler)   // 参数: TaxType (税务类型)
		acctGroup.GET("/taxCode", s.FindAcctByTaxCodeHandler)   // 参数: TaxCode (税务代码)
		acctGroup.GET("/phoneNum", s.FindAcctByPhoneNumHandler) // 参数: PhoneNum (电话号码)
		acctGroup.GET("/email", s.FindAcctByEmailHandler)       // 参数: Email (电子邮件)
		acctGroup.GET("/website", s.FindAcctByWebsiteHandler)   // 参数: Website (网站)
		acctGroup.GET("/regDate", s.FindAcctByRegDateHandler)   // 参数: RegDate (注册日期)
		acctGroup.GET("/notes", s.FindAcctByNotesHandler)       // 参数: Notes (备注)

		// 模糊查找接口
		acctGroup.GET("/fzz/acctAddr", s.FzzFindAcctByAcctAddrHandler) // 参数: AcctAddr (账户地址)
		acctGroup.GET("/fzz/email", s.FzzFindAcctByEmailHandler)       // 参数: Email (电子邮件)
		acctGroup.GET("/fzz/website", s.FzzFindAcctByWebsiteHandler)   // 参数: Website (网站)
		acctGroup.GET("/fzz/regDate", s.FzzFindAcctByRegDateHandler)   // 参数: RegDate (注册日期)
		acctGroup.GET("/fzz/notes", s.FzzFindAcctByNotesHandler)       // 参数: Notes (备注)
	}

	// 创建“acctBank”子分组，用于账户银行相关的查找
	acctBankGroup := findGroup.Group("/acctBank")
	{
		// 精确查找接口
		acctBankGroup.GET("/id", s.FindAcctBankByIdHandler)               // 参数: AcctId (账户ID)
		acctBankGroup.GET("/accName", s.FindAcctBankByAccNameHandler)     // 参数: AccName (账户名称)
		acctBankGroup.GET("/accNum", s.FindAcctBankByAccNumHandler)       // 参数: AccNum (账户编号)
		acctBankGroup.GET("/currency", s.FindAcctBankByCurrencyHandler)   // 参数: Currency (货币)
		acctBankGroup.GET("/bankName", s.FindAcctBankByBankNameHandler)   // 参数: BankName (银行名称)
		acctBankGroup.GET("/swiftCode", s.FindAcctBankBySwiftCodeHandler) // 参数: SwiftCode (SWIFT代码)
		acctBankGroup.GET("/bankAddr", s.FindAcctBankByBankAddrHandler)   // 参数: BankAddr (银行地址)
		acctBankGroup.GET("/notes", s.FindAcctBankByNotesHandler)         // 参数: Notes (备注)

		// 模糊查找接口
		acctBankGroup.GET("/fzz/bankName", s.FzzFindAcctBankByBankNameHandler) // 参数: BankName (银行名称)
		acctBankGroup.GET("/fzz/bankAddr", s.FzzFindAcctBankByBankAddrHandler) // 参数: BankAddr (银行地址)
		acctBankGroup.GET("/fzz/notes", s.FzzFindAcctBankByNotesHandler)       // 参数: Notes (备注)
	}

	saveGroup := r.Group("/save")
	{
		saveGroup.POST("/mercType", s.SaveMercTypeHandler)           // 参数: mercType，mercTypeId
		saveGroup.POST("/suprType", s.SaveSuprTypeHandler)           // 参数: suprType，suprTypeId
		saveGroup.POST("/prdtType", s.SavePrdtTypeHandler)           // 参数: prdtType，prdtTypeId
		saveGroup.POST("/foodAddType", s.SaveFoodAddTypeHandler)     // 参数: foodAddType，foodAddTypeId
		saveGroup.POST("/feedAddType", s.SaveFeedAddTypeHandler)     // 参数: feedAddType，feedAddTypeId
		saveGroup.POST("/unitMeas", s.SaveUnitMeasHandler)           // 参数: unitMeas，unitMeasId
		saveGroup.POST("/packType", s.SavePackTypeHandler)           // 参数: packType，packTypeId
		saveGroup.POST("/conType", s.SaveConTypeHandler)             // 参数: conType，conTypeId
		saveGroup.POST("/currency", s.SaveCurrencyHandler)           // 参数: currency，currencyId
		saveGroup.POST("/tradeTerm", s.SaveTradeTermHandler)         // 参数: tradeTerm，tradeTermId
		saveGroup.POST("/nation", s.SaveNationHandler)               // 参数: nation，nationId
		saveGroup.POST("/port", s.SavePortHandler)                   // 参数: port，portId
		saveGroup.POST("/taxType", s.SaveTaxTypeHandler)             // 参数: taxType，taxTypeId
		saveGroup.POST("/brandType", s.SaveBrandTypeHandler)         // 参数: brandType，brandTypeId
		saveGroup.POST("/degree", s.SaveDegreeHandler)               // 参数: degree，degreeId
		saveGroup.POST("/dept", s.SaveDeptHandler)                   // 参数: dept，deptId
		saveGroup.POST("/post", s.SavePostHandler)                   // 参数: post，postId
		saveGroup.POST("/qualStd", s.SaveQualStdHandler)             // 参数: qualStd，qualStdId
		saveGroup.POST("/invLoc", s.SaveInvLocHandler)               // 参数: invLoc，invLocId
		saveGroup.POST("/docReq", s.SaveDocReqHandler)               // 参数: docReq，docReqId
		saveGroup.POST("/payMth", s.SavePayMthHandler)               // 参数: payMth，payMthId
		saveGroup.POST("/payLimit", s.SavePayLimitHandler)           // 参数: payLimit，payLimitId
		saveGroup.POST("/finaDocStatus", s.SaveFinaDocStatusHandler) // 参数: finaDocStatus，finaDocStatusId
		saveGroup.POST("/finaDocType", s.SaveFinaDocTypeHandler)     // 参数: finaDocType，finaDocTypeId
		saveGroup.POST("/expType", s.SaveExpTypeHandler)             // 参数: expType，expTypeId
		saveGroup.POST("/rates", s.SaveRatesHandler)                 // 参数: rates，ratesId
		saveGroup.POST("/bussOrderSta", s.SaveBussOrderStaHandler)   // 参数: bussOrderSta，bussOrderStaId
	}
	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
