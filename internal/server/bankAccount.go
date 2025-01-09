package server

import (
	"exporter/internal/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) FindBankAccountHandler(c *gin.Context) {
	bankAccounts := &[]models.BankAccount{}
	s.db.FindBankAccount(bankAccounts)
	c.JSON(http.StatusOK, models.Message{BankAccount: *bankAccounts})
}

func (s *Server) DeleteBankAccountHandler(c *gin.Context) {
	bankAccount := &models.BankAccount{}
	bankAccount.ID = s.Str2Uint(c.PostForm("ID"))
	log.Printf("%d\n", bankAccount.ID)

	err := s.db.DeleteBankAccount(bankAccount)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to delete bankAccount",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete bankAccount successfully",
		})
	}
}

func (s *Server) FindBankAccountByIDHandler(c *gin.Context) {
	bankAccounts := &[]models.BankAccount{}
	mercID := c.PostForm("MercID")
	idStr, _ := strconv.Atoi(mercID)
	id := uint(idStr)
	s.db.FindBankAccountByID(bankAccounts, id)
	c.JSON(http.StatusOK, models.Message{BankAccount: *bankAccounts})
}

func (s *Server) SaveBankAccountHandler(c *gin.Context) {
	bankAccount := &models.BankAccount{}
	err := c.ShouldBind(bankAccount)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			// RetMessage: err.Error(),
			RetMessage: "请绑定客商信息",
		})
		return
	}
	log.Printf("%v\n", bankAccount)

	err, bankAccount.FileID, bankAccount.FileName = s.SaveFile(c, "file")
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to save file",
		})
	}
	log.Printf("id   %d\n", bankAccount.MerchantID)
	bankAccount.Merchant.ID = bankAccount.MerchantID
	err = s.db.Save(bankAccount)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to save bankAccount",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved bankAccount successfully",
		})
	}
}
