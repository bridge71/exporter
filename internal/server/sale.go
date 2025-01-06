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

	s.db.FindSalesPrdtInfo(&Sales)
	c.JSON(http.StatusOK, models.Message{Sale: Sales})
}

func (s *Server) FindSalePrdtInfoHandler(c *gin.Context) {
	var Sale models.Sale

	s.db.FindSalePrdtInfo(&Sale)
	PrdtInfos := Sale.PrdtInfos
	c.JSON(http.StatusOK, models.Message{PrdtInfo: PrdtInfos})
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
	Sale := &models.Sale{}
	if err := c.ShouldBind(Sale); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}
	var PrdtInfo models.PrdtInfo
	s.db.FindById(Sale.PrdtInfoId, &PrdtInfo)
	s.db.FindById(Sale.ID, Sale)
	Sale.PrdtInfos = append(Sale.PrdtInfos, PrdtInfo)

	// 保存 Sale 记录
	if err := s.db.DeleteSalePrdtInfo(Sale, &PrdtInfo); err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{
			RetMessage: "保存失败",
		})
		return
	}

	c.JSON(http.StatusOK, models.Message{
		RetMessage: "保存成功",
	})
}

func (s *Server) AddSalePrdtInfo(c *gin.Context) {
	Sale := &models.Sale{}
	if err := c.ShouldBind(Sale); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "绑定数据失败",
		})
		return
	}
	log.Printf("%d   %d\n", Sale.ID, Sale.PrdtInfoId)
	var PrdtInfo models.PrdtInfo
	s.db.FindById(Sale.PrdtInfoId, &PrdtInfo)
	s.db.FindById(Sale.ID, Sale)
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
	s.db.FindSalesPrdtInfo(&Sales)
	l := len(Sales)
	idExitSlice := make([]uint, 0)
	if Sale.ID > 0 {
		for j := 0; j < l; j++ {
			if Sales[j].ID == Sale.ID {
				l2 := len(Sales[j].DocReq)
				for k := 0; k < l2; k++ {
					idExitSlice = append(idExitSlice, Sales[j].DocReq[k].DocReqId)
				}
				break

			}
		}
	}
	i := 0

	idSlice := make([]uint, 0)
	for {
		idStr := c.PostForm("DocReq[" + strconv.Itoa(i) + "][DocReqId]")
		if idStr == "" {
			break // 如果没有更多数据，退出循环
		}
		id_ := s.Str2Uint(idStr)
		if id_ == 0 {
			break
		}
		idSlice = append(idSlice, id_)
		doc := []models.DocReq{}
		s.db.FindDocReqById(id_, &doc)
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
		shouldDelDocReq = append(shouldDelDocReq, models.DocReq{DocReqId: diffInFirst[i]})
	}
	s.db.DeleteSaleDocReq(Sale, &shouldDelDocReq)
	var err error
	err, Sale.FileId, Sale.FileName = s.SaveFile(c)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: err.Error(),
		})
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
