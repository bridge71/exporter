package server

import (
	"exporter/internal/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FindSaleHandler 查询所有 Sale 记录
func (s *Server) FindSaleHandler(c *gin.Context) {
	var Sales []models.Sale

	s.db.FindSalePrdtInfo(&Sales)
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

	// 动态解析 PrdtInfos 数组
	Sale.PrdtInfos = make([]models.PrdtInfo, 0)
	i := 0
	for {
		idStr := c.PostForm("PrdtInfos[" + strconv.Itoa(i) + "][ID]")
		if idStr == "" {
			break // 如果没有更多数据，退出循环
		}

		id_ := s.Str2Uint(idStr)
		if id_ == 0 {
			break
		}
		prdt := []models.PrdtInfo{}
		s.db.FindById(id_, &prdt)
		Sale.PrdtInfos = append(Sale.PrdtInfos, prdt[0])
		i++
	}
	i = 0
	for {
		idStr := c.PostForm("DocReq[" + strconv.Itoa(i) + "][DocReqId]")
		if idStr == "" {
			break // 如果没有更多数据，退出循环
		}

		id_ := s.Str2Uint(idStr)

		if id_ == 0 {
			break
		}
		doc := []models.DocReq{}
		s.db.FindDocReqById(id_, &doc)
		Sale.DocReq = append(Sale.DocReq, doc[0])
		i++
	}
	var err error
	err, Sale.FileId, Sale.FileName = s.SaveFile(c)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: err.Error(),
		})
	}
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
