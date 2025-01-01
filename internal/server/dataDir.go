package server

import (
	"exporter/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) DeleteMercTypeHandler(c *gin.Context) {
	mercType := &models.MercType{}
	err := c.ShouldBind(mercType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of mercType: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteMercType(mercType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete mercType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete mercType successfully",
		})
	}
}

func (s *Server) DeleteSuprTypeHandler(c *gin.Context) {
	suprType := &models.SuprType{}
	err := c.ShouldBind(suprType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of suprType: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteSuprType(suprType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete suprType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete suprType successfully",
		})
	}
}

func (s *Server) DeletePrdtTypeHandler(c *gin.Context) {
	prdtType := &models.PrdtType{}
	err := c.ShouldBind(prdtType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of prdtType: " + err.Error(),
		})
		return
	}

	err = s.db.DeletePrdtType(prdtType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete prdtType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete prdtType successfully",
		})
	}
}

func (s *Server) DeleteFoodAddTypeHandler(c *gin.Context) {
	foodAddType := &models.FoodAddType{}
	err := c.ShouldBind(foodAddType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of foodAddType: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteFoodAddType(foodAddType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete foodAddType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete foodAddType successfully",
		})
	}
}

func (s *Server) DeleteFeedAddTypeHandler(c *gin.Context) {
	feedAddType := &models.FeedAddType{}
	err := c.ShouldBind(feedAddType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of feedAddType: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteFeedAddType(feedAddType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete feedAddType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete feedAddType successfully",
		})
	}
}

func (s *Server) DeleteUnitMeasHandler(c *gin.Context) {
	unitMeas := &models.UnitMeas{}
	err := c.ShouldBind(unitMeas)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of unitMeas: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteUnitMeas(unitMeas)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete unitMeas: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete unitMeas successfully",
		})
	}
}

func (s *Server) DeletePackTypeHandler(c *gin.Context) {
	packType := &models.PackType{}
	err := c.ShouldBind(packType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of packType: " + err.Error(),
		})
		return
	}

	err = s.db.DeletePackType(packType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete packType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete packType successfully",
		})
	}
}

func (s *Server) DeleteConTypeHandler(c *gin.Context) {
	conType := &models.ConType{}
	err := c.ShouldBind(conType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of conType: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteConType(conType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete conType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete conType successfully",
		})
	}
}

func (s *Server) DeleteCurrencyHandler(c *gin.Context) {
	currency := &models.Currency{}
	err := c.ShouldBind(currency)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of currency: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteCurrency(currency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete currency: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete currency successfully",
		})
	}
}

func (s *Server) DeleteTradeTermHandler(c *gin.Context) {
	tradeTerm := &models.TradeTerm{}
	err := c.ShouldBind(tradeTerm)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of tradeTerm: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteTradeTerm(tradeTerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete tradeTerm: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete tradeTerm successfully",
		})
	}
}

func (s *Server) DeleteNationHandler(c *gin.Context) {
	nation := &models.Nation{}
	err := c.ShouldBind(nation)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of nation: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteNation(nation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete nation: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete nation successfully",
		})
	}
}

func (s *Server) DeletePortHandler(c *gin.Context) {
	port := &models.Port{}
	err := c.ShouldBind(port)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of port: " + err.Error(),
		})
		return
	}

	err = s.db.DeletePort(port)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete port: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete port successfully",
		})
	}
}

func (s *Server) DeleteTaxTypeHandler(c *gin.Context) {
	taxType := &models.TaxType{}
	err := c.ShouldBind(taxType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of taxType: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteTaxType(taxType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete taxType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete taxType successfully",
		})
	}
}

func (s *Server) DeleteBrandTypeHandler(c *gin.Context) {
	brandType := &models.BrandType{}
	err := c.ShouldBind(brandType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of brandType: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteBrandType(brandType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete brandType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete brandType successfully",
		})
	}
}

