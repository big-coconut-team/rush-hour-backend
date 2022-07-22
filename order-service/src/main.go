package main

import (
	"log"
	"scalable-final-proj/order-service/controllers"
	"scalable-final-proj/order-service/models"
	"scalable-final-proj/order-service/utils"

	"github.com/gin-gonic/gin"

	// "sync"
	"github.com/confluentinc/confluent-kafka-go/kafka"
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
	// private.POST("/create_order", controllers.CreateOrder) // is called by orchest instead of user
	private.POST("/test", controllers.DummyOrder)
	private.GET("/get_oid", controllers.GetOrderId)
	router.Run("localhost:8099")
	// var wg sync.WaitGroup
	// wg.Add(1)
	// wg.Wait()
}
