package entrypoints

import (
	"net/http"
	"user-service/models"

	"github.com/gin-gonic/gin"
)

type ChangePasswordInput struct {
	UserID      int    `json:"uid" binding:"required"`
	NewPassword string `json:"newpassword" binding:"required"`
}

func ChangePassword(c *gin.Context) {
	var input ChangePasswordInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.UserID = input.UserID
	u.Password = input.NewPassword

	_, err := u.UpdateUser("UserID", "Password")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Change of password was successful."})
}