func (s *Server) DeleteEduLevelHandler(c *gin.Context) {
	degree := &models.EduLevel{}
	err := c.ShouldBind(degree)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of degree: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteEduLevel(degree)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete degree: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete degree successfully",
		})
	}
}

func (s *Server) DeleteDeptHandler(c *gin.Context) {
	dept := &models.Dept{}
	err := c.ShouldBind(dept)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of dept: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteDept(dept)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete dept: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete dept successfully",
		})
	}
}

func (s *Server) DeletePositionHandler(c *gin.Context) {
	post := &models.Position{}
	err := c.ShouldBind(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of post: " + err.Error(),
		})
		return
	}

	err = s.db.DeletePosition(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete post: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete post successfully",
		})
	}
}

func (s *Server) DeleteQualStdHandler(c *gin.Context) {
	qualStd := &models.QualStd{}
	err := c.ShouldBind(qualStd)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of qualStd: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteQualStd(qualStd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete qualStd: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete qualStd successfully",
		})
	}
}

func (s *Server) DeleteInvLocHandler(c *gin.Context) {
	invLoc := &models.InvLoc{}
	err := c.ShouldBind(invLoc)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of invLoc: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteInvLoc(invLoc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete invLoc: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete invLoc successfully",
		})
	}
}

func (s *Server) DeleteDocReqHandler(c *gin.Context) {
	docReq := &models.DocReq{}
	err := c.ShouldBind(docReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of docReq: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteDocReq(docReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete docReq: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete docReq successfully",
		})
	}
}

func (s *Server) DeletePayMthHandler(c *gin.Context) {
	payMth := &models.PayMth{}
	err := c.ShouldBind(payMth)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of payMth: " + err.Error(),
		})
		return
	}

	err = s.db.DeletePayMth(payMth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete payMth: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete payMth successfully",
		})
	}
}

func (s *Server) DeletePayLimitHandler(c *gin.Context) {
	payLimit := &models.PayLimit{}
	err := c.ShouldBind(payLimit)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of payLimit: " + err.Error(),
		})
		return
	}

	err = s.db.DeletePayLimit(payLimit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete payLimit: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete payLimit successfully",
		})
	}
}

func (s *Server) DeleteFinaDocStatusHandler(c *gin.Context) {
	finaDocStatus := &models.FinaDocStatus{}
	err := c.ShouldBind(finaDocStatus)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of finaDocStatus: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteFinaDocStatus(finaDocStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete finaDocStatus: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete finaDocStatus successfully",
		})
	}
}

func (s *Server) DeleteFinaDocTypeHandler(c *gin.Context) {
	finaDocType := &models.FinaDocType{}
	err := c.ShouldBind(finaDocType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of finaDocType: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteFinaDocType(finaDocType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete finaDocType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete finaDocType successfully",
		})
	}
}

func (s *Server) DeleteExpTypeHandler(c *gin.Context) {
	expType := &models.ExpType{}
	err := c.ShouldBind(expType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of expType: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteExpType(expType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete expType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete expType successfully",
		})
	}
}

func (s *Server) DeleteRatesHandler(c *gin.Context) {
	rates := &models.Rates{}
	err := c.ShouldBind(rates)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of rates: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteRates(rates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete rates: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete rates successfully",
		})
	}
}

func (s *Server) DeleteBussOrderStaHandler(c *gin.Context) {
	bussOrderSta := &models.BussOrderSta{}
	err := c.ShouldBind(bussOrderSta)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of bussOrderSta: " + err.Error(),
		})
		return
	}

	err = s.db.DeleteBussOrderSta(bussOrderSta)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to delete bussOrderSta: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete bussOrderSta successfully",
		})
	}
}

func (s *Server) FindMercTypeHandler(c *gin.Context) {
	mercTypes := &[]models.MercType{}
	s.db.FindMercType(mercTypes)
	c.JSON(http.StatusOK, models.Message{
		MercType: *mercTypes,
	})
}

