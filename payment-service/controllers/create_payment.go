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