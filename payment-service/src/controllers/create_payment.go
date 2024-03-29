package controllers

import (
	// "context"
	// "net/http"
	// "github.com/gin-gonic/gin"
	"encoding/json"
	// "bytes"
	"fmt"
	"log"
	"scalable-final-proj/payment-service/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type InputPayment struct {
	// gorm.Model
<<<<<<< HEAD
	MadeByUserID int            `json:"made_by_id" binding:"required"`
	TotalPrice   int            `json:"total_price" binding:"required"`
	Paid         bool           `json:"paid" binding:"required"`
	ProductDict  map[string]int `json:"prod_dict" binding:"required"`
	PaymentID    int            `json:"payment_id" binding:"required"`
=======
	MadeByUserID   	int 	`json:"made_by_id" binding:"required"`
	TotalPrice		int 	`json:"total_price" binding:"required"`
	Paid			bool	`json:"paid" binding:"required"`
	ProductDict		map[string]int	`json:"prod_dict" binding:"required"`
	PaymentID		int		`json:"payment_id" binding:"required"`
>>>>>>> 7f55db27b2f4abadba5fdbb62d32a456d7ded1f7
}

func CreateNewPayment(input []byte) {

	var ip InputPayment
	err := json.Unmarshal([]byte(input), &ip)
	var p models.Payment

	p.PaymentID = ip.PaymentID
	p.MadeByUserID = ip.MadeByUserID
	p.Amount = ip.TotalPrice
	p.Paid = ip.Paid
	p.ProductDict = fmt.Sprintf("%v", ip.ProductDict)

	p.Paid = false
	if err != nil {
		log.Panic(err)
	}
	_, err = p.SavePayment()
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("PAYMENT ID: %d\n", p.PaymentID)
}

func MakePayment(input []byte) {
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
