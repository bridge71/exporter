package server

import (
	"exporter/internal/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FindBuyHandler 查询所有 Buy 记录
func (s *Server) FindBuyHandler(c *gin.Context) {
	var Buys []models.Buy
	s.db.FindBuys(&Buys)
	c.JSON(http.StatusOK, models.Message{Buy: Buys})
}

func (s *Server) DeleteBuyOut(c *gin.Context) {
	Buy := models.Buy{}
	Buy.ID = s.Str2Uint(c.PostForm("ID"))
	var Out models.Out
	Out.ID = s.Str2Uint(c.PostForm("OutID"))
	Buy.Outs = append(Buy.Outs, Out)

	if Buy.ID == 0 || Out.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Buy 记录
	if err := s.db.DeleteBuyOuts(&Buy, &Out); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) DeleteBuyShouldOut(c *gin.Context) {
	Buy := models.Buy{}
	Buy.ID = s.Str2Uint(c.PostForm("ID"))
	var ShouldOut models.ShouldOut
	ShouldOut.ID = s.Str2Uint(c.PostForm("ShouldOutID"))
	Buy.ShouldOuts = append(Buy.ShouldOuts, ShouldOut)

	if Buy.ID == 0 || ShouldOut.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Buy 记录
	if err := s.db.DeleteBuyShouldOuts(&Buy, &ShouldOut); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) AddBuyShouldOut(c *gin.Context) {
	Buy := models.Buy{}
	Buy.ID = s.Str2Uint(c.PostForm("ID"))
	var ShouldOut models.ShouldOut
	ShouldOut.ID = s.Str2Uint(c.PostForm("ShouldOutID"))

	if Buy.ID == 0 || ShouldOut.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(ShouldOut.ID, &ShouldOut)
	s.db.FindByID(Buy.ID, &Buy)
	Buy.ShouldOuts = append(Buy.ShouldOuts, ShouldOut)

	if ShouldOut.BillReceNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Buy 记录
	if err := s.db.Save(Buy); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddBuyOut(c *gin.Context) {
	Buy := models.Buy{}
	Buy.ID = s.Str2Uint(c.PostForm("ID"))
	var Out models.Out
	Out.ID = s.Str2Uint(c.PostForm("OutID"))

	if Buy.ID == 0 || Out.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(Out.ID, &Out)
	s.db.FindByID(Buy.ID, &Buy)
	Buy.Outs = append(Buy.Outs, Out)

	if Out.ReceNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Buy 记录
	if err := s.db.Save(Buy); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) FindBuyShouldOutHandler(c *gin.Context) {
	Buy := models.Buy{}
	Buy.ID = s.Str2Uint(c.PostForm("ID"))
	if Buy.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindBuyShouldOuts(&Buy)
	ShouldOuts := Buy.ShouldOuts
	c.JSON(http.StatusOK, models.Message{ShouldOut: ShouldOuts})
}

func (s *Server) FindBuyOutHandler(c *gin.Context) {
	Buy := models.Buy{}
	Buy.ID = s.Str2Uint(c.PostForm("ID"))
	if Buy.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindBuyOuts(&Buy)
	Outs := Buy.Outs
	c.JSON(http.StatusOK, models.Message{Out: Outs})
}

func (s *Server) FindBuyPrdtInfoHandler(c *gin.Context) {
	Buy := models.Buy{}
	Buy.ID = s.Str2Uint(c.PostForm("ID"))
	if Buy.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindBuyPrdtInfo(&Buy)
	PrdtInfos := Buy.PrdtInfos
	c.JSON(http.StatusOK, models.Message{PrdtInfo: PrdtInfos})
}

func (s *Server) FindBuyPurrecHandler(c *gin.Context) {
	Buy := models.Buy{}
	Buy.ID = s.Str2Uint(c.PostForm("ID"))
	s.db.FindBuyPurrec(&Buy)
	Purrecs := Buy.Purrecs
	c.JSON(http.StatusOK, models.Message{Purrec: Purrecs})
}

// DeleteBuyHandler 删除 Buy 记录
func (s *Server) DeleteBuyHandler(c *gin.Context) {
	Buy := &models.Buy{}
	if err := c.ShouldBind(Buy); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("删除 Buy: %+v\n", Buy)

	if err := s.db.Delete(Buy); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) DeleteBuyPrdtInfo(c *gin.Context) {
	// Buy := models.Buy{}
	// if err := c.ShouldBind(&Buy); err != nil {
	// 	c.JSON(http.StatusBadRequest, models.Message{
	// 		RetMessage: err.Error(),
	// 		// RetMessage: "绑定数据失败",
	// 	})
	// 	return
	// }
	// PrdtInfo := models.PrdtInfo{}
	// PrdtInfo.ID = Buy.PrdtInfoID
	Buy := models.Buy{}
	Buy.ID = s.Str2Uint(c.PostForm("ID"))
	var PrdtInfo models.PrdtInfo
	PrdtInfo.ID = s.Str2Uint(c.PostForm("PrdtInfoID"))
	if Buy.ID == 0 || PrdtInfo.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	Buy.PrdtInfos = append(Buy.PrdtInfos, PrdtInfo)

	if err := s.db.DeleteBuyPrdtInfo(&Buy, &PrdtInfo); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) DeleteBuyPurrec(c *gin.Context) {
	Buy := models.Buy{}
	Buy.ID = s.Str2Uint(c.PostForm("ID"))
	var Purrec models.Purrec
	Purrec.ID = s.Str2Uint(c.PostForm("PurrecID"))
	Buy.Purrecs = append(Buy.Purrecs, Purrec)

	if Buy.ID == 0 || Purrec.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Buy 记录
	if err := s.db.DeleteBuyPurrec(&Buy, &Purrec); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "删除成功",
	})
}

