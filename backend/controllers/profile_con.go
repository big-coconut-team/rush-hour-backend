package controllers

import (
	"errors"
	"net/http"
	"scalable-final-proj/backend/models"
	"scalable-final-proj/backend/product_svc/p_models"
	"scalable-final-proj/backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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

	p := GetProductsByID(models.DB, uint(user_id))

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u, "product": p})
}

func GetUserByID(uid uint) (models.User, error) {

	var u models.User

	if err := models.DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.Password = ""

	return u, nil

}

func GetProductsByID(db *gorm.DB, uid uint) []p_models.Product {
	orders := make([]p_models.Product, 0)
	db.Where("user_id=?", uid).Find(&p_models.Product{}).Scan(&orders)
	return orders
}
