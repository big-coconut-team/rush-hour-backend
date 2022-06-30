package controllers

import (
	"errors"
	"net/http"
	"scalable-final-proj/backend/models"
	"scalable-final-proj/backend/utils"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {

	user_id, err := utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, p, err := GetUserByID(uint(user_id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u, "product": p})
}

func GetUserByID(uid uint) (models.User, models.Product, error) {

	var u models.User
	var p models.Product

	if err := models.DB.First(&u, uid).Error; err != nil {
		return u, p, errors.New("User not found!")
	}

	u.Password = ""

	return u, p, nil

}
