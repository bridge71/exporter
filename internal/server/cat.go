package server

import (
	"exporter/internal/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FindCatHandler 查询所有 Cat 记录
func (s *Server) FindCatHandler(c *gin.Context) {
	var cats []models.Cat
	s.db.FindCat(&cats)
	c.JSON(http.StatusOK, models.Message{Cat: cats})
}

// DeleteCatHandler 删除 Cat 记录
func (s *Server) DeleteCatHandler(c *gin.Context) {
	cat := &models.Cat{}
	if err := c.ShouldBind(cat); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("删除 Cat: %+v\n", cat)

	if err := s.db.DeleteCat(cat); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

// FindCatByIDHandler 根据 CatID 查询 Cat 记录
func (s *Server) FindCatByIDHandler(c *gin.Context) {
	catIDStr := c.PostForm("CatID")
	catID, err := strconv.Atoi(catIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "CatID 格式错误",
		})
		return
	}

	var cat models.Cat
	if err := s.db.FindCatByID(&cat, uint(catID)); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "查询失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{Cat: []models.Cat{cat}})
}

// SaveCatHandler 保存或更新 Cat 记录
func (s *Server) SaveCatHandler(c *gin.Context) {
	cat := &models.Cat{}
	if err := c.ShouldBind(cat); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("保存 Cat: %+v\n", cat)

	// 处理文件上传
	err, fileID, fileName := s.SaveFile(c, "file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "文件上传失败",
		})
		return
	}
	cat.FileID = fileID
	cat.FileName = fileName

	// 保存 Cat 记录
	if err := s.db.SaveCat(cat); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}
