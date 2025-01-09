package server

import (
	"exporter/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindLoadingInfoHandler 查询所有 LoadingInfo 记录
func (s *Server) FindLoadingInfoHandler(c *gin.Context) {
	var LoadingInfos []models.LoadingInfo
	s.db.FindLoadingInfo(&LoadingInfos)
	c.JSON(http.StatusOK, models.Message{LoadingInfo: LoadingInfos})
}

// DeleteLoadingInfoHandler 删除 LoadingInfo 记录
func (s *Server) DeleteLoadingInfoHandler(c *gin.Context) {
	LoadingInfo := &models.LoadingInfo{}
	if err := c.ShouldBind(LoadingInfo); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("删除 LoadingInfo: %+v\n", LoadingInfo)

	if err := s.db.Delete(LoadingInfo); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

// SaveLoadingInfoHandler 保存或更新 LoadingInfo 记录
func (s *Server) SaveLoadingInfoHandler(c *gin.Context) {
	LoadingInfo := &models.LoadingInfo{}
	LoadingInfo.BrandID = s.Str2Uint(c.PostForm("BrandID"))
	LoadingInfo.CatID = s.Str2Uint(c.PostForm("CatID"))
	LoadingInfo.PackSpecID = s.Str2Uint(c.PostForm("PackSpecID"))
	if LoadingInfo.BrandID == 0 || LoadingInfo.CatID == 0 || LoadingInfo.PackSpecID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败",
		})
		return
	}
	LoadingInfo.ID = s.Str2Uint(c.PostForm("ID"))
	if LoadingInfo.ID != 0 {
		s.db.FindByID(LoadingInfo.ID, LoadingInfo)
	}
	LoadingInfo.Brand.ID = LoadingInfo.BrandID
	LoadingInfo.Cat.ID = LoadingInfo.CatID
	LoadingInfo.PackSpec.ID = LoadingInfo.PackSpecID

	LoadingInfo.Unit = c.PostForm("Unit")
	LoadingInfo.ItemNum = s.Str2Uint(c.PostForm("ItemNum"))
	LoadingInfo.NetWeight = s.Str2Uint(c.PostForm("NetWeight"))
	LoadingInfo.PrdtPlant = c.PostForm("PrdtPlant")
	LoadingInfo.BatNum = c.PostForm("BatNum")
	LoadingInfo.VehNum = c.PostForm("VehNum")
	LoadingInfo.VehNum = c.PostForm("VehNum")
	LoadingInfo.SealNum = c.PostForm("SealNum")
	LoadingInfo.CnrNum = c.PostForm("CnrNum")
	log.Printf("保存 LoadingInfo: %+v\n", LoadingInfo)

	// 保存 LoadingInfo 记录
	if err := s.db.Save(LoadingInfo); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}
