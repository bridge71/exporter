package server

import (
	"exporter/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) FindEmplHandler(c *gin.Context) {
	Empl := []models.Empl{}
	s.db.FindEmpl(&Empl)
	c.JSON(http.StatusOK, models.Message{
		Empl: Empl,
	})
}

func (s *Server) SaveEmplHandler(c *gin.Context) {
	Empl := &models.Empl{}
	err := c.ShouldBind(Empl)
	log.Printf("%v\n", *Empl)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: err.Error(),
			// RetMessage: "error in bind of Empl",
		})
		return
	}
	log.Printf("%v\n", Empl)

	err, Empl.FileId, Empl.FileName = s.SaveFile(c, "file")
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to save file",
		})
	}
	err = s.db.SaveEmpl(Empl)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to save Empl",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved Empl successfully",
		})
	}
}

func (s *Server) DeleteEmplHandler(c *gin.Context) {
	Empl := &models.Empl{}
	err := c.ShouldBind(Empl)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of Empl",
		})
		return
	}
	log.Printf("%v\n", Empl)

	err = s.db.DeleteEmpl(Empl)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to delete Empl",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete Empl successfully",
		})
	}
}
