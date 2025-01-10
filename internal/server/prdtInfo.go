package server

import (
	"exporter/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindPrdtInfoHandler 查询所有 PrdtInfo 记录
func (s *Server) FindPrdtInfoHandler(c *gin.Context) {
	var PrdtInfos []models.PrdtInfo
	s.db.FindPrdtInfo(&PrdtInfos)
	c.JSON(http.StatusOK, models.Message{PrdtInfo: PrdtInfos})
}

// DeletePrdtInfoHandler 删除 PrdtInfo 记录
func (s *Server) DeletePrdtInfoHandler(c *gin.Context) {
	PrdtInfo := &models.PrdtInfo{}
	if err := c.ShouldBind(PrdtInfo); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("删除 PrdtInfo: %+v\n", PrdtInfo)

	if err := s.db.Delete(PrdtInfo); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

// SavePrdtInfoHandler 保存或更新 PrdtInfo 记录
func (s *Server) SavePrdtInfoHandler(c *gin.Context) {
	PrdtInfo := &models.PrdtInfo{}

	PrdtInfo.ID = s.Str2Uint(c.PostForm("ID"))
	if PrdtInfo.ID != 0 {
		s.db.FindByID(PrdtInfo.ID, PrdtInfo)
	}
	PrdtInfo.BrandID = s.Str2Uint(c.PostForm("BrandID"))
	PrdtInfo.CatID = s.Str2Uint(c.PostForm("CatID"))
	PrdtInfo.PackSpecID = s.Str2Uint(c.PostForm("PackSpecID"))
	PrdtInfo.SpotID = s.Str2Uint(c.PostForm("SpotID"))
	if PrdtInfo.BrandID == 0 || PrdtInfo.CatID == 0 || PrdtInfo.PackSpecID == 0 || PrdtInfo.SpotID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败",
		})
		return
	}
	PrdtInfo.Brand.ID = PrdtInfo.BrandID
	PrdtInfo.Cat.ID = PrdtInfo.CatID
	PrdtInfo.Spot.ID = PrdtInfo.SpotID
	PrdtInfo.PackSpec.ID = PrdtInfo.PackSpecID

	PrdtInfo.Amount = s.Str2Uint(c.PostForm("Amount"))
	PrdtInfo.Factory = c.PostForm("Factory")
	PrdtInfo.Currency = c.PostForm("Currency")
	PrdtInfo.Unit = c.PostForm("Unit")
	PrdtInfo.UnitPrice = s.Str2Uint(c.PostForm("UnitPrice"))
	PrdtInfo.Weight = s.Str2Uint(c.PostForm("Weight"))
	PrdtInfo.ItemNum = s.Str2Uint(c.PostForm("ItemNum"))
	PrdtInfo.WeightUnit = c.PostForm("WeightUnit")
	PrdtInfo.TradeTerm = c.PostForm("TradeTerm")
	PrdtInfo.Factory = c.PostForm("Factory")
	log.Printf("保存 PrdtInfo: %+v\n", PrdtInfo)

	// 保存 PrdtInfo 记录
	if err := s.db.Save(PrdtInfo); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}
