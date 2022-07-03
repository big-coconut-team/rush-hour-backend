package controllers

import (
	// "context"
	"net/http"
	"github.com/gin-gonic/gin"

	"scalable-final-proj/payment-service/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"	
)

type CreatePaymentInput struct {
	MadeByUserID	int		`json:"made_by_id" binding:"required"`
	Amount			int		`json:"amount" binding:"required"`
}

func CreatePayment(c *gin.Context) {

	var input CreatePaymentInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := models.Payment{}

	p.MadeByUserID = input.MadeByUserID
	p.Amount = input.Amount
	p.Paid = false

	_, err = p.SavePayment()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "payment created"})
}


type MakePaymentInput struct{
	PaymentID 	int 	`json:"payment_id" binding:"required"`
}

func MakePayment(c *gin.Context) {

	var input MakePaymentInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := models.Payment{}

	_, err = p.UpdatePayment(input.PaymentID, true)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "paid"})
}
