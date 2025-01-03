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
	s.db.Find(&PrdtInfos)
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
	if err := c.ShouldBind(PrdtInfo); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}

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

func (s *Server) CreatePrdtInfoHandler(c *gin.Context) {
	PrdtInfo := &models.PrdtInfo{}
	if err := c.ShouldBind(PrdtInfo); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("保存 PrdtInfo: %+v\n", PrdtInfo)

	// 保存 PrdtInfo 记录
	if err := s.db.Create(PrdtInfo); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}
