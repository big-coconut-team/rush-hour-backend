package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"scalable-final-proj/order-service/utils"
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

func DummyOrder(c *gin.Context) {
	data := fmt.Sprintf(
		`{
			"send_from": "controller",
			"action" : "CreateOrder",
			"data": {
				"made_by_id": 3,
				"total_price": 1540,
				"prod_dict": {"1":2,"4":3,"7":25},
			}
		}`)
		// total price sent from controllers
	SendMSG("orchest", []byte(data))
}