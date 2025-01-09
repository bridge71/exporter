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
	Send.ID = s.Str2Uint(c.PostForm("SendID"))
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
	In.ID = s.Str2Uint(c.PostForm("InID"))
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
	ShouldIn.ID = s.Str2Uint(c.PostForm("ID"))

	if ShouldIn.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
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
	Sale.ID = s.Str2Uint(c.PostForm("SaleID"))
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
	Sale.ID = s.Str2Uint(c.PostForm("SaleID"))

	if ShouldIn.ID == 0 || Sale.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(ShouldIn.ID, &ShouldIn)
	s.db.FindByID(Sale.ID, &Sale)

	if Sale.ID == 0 {
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
	In.ID = s.Str2Uint(c.PostForm("InID"))

	if ShouldIn.ID == 0 || In.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(ShouldIn.ID, &ShouldIn)
	s.db.FindByID(In.ID, &In)

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
	Send.ID = s.Str2Uint(c.PostForm("SendID"))

	if ShouldIn.ID == 0 || Send.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(ShouldIn.ID, &ShouldIn)
	s.db.FindByID(Send.ID, &Send)

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

	ShouldIn.ID = s.Str2Uint(c.PostForm("ID"))
	if ShouldIn.ID != 0 {
		s.db.FindByID(ShouldIn.ID, ShouldIn)
	}
	ShouldIn.MerchantID = s.Str2Uint(c.PostForm("MerchantID"))
	ShouldIn.AcctID = s.Str2Uint(c.PostForm("AcctID"))
	ShouldIn.BankAccountID = s.Str2Uint(c.PostForm("BankAccountID"))
	ShouldIn.AcctBankID = s.Str2Uint(c.PostForm("AcctBankID"))

	// 验证必填字段
	if ShouldIn.MerchantID == 0 || ShouldIn.AcctID == 0 || ShouldIn.BankAccountID == 0 || ShouldIn.AcctBankID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败，必填字段缺失",
		})
		return
	}

	// 绑定嵌套结构体
	ShouldIn.Merchant.ID = ShouldIn.MerchantID
	ShouldIn.Acct.ID = ShouldIn.AcctID
	ShouldIn.BankAccount.ID = ShouldIn.BankAccountID
	ShouldIn.AcctBank.ID = ShouldIn.AcctBankID

	// 绑定其他字段
	ShouldIn.BillReceNum = c.PostForm("BillReceNum")
	ShouldIn.DocDate = c.PostForm("DocDate")
	ShouldIn.ExpReceDate = c.PostForm("ExpReceDate")
	ShouldIn.FinaDocType = c.PostForm("FinaDocType")
	ShouldIn.FinaDocStatus = c.PostForm("FinaDocStatus")
	ShouldIn.TotAmt = s.Str2Uint(c.PostForm("TotAmt"))
	ShouldIn.Currency = c.PostForm("Currency")
	ShouldIn.Notes = c.PostForm("Notes")
	ShouldIn.FileID = s.Str2Uint(c.PostForm("FileID"))
	ShouldIn.FileName = c.PostForm("FileName")

	log.Printf("保存 ShouldIn: %+v\n", ShouldIn)

	_, ShouldIn.FileID, ShouldIn.FileName = s.SaveFile(c, "file")
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
