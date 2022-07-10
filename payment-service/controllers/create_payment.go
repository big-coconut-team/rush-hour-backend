package controllers

import (
	// "context"
	// "net/http"
	// "github.com/gin-gonic/gin"
	"encoding/json"
	// "bytes"
	"log"
	// "fmt"
	"scalable-final-proj/payment-service/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"	
)

// type CreatePaymentInput struct {
// 	MadeByUserID	int		`json:"made_by_id" binding:"required"`
// 	Amount			int		`json:"amount" binding:"required"`
// }

// func CreatePayment(c *gin.Context) {

// 	var input CreatePaymentInput

// 	err := c.ShouldBindJSON(&input)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

	// p := models.Payment{}

	// p.MadeByUserID = input.MadeByUserID
	// p.Amount = input.Amount
	// p.Paid = false

// 	_, err = p.SavePayment()

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	USER_SERVICE_ADDR := "localhost"
// 	USER_SERVICE_PORT := "3333"
// 	res, err :=json.Marshal(p)

// 	resp, err := http.Post("http://"+USER_SERVICE_ADDR+":"+USER_SERVICE_PORT+"/start_pay", "application/json", bytes.NewBuffer(res))

// 	c.JSON(http.StatusOK, gin.H{"message": resp.Body,
// "data sent": p.Paid})
	
// 	defer resp.Body.Close()

// 	// c.JSON(http.StatusOK, gin.H{"message": "payment created"})
// }


// type MakePaymentInput struct{
// 	PaymentID 	int 	`json:"payment_id" binding:"required"`
// }

// func MakePayment(c *gin.Context) {

// 	var input MakePaymentInput

// 	err := c.ShouldBindJSON(&input)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	p := models.Payment{}

// 	_, err = p.UpdatePayment(input.PaymentID, true)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "paid"})
// }

func CreateNewPayment(input []byte) {
	var p models.Payment
	err := json.Unmarshal([]byte(input), &p)
	p.Paid = false
	if err != nil {
		log.Panic(err)
	}
	_, err = p.SavePayment()
	if err != nil {
		log.Panic(err)
	}
}

func MakePayment(input []byte){
	var p models.Payment
	var tempData map[string]interface{}
	err := json.Unmarshal(input, &tempData)
	if err != nil {
		log.Panic(err)
	}
	id := int(tempData["payment_id"].(float64))
	// fmt.Printf("ID TYPE: \n%T\n", id)
	_, err = p.UpdatePayment(id, true)	
	if err != nil {
		log.Panic(err)
	}
}