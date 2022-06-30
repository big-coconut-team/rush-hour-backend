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

	u, err := GetUserByID(uint(user_id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func GetUserByID(uid uint) (models.User, error) {

	var u models.User

	if err := models.DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.Password = ""

	return u, nil

}