func (s *Server) FindSuprTypeHandler(c *gin.Context) {
	suprTypes := &[]models.SuprType{}
	s.db.FindSuprType(suprTypes)
	c.JSON(http.StatusOK, models.Message{
		SuprType: *suprTypes,
	})
}

func (s *Server) FindPrdtTypeHandler(c *gin.Context) {
	prdtTypes := &[]models.PrdtType{}
	s.db.FindPrdtType(prdtTypes)
	c.JSON(http.StatusOK, models.Message{
		PrdtType: *prdtTypes,
	})
}

func (s *Server) FindFoodAddTypeHandler(c *gin.Context) {
	foodAddTypes := &[]models.FoodAddType{}
	s.db.FindFoodAddType(foodAddTypes)
	c.JSON(http.StatusOK, models.Message{
		FoodAddType: *foodAddTypes,
	})
}

func (s *Server) FindFeedAddTypeHandler(c *gin.Context) {
	feedAddTypes := &[]models.FeedAddType{}
	s.db.FindFeedAddType(feedAddTypes)
	c.JSON(http.StatusOK, models.Message{
		FeedAddType: *feedAddTypes,
	})
}

func (s *Server) FindUnitMeasHandler(c *gin.Context) {
	unitMeas := &[]models.UnitMeas{}
	s.db.FindUnitMeas(unitMeas)
	c.JSON(http.StatusOK, models.Message{
		UnitMeas: *unitMeas,
	})
}

func (s *Server) FindPackTypeHandler(c *gin.Context) {
	packTypes := &[]models.PackType{}
	s.db.FindPackType(packTypes)
	c.JSON(http.StatusOK, models.Message{
		PackType: *packTypes,
	})
}

func (s *Server) FindConTypeHandler(c *gin.Context) {
	conTypes := &[]models.ConType{}
	s.db.FindConType(conTypes)
	c.JSON(http.StatusOK, models.Message{
		ConType: *conTypes,
	})
}

func (s *Server) FindCurrencyHandler(c *gin.Context) {
	currencies := &[]models.Currency{}
	s.db.FindCurrency(currencies)
	c.JSON(http.StatusOK, models.Message{
		Currency: *currencies,
	})
}

func (s *Server) FindTradeTermHandler(c *gin.Context) {
	tradeTerms := &[]models.TradeTerm{}
	s.db.FindTradeTerm(tradeTerms)
	c.JSON(http.StatusOK, models.Message{
		TradeTerm: *tradeTerms,
	})
}

func (s *Server) FindNationHandler(c *gin.Context) {
	nations := &[]models.Nation{}
	s.db.FindNation(nations)
	c.JSON(http.StatusOK, models.Message{
		Nation: *nations,
	})
}

func (s *Server) FindPortHandler(c *gin.Context) {
	ports := &[]models.Port{}
	s.db.FindPort(ports)
	c.JSON(http.StatusOK, models.Message{
		Port: *ports,
	})
}

func (s *Server) FindTaxTypeHandler(c *gin.Context) {
	taxTypes := &[]models.TaxType{}
	s.db.FindTaxType(taxTypes)
	c.JSON(http.StatusOK, models.Message{
		TaxType: *taxTypes,
	})
}

func (s *Server) FindBrandTypeHandler(c *gin.Context) {
	brandTypes := &[]models.BrandType{}
	s.db.FindBrandType(brandTypes)
	c.JSON(http.StatusOK, models.Message{
		BrandType: *brandTypes,
	})
}

func (s *Server) FindEduLevelHandler(c *gin.Context) {
	degrees := &[]models.EduLevel{}
	s.db.FindEduLevel(degrees)
	c.JSON(http.StatusOK, models.Message{
		EduLevel: *degrees,
	})
}

func (s *Server) FindDeptHandler(c *gin.Context) {
	depts := &[]models.Dept{}
	s.db.FindDept(depts)
	c.JSON(http.StatusOK, models.Message{
		Dept: *depts,
	})
}

