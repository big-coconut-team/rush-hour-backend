package controllers

import (
	"net/http"
	"scalable-final-proj/backend/models"
	"scalable-final-proj/backend/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductInput struct {
	ProdName        string    `json:"prod_name" binding:"required"`
	Details         string    `json:"details" binding:"required"`
	StartTime       time.Time `json:"start_time" binding:"required"`
	EndTime         time.Time `json:"end_time" binding:"required"`
	InitialPrice    int       `json:"initial_price" binding:"required"`
	DiscountedPrice int       `json:"discounted_price" binding:"required"`
	Stock           int       `json:"stock" binding:"required"`
}

func AddProduct(c *gin.Context) {
	var input ProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := models.Product{}

	p.ProdName = input.ProdName
	p.Details = input.Details
	p.StartTime = input.StartTime
	p.EndTime = input.EndTime
	p.InitialPrice = input.InitialPrice
	p.DiscountedPrice = input.DiscountedPrice
	p.Stock = input.Stock
	p.NumSold = 0
	p.UserID = id

	_, err = p.SaveProd()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success adding product"})

}
