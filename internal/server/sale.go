package server

import (
	"exporter/internal/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FindSaleHandler 查询所有 Sale 记录
func (s *Server) FindSaleHandler(c *gin.Context) {
	var Sales []models.Sale
	s.db.FindSales(&Sales)
	c.JSON(http.StatusOK, models.Message{Sale: Sales})
}

func (s *Server) DeleteSaleIn(c *gin.Context) {
	Sale := models.Sale{}
	Sale.ID = s.Str2Uint(c.PostForm("ID"))
	var In models.In
	In.ID = s.Str2Uint(c.PostForm("InID"))
	Sale.Ins = append(Sale.Ins, In)

	if Sale.ID == 0 || In.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Sale 记录
	if err := s.db.DeleteSaleIns(&Sale, &In); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) DeleteSaleShouldIn(c *gin.Context) {
	Sale := models.Sale{}
	Sale.ID = s.Str2Uint(c.PostForm("ID"))
	var ShouldIn models.ShouldIn
	ShouldIn.ID = s.Str2Uint(c.PostForm("ShouldInID"))
	Sale.ShouldIns = append(Sale.ShouldIns, ShouldIn)

	if Sale.ID == 0 || ShouldIn.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Sale 记录
	if err := s.db.DeleteSaleShouldIns(&Sale, &ShouldIn); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) AddSaleShouldIn(c *gin.Context) {
	Sale := models.Sale{}
	Sale.ID = s.Str2Uint(c.PostForm("ID"))
	var ShouldIn models.ShouldIn
	ShouldIn.ID = s.Str2Uint(c.PostForm("ShouldInID"))

	if Sale.ID == 0 || ShouldIn.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(ShouldIn.ID, &ShouldIn)
	s.db.FindByID(Sale.ID, &Sale)
	Sale.ShouldIns = append(Sale.ShouldIns, ShouldIn)

	if ShouldIn.BillReceNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Sale 记录
	if err := s.db.Save(Sale); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddSaleIn(c *gin.Context) {
	Sale := models.Sale{}
	Sale.ID = s.Str2Uint(c.PostForm("ID"))
	var In models.In
	In.ID = s.Str2Uint(c.PostForm("InID"))

	if Sale.ID == 0 || In.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(In.ID, &In)
	s.db.FindByID(Sale.ID, &Sale)
	Sale.Ins = append(Sale.Ins, In)

	if In.ReceNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Sale 记录
	if err := s.db.Save(Sale); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) FindSaleShouldInHandler(c *gin.Context) {
	Sale := models.Sale{}
	Sale.ID = s.Str2Uint(c.PostForm("ID"))
	if Sale.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindSaleShouldIns(&Sale)
	ShouldIns := Sale.ShouldIns
	c.JSON(http.StatusOK, models.Message{ShouldIn: ShouldIns})
}

func (s *Server) FindSaleInHandler(c *gin.Context) {
	Sale := models.Sale{}
	Sale.ID = s.Str2Uint(c.PostForm("ID"))
	if Sale.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindSaleIns(&Sale)
	Ins := Sale.Ins
	c.JSON(http.StatusOK, models.Message{In: Ins})
}

func (s *Server) FindSalePrdtInfoHandler(c *gin.Context) {
	Sale := models.Sale{}
	Sale.ID = s.Str2Uint(c.PostForm("ID"))
	if Sale.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindSalePrdtInfo(&Sale)
	PrdtInfos := Sale.PrdtInfos
	c.JSON(http.StatusOK, models.Message{PrdtInfo: PrdtInfos})
}

func (s *Server) FindSaleSendHandler(c *gin.Context) {
	Sale := models.Sale{}
	Sale.ID = s.Str2Uint(c.PostForm("ID"))
	s.db.FindSaleSend(&Sale)
	sends := Sale.Sends
	c.JSON(http.StatusOK, models.Message{Send: sends})
}

// DeleteSaleHandler 删除 Sale 记录
func (s *Server) DeleteSaleHandler(c *gin.Context) {
	Sale := &models.Sale{}
	if err := c.ShouldBind(Sale); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("删除 Sale: %+v\n", Sale)

	if err := s.db.Delete(Sale); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) DeleteSalePrdtInfo(c *gin.Context) {
	// Sale := models.Sale{}
	// if err := c.ShouldBind(&Sale); err != nil {
	// 	c.JSON(http.StatusBadRequest, models.Message{
	// 		RetMessage: err.Error(),
	// 		// RetMessage: "绑定数据失败",
	// 	})
	// 	return
	// }
	// PrdtInfo := models.PrdtInfo{}
	// PrdtInfo.ID = Sale.PrdtInfoID
	Sale := models.Sale{}
	Sale.ID = s.Str2Uint(c.PostForm("ID"))
	var PrdtInfo models.PrdtInfo
	PrdtInfo.ID = s.Str2Uint(c.PostForm("PrdtInfoID"))
	if Sale.ID == 0 || PrdtInfo.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	Sale.PrdtInfos = append(Sale.PrdtInfos, PrdtInfo)

	if err := s.db.DeleteSalePrdtInfo(&Sale, &PrdtInfo); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) DeleteSaleSend(c *gin.Context) {
	Sale := models.Sale{}
	Sale.ID = s.Str2Uint(c.PostForm("ID"))
	var send models.Send
	send.ID = s.Str2Uint(c.PostForm("SendID"))
	Sale.Sends = append(Sale.Sends, send)

	if Sale.ID == 0 || send.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Sale 记录
	if err := s.db.DeleteSaleSend(&Sale, &send); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) AddSalePrdtInfo(c *gin.Context) {
	Sale := models.Sale{}
	Sale.ID = s.Str2Uint(c.PostForm("ID"))
	var PrdtInfo models.PrdtInfo
	PrdtInfo.ID = s.Str2Uint(c.PostForm("PrdtInfoID"))

	if Sale.ID == 0 || PrdtInfo.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(PrdtInfo.ID, &PrdtInfo)
	s.db.FindByID(Sale.ID, &Sale)
	if PrdtInfo.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	Sale.PrdtInfos = append(Sale.PrdtInfos, PrdtInfo)

	// 保存 Sale 记录
	if err := s.db.Save(Sale); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddSaleSend(c *gin.Context) {
	Sale := models.Sale{}
	Sale.ID = s.Str2Uint(c.PostForm("ID"))
	var Send models.Send
	Send.ID = s.Str2Uint(c.PostForm("SendID"))

	if Sale.ID == 0 || Send.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(Send.ID, &Send)
	s.db.FindByID(Sale.ID, &Sale)
	Sale.Sends = append(Sale.Sends, Send)

	if Send.SaleInvNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Sale 记录
	if err := s.db.Save(Sale); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) SaveSaleHandler(c *gin.Context) {
	Sale := &models.Sale{}
	if err := c.ShouldBind(Sale); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("保存 Sale: %+v\n", Sale)

	var Sales []models.Sale
	s.db.FindSales(&Sales)
	l := len(Sales)
	idExitSlice := make([]uint, 0)
	if Sale.ID > 0 {
		for j := 0; j < l; j++ {
			if Sales[j].ID == Sale.ID {
				l2 := len(Sales[j].DocReq)
				for k := 0; k < l2; k++ {
					idExitSlice = append(idExitSlice, Sales[j].DocReq[k].DocReqID)
				}
				break

			}
		}
	}
	i := 0

	idSlice := make([]uint, 0)
	for {
		idStr := c.PostForm("DocReq[" + strconv.Itoa(i) + "][DocReqID]")
		if idStr == "" {
			break // 如果没有更多数据，退出循环
		}
		id_ := s.Str2Uint(idStr)
		if id_ == 0 {
			break
		}
		idSlice = append(idSlice, id_)
		doc := []models.DocReq{}
		s.db.FindDocReqByID(id_, &doc)
		Sale.DocReq = append(Sale.DocReq, doc[0])
		i++
	}

	// 找出第二个切片中有但第一个切片中没有的元素
	diffInSecond := make([]uint, 0)
	// 找出第一个切片中有但第二个切片中没有的元素
	diffInFirst := make([]uint, 0)

	// 使用 map 记录第一个切片的元素
	existMap := make(map[uint]bool)
	for _, id := range idExitSlice {
		existMap[id] = true
	}

	// 遍历第二个切片，找出第一个切片中没有的元素
	for _, id := range idSlice {
		if !existMap[id] {
			diffInSecond = append(diffInSecond, id)
		}
	}

	// 使用 map 记录第二个切片的元素
	existMap = make(map[uint]bool)
	for _, id := range idSlice {
		existMap[id] = true
	}

	// 遍历第一个切片，找出第二个切片中没有的元素
	for _, id := range idExitSlice {
		if !existMap[id] {
			diffInFirst = append(diffInFirst, id)
		}
	}

	shouldDelDocReq := []models.DocReq{}
	for i := 0; i < len(diffInFirst); i++ {
		shouldDelDocReq = append(shouldDelDocReq, models.DocReq{DocReqID: diffInFirst[i]})
	}
	// s.db.DeleteSaleDocReq(Sale, &shouldDelDocReq)
	var err error
	err, Sale.FileID, Sale.FileName = s.SaveFile(c, "file")
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: err.Error(),
		})
	}
	// 保存 Sale 记录
	if err := s.db.TransSale(Sale, &shouldDelDocReq); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}
