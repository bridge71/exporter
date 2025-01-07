package server

import (
	"exporter/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) FindPackSpecHandler(c *gin.Context) {
	PackSpec := []models.PackSpec{}
	s.db.FindPackSpec(&PackSpec)
	c.JSON(http.StatusOK, models.Message{
		PackSpec: PackSpec,
	})
}

func (s *Server) SavePackSpecHandler(c *gin.Context) {
	PackSpec := &models.PackSpec{}
	err := c.ShouldBind(PackSpec)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of PackSpec",
		})
		return
	}
	log.Printf("%v\n", PackSpec)

	err, PackSpec.FileId, PackSpec.FileName = s.SaveFile(c, "file")
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to save file",
		})
	}
	err = s.db.SavePackSpec(PackSpec)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to save PackSpec",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved PackSpec successfully",
		})
	}
}

func (s *Server) DeletePackSpecHandler(c *gin.Context) {
	PackSpec := &models.PackSpec{}
	err := c.ShouldBind(PackSpec)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of PackSpec",
		})
		return
	}
	log.Printf("%v\n", PackSpec)

	err = s.db.DeletePackSpec(PackSpec)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to delete PackSpec",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete PackSpec successfully",
		})
	}
}
