package controllers

import (
	// "context"
	"net/http"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"bytes"

	"scalable-final-proj/order-service/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"	
)

type CreateOrderInput struct {
	MadeByUserID	int		`json:"made_by_id" binding:"required"`
	ProdIDs			string	`json:"prod_list" binding:"required"`
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

	USER_SERVICE_ADDR := "localhost"
	USER_SERVICE_PORT := "3333"
	res, err :=json.Marshal(o)

	resp, err := http.Post("http://"+USER_SERVICE_ADDR+":"+USER_SERVICE_PORT+"/start_order", "application/json", bytes.NewBuffer(res))

	c.JSON(http.StatusOK, gin.H{"message": resp.Body})
	
	defer resp.Body.Close()
}