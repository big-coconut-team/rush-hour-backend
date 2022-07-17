package controllers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"controller_svc/utils"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func SendMSG(topic string, data []byte) (error){
	delivery_chan := make(chan kafka.Event, 10000)
	err := utils.Getp_client().Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data},
		delivery_chan,
	)
	return err
}

type PlaceOrderInput struct {
	ProdIDs			map[string]float64	`json:"prod_dict" binding:"required"`
}

func PlaceOrder(c *gin.Context) {
	var input PlaceOrderInput
    
	err := c.ShouldBindJSON(&input)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	id, err := utils.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	prod_dict, err := json.Marshal(input.ProdIDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := fmt.Sprintf(
		`{
			"send_from": "controller",
			"action" : "CreateOrder",
			"data": {
				"made_by_id": %d,
				"total_price": 1540,
				"prod_dict": %s,
			}
		}`,id ,prod_dict)
		// total price sent from controllers
	err = SendMSG("orchest", []byte(data))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

type PayInput struct {
	PaymentID 	int 	`json:"payment_id" binding:"required"`
}

func Pay(c *gin.Context) {
	var input PayInput
    
	err := c.ShouldBindJSON(&input)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	pay_id := input.PaymentID

	data := fmt.Sprintf(
		`{
			"send_from": "order",
			"action" : "MakePayment",
			"data": {
				"prod_dict": {"1":2,"4":3,"7":25},
				"payment_id": %d
			}
		}`, pay_id )
		// total price sent from controllers
	err = SendMSG("orchest", []byte(data))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

