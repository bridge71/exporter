package server

import (
	"exporter/internal/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) FindAcctBankHandler(c *gin.Context) {
	acctBanks := []models.AcctBank{}
	s.db.FindAcctBank(&acctBanks)
	length := len(acctBanks)
	var acct []models.Acct
	for i := 0; i < length; i++ {
		s.db.FirstAcct(acctBanks[i].AcctId, &acct)
		acctBanks[i].AcctName = acct[0].AcctName
	}
	c.JSON(http.StatusOK, models.Message{AcctBank: acctBanks})
}

func (s *Server) DeleteAcctBankHandler(c *gin.Context) {
	acctBank := &models.AcctBank{}
	err := c.ShouldBind(acctBank)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of acctBank",
		})
		return
	}
	log.Printf("%v\n", acctBank)

	err = s.db.DeleteAcctBank(acctBank)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to delete acctBank",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete acctBank successfully",
		})
	}
}

func (s *Server) FindAcctBankByIdHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	acctId := c.PostForm("AcctId")
	idStr, _ := strconv.Atoi(acctId)
	id := uint(idStr)
	s.db.FindAcctBankById(acctBanks, id)
	c.JSON(http.StatusOK, models.Message{AcctBank: *acctBanks})
}

func (s *Server) SaveAcctBankHandler(c *gin.Context) {
	acctBank := &models.AcctBank{}
	err := c.ShouldBind(acctBank)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: err.Error(),
		})
		return
	}
	log.Printf("%v\n", acctBank)

	err, acctBank.FileId, acctBank.FileName = s.SaveFile(c, "file")
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to save file",
		})
	}
	err = s.db.SaveAcctBank(acctBank)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to save acctBank",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved acctBank successfully",
		})
	}
}
