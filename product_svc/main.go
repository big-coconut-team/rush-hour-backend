package main

import (
	"log"
	"product_svc/p_controllers"
	"product_svc/p_models"
	"product_svc/p_utils"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
)

func main() {
	p_utils.Initp_client()

	go func() {
		for e := range p_utils.Getp_client().Events() {
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

	go p_controllers.RunQueue()

	p_models.ConnectDataBase()

	router := gin.Default()
	protected := router.Group("/api/user")
	protected.POST("/add_product", p_controllers.AddProduct)
	protected.POST("/list_product", p_controllers.DownloadPhoto)
	// protected.POST("/update_stock", p_controllers.GetStockUpdate)

	// router.GET("/albums", getAlbums)
	// router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)

	router.Run("localhost:8001")
}
