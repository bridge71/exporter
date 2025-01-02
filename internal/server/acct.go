package server

import (
	"exporter/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) FindAcctHandler(c *gin.Context) {
	accts := []models.Acct{}
	s.db.Find(&accts)
	for i := range accts {
		acctBank := []models.AcctBank{}
		s.db.FindAcctBankById(&acctBank, accts[i].AcctId)
		accts[i].AcctBanks = append(accts[i].AcctBanks, acctBank...)
	}
	c.JSON(http.StatusOK, models.Message{
		Acct: accts,
	})
}

func (s *Server) SaveAcctHandler(c *gin.Context) {
	acct := &models.Acct{}
	err := c.ShouldBind(acct)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of acct",
		})
		return
	}
	log.Printf("%v\n", acct)

	err, acct.FileId, acct.FileName = s.SaveFile(c)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to save file",
		})
	}
	err = s.db.Save(acct)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to save acct",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved acct successfully",
		})
	}
}

func (s *Server) DeleteAcctHandler(c *gin.Context) {
	acct := &models.Acct{}
	err := c.ShouldBind(acct)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of acct",
		})
		return
	}
	log.Printf("%v\n", acct)

	err = s.db.Delete(acct)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to delete acct",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete acct successfully",
		})
	}
}
