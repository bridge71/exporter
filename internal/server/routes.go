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

	r.POST("/create/user", s.CreateUserHandler)
	r.POST("/auth", s.LoginHandler)
	r.POST("/modify/user/password", s.ModifyUserPassHandler)

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
