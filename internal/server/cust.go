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
	err := c.ShouldBind(cust)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of cust",
		})
		return
	}
	log.Printf("%v\n", cust)

	err = s.db.DeleteCust(cust)
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

func (s *Server) FindCustByIdHandler(c *gin.Context) {
	custs := &[]models.Cust{}
	mercId := c.PostForm("MercId")
	idStr, _ := strconv.Atoi(mercId)
	id := uint(idStr)
	s.db.FindCustById(custs, id)
	c.JSON(http.StatusOK, models.Message{Cust: *custs})
}

func (s *Server) SaveCustHandler(c *gin.Context) {
	cust := &models.Cust{}
	err := c.ShouldBind(cust)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of cust",
		})
		return
	}
	log.Printf("%v\n", cust)

	err, cust.FileId, cust.FileName = s.SaveFile(c, "file")
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to save file",
		})
	}
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