func (s *Server) FindPositionHandler(c *gin.Context) {
	posts := &[]models.Position{}
	s.db.FindPosition(posts)
	c.JSON(http.StatusOK, models.Message{
		Position: *posts,
	})
}

func (s *Server) FindQualStdHandler(c *gin.Context) {
	qualStds := &[]models.QualStd{}
	s.db.FindQualStd(qualStds)
	c.JSON(http.StatusOK, models.Message{
		QualStd: *qualStds,
	})
}

func (s *Server) FindInvLocHandler(c *gin.Context) {
	invLocs := &[]models.InvLoc{}
	s.db.FindInvLoc(invLocs)
	c.JSON(http.StatusOK, models.Message{
		InvLoc: *invLocs,
	})
}

func (s *Server) FindDocReqHandler(c *gin.Context) {
	docReqs := &[]models.DocReq{}
	s.db.FindDocReq(docReqs)
	c.JSON(http.StatusOK, models.Message{
		DocReq: *docReqs,
	})
}

func (s *Server) FindPayMthHandler(c *gin.Context) {
	payMths := &[]models.PayMth{}
	s.db.FindPayMth(payMths)
	c.JSON(http.StatusOK, models.Message{
		PayMth: *payMths,
	})
}

func (s *Server) FindPayLimitHandler(c *gin.Context) {
	payLimits := &[]models.PayLimit{}
	s.db.FindPayLimit(payLimits)
	c.JSON(http.StatusOK, models.Message{
		PayLimit: *payLimits,
	})
}

func (s *Server) FindFinaDocStatusHandler(c *gin.Context) {
	finaDocStatuses := &[]models.FinaDocStatus{}
	s.db.FindFinaDocStatus(finaDocStatuses)
	c.JSON(http.StatusOK, models.Message{
		FinaDocStatus: *finaDocStatuses,
	})
}

func (s *Server) FindFinaDocTypeHandler(c *gin.Context) {
	finaDocTypes := &[]models.FinaDocType{}
	s.db.FindFinaDocType(finaDocTypes)
	c.JSON(http.StatusOK, models.Message{
		FinaDocType: *finaDocTypes,
	})
}

func (s *Server) FindExpTypeHandler(c *gin.Context) {
	expTypes := &[]models.ExpType{}
	s.db.FindExpType(expTypes)
	c.JSON(http.StatusOK, models.Message{
		ExpType: *expTypes,
	})
}

func (s *Server) FindRatesHandler(c *gin.Context) {
	rates := &[]models.Rates{}
	s.db.FindRates(rates)
	c.JSON(http.StatusOK, models.Message{
		Rates: *rates,
	})
}

func (s *Server) FindBussOrderStaHandler(c *gin.Context) {
	bussOrderStas := &[]models.BussOrderSta{}
	s.db.FindBussOrderSta(bussOrderStas)
	c.JSON(http.StatusOK, models.Message{
		BussOrderSta: *bussOrderStas,
	})
}

func (s *Server) SaveMercTypeHandler(c *gin.Context) {
	mercType := &models.MercType{}
	err := c.ShouldBind(mercType)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of mercType",
		})
		return
	}

	err = s.db.SaveMercType(mercType)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to save mercType",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved mercType successfully",
		})
	}
}

func (s *Server) SaveSuprTypeHandler(c *gin.Context) {
	suprType := &models.SuprType{}
	err := c.ShouldBind(suprType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of suprType: " + err.Error(),
		})
		return
	}

	err = s.db.SaveSuprType(suprType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save suprType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved suprType successfully",
		})
	}
}

func (s *Server) SavePrdtTypeHandler(c *gin.Context) {
	prdtType := &models.PrdtType{}
	err := c.ShouldBind(prdtType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of prdtType: " + err.Error(),
		})
		return
	}

	err = s.db.SavePrdtType(prdtType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save prdtType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved prdtType successfully",
		})
	}
}

