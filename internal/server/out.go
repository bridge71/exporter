package server

import (
	"exporter/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindOutHandler 查询所有 Out 记录
func (s *Server) FindOutHandler(c *gin.Context) {
	var Outs []models.Out
	s.db.FindOuts(&Outs)
	c.JSON(http.StatusOK, models.Message{Out: Outs})
}

func (s *Server) DeleteOutShouldOut(c *gin.Context) {
	Out := models.Out{}
	Out.ID = s.Str2Uint(c.PostForm("ID"))
	var ShouldOut models.ShouldOut
	ShouldOut.ID = s.Str2Uint(c.PostForm("ShouldOutID"))
	Out.ShouldOuts = append(Out.ShouldOuts, ShouldOut)

	if Out.ID == 0 || ShouldOut.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Out 记录
	if err := s.db.DeleteOutShouldOut(&Out, &ShouldOut); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) DeleteOutPurrec(c *gin.Context) {
	Out := models.Out{}
	Out.ID = s.Str2Uint(c.PostForm("ID"))
	var Purrec models.Purrec
	Purrec.ID = s.Str2Uint(c.PostForm("PurrecID"))
	Out.Purrecs = append(Out.Purrecs, Purrec)

	if Out.ID == 0 || Purrec.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Out 记录
	if err := s.db.DeleteOutPurrec(&Out, &Purrec); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) FindOutPurrecHandler(c *gin.Context) {
	Out := models.Out{}
	Out.ID = s.Str2Uint(c.PostForm("ID"))
	if Out.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindOutPurrec(&Out)
	Purrec := Out.Purrecs
	c.JSON(http.StatusOK, models.Message{Purrec: Purrec})
}

func (s *Server) FindOutShouldOutHandler(c *gin.Context) {
	Out := models.Out{}
	Out.ID = s.Str2Uint(c.PostForm("ID"))
	if Out.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindOutShouldOut(&Out)
	ShouldOut := Out.ShouldOuts
	c.JSON(http.StatusOK, models.Message{ShouldOut: ShouldOut})
}

func (s *Server) FindOutBuyHandler(c *gin.Context) {
	Out := models.Out{}
	Out.ID = s.Str2Uint(c.PostForm("ID"))
	s.db.FindOutBuy(&Out)
	Buy := Out.Buys
	c.JSON(http.StatusOK, models.Message{Buy: Buy})
}

// DeleteOutHandler 删除 Out 记录
func (s *Server) DeleteOutHandler(c *gin.Context) {
	Out := &models.Out{}
	Out.ID = s.Str2Uint(c.PostForm("ID"))
	if Out.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("删除 Out: %+v\n", Out)

	if err := s.db.Delete(Out); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) DeleteOutBuy(c *gin.Context) {
	Out := models.Out{}
	Out.ID = s.Str2Uint(c.PostForm("ID"))
	var Buy models.Buy
	Buy.ID = s.Str2Uint(c.PostForm("BuyID"))
	Out.Buys = append(Out.Buys, Buy)

	if Out.ID == 0 || Buy.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Out 记录
	if err := s.db.DeleteOutBuy(&Out, &Buy); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) AddOutBuy(c *gin.Context) {
	Out := models.Out{}
	Out.ID = s.Str2Uint(c.PostForm("ID"))
	var Buy models.Buy
	Buy.ID = s.Str2Uint(c.PostForm("BuyID"))

	if Out.ID == 0 || Buy.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(Out.ID, &Out)
	s.db.FindByID(Buy.ID, &Buy)

	if Buy.OrderNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	Out.Buys = append(Out.Buys, Buy)

	// 保存 Out 记录
	if err := s.db.Save(Out); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddOutShouldOuts(c *gin.Context) {
	Out := models.Out{}
	Out.ID = s.Str2Uint(c.PostForm("ID"))
	var ShouldOut models.ShouldOut
	ShouldOut.ID = s.Str2Uint(c.PostForm("ShouldOutID"))

	if Out.ID == 0 || ShouldOut.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(Out.ID, &Out)
	s.db.FindByID(ShouldOut.ID, &ShouldOut)

	if ShouldOut.BillReceNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	Out.ShouldOuts = append(Out.ShouldOuts, ShouldOut)

	// 保存 Out 记录
	if err := s.db.Save(Out); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddOutPurrecs(c *gin.Context) {
	Out := models.Out{}
	Out.ID = s.Str2Uint(c.PostForm("ID"))
	var Purrec models.Purrec
	Purrec.ID = s.Str2Uint(c.PostForm("PurrecID"))

	if Out.ID == 0 || Purrec.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(Out.ID, &Out)
	s.db.FindByID(Purrec.ID, &Purrec)

	if Purrec.SaleInvNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	Out.Purrecs = append(Out.Purrecs, Purrec)

	// 保存 Out 记录
	if err := s.db.Save(Out); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) SaveOutHandler(c *gin.Context) {
	Out := &models.Out{}

	// 如果 ID 存在，查找已有记录
	Out.ID = s.Str2Uint(c.PostForm("ID"))
	if Out.ID != 0 {
		s.db.FindByID(Out.ID, Out)
	}

	// 绑定必填字段
	Out.MerchantID = s.Str2Uint(c.PostForm("MerchantID"))
	Out.AcctID = s.Str2Uint(c.PostForm("AcctID"))
	Out.BankAccountID = s.Str2Uint(c.PostForm("BankAccountID"))
	Out.AcctBankID = s.Str2Uint(c.PostForm("AcctBankID"))

	// 验证必填字段
	if Out.MerchantID == 0 || Out.AcctID == 0 || Out.BankAccountID == 0 || Out.AcctBankID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败，必填字段缺失",
		})
		return
	}

	// 绑定嵌套结构体
	Out.Merchant.ID = Out.MerchantID
	Out.Acct.ID = Out.AcctID
	Out.BankAccount.ID = Out.BankAccountID
	Out.AcctBank.ID = Out.AcctBankID

	// 绑定其他字段
	Out.ReceNum = c.PostForm("ReceNum")
	Out.RealReceDate = c.PostForm("RealReceDate")
	Out.ExpReceDate = c.PostForm("ExpReceDate")
	Out.FinaDocType = c.PostForm("FinaDocType")
	Out.FinaDocStatus = c.PostForm("FinaDocStatus")
	Out.TotAmt = s.Str2Uint(c.PostForm("TotAmt"))
	Out.Currency = c.PostForm("Currency")
	Out.Notes = c.PostForm("Notes")
	Out.FileID = s.Str2Uint(c.PostForm("FileID"))
	Out.FileName = c.PostForm("FileName")

	log.Printf("保存 Out: %+v\n", Out)

	_, Out.FileID, Out.FileName = s.SaveFile(c, "file")
	// 保存 Out 记录
	if err := s.db.SaveOut(Out); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}
