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
	ShouldOut.ID = s.Str2Uint(c.PostForm("ShouldOutId"))
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
	Purrec.ID = s.Str2Uint(c.PostForm("PurrecId"))
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
	if err := c.ShouldBind(Out); err != nil {
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
	Buy.ID = s.Str2Uint(c.PostForm("BuyId"))
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
	Buy.ID = s.Str2Uint(c.PostForm("BuyId"))

	if Out.ID == 0 || Buy.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindById(Out.ID, &Out)
	s.db.FindById(Buy.ID, &Buy)

	if Buy.AcctName == "" {
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
	ShouldOut.ID = s.Str2Uint(c.PostForm("ShouldOutId"))

	if Out.ID == 0 || ShouldOut.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindById(Out.ID, &Out)
	s.db.FindById(ShouldOut.ID, &ShouldOut)

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
	Purrec.ID = s.Str2Uint(c.PostForm("PurrecId"))

	if Out.ID == 0 || Purrec.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindById(Out.ID, &Out)
	s.db.FindById(Purrec.ID, &Purrec)

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
	if err := c.ShouldBind(Out); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("保存 Out: %+v\n", Out)

	_, Out.FileId, Out.FileName = s.SaveFile(c, "file")
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
