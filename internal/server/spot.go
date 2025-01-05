package server

import (
	"exporter/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindSpotHandler 查询所有 Spot 记录
func (s *Server) FindSpotHandler(c *gin.Context) {
	var Spots []models.Spot
	s.db.Find(&Spots)
	c.JSON(http.StatusOK, models.Message{Spot: Spots})
}

// DeleteSpotHandler 删除 Spot 记录
func (s *Server) DeleteSpotHandler(c *gin.Context) {
	Spot := &models.Spot{}
	if err := c.ShouldBind(Spot); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("删除 Spot: %+v\n", Spot)

	if err := s.db.Delete(Spot); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

// SaveSpotHandler 保存或更新 Spot 记录
func (s *Server) SaveSpotHandler(c *gin.Context) {
	Spot := &models.Spot{}
	if err := c.ShouldBind(Spot); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("保存 Spot: %+v\n", Spot)

	// 保存 Spot 记录
	if err := s.db.Save(Spot); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}
