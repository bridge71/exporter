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

	r.GET("/add/dicdata", s.AddDicData)

	r.GET("/find/shouldIn", s.FindShouldInHandler)
	r.POST("/find/shouldIn/send", s.FindShouldInSendHandler)
	r.POST("/find/shouldIn/in", s.FindShouldInInHandler)
	r.POST("/find/shouldIn/sale", s.FindShouldInSaleHandler)
	r.POST("/add/shouldIn/in", s.AddShouldInIns)
	r.POST("/add/shouldIn/sale", s.AddShouldInSale)
	r.POST("/add/shouldIn/send", s.AddShouldInSends)
	r.POST("/delete/shouldIn/in", s.DeleteShouldInIn)
	r.POST("/delete/shouldIn/sale", s.DeleteShouldInSale)
	r.POST("/delete/shouldIn/send", s.DeleteShouldInSend)
	r.POST("/delete/shouldIn", s.DeleteShouldInHandler)
	r.POST("/save/shouldIn", s.SaveShouldInHandler)

	r.GET("/find/in", s.FindInHandler)
	r.POST("/find/in/send", s.FindInSendHandler)
	r.POST("/find/in/shouldIn", s.FindInShouldInHandler)
	r.POST("/find/in/sale", s.FindInSaleHandler)
	r.POST("/add/in/shouldIn", s.AddInShouldIns)
	r.POST("/add/in/sale", s.AddInSale)
	r.POST("/add/in/send", s.AddInSends)
	r.POST("/delete/in/shouldIn", s.DeleteInShouldIn)
	r.POST("/delete/in/sale", s.DeleteInSale)
	r.POST("/delete/in/send", s.DeleteInSend)
	r.POST("/delete/in", s.DeleteInHandler)
	r.POST("/save/in", s.SaveInHandler)

	r.GET("/find/send", s.FindSendHandler)
	r.POST("/find/send/prdtInfo", s.FindSendPrdtInfoHandler)
	r.POST("/find/send/sale", s.FindSendSaleHandler)
	r.POST("/find/send/in", s.FindSendInHandler)
	r.POST("/find/send/shouldIn", s.FindSendShouldInHandler)
	r.POST("/find/send/loadingInfo", s.FindSendLoadingInfoHandler)
	r.POST("/add/send/prdtInfo", s.AddSendPrdtInfo)
	r.POST("/add/send/in", s.AddSendIns)
	r.POST("/add/send/shouldIn", s.AddSendShouldIns)
	r.POST("/add/send/loadingInfo", s.AddSendLoadingInfo)
	r.POST("/add/send/sale", s.AddSendSale)
	r.POST("/delete/send/prdtInfo", s.DeleteSendPrdtInfo)
	r.POST("/delete/send/in", s.DeleteSendIn)
	r.POST("/delete/send/shouldIn", s.DeleteSendShouldIn)
	r.POST("/delete/send/sale", s.DeleteSendSale)
	r.POST("/delete/send", s.DeleteSendHandler)
	r.POST("/save/send", s.SaveSendHandler)

	r.GET("/find/sale", s.FindSaleHandler)
	r.POST("/find/sale/prdtInfo", s.FindSalePrdtInfoHandler)
	r.POST("/find/sale/send", s.FindSaleSendHandler)
	r.POST("/find/sale/shouldIn", s.FindSaleShouldInHandler)
	r.POST("/find/sale/in", s.FindSaleInHandler)
	r.POST("/add/sale/prdtInfo", s.AddSalePrdtInfo)
	r.POST("/add/sale/shouldIn", s.AddSaleShouldIn)
	r.POST("/add/sale/in", s.AddSaleIn)
	r.POST("/add/sale/send", s.AddSaleSend)
	r.POST("/delete/sale/prdtInfo", s.DeleteSalePrdtInfo)
	r.POST("/delete/sale/shouldIn", s.DeleteSaleShouldIn)
	r.POST("/delete/sale/in", s.DeleteSaleIn)
	r.POST("/delete/sale/send", s.DeleteSaleSend)
	r.POST("/delete/sale", s.DeleteSaleHandler)
	r.POST("/save/sale", s.SaveSaleHandler)

	r.GET("/find/spot", s.FindSpotHandler)
	r.POST("/delete/spot", s.DeleteSpotHandler)
	r.POST("/save/spot", s.SaveSpotHandler)

	r.GET("/find/costInfo", s.FindCostInfoHandler)
	r.POST("/delete/costInfo", s.DeleteCostInfoHandler)
	r.POST("/save/costInfo", s.SaveCostInfoHandler)

	r.GET("/find/loadingInfo", s.FindLoadingInfoHandler)
	r.POST("/delete/loadingInfo", s.DeleteLoadingInfoHandler)
	r.POST("/save/loadingInfo", s.SaveLoadingInfoHandler)

	r.GET("/find/prdtInfo", s.FindPrdtInfoHandler)
	r.POST("/delete/prdtInfo", s.DeletePrdtInfoHandler)
	r.POST("/save/prdtInfo", s.SavePrdtInfoHandler)

	r.GET("/find/empl", s.FindEmplHandler)
	r.POST("/delete/empl", s.DeleteEmplHandler)
	r.POST("/save/empl", s.SaveEmplHandler)

	r.GET("/find/packSpec", s.FindPackSpecHandler)
	r.POST("/delete/packSpec", s.DeletePackSpecHandler)
	r.POST("/save/packSpec", s.SavePackSpecHandler)

	r.GET("/find/brand", s.FindBrandHandler)
	r.POST("/delete/brand", s.DeleteBrandHandler)
	r.POST("/find/brand/id", s.FindBrandByIdHandler)
	r.POST("/save/brand", s.SaveBrandHandler)

	r.GET("/find/cat", s.FindCatHandler)
	r.POST("/delete/cat", s.DeleteCatHandler)
	r.POST("/find/cat/id", s.FindCatByIdHandler)
	r.POST("/save/cat", s.SaveCatHandler)

	r.POST("/save/payMentMethod", s.SavePayMentMethodHandler)
	r.POST("/delete/payMentMethod", s.DeletePayMentMethodHandler)
	r.GET("/find/payMentMethod", s.FindPayMentMethodHandler)

	r.POST("/save/merchant", s.SaveMerchantHandler)
	r.POST("/delete/merchant", s.DeleteMerchantHandler)
	r.GET("/find/merchant", s.FindMerchantHandler)

	r.POST("/save/acct", s.SaveAcctHandler)
	r.POST("/delete/acct", s.DeleteAcctHandler)
	r.GET("/find/acct", s.FindAcctHandler)

	r.POST("/save/acctBank", s.SaveAcctBankHandler)
	r.POST("/delete/acctBank", s.DeleteAcctBankHandler)
	r.GET("/find/acctBank", s.FindAcctBankHandler)
	r.GET("/find/acctBank/id", s.FindAcctBankByIdHandler)

	r.POST("/save/cust", s.SaveCustHandler)
	r.POST("/delete/cust", s.DeleteCustHandler)
	r.GET("/find/cust", s.FindCustHandler)
	r.GET("/find/cust/id", s.FindCustByIdHandler)

	r.POST("/save/bankAccount", s.SaveBankAccountHandler)
	r.POST("/delete/bankAccount", s.DeleteBankAccountHandler)
	r.GET("/find/bankAccount", s.FindBankAccountHandler)
	r.GET("/find/bankAccount/id", s.FindBankAccountByIdHandler)

	r.POST("/save/user", s.SaveUserHandler)
	r.POST("/auth", s.LoginHandler)
	r.GET("/qut/root", s.CreateRootHandler)
	r.GET("/find/user", s.FindUserHandler)

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
		deleteGroup.POST("/eduLevel", s.DeleteEduLevelHandler)
		deleteGroup.POST("/dept", s.DeleteDeptHandler)
		deleteGroup.POST("/position", s.DeletePositionHandler)
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
		findGroup.GET("/eduLevel", s.FindEduLevelHandler)
		findGroup.GET("/dept", s.FindDeptHandler)
		findGroup.GET("/position", s.FindPositionHandler)
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

	}

	saveGroup := r.Group("/save")
	{
		saveGroup.POST("/mercType", s.SaveMercTypeHandler)
		saveGroup.POST("/suprType", s.SaveSuprTypeHandler)
		saveGroup.POST("/prdtType", s.SavePrdtTypeHandler)
		saveGroup.POST("/foodAddType", s.SaveFoodAddTypeHandler)
		saveGroup.POST("/feedAddType", s.SaveFeedAddTypeHandler)
		saveGroup.POST("/unitMeas", s.SaveUnitMeasHandler)
		saveGroup.POST("/packType", s.SavePackTypeHandler)
		saveGroup.POST("/conType", s.SaveConTypeHandler)
		saveGroup.POST("/currency", s.SaveCurrencyHandler)
		saveGroup.POST("/tradeTerm", s.SaveTradeTermHandler)
		saveGroup.POST("/nation", s.SaveNationHandler)
		saveGroup.POST("/port", s.SavePortHandler)
		saveGroup.POST("/taxType", s.SaveTaxTypeHandler)
		saveGroup.POST("/brandType", s.SaveBrandTypeHandler)
		saveGroup.POST("/eduLevel", s.SaveEduLevelHandler)
		saveGroup.POST("/dept", s.SaveDeptHandler)
		saveGroup.POST("/position", s.SavePositionHandler)
		saveGroup.POST("/qualStd", s.SaveQualStdHandler)
		saveGroup.POST("/invLoc", s.SaveInvLocHandler)
		saveGroup.POST("/docReq", s.SaveDocReqHandler)
		saveGroup.POST("/payMth", s.SavePayMthHandler)
		saveGroup.POST("/payLimit", s.SavePayLimitHandler)
		saveGroup.POST("/finaDocStatus", s.SaveFinaDocStatusHandler)
		saveGroup.POST("/finaDocType", s.SaveFinaDocTypeHandler)
		saveGroup.POST("/expType", s.SaveExpTypeHandler)
		saveGroup.POST("/rates", s.SaveRatesHandler)
		saveGroup.POST("/bussOrderSta", s.SaveBussOrderStaHandler)
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
