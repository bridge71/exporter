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

	r.GET("/find/costInfo", s.FindCostInfoHandler)
	r.POST("/delete/costInfo", s.DeleteCostInfoHandler)
	r.POST("/save/costInfo", s.SaveCostInfoHandler)

	r.GET("/find/loadingInfo", s.FindLoadingInfoHandler)
	r.POST("/delete/loadingInfo", s.DeleteLoadingInfoHandler)
	r.POST("/save/loadingInfo", s.SaveLoadingInfoHandler)

	r.GET("/find/prdtInfo", s.FindPrdtInfoHandler)
	r.POST("/delete/prdtInfo", s.DeletePrdtInfoHandler)
	r.POST("/save/prdtInfo", s.SavePrdtInfoHandler)

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
