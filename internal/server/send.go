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
	In.ID = s.Str2Uint(c.PostForm("InID"))
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
	ShouldIn.ID = s.Str2Uint(c.PostForm("ShouldInID"))
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
	Send.ID = s.Str2Uint(c.PostForm("ID"))
	if Send.ID == 0 {
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
	PrdtInfo.ID = s.Str2Uint(c.PostForm("PrdtInfoID"))
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
	Sale.ID = s.Str2Uint(c.PostForm("SaleID"))
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
	LoadingInfo.ID = s.Str2Uint(c.PostForm("LoadingInfoID"))

	if Send.ID == 0 || LoadingInfo.ID == 0 {
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
	s.db.FindByID(Send.ID, &Send)
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
	PrdtInfo.ID = s.Str2Uint(c.PostForm("PrdtInfoID"))

	if Send.ID == 0 || PrdtInfo.ID == 0 {
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
	s.db.FindByID(Send.ID, &Send)
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
	Sale.ID = s.Str2Uint(c.PostForm("SaleID"))

	if Send.ID == 0 || Sale.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(Send.ID, &Send)
	s.db.FindByID(Sale.ID, &Sale)

	if Sale.ID == 0 {
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
	In.ID = s.Str2Uint(c.PostForm("InID"))

	if Send.ID == 0 || In.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(Send.ID, &Send)
	s.db.FindByID(In.ID, &In)

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
	ShouldIn.ID = s.Str2Uint(c.PostForm("ShouldInID"))

	if Send.ID == 0 || ShouldIn.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(Send.ID, &Send)
	s.db.FindByID(ShouldIn.ID, &ShouldIn)

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

	Send.ID = s.Str2Uint(c.PostForm("ID"))
	if Send.ID != 0 {
		s.db.FindByID(Send.ID, Send)
	}
	Send.AcctID = s.Str2Uint(c.PostForm("AcctID"))
	Send.PackSpecID = s.Str2Uint(c.PostForm("PackSpecID"))
	Send.PayMentMethodID = s.Str2Uint(c.PostForm("PayMentMethodID"))
	Send.AcctBankID = s.Str2Uint(c.PostForm("AcctBankID"))

	log.Printf("ssss ss %d %d %d\n", Send.Merchant1ID, Send.Merchant2ID, Send.Merchant3ID)
	// 验证必填字段
	if Send.AcctID == 0 || Send.PackSpecID == 0 || Send.PayMentMethodID == 0 || Send.AcctBankID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败，必填字段缺失",
		})
		return
	}

	// 如果 ID 存在，查找已有记录

	// 绑定嵌套结构体
	Send.Acct.ID = Send.AcctID
	Send.PackSpec.ID = Send.PackSpecID
	Send.PayMentMethod.ID = Send.PayMentMethodID
	Send.AcctBank.ID = Send.AcctBankID
	Send.Merchant1ID = s.Str2Uint(c.PostForm("Merchant1ID"))
	Send.Merchant2ID = s.Str2Uint(c.PostForm("Merchant2ID"))
	Send.Merchant3ID = s.Str2Uint(c.PostForm("Merchant3ID"))

	// 绑定其他字段
	Send.SaleInvNum = c.PostForm("SaleInvNum")
	Send.SaleInvDate = c.PostForm("SaleInvDate")
	Send.Merchant1Name = c.PostForm("Merchant1Name")
	Send.Merchant2Name = c.PostForm("Merchant2Name")
	Send.Merchant3Name = c.PostForm("Merchant3Name")
	Send.SrcPlace = c.PostForm("SrcPlace")
	Send.Des = c.PostForm("Des")
	Send.ShipName = c.PostForm("ShipName")
	Send.Voyage = c.PostForm("Voyage")
	Send.TotNum = s.Str2Uint(c.PostForm("TotNum"))
	Send.TotalNetWeight = s.Str2Uint(c.PostForm("TotalNetWeight"))
	Send.UnitMeas1 = c.PostForm("UnitMeas1")
	Send.GrossWt = s.Str2Uint(c.PostForm("GrossWt"))
	Send.UnitMeas2 = c.PostForm("UnitMeas2")
	Send.TotVol = c.PostForm("TotVol")
	Send.UnitMeas3 = c.PostForm("UnitMeas3")
	Send.BillLadNum = c.PostForm("BillLadNum")
	Send.DateOfShip = c.PostForm("DateOfShip")
	Send.Note1 = c.PostForm("Note1")
	Send.Note2 = c.PostForm("Note2")
	Send.File1ID = s.Str2Uint(c.PostForm("File1ID"))
	Send.File2ID = s.Str2Uint(c.PostForm("File2ID"))
	Send.File1Name = c.PostForm("File1Name")
	Send.File2Name = c.PostForm("File2Name")

	// log.Printf("保存 Send: %+v\n", Send)
	log.Printf("保存 %s\n", Send.Merchant1Name)
	var err error

	err, Send.File1ID, Send.File1Name = s.SaveFile(c, "file1")
	err, Send.File2ID, Send.File2Name = s.SaveFile(c, "file2")
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
