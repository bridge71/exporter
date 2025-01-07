package server

import (
	"exporter/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindSendHandler 查询所有 Send 记录
func (s *Server) FindSendHandler(c *gin.Context) {
	var Sends []models.Send
	s.db.FindSends(&Sends)
	c.JSON(http.StatusOK, models.Message{Send: Sends})
}

func (s *Server) FindSendPrdtInfoHandler(c *gin.Context) {
	Send := models.Send{}
	Send.ID = s.Str2Uint(c.PostForm("ID"))
	if Send.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindSendPrdtInfo(&Send)
	PrdtInfos := Send.PrdtInfos
	c.JSON(http.StatusOK, models.Message{PrdtInfo: PrdtInfos})
}

func (s *Server) FindSendSaleHandler(c *gin.Context) {
	Send := models.Send{}
	Send.ID = s.Str2Uint(c.PostForm("ID"))
	s.db.FindSendSale(&Send)
	Sale := Send.Sales
	c.JSON(http.StatusOK, models.Message{Sale: Sale})
}

// DeleteSendHandler 删除 Send 记录
func (s *Server) DeleteSendHandler(c *gin.Context) {
	Send := &models.Send{}
	if err := c.ShouldBind(Send); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("删除 Send: %+v\n", Send)

	if err := s.db.Delete(Send); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) DeleteSendPrdtInfo(c *gin.Context) {
	Send := models.Send{}
	Send.ID = s.Str2Uint(c.PostForm("ID"))
	var PrdtInfo models.PrdtInfo
	PrdtInfo.ID = s.Str2Uint(c.PostForm("PrdtInfoId"))
	if Send.ID == 0 || PrdtInfo.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	Send.PrdtInfos = append(Send.PrdtInfos, PrdtInfo)

	if err := s.db.DeleteSendPrdtInfo(&Send, &PrdtInfo); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) DeleteSendSale(c *gin.Context) {
	Send := models.Send{}
	Send.ID = s.Str2Uint(c.PostForm("ID"))
	var Sale models.Sale
	Sale.ID = s.Str2Uint(c.PostForm("SaleId"))
	Send.Sales = append(Send.Sales, Sale)

	if Send.ID == 0 || Sale.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Send 记录
	if err := s.db.DeleteSendSale(&Send, &Sale); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) AddSendPrdtInfo(c *gin.Context) {
	Send := models.Send{}
	Send.ID = s.Str2Uint(c.PostForm("ID"))
	var PrdtInfo models.PrdtInfo
	PrdtInfo.ID = s.Str2Uint(c.PostForm("PrdtInfoId"))

	if Send.ID == 0 || PrdtInfo.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindById(PrdtInfo.ID, &PrdtInfo)
	s.db.FindById(Send.ID, &Send)
	log.Printf("%d\n", PrdtInfo.ID)
	Send.PrdtInfos = append(Send.PrdtInfos, PrdtInfo)

	// 保存 Send 记录
	if err := s.db.Save(Send); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddSendSale(c *gin.Context) {
	Send := models.Send{}
	Send.ID = s.Str2Uint(c.PostForm("ID"))
	var Sale models.Sale
	Sale.ID = s.Str2Uint(c.PostForm("SendId"))

	if Send.ID == 0 || Send.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindById(Send.ID, &Send)
	s.db.FindById(Send.ID, &Send)
	Send.Sales = append(Send.Sales, Sale)

	// 保存 Send 记录
	if err := s.db.Save(Send); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) SaveSendHandler(c *gin.Context) {
	Send := &models.Send{}
	if err := c.ShouldBind(Send); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("保存 Send: %+v\n", Send)
	var err error

	err, Send.File1Id, Send.File1Name = s.SaveFile(c, "file1")
	err, Send.File2Id, Send.File2Name = s.SaveFile(c, "file2")
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: err.Error(),
		})
	}
	// 保存 Send 记录
	if err := s.db.Save(Send); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}
