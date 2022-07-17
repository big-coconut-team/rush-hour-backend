package controllers

import (
	// "context"
	"net/http"
	"scalable-final-proj/order-service/models"

	"github.com/gin-gonic/gin"

	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/go-redis/redis"
	// "context"
	// "bytes"
	// "fmt"
)

type CreateOrderInput struct {
	MadeByUserID int    `json:"made_by_id" binding:"required"`
	ProdIDs      string `json:"prod_list" binding:"required"`
}

func CreateOrder(c *gin.Context) {

	var input CreateOrderInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	o := models.Order{}

	o.MadeByUserID = input.MadeByUserID
	o.ProductIDs = input.ProdIDs

	_, err = o.SaveOrder()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order created"})
}
