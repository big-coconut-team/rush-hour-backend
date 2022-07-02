package main

import (
	"user-service/entrypoints"
	"user-service/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	router := gin.Default()

	router.POST("/signup", entrypoints.Register)
	router.POST("/verifypassword", entrypoints.VerifyPassword)

	router.Run("localhost:8000")
}
