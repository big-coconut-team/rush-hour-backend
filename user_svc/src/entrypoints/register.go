package entrypoints

import (
	"net/http"
	"net/mail"
	"regexp"
	"user-service/models"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if _, err := mail.ParseAddress(input.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isAlpha := regexp.MustCompile(`^[a-zA-Z0-9 ]+$`).MatchString
	if !isAlpha(input.Username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username should contain alphanumeric or spaces only."})
		return
	}

	r, sqlerr := models.GetExistingUser(input.Username, input.Email, c)
	if sqlerr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": sqlerr.Error()})
		return
	} else if len(r.Username) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User/Email already exists!"})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password
	u.Email = input.Email
	u.Coin = 0

	_, err := u.SaveUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration success."})
}
