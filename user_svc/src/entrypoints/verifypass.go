package entrypoints

import (
	"net/http"
	"net/mail"
	"regexp"
	"user-service/models"

	"github.com/gin-gonic/gin"
)

type LoginInput struct { // allow sign in with either username or email -- pass in blank string for the unused option
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password" binding:"required"`
}

func VerifyPassword(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if _, err := mail.ParseAddress(input.Email); err != nil && len(input.Email) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err error

	if len(input.Username) == 0 && len(input.Email) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No username/email is given..."})
		return
	} else if len(input.Username) > 0 && len(input.Email) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Both username and email provided..."})
		return
	} else if len(input.Username) > 0 {
		isAlpha := regexp.MustCompile(`^[a-zA-Z0-9 ]+$`).MatchString
		if !isAlpha(input.Username) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username should contain alphanumeric or spaces only."})
			return
		}
		err = models.LoginCheck(input.Username, input.Password, false)
	} else {
		err = models.LoginCheck(input.Email, input.Password, true)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username/email or password is incorrect!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Correct password entered."})
}
