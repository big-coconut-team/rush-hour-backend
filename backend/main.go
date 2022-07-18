package main

import (
	// "net/http"

	"controller_svc/controllers"
	"controller_svc/middlewares"
	"controller_svc/utils"
	"log"
	// "scalable-final-proj/backend/product_svc/p_controllers"
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
					log.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					log.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n, order created",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

	router := gin.Default()

	public := router.Group("/api")

	public.POST("/signup", controllers.Register)
	public.POST("/login", controllers.Login)
	
	protected := router.Group("/api/user")

	protected.Use(middlewares.JwtAuthMiddleware())

	// protected.POST("/changepass", controllers.ChangePassword)
	// protected.GET("/profile", controllers.CurrentUser)

	protected.POST("/add_product", controllers.AddProduct)
	protected.GET("/list_product", controllers.DownloadPhoto)
	
	public.POST("/place_order", controllers.PlaceOrder)
	protected.POST("/make_payment", controllers.Pay)
	// router.GET("/albums", getAlbums)
	// router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)

	router.Run("localhost:8088")
}
