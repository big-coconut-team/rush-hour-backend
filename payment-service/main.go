package main

import (
	"scalable-final-proj/payment-service/controllers"
	"scalable-final-proj/payment-service/models"
	"scalable-final-proj/payment-service/utils"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	utils.Initp_client()

	go func() {
		for e := range utils.Getp_client().Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n, order created",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

	go controllers.RunQueue()

	models.ConnectDatabase()

	router := gin.Default()

	// public := router.Group("/api")
	private := router.Group("/api")
	
	private.POST("/test_create", controllers.DummyCreatePayment)
	private.POST("/test_pay", controllers.DummyMakePayment)
	// private.POST("/create_payment", controllers.CreatePayment)
	// private.POST("/make_payment", controllers.MakePayment)

	router.Run("localhost:8091")
}
