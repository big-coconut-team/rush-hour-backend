package models

import (
	// "github.com/jinzhu/gorm"

	"errors"
)

func GetUserByID(uid uint) (User, error) {

	var u User

	if err := DB.First(&u, uid).Error; err != nil {
		return u, errors.New("User not found!")
	}

	u.PrepareGive()

	return u, nil

}
func (u *User) PrepareGive() {
	u.Password = ""
}
