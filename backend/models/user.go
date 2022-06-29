package models

import (
	// "github.com/jinzhu/gorm"

	"html"
	"scalable-final-proj/backend/utils"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	// gorm.Model
	UserID   int    `gorm:"primary_key;size:11;not null;" json:"uid"`
	Username string `gorm:"size:30;not null;" json:"username"`
	Password string `gorm:"size:256;not null;" json:"password"`
	Email    string `gorm:"size:256;not null;" json:"email"`
	Coin     int    `gorm:"size:10;not null;" json:"coin"`
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {

	var err error

	u := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := util.GenerateToken(u.ID) // generate token

	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *User) SaveUser() (*User, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave() error {

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
