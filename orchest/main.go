package main

import (
	// "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"orchrest/utils"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.Initp_client()

	go func() {
		for e := range utils.Getp_client().Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					// resp = fmt.Sprintf("Failed to deliver message: %v\n, order failed to create", ev.TopicPartition)
					log.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					// resp = fmt.Sprintf("Successfully produced record to topic %s partition [%d] @ offset %v\n, order created",
					// *ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
					log.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n, order created",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

	go utils.ListenOrder()

	router := gin.Default()
	// public := router.Group("/api")

	router.POST("/start_order", utils.StartOrder)
	router.POST("/start_pay", utils.SendMSGPayment)

	router.Run("localhost:3333")
}
