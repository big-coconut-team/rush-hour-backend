package main

import (
	"scalable-final-proj/payment-service/controllers"
	"scalable-final-proj/payment-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	router := gin.Default()

	// public := router.Group("/api")
	private := router.Group("/api")
	private.POST("/create_payment", controllers.CreatePayment)
	private.POST("/make_payment", controllers.MakePayment)

	router.Run("localhost:8099")
}
