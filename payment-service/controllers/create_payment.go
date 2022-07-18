package controllers

import (
	// "context"
	// "net/http"
	// "github.com/gin-gonic/gin"
	"encoding/json"
	// "bytes"
	"log"
	"fmt"
	"scalable-final-proj/payment-service/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"	
)


type InputPayment struct {
	// gorm.Model
	MadeByUserID   	int 	`json:"made_by_id" binding:"required"`
	TotalPrice		int 	`json:"total_price" binding:"required"`
	Paid			bool	`json:"paid" binding:"required"`
	ProductDict		map[string]int	`json:"prod_dict" binding:"required"`
}


func CreateNewPayment(input []byte) {

	var ip InputPayment
	err := json.Unmarshal([]byte(input), &ip)
	var p models.Payment

	p.MadeByUserID = ip.MadeByUserID
	p.Amount = ip.TotalPrice
	p.Paid = ip.Paid
	p.ProductDict = fmt.Sprintf("%v",ip.ProductDict)

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