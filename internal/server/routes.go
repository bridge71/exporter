package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))
	r.GET("/", s.HelloWorldHandler)

	r.GET("/health", s.healthHandler)

	r.POST("/create/acct", s.CreateAcctHandler)
	r.POST("/modify/acct", s.CreateAcctHandler)
	r.GET("/find/acct", s.FindAcctHandler)

	r.POST("/create/user", s.CreateUserHandler)
	r.POST("/auth", s.LoginHandler)
	r.POST("/modify/user/password", s.ModifyUserPassHandler)
	// 创建“find”分组
	findGroup := r.Group("/find")

	// 创建“acct”子分组，用于账户相关的查找
	acctGroup := findGroup.Group("/acct")
	{
		// 精确查找接口
		acctGroup.GET("/acctCode", s.FindAcctByAcctCodeHandler)
		acctGroup.GET("/acctAbbr", s.FindAcctByAcctAbbrHandler)
		acctGroup.GET("/EtyAbbr", s.FindAcctByEtyAbbrHandler)
		acctGroup.GET("/acctName", s.FindAcctByAcctNameHandler)
		acctGroup.GET("/acctAddr", s.FindAcctByAcctAddrHandler)
		acctGroup.GET("/nation", s.FindAcctByNationHandler)
		acctGroup.GET("/taxType", s.FindAcctByTaxTypeHandler)
		acctGroup.GET("/taxCode", s.FindAcctByTaxCodeHandler)
		acctGroup.GET("/phoneNum", s.FindAcctByPhoneNumHandler)
		acctGroup.GET("/email", s.FindAcctByEmailHandler)
		acctGroup.GET("/website", s.FindAcctByWebsiteHandler)
		acctGroup.GET("/regDate", s.FindAcctByRegDateHandler)
		acctGroup.GET("/notes", s.FindAcctByNotesHandler)

		// 模糊查找接口
		acctGroup.GET("/fzz/acctAddr", s.FzzFindAcctByAcctAddrHandler)
		acctGroup.GET("/fzz/email", s.FzzFindAcctByEmailHandler)
		acctGroup.GET("/fzz/website", s.FzzFindAcctByWebsiteHandler)
		acctGroup.GET("/fzz/regDate", s.FzzFindAcctByRegDateHandler)
		acctGroup.GET("/fzz/notes", s.FzzFindAcctByNotesHandler)
	}

	// 创建“acctBank”子分组，用于账户银行相关的查找
	acctBankGroup := findGroup.Group("/acctBank")
	{
		// 精确查找接口
		acctBankGroup.GET("/id", s.FindAcctBankByIdHandler)
		acctBankGroup.GET("/accName", s.FindAcctBankByAccNameHandler)
		acctBankGroup.GET("/accNum", s.FindAcctBankByAccNumHandler)
		acctBankGroup.GET("/currency", s.FindAcctBankByCurrencyHandler)
		acctBankGroup.GET("/bankName", s.FindAcctBankByBankNameHandler)
		acctBankGroup.GET("/swiftCode", s.FindAcctBankBySwiftCodeHandler)
		acctBankGroup.GET("/bankAddr", s.FindAcctBankByBankAddrHandler)
		acctBankGroup.GET("/notes", s.FindAcctBankByNotesHandler)

		// 模糊查找接口
		acctBankGroup.GET("/fzz/bankName", s.FzzFindAcctBankByBankNameHandler)
		acctBankGroup.GET("/fzz/bankAddr", s.FzzFindAcctBankByBankAddrHandler)
		acctBankGroup.GET("/fzz/notes", s.FzzFindAcctBankByNotesHandler)
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