func (s *Server) SaveFoodAddTypeHandler(c *gin.Context) {
	foodAddType := &models.FoodAddType{}
	err := c.ShouldBind(foodAddType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of foodAddType: " + err.Error(),
		})
		return
	}

	err = s.db.SaveFoodAddType(foodAddType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save foodAddType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved foodAddType successfully",
		})
	}
}

func (s *Server) SaveFeedAddTypeHandler(c *gin.Context) {
	feedAddType := &models.FeedAddType{}
	err := c.ShouldBind(feedAddType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of feedAddType: " + err.Error(),
		})
		return
	}

	err = s.db.SaveFeedAddType(feedAddType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save feedAddType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved feedAddType successfully",
		})
	}
}

func (s *Server) SaveUnitMeasHandler(c *gin.Context) {
	unitMeas := &models.UnitMeas{}
	err := c.ShouldBind(unitMeas)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of unitMeas: " + err.Error(),
		})
		return
	}

	err = s.db.SaveUnitMeas(unitMeas)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save unitMeas: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved unitMeas successfully",
		})
	}
}

func (s *Server) SavePackTypeHandler(c *gin.Context) {
	packType := &models.PackType{}
	err := c.ShouldBind(packType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of packType: " + err.Error(),
		})
		return
	}

	err = s.db.SavePackType(packType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save packType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved packType successfully",
		})
	}
}

func (s *Server) SaveConTypeHandler(c *gin.Context) {
	conType := &models.ConType{}
	err := c.ShouldBind(conType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of conType: " + err.Error(),
		})
		return
	}

	err = s.db.SaveConType(conType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save conType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved conType successfully",
		})
	}
}

func (s *Server) SaveCurrencyHandler(c *gin.Context) {
	currency := &models.Currency{}
	err := c.ShouldBind(currency)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of currency: " + err.Error(),
		})
		return
	}

	err = s.db.SaveCurrency(currency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save currency: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved currency successfully",
		})
	}
}

func (s *Server) SaveTradeTermHandler(c *gin.Context) {
	tradeTerm := &models.TradeTerm{}
	err := c.ShouldBind(tradeTerm)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of tradeTerm: " + err.Error(),
		})
		return
	}

	err = s.db.SaveTradeTerm(tradeTerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save tradeTerm: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved tradeTerm successfully",
		})
	}
}

func (s *Server) SaveNationHandler(c *gin.Context) {
	nation := &models.Nation{}
	err := c.ShouldBind(nation)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of nation: " + err.Error(),
		})
		return
	}

	err = s.db.SaveNation(nation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save nation: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved nation successfully",
		})
	}
}

func (s *Server) SavePortHandler(c *gin.Context) {
	port := &models.Port{}
	err := c.ShouldBind(port)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of port: " + err.Error(),
		})
		return
	}

	err = s.db.SavePort(port)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save port: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved port successfully",
		})
	}
}

func (s *Server) SaveTaxTypeHandler(c *gin.Context) {
	taxType := &models.TaxType{}
	err := c.ShouldBind(taxType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of taxType: " + err.Error(),
		})
		return
	}

	err = s.db.SaveTaxType(taxType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save taxType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved taxType successfully",
		})
	}
}

func (s *Server) SaveBrandTypeHandler(c *gin.Context) {
	brandType := &models.BrandType{}
	err := c.ShouldBind(brandType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of brandType: " + err.Error(),
		})
		return
	}

	err = s.db.SaveBrandType(brandType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save brandType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved brandType successfully",
		})
	}
}

func (s *Server) SaveEduLevelHandler(c *gin.Context) {
	degree := &models.EduLevel{}
	err := c.ShouldBind(degree)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of degree: " + err.Error(),
		})
		return
	}

	err = s.db.SaveEduLevel(degree)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save degree: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved degree successfully",
		})
	}
}

