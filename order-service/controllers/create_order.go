package controllers

import (
	// "context"
	// "net/http"
	"github.com/gin-gonic/gin"
	"encoding/json"
	// "bytes"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"scalable-final-proj/order-service/utils"
	"scalable-final-proj/order-service/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

// type CreateOrderInput struct {
// 	MadeByUserID	int		`json:"made_by_id" binding:"required"`
// 	ProdIDs			string	`json:"prod_list" binding:"required"`
// }

// func CreateOrder(c *gin.Context) {

// 	var input CreateOrderInput

// 	err := c.ShouldBindJSON(&input)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	o := models.Order{}

// 	o.MadeByUserID = input.MadeByUserID
// 	o.ProductIDs = input.ProdIDs

// 	_, err = o.SaveOrder()

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// USER_SERVICE_ADDR := "localhost"
// 	// USER_SERVICE_PORT := "3333"
// 	data, err :=json.Marshal(o)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}	

// 	res,err := json.Marshal(fmt.Sprintf(`{
// 		"send_from": "order",
// 		"data": %s
// 	}`, data))


// 	// resp, err := http.Post("http://"+USER_SERVICE_ADDR+":"+USER_SERVICE_PORT+"/start_order", "application/json", bytes.NewBuffer(res))
// 	SendMSG("orchest", res)

// 	c.JSON(http.StatusOK, gin.H{"message": "order created"})
	
// 	// defer resp.Body.Close()
// }

func SendMSG(topic string, data []byte) (error){
	delivery_chan := make(chan kafka.Event, 10000)
	err := utils.Getp_client().Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data},
		delivery_chan,
	)
	return err
}

func CreateNewOrder(input []byte) {
	var o models.Order
	err := json.Unmarshal([]byte(input), &o)
	if err != nil {
		log.Panic(err)
	}
	_, err = o.SaveOrder()
	if err != nil {
		log.Panic(err)
	}
}

func DummyOrder(c *gin.Context) {
	data := fmt.Sprintf(
		`{
			"send_from": "controller",
			"data": {
				"made_by_id": 3,
				"prod_list": "[1,6,8,5]"
			}
		}`)

	SendMSG("orchest", []byte(data))
}