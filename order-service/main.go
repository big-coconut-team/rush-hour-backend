package main

import (
	"scalable-final-proj/order-service/controllers"
	"scalable-final-proj/order-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	router := gin.Default()

	// public := router.Group("/api")
	private := router.Group("/api")
	private.POST("/create_order", controllers.CreateOrder)

	router.Run("localhost:8099")
}
