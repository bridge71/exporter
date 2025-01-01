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

// FindPayMentMethodByIdHandler 根据 PayMentMethodId 查询 PayMentMethod 记录
func (s *Server) FindPayMentMethodByIdHandler(c *gin.Context) {
	PayMentMethodIdStr := c.PostForm("PayMentMethodId")
	PayMentMethodId, err := strconv.Atoi(PayMentMethodIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "PayMentMethodId 格式错误",
		})
		return
	}

	var PayMentMethod models.PayMentMethod
	if err := s.db.FindPayMentMethodByID(&PayMentMethod, uint(PayMentMethodId)); err != nil {
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
