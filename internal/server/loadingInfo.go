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
	s.db.Find(&LoadingInfos)
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
	if err := c.ShouldBind(LoadingInfo); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}

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
