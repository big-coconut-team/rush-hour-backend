package models

import (
	"html"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserID   int    `gorm:"primary_key;size:11;not null;" json:"uid"`
	Username string `gorm:"size:30;not null;" json:"username"`
	Password string `gorm:"size:256;not null;" json:"password"`
	Email    string `gorm:"size:256;not null;" json:"email"`
	Coin     int    `gorm:"size:10;not null;" json:"coin"`
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(search string, password string, isEmail bool, c *gin.Context) (int, error) {

	var err error

	u := User{}

	if isEmail {
		err = DB.WithContext(c.Request.Context()).Model(User{}).Where("email = ?", search).Take(&u).Error
	} else {
		err = DB.WithContext(c.Request.Context()).Model(User{}).Where("LOWER(REPLACE(TRIM(username), \" \", \"\")) = LOWER(REPLACE(TRIM(?), \" \", \"\"))", search).Take(&u).Error
	}

	if err != nil {
		return -1, err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return -1, err
	}

	return u.UserID, nil
}

func GetExistingUserByID(uid int, c *gin.Context) (User, error) {
	var err error

	u := User{}

	err = DB.WithContext(c.Request.Context()).Model(&User{}).Where("user_id = ?", uid).First(&u).Error

	return u, err
}

func GetExistingUser(username string, email string, c *gin.Context) (User, error) {
	var err error

	u := User{}

	err = DB.WithContext(c.Request.Context()).Model(&User{}).Where("LOWER(REPLACE(TRIM(username), \" \", \"\")) = LOWER(REPLACE(TRIM(?), \" \", \"\"))", username).Or("email = ?", email).First(&u).Error

	return u, err
}

func (u *User) SaveUser(c *gin.Context) (*User, error) {

	var err error
	err = DB.WithContext(c.Request.Context()).Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) UpdateUser(c *gin.Context) (*User, error) {
	var err error

	//turn password into hash
	if len(u.Password) > 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return &User{}, err
		}
		u.Password = string(hashedPassword)
	}

	err = DB.WithContext(c.Request.Context()).Model(&User{}).Where("user_id", u.UserID).Updates(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}
