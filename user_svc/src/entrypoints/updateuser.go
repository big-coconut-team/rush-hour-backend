package entrypoints

import (
	"net/http"
	"net/mail"
	"user-service/models"

	"github.com/gin-gonic/gin"
)

type ChangePasswordInput struct {
	UserID      int    `json:"uid" binding:"required"`
	NewPassword string `json:"newpassword"`
	NewEmail    string `json:"newemail"`
}

func ChangePassword(c *gin.Context) {
	var input ChangePasswordInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.NewPassword == "" && input.NewEmail == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No password or email given."})
		return
	}

	u := models.User{}

	u.UserID = input.UserID
	if len(input.NewPassword) > 0 {
		u.Password = input.NewPassword
	}

	if len(input.NewEmail) > 0 {
		if _, err := mail.ParseAddress(input.NewEmail); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		u.Email = input.NewEmail
	}

	_, err := u.UpdateUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Change of password was successful."})
}
