package entrypoints

import (
	"net/http"
	"user-service/models"

	"github.com/gin-gonic/gin"
)

type UserInput struct {
	UserID int `json:"uid" binding:"required"`
}

func GetUser(c *gin.Context) {
	var input UserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetExistingUserByID(input.UserID, c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.Password = ""

	c.JSON(http.StatusOK, u)
}
