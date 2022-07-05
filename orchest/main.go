package main

import (
    // "github.com/confluentinc/confluent-kafka-go/kafka"
	"orchrest/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	
	router := gin.Default()
	// public := router.Group("/api")

	router.POST("/start_order", utils.StartOrder)
	router.POST("/listen_order", utils.ListenOrder)

	router.Run("localhost:3333")

}