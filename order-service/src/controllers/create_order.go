package controllers

import (
	// "context"
	"net/http"
	// "github.com/gin-gonic/gin"
	"encoding/json"
	// "bytes"
	// "fmt"
	// "github.com/confluentinc/confluent-kafka-go/kafka"
	// "scalable-final-proj/order-service/utils"
	"fmt"
	"log"
	"scalable-final-proj/order-service/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var order_id int

type InputOrder struct {
	MadeByUserID int            `json:"made_by_id" binding:"required"`
	ProductDict  map[string]int `json:"prod_dict" binding:"required"`
	TotalPrice   int            `json:"total_price" binding:"required"`
}

<<<<<<< HEAD
func CreateOrder(input []byte) int {
=======
func CreateOrder(input []byte) (int) {
>>>>>>> 7f55db27b2f4abadba5fdbb62d32a456d7ded1f7
	var io InputOrder
	err := json.Unmarshal(input, &io)

	var o models.Order
	o.MadeByUserID = io.MadeByUserID
	o.ProductDict = fmt.Sprintf("%v", io.ProductDict)
	o.TotalPrice = io.TotalPrice

	if err != nil {
		log.Panic(err)
	}
	ord, err := o.SaveOrder()
	if err != nil {
		log.Panic(err)
	}
<<<<<<< HEAD
	order_id = o.OrderID
	return o.OrderID
}

func GetOrderId(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"order_id": order_id})
	return
}
=======
	return ord.OrderID
}
>>>>>>> 7f55db27b2f4abadba5fdbb62d32a456d7ded1f7
