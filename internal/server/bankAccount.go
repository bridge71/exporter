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
	err := c.ShouldBind(bankAccount)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of bankAccount",
		})
		return
	}
	log.Printf("%v\n", bankAccount)

	err = s.db.DeleteBankAccount(bankAccount)
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

func (s *Server) FindBankAccountByIdHandler(c *gin.Context) {
	bankAccounts := &[]models.BankAccount{}
	mercId := c.PostForm("MercId")
	idStr, _ := strconv.Atoi(mercId)
	id := uint(idStr)
	s.db.FindBankAccountById(bankAccounts, id)
	c.JSON(http.StatusOK, models.Message{BankAccount: *bankAccounts})
}

func (s *Server) SaveBankAccountHandler(c *gin.Context) {
	bankAccount := &models.BankAccount{}
	err := c.ShouldBind(bankAccount)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of bankAccount",
		})
		return
	}
	log.Printf("%v\n", bankAccount)

	err, bankAccount.FileId, bankAccount.FileName = s.SaveFile(c)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to save file",
		})
	}
	err = s.db.SaveBankAccount(bankAccount)
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
