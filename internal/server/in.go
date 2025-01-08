package server

import (
	"exporter/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindInHandler 查询所有 In 记录
func (s *Server) FindInHandler(c *gin.Context) {
	var Ins []models.In
	s.db.FindIns(&Ins)
	c.JSON(http.StatusOK, models.Message{In: Ins})
}

func (s *Server) DeleteInShouldIn(c *gin.Context) {
	In := models.In{}
	In.ID = s.Str2Uint(c.PostForm("ID"))
	var ShouldIn models.ShouldIn
	ShouldIn.ID = s.Str2Uint(c.PostForm("ShouldInId"))
	In.ShouldIns = append(In.ShouldIns, ShouldIn)

	if In.ID == 0 || ShouldIn.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 In 记录
	if err := s.db.DeleteInShouldIn(&In, &ShouldIn); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) DeleteInSend(c *gin.Context) {
	In := models.In{}
	In.ID = s.Str2Uint(c.PostForm("ID"))
	var Send models.Send
	Send.ID = s.Str2Uint(c.PostForm("SendId"))
	In.Sends = append(In.Sends, Send)

	if In.ID == 0 || Send.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 In 记录
	if err := s.db.DeleteInSend(&In, &Send); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) FindInSendHandler(c *gin.Context) {
	In := models.In{}
	In.ID = s.Str2Uint(c.PostForm("ID"))
	if In.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindInSend(&In)
	Send := In.Sends
	c.JSON(http.StatusOK, models.Message{Send: Send})
}

func (s *Server) FindInShouldInHandler(c *gin.Context) {
	In := models.In{}
	In.ID = s.Str2Uint(c.PostForm("ID"))
	if In.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindInShouldIn(&In)
	ShouldIn := In.ShouldIns
	c.JSON(http.StatusOK, models.Message{ShouldIn: ShouldIn})
}

func (s *Server) FindInSaleHandler(c *gin.Context) {
	In := models.In{}
	In.ID = s.Str2Uint(c.PostForm("ID"))
	s.db.FindInSale(&In)
	Sale := In.Sales
	c.JSON(http.StatusOK, models.Message{Sale: Sale})
}

// DeleteInHandler 删除 In 记录
func (s *Server) DeleteInHandler(c *gin.Context) {
	In := &models.In{}
	if err := c.ShouldBind(In); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("删除 In: %+v\n", In)

	if err := s.db.Delete(In); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) DeleteInSale(c *gin.Context) {
	In := models.In{}
	In.ID = s.Str2Uint(c.PostForm("ID"))
	var Sale models.Sale
	Sale.ID = s.Str2Uint(c.PostForm("SaleId"))
	In.Sales = append(In.Sales, Sale)

	if In.ID == 0 || Sale.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 In 记录
	if err := s.db.DeleteInSale(&In, &Sale); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) AddInSale(c *gin.Context) {
	In := models.In{}
	In.ID = s.Str2Uint(c.PostForm("ID"))
	var Sale models.Sale
	Sale.ID = s.Str2Uint(c.PostForm("SaleId"))

	if In.ID == 0 || Sale.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindById(In.ID, &In)
	s.db.FindById(Sale.ID, &Sale)

	if Sale.AcctName == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	In.Sales = append(In.Sales, Sale)

	// 保存 In 记录
	if err := s.db.Save(In); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddInShouldIns(c *gin.Context) {
	In := models.In{}
	In.ID = s.Str2Uint(c.PostForm("ID"))
	var ShouldIn models.ShouldIn
	ShouldIn.ID = s.Str2Uint(c.PostForm("ShouldInId"))

	if In.ID == 0 || ShouldIn.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindById(In.ID, &In)
	s.db.FindById(ShouldIn.ID, &ShouldIn)

	if ShouldIn.BillReceNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	In.ShouldIns = append(In.ShouldIns, ShouldIn)

	// 保存 In 记录
	if err := s.db.Save(In); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddInSends(c *gin.Context) {
	In := models.In{}
	In.ID = s.Str2Uint(c.PostForm("ID"))
	var Send models.Send
	Send.ID = s.Str2Uint(c.PostForm("SendId"))

	if In.ID == 0 || Send.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindById(In.ID, &In)
	s.db.FindById(Send.ID, &Send)

	if Send.SaleInvNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	In.Sends = append(In.Sends, Send)

	// 保存 In 记录
	if err := s.db.Save(In); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) SaveInHandler(c *gin.Context) {
	In := &models.In{}
	if err := c.ShouldBind(In); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("保存 In: %+v\n", In)

	_, In.FileId, In.FileName = s.SaveFile(c, "file")
	// 保存 In 记录
	if err := s.db.SaveIn(In); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}
