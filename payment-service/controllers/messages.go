package controllers

import (
	"github.com/gin-gonic/gin"
	"scalable-final-proj/payment-service/utils"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"fmt"
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

func DummyCreatePayment(c *gin.Context) {
	data := fmt.Sprintf(
		`{
			"send_from": "order",
			"action" : "CreatePayment",
			"data": {
				"made_by_id" : 4,
				"prod_dict": "{1:2,4:3,7:25}",
				"total_price": 256
			}
		}`)

	SendMSG("orchest", []byte(data))
}

func DummyMakePayment(c *gin.Context) {
	data := fmt.Sprintf(
		`{
			"send_from": "order",
			"action" : "MakePayment",
			"data": {
				"payment_id": 38
			}
		}`)

	SendMSG("orchest", []byte(data))
}