func (s *Server) SaveDeptHandler(c *gin.Context) {
	dept := &models.Dept{}
	err := c.ShouldBind(dept)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of dept: " + err.Error(),
		})
		return
	}

	err = s.db.SaveDept(dept)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save dept: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved dept successfully",
		})
	}
}

func (s *Server) SavePositionHandler(c *gin.Context) {
	post := &models.Position{}
	err := c.ShouldBind(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of post: " + err.Error(),
		})
		return
	}

	err = s.db.SavePosition(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save post: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved post successfully",
		})
	}
}

func (s *Server) SaveQualStdHandler(c *gin.Context) {
	qualStd := &models.QualStd{}
	err := c.ShouldBind(qualStd)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of qualStd: " + err.Error(),
		})
		return
	}

	err = s.db.SaveQualStd(qualStd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save qualStd: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved qualStd successfully",
		})
	}
}

func (s *Server) SaveInvLocHandler(c *gin.Context) {
	invLoc := &models.InvLoc{}
	err := c.ShouldBind(invLoc)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of invLoc: " + err.Error(),
		})
		return
	}

	err = s.db.SaveInvLoc(invLoc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save invLoc: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved invLoc successfully",
		})
	}
}

func (s *Server) SaveDocReqHandler(c *gin.Context) {
	docReq := &models.DocReq{}
	err := c.ShouldBind(docReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of docReq: " + err.Error(),
		})
		return
	}

	err = s.db.SaveDocReq(docReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save docReq: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved docReq successfully",
		})
	}
}

func (s *Server) SavePayMthHandler(c *gin.Context) {
	payMth := &models.PayMth{}
	err := c.ShouldBind(payMth)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of payMth: " + err.Error(),
		})
		return
	}

	err = s.db.SavePayMth(payMth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save payMth: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved payMth successfully",
		})
	}
}

func (s *Server) SavePayLimitHandler(c *gin.Context) {
	payLimit := &models.PayLimit{}
	err := c.ShouldBind(payLimit)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of payLimit: " + err.Error(),
		})
		return
	}

	err = s.db.SavePayLimit(payLimit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save payLimit: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved payLimit successfully",
		})
	}
}

func (s *Server) SaveFinaDocStatusHandler(c *gin.Context) {
	finaDocStatus := &models.FinaDocStatus{}
	err := c.ShouldBind(finaDocStatus)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of finaDocStatus: " + err.Error(),
		})
		return
	}

	err = s.db.SaveFinaDocStatus(finaDocStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save finaDocStatus: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved finaDocStatus successfully",
		})
	}
}

func (s *Server) SaveFinaDocTypeHandler(c *gin.Context) {
	finaDocType := &models.FinaDocType{}
	err := c.ShouldBind(finaDocType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of finaDocType: " + err.Error(),
		})
		return
	}

	err = s.db.SaveFinaDocType(finaDocType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save finaDocType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved finaDocType successfully",
		})
	}
}

func (s *Server) SaveExpTypeHandler(c *gin.Context) {
	expType := &models.ExpType{}
	err := c.ShouldBind(expType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of expType: " + err.Error(),
		})
		return
	}

	err = s.db.SaveExpType(expType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save expType: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved expType successfully",
		})
	}
}

func (s *Server) SaveRatesHandler(c *gin.Context) {
	rates := &models.Rates{}
	err := c.ShouldBind(rates)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of rates: " + err.Error(),
		})
		return
	}

	err = s.db.SaveRates(rates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save rates: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved rates successfully",
		})
	}
}

func (s *Server) SaveBussOrderStaHandler(c *gin.Context) {
	bussOrderSta := &models.BussOrderSta{}
	err := c.ShouldBind(bussOrderSta)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "error in bind of bussOrderSta: " + err.Error(),
		})
		return
	}

	err = s.db.SaveBussOrderSta(bussOrderSta)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "failed to save bussOrderSta: " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved bussOrderSta successfully",
		})
	}
}
