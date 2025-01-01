package server

import (
	"exporter/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) FindMerchantHandler(c *gin.Context) {
	merchants := []models.Merchant{}
	s.db.FindMerchant(&merchants)
	for i := range merchants {
		custs := []models.Cust{}
		s.db.FindCustById(&custs, merchants[i].MercId)
		merchants[i].Custs = append(merchants[i].Custs, custs...)

		bankAccounts := []models.BankAccount{}
		s.db.FindBankAccountById(&bankAccounts, merchants[i].MercId)
		merchants[i].BankAccounts = append(merchants[i].BankAccounts, bankAccounts...)
	}
	c.JSON(http.StatusOK, models.Message{
		Merchant: merchants,
	})
}

func (s *Server) SaveMerchantHandler(c *gin.Context) {
	merchant := &models.Merchant{}
	err := c.ShouldBind(merchant)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of merchant",
		})
		return
	}
	log.Printf("%v\n", merchant)

	err, merchant.FileId, merchant.FileName = s.SaveFile(c)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to save file",
		})
	}
	err = s.db.SaveMerchant(merchant)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to save merchant",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "saved merchant successfully",
		})
	}
}

func (s *Server) DeleteMerchantHandler(c *gin.Context) {
	merchant := &models.Merchant{}
	err := c.ShouldBind(merchant)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "error in bind of merchant",
		})
		return
	}
	log.Printf("%v\n", merchant)

	err = s.db.DeleteMerchant(merchant)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Message{
			RetMessage: "failed to delete merchant",
		})
	} else {
		c.JSON(http.StatusOK, models.Message{
			RetMessage: "delete merchant successfully",
		})
	}
}
