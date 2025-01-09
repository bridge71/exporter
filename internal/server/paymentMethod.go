package server

import (
	"exporter/internal/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) FindPayMentMethodHandler(c *gin.Context) {
	var PayMentMethods []models.PayMentMethod
	s.db.FindPayMentMethods(&PayMentMethods)
	c.JSON(http.StatusOK, models.Message{PayMentMethod: PayMentMethods})
}

func (s *Server) DeletePayMentMethodHandler(c *gin.Context) {
	PayMentMethod := &models.PayMentMethod{}
	if err := c.ShouldBind(PayMentMethod); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("删除 PayMentMethod: %+v\n", PayMentMethod)

	if err := s.db.DeletePayMentMethod(PayMentMethod); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

// FindPayMentMethodByIDHandler 根据 PayMentMethodID 查询 PayMentMethod 记录
func (s *Server) FindPayMentMethodByIDHandler(c *gin.Context) {
	PayMentMethodIDStr := c.PostForm("PayMentMethodID")
	PayMentMethodID, err := strconv.Atoi(PayMentMethodIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "PayMentMethodID 格式错误",
		})
		return
	}

	var PayMentMethod models.PayMentMethod
	if err := s.db.FindPayMentMethodByID(&PayMentMethod, uint(PayMentMethodID)); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "查询失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{PayMentMethod: []models.PayMentMethod{PayMentMethod}})
}

// SavePayMentMethodHandler 保存或更新 PayMentMethod 记录
func (s *Server) SavePayMentMethodHandler(c *gin.Context) {
	PayMentMethod := &models.PayMentMethod{}
	if err := c.ShouldBind(PayMentMethod); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("保存 PayMentMethod: %+v\n", PayMentMethod)

	if err := s.db.SavePayMentMethod(PayMentMethod); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}
