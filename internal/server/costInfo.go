package server

import (
	"exporter/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindCostInfoHandler 查询所有 CostInfo 记录
func (s *Server) FindCostInfoHandler(c *gin.Context) {
	var CostInfos []models.CostInfo
	s.db.Find(&CostInfos)
	c.JSON(http.StatusOK, models.Message{CostInfo: CostInfos})
}

// DeleteCostInfoHandler 删除 CostInfo 记录
func (s *Server) DeleteCostInfoHandler(c *gin.Context) {
	CostInfo := &models.CostInfo{}
	if err := c.ShouldBind(CostInfo); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("删除 CostInfo: %+v\n", CostInfo)

	if err := s.db.Delete(CostInfo); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

// SaveCostInfoHandler 保存或更新 CostInfo 记录
func (s *Server) SaveCostInfoHandler(c *gin.Context) {
	CostInfo := &models.CostInfo{}
	if err := c.ShouldBind(CostInfo); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("保存 CostInfo: %+v\n", CostInfo)

	// 保存 CostInfo 记录
	if err := s.db.Save(CostInfo); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}
