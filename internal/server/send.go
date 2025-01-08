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

func (s *Server) DeleteSendIn(c *gin.Context) {
	Send := models.Send{}
	Send.ID = s.Str2Uint(c.PostForm("ID"))
	var In models.In
	In.ID = s.Str2Uint(c.PostForm("InId"))
	Send.Ins = append(Send.Ins, In)

	if Send.ID == 0 || In.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Send 记录
	if err := s.db.DeleteSendIns(&Send, &In); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) DeleteSendShouldIn(c *gin.Context) {
	Send := models.Send{}
	Send.ID = s.Str2Uint(c.PostForm("ID"))
	var ShouldIn models.ShouldIn
	ShouldIn.ID = s.Str2Uint(c.PostForm("ShouldInId"))
	Send.ShouldIns = append(Send.ShouldIns, ShouldIn)

	if Send.ID == 0 || ShouldIn.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Send 记录
	if err := s.db.DeleteSendShouldIns(&Send, &ShouldIn); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) FindSendShouldInHandler(c *gin.Context) {
	Send := models.Send{}
	Send.ID = s.Str2Uint(c.PostForm("ID"))
	if Send.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindSendShouldIns(&Send)
	ShouldIns := Send.ShouldIns
	c.JSON(http.StatusOK, models.Message{ShouldIn: ShouldIns})
}

func (s *Server) FindSendInHandler(c *gin.Context) {
	Send := models.Send{}
	Send.ID = s.Str2Uint(c.PostForm("ID"))
	if Send.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindSendIns(&Send)
	Ins := Send.Ins
	c.JSON(http.StatusOK, models.Message{In: Ins})
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

func (s *Server) FindSendLoadingInfoHandler(c *gin.Context) {
	Send := models.Send{}
	Send.ID = s.Str2Uint(c.PostForm("ID"))
	if Send.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindSendLoadingInfo(&Send)
	LoadingInfo := Send.LoadingInfos
	c.JSON(http.StatusOK, models.Message{LoadingInfo: LoadingInfo})
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

func (s *Server) AddSendLoadingInfo(c *gin.Context) {
	Send := models.Send{}
	Send.ID = s.Str2Uint(c.PostForm("ID"))
	var LoadingInfo models.LoadingInfo
	LoadingInfo.ID = s.Str2Uint(c.PostForm("LoadingInfoId"))

	if Send.ID == 0 || LoadingInfo.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindById(LoadingInfo.ID, &LoadingInfo)

	if LoadingInfo.PackSpec == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindById(Send.ID, &Send)
	log.Printf("%d\n", LoadingInfo.ID)
	Send.LoadingInfos = append(Send.LoadingInfos, LoadingInfo)

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

	if PrdtInfo.CatEngName == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
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
	Sale.ID = s.Str2Uint(c.PostForm("SaleId"))

	if Send.ID == 0 || Sale.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindById(Send.ID, &Send)
	s.db.FindById(Sale.ID, &Sale)

	if Sale.AcctName == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
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

func (s *Server) AddSendIns(c *gin.Context) {
	Send := models.Send{}
	Send.ID = s.Str2Uint(c.PostForm("ID"))
	var In models.In
	In.ID = s.Str2Uint(c.PostForm("InId"))

	if Send.ID == 0 || In.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindById(Send.ID, &Send)
	s.db.FindById(In.ID, &In)

	if In.ReceNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	Send.Ins = append(Send.Ins, In)

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

func (s *Server) AddSendShouldIns(c *gin.Context) {
	Send := models.Send{}
	Send.ID = s.Str2Uint(c.PostForm("ID"))
	var ShouldIn models.ShouldIn
	ShouldIn.ID = s.Str2Uint(c.PostForm("ShouldInId"))

	if Send.ID == 0 || ShouldIn.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindById(Send.ID, &Send)
	s.db.FindById(ShouldIn.ID, &ShouldIn)

	if ShouldIn.BillReceNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	Send.ShouldIns = append(Send.ShouldIns, ShouldIn)

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
	if err := s.db.SaveSend(Send); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}
