package server

import (
	"exporter/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindPurrecHandler 查询所有 Purrec 记录
func (s *Server) FindPurrecHandler(c *gin.Context) {
	var Purrecs []models.Purrec
	s.db.FindPurrecs(&Purrecs)
	c.JSON(http.StatusOK, models.Message{Purrec: Purrecs})
}

func (s *Server) DeletePurrecOut(c *gin.Context) {
	Purrec := models.Purrec{}
	Purrec.ID = s.Str2Uint(c.PostForm("ID"))
	var Out models.Out
	Out.ID = s.Str2Uint(c.PostForm("OutID"))
	Purrec.Outs = append(Purrec.Outs, Out)

	if Purrec.ID == 0 || Out.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Purrec 记录
	if err := s.db.DeletePurrecOuts(&Purrec, &Out); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) DeletePurrecShouldOut(c *gin.Context) {
	Purrec := models.Purrec{}
	Purrec.ID = s.Str2Uint(c.PostForm("ID"))
	var ShouldOut models.ShouldOut
	ShouldOut.ID = s.Str2Uint(c.PostForm("ShouldOutID"))
	Purrec.ShouldOuts = append(Purrec.ShouldOuts, ShouldOut)

	if Purrec.ID == 0 || ShouldOut.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Purrec 记录
	if err := s.db.DeletePurrecShouldOuts(&Purrec, &ShouldOut); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) FindPurrecShouldOutHandler(c *gin.Context) {
	Purrec := models.Purrec{}
	Purrec.ID = s.Str2Uint(c.PostForm("ID"))
	if Purrec.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindPurrecShouldOuts(&Purrec)
	ShouldOuts := Purrec.ShouldOuts
	c.JSON(http.StatusOK, models.Message{ShouldOut: ShouldOuts})
}

func (s *Server) FindPurrecOutHandler(c *gin.Context) {
	Purrec := models.Purrec{}
	Purrec.ID = s.Str2Uint(c.PostForm("ID"))
	if Purrec.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindPurrecOuts(&Purrec)
	Outs := Purrec.Outs
	c.JSON(http.StatusOK, models.Message{Out: Outs})
}

func (s *Server) FindPurrecPrdtInfoHandler(c *gin.Context) {
	Purrec := models.Purrec{}
	Purrec.ID = s.Str2Uint(c.PostForm("ID"))
	if Purrec.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindPurrecPrdtInfo(&Purrec)
	PrdtInfos := Purrec.PrdtInfos
	c.JSON(http.StatusOK, models.Message{PrdtInfo: PrdtInfos})
}

func (s *Server) FindPurrecLoadingInfoHandler(c *gin.Context) {
	Purrec := models.Purrec{}
	Purrec.ID = s.Str2Uint(c.PostForm("ID"))
	if Purrec.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindPurrecLoadingInfo(&Purrec)
	LoadingInfo := Purrec.LoadingInfos
	c.JSON(http.StatusOK, models.Message{LoadingInfo: LoadingInfo})
}

func (s *Server) FindPurrecBuyHandler(c *gin.Context) {
	Purrec := models.Purrec{}
	Purrec.ID = s.Str2Uint(c.PostForm("ID"))
	s.db.FindPurrecBuy(&Purrec)
	Buy := Purrec.Buys
	c.JSON(http.StatusOK, models.Message{Buy: Buy})
}

// DeletePurrecHandler 删除 Purrec 记录
func (s *Server) DeletePurrecHandler(c *gin.Context) {
	Purrec := &models.Purrec{}
	if err := c.ShouldBind(Purrec); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("删除 Purrec: %+v\n", Purrec)

	if err := s.db.Delete(Purrec); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) DeletePurrecPrdtInfo(c *gin.Context) {
	Purrec := models.Purrec{}
	Purrec.ID = s.Str2Uint(c.PostForm("ID"))
	var PrdtInfo models.PrdtInfo
	PrdtInfo.ID = s.Str2Uint(c.PostForm("PrdtInfoID"))
	if Purrec.ID == 0 || PrdtInfo.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	Purrec.PrdtInfos = append(Purrec.PrdtInfos, PrdtInfo)

	if err := s.db.DeletePurrecPrdtInfo(&Purrec, &PrdtInfo); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) DeletePurrecBuy(c *gin.Context) {
	Purrec := models.Purrec{}
	Purrec.ID = s.Str2Uint(c.PostForm("ID"))
	var Buy models.Buy
	Buy.ID = s.Str2Uint(c.PostForm("BuyID"))
	Purrec.Buys = append(Purrec.Buys, Buy)

	if Purrec.ID == 0 || Buy.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Purrec 记录
	if err := s.db.DeletePurrecBuy(&Purrec, &Buy); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) AddPurrecLoadingInfo(c *gin.Context) {
	Purrec := models.Purrec{}
	Purrec.ID = s.Str2Uint(c.PostForm("ID"))
	var LoadingInfo models.LoadingInfo
	LoadingInfo.ID = s.Str2Uint(c.PostForm("LoadingInfoID"))

	if Purrec.ID == 0 || LoadingInfo.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(LoadingInfo.ID, &LoadingInfo)

	if LoadingInfo.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(Purrec.ID, &Purrec)
	log.Printf("%d\n", LoadingInfo.ID)
	Purrec.LoadingInfos = append(Purrec.LoadingInfos, LoadingInfo)

	// 保存 Purrec 记录
	if err := s.db.Save(Purrec); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddPurrecPrdtInfo(c *gin.Context) {
	Purrec := models.Purrec{}
	Purrec.ID = s.Str2Uint(c.PostForm("ID"))
	var PrdtInfo models.PrdtInfo
	PrdtInfo.ID = s.Str2Uint(c.PostForm("PrdtInfoID"))

	if Purrec.ID == 0 || PrdtInfo.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(PrdtInfo.ID, &PrdtInfo)

	if PrdtInfo.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(Purrec.ID, &Purrec)
	log.Printf("%d\n", PrdtInfo.ID)
	Purrec.PrdtInfos = append(Purrec.PrdtInfos, PrdtInfo)

	// 保存 Purrec 记录
	if err := s.db.Save(Purrec); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddPurrecBuy(c *gin.Context) {
	Purrec := models.Purrec{}
	Purrec.ID = s.Str2Uint(c.PostForm("ID"))
	var Buy models.Buy
	Buy.ID = s.Str2Uint(c.PostForm("BuyID"))

	if Purrec.ID == 0 || Buy.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(Purrec.ID, &Purrec)
	s.db.FindByID(Buy.ID, &Buy)

	if Buy.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	Purrec.Buys = append(Purrec.Buys, Buy)

	// 保存 Purrec 记录
	if err := s.db.Save(Purrec); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddPurrecOuts(c *gin.Context) {
	Purrec := models.Purrec{}
	Purrec.ID = s.Str2Uint(c.PostForm("ID"))
	var Out models.Out
	Out.ID = s.Str2Uint(c.PostForm("OutID"))

	if Purrec.ID == 0 || Out.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(Purrec.ID, &Purrec)
	s.db.FindByID(Out.ID, &Out)

	if Out.ReceNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	Purrec.Outs = append(Purrec.Outs, Out)

	// 保存 Purrec 记录
	if err := s.db.Save(Purrec); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddPurrecShouldOuts(c *gin.Context) {
	Purrec := models.Purrec{}
	Purrec.ID = s.Str2Uint(c.PostForm("ID"))
	var ShouldOut models.ShouldOut
	ShouldOut.ID = s.Str2Uint(c.PostForm("ShouldOutID"))

	if Purrec.ID == 0 || ShouldOut.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(Purrec.ID, &Purrec)
	s.db.FindByID(ShouldOut.ID, &ShouldOut)

	if ShouldOut.BillReceNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	Purrec.ShouldOuts = append(Purrec.ShouldOuts, ShouldOut)

	// 保存 Purrec 记录
	if err := s.db.Save(Purrec); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) SavePurrecHandler(c *gin.Context) {
	Purrec := &models.Purrec{}
	if err := c.ShouldBind(Purrec); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("保存 Purrec: %+v\n", Purrec)
	var err error

	err, Purrec.File1ID, Purrec.File1Name = s.SaveFile(c, "file1")
	err, Purrec.File2ID, Purrec.File2Name = s.SaveFile(c, "file2")
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: err.Error(),
		})
	}
	// 保存 Purrec 记录
	if err := s.db.SavePurrec(Purrec); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}
