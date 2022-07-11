package entrypoints

import (
	"errors"
	"net/http"
	"os"
	"user-service/models"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {

	user_id := GetUserID()

	u, err := GetUserByID(uint(user_id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// p := GetProductsByID(models.DB, uint(user_id))

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func GetUserID() int {
	uid := os.Getuid()
	return uid
}

func GetUserByID(uid uint) (models.User, error) {

	var u models.User

	if err := models.DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.Password = ""

	return u, nil

}

// func Current() (*User, error) {
// 	cache.Do(func() { cache.u, cache.err = current() })
// 	if cache.err != nil {
// 		return nil, cache.err
// 	}
// 	u := *cache.u // copy
// 	return &u, nil
// }

// var cache struct {
// 	sync.Once
// 	u   *User
// 	err error
// }

// func LookupId(uid string) (*User, error) {
// 	if u, err := Current(); err == nil && u.Uid == uid {
// 		return u, err
// 	}
// 	return lookupUserId(uid)
// }
