package server

import (
	"exporter/internal/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) FindBrandHandler(c *gin.Context) {
	var brands []models.Brand
	s.db.FindBrand(&brands)
	c.JSON(http.StatusOK, models.Message{Brand: brands})
}

func (s *Server) DeleteBrandHandler(c *gin.Context) {
	brand := &models.Brand{}
	if err := c.ShouldBind(brand); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("删除 brand: %+v\n", brand)

	if err := s.db.DeleteBrand(brand); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) FindBrandByIDHandler(c *gin.Context) {
	brandIDStr := c.PostForm("BrandID")
	brandID, err := strconv.Atoi(brandIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "brandID 格式错误",
		})
		return
	}

	var brand models.Brand
	if err := s.db.FindBrandByID(&brand, uint(brandID)); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "查询失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{Brand: []models.Brand{brand}})
}

func (s *Server) SaveBrandHandler(c *gin.Context) {
	brand := &models.Brand{}
	if err := c.ShouldBind(brand); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("保存 brand: %+v\n", brand)

	// 处理文件上传
	err, fileID, fileName := s.SaveFile(c, "file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "文件上传失败",
		})
		return
	}
	brand.FileID = fileID
	brand.FileName = fileName

	// 保存 brand 记录
	if err := s.db.SaveBrand(brand); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}
