package server

import (
	"exporter/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) DeleteShouldInSend(c *gin.Context) {
	ShouldIn := models.ShouldIn{}
	ShouldIn.ID = s.Str2Uint(c.PostForm("ID"))
	var Send models.Send
	ShouldIn.ID = s.Str2Uint(c.PostForm("SendId"))
	ShouldIn.Sends = append(ShouldIn.Sends, Send)

	if Send.ID == 0 || ShouldIn.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Send 记录
	if err := s.db.DeleteShouldInSend(&ShouldIn, &Send); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

// FindShouldInHandler 查询所有 ShouldIn 记录
func (s *Server) FindShouldInHandler(c *gin.Context) {
	var ShouldIns []models.ShouldIn
	s.db.FindShouldIns(&ShouldIns)
	c.JSON(http.StatusOK, models.Message{ShouldIn: ShouldIns})
}

func (s *Server) DeleteShouldInIn(c *gin.Context) {
	ShouldIn := models.ShouldIn{}
	ShouldIn.ID = s.Str2Uint(c.PostForm("ID"))
	var In models.In
	ShouldIn.ID = s.Str2Uint(c.PostForm("InId"))
	ShouldIn.Ins = append(ShouldIn.Ins, In)

	if In.ID == 0 || ShouldIn.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 In 记录
	if err := s.db.DeleteShouldInIn(&ShouldIn, &In); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) FindShouldInSendHandler(c *gin.Context) {
	ShouldIn := models.ShouldIn{}
	ShouldIn.ID = s.Str2Uint(c.PostForm("ID"))
	if ShouldIn.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindShouldInSend(&ShouldIn)
	Send := ShouldIn.Sends
	c.JSON(http.StatusOK, models.Message{Send: Send})
}

func (s *Server) FindShouldInInHandler(c *gin.Context) {
	ShouldIn := models.ShouldIn{}
	ShouldIn.ID = s.Str2Uint(c.PostForm("ID"))
	if ShouldIn.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindShouldInIn(&ShouldIn)
	Ins := ShouldIn.Ins
	c.JSON(http.StatusOK, models.Message{In: Ins})
}

func (s *Server) FindShouldInSaleHandler(c *gin.Context) {
	ShouldIn := models.ShouldIn{}
	ShouldIn.ID = s.Str2Uint(c.PostForm("ID"))
	s.db.FindShouldInSale(&ShouldIn)
	Sale := ShouldIn.Sales
	c.JSON(http.StatusOK, models.Message{Sale: Sale})
}

// DeleteShouldInHandler 删除 ShouldIn 记录
func (s *Server) DeleteShouldInHandler(c *gin.Context) {
	ShouldIn := &models.ShouldIn{}
	if err := c.ShouldBind(ShouldIn); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("删除 ShouldIn: %+v\n", ShouldIn)

	if err := s.db.Delete(ShouldIn); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) DeleteShouldInSale(c *gin.Context) {
	ShouldIn := models.ShouldIn{}
	ShouldIn.ID = s.Str2Uint(c.PostForm("ID"))
	var Sale models.Sale
	Sale.ID = s.Str2Uint(c.PostForm("SaleId"))
	ShouldIn.Sales = append(ShouldIn.Sales, Sale)

	if ShouldIn.ID == 0 || Sale.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 ShouldIn 记录
	if err := s.db.DeleteShouldInSale(&ShouldIn, &Sale); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) AddShouldInSale(c *gin.Context) {
	ShouldIn := models.ShouldIn{}
	ShouldIn.ID = s.Str2Uint(c.PostForm("ID"))
	var Sale models.Sale
	Sale.ID = s.Str2Uint(c.PostForm("SaleId"))

	if ShouldIn.ID == 0 || Sale.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindById(ShouldIn.ID, &ShouldIn)
	s.db.FindById(Sale.ID, &Sale)

	if Sale.AcctName == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	ShouldIn.Sales = append(ShouldIn.Sales, Sale)

	// 保存 ShouldIn 记录
	if err := s.db.Save(ShouldIn); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddShouldInIns(c *gin.Context) {
	ShouldIn := models.ShouldIn{}
	ShouldIn.ID = s.Str2Uint(c.PostForm("ID"))
	var In models.In
	In.ID = s.Str2Uint(c.PostForm("InId"))

	if ShouldIn.ID == 0 || In.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindById(ShouldIn.ID, &ShouldIn)
	s.db.FindById(In.ID, &In)

	if In.ReceNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	ShouldIn.Ins = append(ShouldIn.Ins, In)

	// 保存 ShouldIn 记录
	if err := s.db.Save(ShouldIn); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddShouldInSends(c *gin.Context) {
	ShouldIn := models.ShouldIn{}
	ShouldIn.ID = s.Str2Uint(c.PostForm("ID"))
	var Send models.Send
	Send.ID = s.Str2Uint(c.PostForm("SendId"))

	if ShouldIn.ID == 0 || Send.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindById(ShouldIn.ID, &ShouldIn)
	s.db.FindById(Send.ID, &Send)

	if ShouldIn.BillReceNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	ShouldIn.Sends = append(ShouldIn.Sends, Send)

	// 保存 ShouldIn 记录
	if err := s.db.Save(ShouldIn); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) SaveShouldInHandler(c *gin.Context) {
	ShouldIn := &models.ShouldIn{}
	if err := c.ShouldBind(ShouldIn); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("保存 ShouldIn: %+v\n", ShouldIn)

	// 保存 ShouldIn 记录
	if err := s.db.SaveShouldIn(ShouldIn); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}