func (s *Server) AddBuyPrdtInfo(c *gin.Context) {
	Buy := models.Buy{}
	Buy.ID = s.Str2Uint(c.PostForm("ID"))
	var PrdtInfo models.PrdtInfo
	PrdtInfo.ID = s.Str2Uint(c.PostForm("PrdtInfoID"))

	if Buy.ID == 0 || PrdtInfo.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(PrdtInfo.ID, &PrdtInfo)
	s.db.FindByID(Buy.ID, &Buy)
	if PrdtInfo.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	Buy.PrdtInfos = append(Buy.PrdtInfos, PrdtInfo)

	// 保存 Buy 记录
	if err := s.db.Save(Buy); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddBuyPurrec(c *gin.Context) {
	Buy := models.Buy{}
	Buy.ID = s.Str2Uint(c.PostForm("ID"))
	var Purrec models.Purrec
	Purrec.ID = s.Str2Uint(c.PostForm("PurrecID"))

	if Buy.ID == 0 || Purrec.ID == 0 {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非数字",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	s.db.FindByID(Purrec.ID, &Purrec)
	s.db.FindByID(Buy.ID, &Buy)
	Buy.Purrecs = append(Buy.Purrecs, Purrec)

	if Purrec.SaleInvNum == "" {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: "非法ID",
			// RetMessage: "绑定数据失败",
		})
		return
	}
	// 保存 Buy 记录
	if err := s.db.Save(Buy); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) SaveBuyHandler(c *gin.Context) {
	Buy := &models.Buy{}
	if err := c.ShouldBind(Buy); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}

	log.Printf("保存 Buy: %+v\n", Buy)

	var Buys []models.Buy
	s.db.FindBuys(&Buys)
	l := len(Buys)
	idExitSlice := make([]uint, 0)
	if Buy.ID > 0 {
		for j := 0; j < l; j++ {
			if Buys[j].ID == Buy.ID {
				l2 := len(Buys[j].DocReq)
				for k := 0; k < l2; k++ {
					idExitSlice = append(idExitSlice, Buys[j].DocReq[k].DocReqID)
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
		Buy.DocReq = append(Buy.DocReq, doc[0])
		i++
	}

	// 找出第二个切片中有但第一个切片中没有的元素
	diffOutSecond := make([]uint, 0)
	// 找出第一个切片中有但第二个切片中没有的元素
	diffOutFirst := make([]uint, 0)

	// 使用 map 记录第一个切片的元素
	existMap := make(map[uint]bool)
	for _, id := range idExitSlice {
		existMap[id] = true
	}

	// 遍历第二个切片，找出第一个切片中没有的元素
	for _, id := range idSlice {
		if !existMap[id] {
			diffOutSecond = append(diffOutSecond, id)
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
			diffOutFirst = append(diffOutFirst, id)
		}
	}

	shouldDelDocReq := []models.DocReq{}
	for i := 0; i < len(diffOutFirst); i++ {
		shouldDelDocReq = append(shouldDelDocReq, models.DocReq{DocReqID: diffOutFirst[i]})
	}
	// s.db.DeleteBuyDocReq(Buy, &shouldDelDocReq)
	var err error
	err, Buy.FileID, Buy.FileName = s.SaveFile(c, "file")
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: err.Error(),
		})
	}
	// 保存 Buy 记录
	if err := s.db.TransBuy(Buy, &shouldDelDocReq); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}
