package server

import (
	"exporter/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) FindAcctHandler(c *gin.Context) {
	var accts []models.Acct
	// accts := []models.Acct{}
	s.db.FindAcct(&accts)
	// for i := range accts {
	// 	acctBank := []models.AcctBank{}
	// 	s.db.FindAcctBankByID(&acctBank, accts[i].ID)
	// 	accts[i].AcctBanks = append(accts[i].AcctBanks, acctBank...)
	// }
	c.JSON(http.StatusOK, models.Message{
		Acct: accts,
	})
}

func (s *Server) SaveAcctHandler(c *gin.Context) {
	acct := &models.Acct{}
	err := c.ShouldBind(acct)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: err.Error(),
		})
		return
	}
	log.Printf("%v\n", acct)

	err, acct.FileID, acct.FileName = s.SaveFile(c, "file")
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: err.Error(),
		})
	}
	err = s.db.Save(acct)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: err.Error(),
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
