package server

import (
	"exporter/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) DeleteShouldOutPurrec(c *gin.Context) {
	ShouldOut := models.ShouldOut{}
	ShouldOut.ID = s.Str2Uint(c.PostForm("ID"))
	var Purrec models.Purrec
	Purrec.ID = s.Str2Uint(c.PostForm("PurrecID"))
	ShouldOut.Purrecs = append(ShouldOut.Purrecs, Purrec)

	if Purrec.ID == 0 || ShouldOut.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Purrec 记录
	if err := s.db.DeleteShouldOutPurrec(&ShouldOut, &Purrec); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

// FindShouldOutHandler 查询所有 ShouldOut 记录
func (s *Server) FindShouldOutHandler(c *gin.Context) {
	var ShouldOuts []models.ShouldOut
	s.db.FindShouldOuts(&ShouldOuts)
	c.JSON(http.StatusOK, models.Message{ShouldOut: ShouldOuts})
}

func (s *Server) DeleteShouldOutOut(c *gin.Context) {
	ShouldOut := models.ShouldOut{}
	ShouldOut.ID = s.Str2Uint(c.PostForm("ID"))
	var Out models.Out
	Out.ID = s.Str2Uint(c.PostForm("OutID"))
	ShouldOut.Outs = append(ShouldOut.Outs, Out)

	if Out.ID == 0 || ShouldOut.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Out 记录
	if err := s.db.DeleteShouldOutOut(&ShouldOut, &Out); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) FindShouldOutPurrecHandler(c *gin.Context) {
	ShouldOut := models.ShouldOut{}
	ShouldOut.ID = s.Str2Uint(c.PostForm("ID"))
	if ShouldOut.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindShouldOutPurrec(&ShouldOut)
	Purrec := ShouldOut.Purrecs
	c.JSON(http.StatusOK, models.Message{Purrec: Purrec})
}

func (s *Server) FindShouldOutOutHandler(c *gin.Context) {
	ShouldOut := models.ShouldOut{}
	ShouldOut.ID = s.Str2Uint(c.PostForm("ID"))
	if ShouldOut.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindShouldOutOut(&ShouldOut)
	Outs := ShouldOut.Outs
	c.JSON(http.StatusOK, models.Message{Out: Outs})
}

func (s *Server) FindShouldOutBuyHandler(c *gin.Context) {
	ShouldOut := models.ShouldOut{}
	ShouldOut.ID = s.Str2Uint(c.PostForm("ID"))
	s.db.FindShouldOutBuy(&ShouldOut)
	Buy := ShouldOut.Buys
	c.JSON(http.StatusOK, models.Message{Buy: Buy})
}

// DeleteShouldOutHandler 删除 ShouldOut 记录
func (s *Server) DeleteShouldOutHandler(c *gin.Context) {
	ShouldOut := &models.ShouldOut{}
	ShouldOut.ID = s.Str2Uint(c.PostForm("ID"))

	if ShouldOut.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	log.Printf("删除 ShouldOut: %+v\n", ShouldOut)

	if err := s.db.Delete(ShouldOut); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) DeleteShouldOutBuy(c *gin.Context) {
	ShouldOut := models.ShouldOut{}
	ShouldOut.ID = s.Str2Uint(c.PostForm("ID"))
	var Buy models.Buy
	Buy.ID = s.Str2Uint(c.PostForm("BuyID"))
	ShouldOut.Buys = append(ShouldOut.Buys, Buy)

	if ShouldOut.ID == 0 || Buy.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 ShouldOut 记录
	if err := s.db.DeleteShouldOutBuy(&ShouldOut, &Buy); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) AddShouldOutBuy(c *gin.Context) {
	ShouldOut := models.ShouldOut{}
	ShouldOut.ID = s.Str2Uint(c.PostForm("ID"))
	var Buy models.Buy
	Buy.ID = s.Str2Uint(c.PostForm("BuyID"))

	if ShouldOut.ID == 0 || Buy.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(ShouldOut.ID, &ShouldOut)
	s.db.FindByID(Buy.ID, &Buy)

	if Buy.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	ShouldOut.Buys = append(ShouldOut.Buys, Buy)

	// 保存 ShouldOut 记录
	if err := s.db.Save(ShouldOut); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddShouldOutOuts(c *gin.Context) {
	ShouldOut := models.ShouldOut{}
	ShouldOut.ID = s.Str2Uint(c.PostForm("ID"))
	var Out models.Out
	Out.ID = s.Str2Uint(c.PostForm("OutID"))

	if ShouldOut.ID == 0 || Out.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(ShouldOut.ID, &ShouldOut)
	s.db.FindByID(Out.ID, &Out)

	if Out.ReceNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	ShouldOut.Outs = append(ShouldOut.Outs, Out)

	// 保存 ShouldOut 记录
	if err := s.db.Save(ShouldOut); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddShouldOutPurrecs(c *gin.Context) {
	ShouldOut := models.ShouldOut{}
	ShouldOut.ID = s.Str2Uint(c.PostForm("ID"))
	var Purrec models.Purrec
	Purrec.ID = s.Str2Uint(c.PostForm("PurrecID"))

	if ShouldOut.ID == 0 || Purrec.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(ShouldOut.ID, &ShouldOut)
	s.db.FindByID(Purrec.ID, &Purrec)

	if ShouldOut.BillReceNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	ShouldOut.Purrecs = append(ShouldOut.Purrecs, Purrec)

	// 保存 ShouldOut 记录
	if err := s.db.Save(ShouldOut); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) SaveShouldOutHandler(c *gin.Context) {
	ShouldOut := &models.ShouldOut{}
	if err := c.ShouldBind(ShouldOut); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("保存 ShouldOut: %+v\n", ShouldOut)

	_, ShouldOut.FileID, ShouldOut.FileName = s.SaveFile(c, "file")
	// 保存 ShouldOut 记录
	if err := s.db.SaveShouldOut(ShouldOut); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}
