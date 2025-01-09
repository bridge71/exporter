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
	c.JSON(http.StatusOK, models.Message{AcctBank: acctBanks})
}

func (s *Server) DeleteAcctBankHandler(c *gin.Context) {
	acctBank := &models.AcctBank{}
	acctBank.ID = s.Str2Uint(c.PostForm("ID"))
	log.Printf("%v\n", acctBank)

	err := s.db.DeleteAcctBank(acctBank)
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

func (s *Server) FindAcctBankByIDHandler(c *gin.Context) {
	acctBanks := &[]models.AcctBank{}
	acctID := c.PostForm("AcctID")
	idStr, _ := strconv.Atoi(acctID)
	id := uint(idStr)
	s.db.FindAcctBankByID(acctBanks, id)
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

	err, acctBank.FileID, acctBank.FileName = s.SaveFile(c, "file")
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to save file",
		})
	}
	// Acct := models.Acct{}
	// log.Printf("sss %d\n", acctBank.AcctID)
	// s.db.FindByID(acctBank.AcctID, &Acct)
	acctBank.Acct.ID = acctBank.AcctID
	// log.Printf("??? %v\n", Acct)
	// log.Printf(" 23123  %d\n", Acct.ID)
	if acctBank.AcctID != 0 {
		err = s.db.Save(acctBank)
		if err != nil {
			c.JSON(http.StatusForbidden, models.Message{
				RetMessage: "请添加会计实体",
			})
			return
		}
	} else {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "请添加会计实体",
		})
		return
	}
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
