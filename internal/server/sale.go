package server

import (
	"exporter/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindSaleHandler 查询所有 Sale 记录
func (s *Server) FindSaleHandler(c *gin.Context) {
	var Sales []models.Sale
	s.db.Find(&Sales)
	c.JSON(http.StatusOK, models.Message{Sale: Sales})
}

// DeleteSaleHandler 删除 Sale 记录
func (s *Server) DeleteSaleHandler(c *gin.Context) {
	Sale := &models.Sale{}
	if err := c.ShouldBind(Sale); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("删除 Sale: %+v\n", Sale)

	if err := s.db.Delete(Sale); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

// SaveSaleHandler 保存或更新 Sale 记录
func (s *Server) SaveSaleHandler(c *gin.Context) {
	Sale := &models.Sale{}
	if err := c.ShouldBind(Sale); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("保存 Sale: %+v\n", Sale)

	// 保存 Sale 记录
	if err := s.db.Save(Sale); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}
