package models

// import (
// 	"github.com/jinzhu/gorm"
// )

type User struct {
	// gorm.Model
	UserID int `gorm:"primary_key;size:11;not null;" json:"uid"`
	Username string `gorm:"size:30;not null;" json:"username"`
	Password string `gorm:"size:256;not null;" json:"password"`
	Email string `gorm:"size:256;not null;" json:"email"`
	Coin int `gorm:"size:10;not null;" json:"coin"`
}