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
	l := len(Sends)
	var merc models.Merchant
	for i := 0; i < l; i++ {
		if Sends[i].Merchant1Id != 0 {
			s.db.FindById(Sends[i].Merchant1Id, &merc)
			Sends[i].Merchant1Name = merc.Merc
		}
		if Sends[i].Merchant2Id != 0 {
			s.db.FindById(Sends[i].Merchant2Id, &merc)
			Sends[i].Merchant2Name = merc.Merc
		}
		if Sends[i].Merchant3Id != 0 {
			s.db.FindById(Sends[i].Merchant3Id, &merc)
			Sends[i].Merchant3Name = merc.Merc
		}
	}
	c.JSON(http.StatusOK, models.Message{Send: Sends})
}

func (s *Server) FindSendPrdtInfoHandler(c *gin.Context) {
	var Send models.Send

	s.db.FindSendPrdtInfo(&Send)
	PrdtInfos := Send.PrdtInfos
	c.JSON(http.StatusOK, models.Message{PrdtInfo: PrdtInfos})
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
	Send := &models.Send{}
	if err := c.ShouldBind(Send); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}
	var PrdtInfo models.PrdtInfo
	s.db.FindById(Send.PrdtInfoId, &PrdtInfo)
	s.db.FindById(Send.ID, Send)
	Send.PrdtInfos = append(Send.PrdtInfos, PrdtInfo)

	// 保存 Send 记录
	if err := s.db.DeleteSendPrdtInfo(Send, &PrdtInfo); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddSendPrdtInfo(c *gin.Context) {
	Send := &models.Send{}
	if err := c.ShouldBind(Send); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}
	log.Printf("%d   %d\n", Send.ID, Send.PrdtInfoId)
	var PrdtInfo models.PrdtInfo
	s.db.FindById(Send.PrdtInfoId, &PrdtInfo)
	s.db.FindById(Send.ID, Send)
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
