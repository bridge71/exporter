package server

import (
	"exporter/internal/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) FindCustHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	s.db.FindCust(custs)
	c.JSON(http.StatusOK, models.Message{Cust: *custs})
}

func (s *Server) DeleteCustHandler(c *gin.Context) {
	cust := &models.Cust{}
	cust.ID = s.Str2Uint(c.PostForm("ID"))
	log.Printf("%v\n", cust)

	err := s.db.DeleteCust(cust)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to delete cust",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete cust successfully",
		})
	}
}

func (s *Server) FindCustByIDHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	mercID := c.PostForm("MercID")
	idStr, _ := strconv.Atoi(mercID)
	id := uint(idStr)
	s.db.FindCustByID(custs, id)
	c.JSON(http.StatusOK, models.Message{Cust: *custs})
}

func (s *Server) SaveCustHandler(c *gin.Context) {
	cust := &models.Cust{}
	err := c.ShouldBind(cust)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "请选择客商信息",
		})
		return
	}
	log.Printf("%v\n", cust)

	err, cust.FileID, cust.FileName = s.SaveFile(c, "file")
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to save file",
		})
	}
	cust.Merchant.ID = cust.MerchantID
	err = s.db.SaveCust(cust)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to save cust",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved cust successfully",
		})
	}
}
