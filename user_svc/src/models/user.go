package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
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

func LoginCheck(search string, password string, isEmail bool) error {

	var err error

	u := User{}

	if isEmail {
		err = DB.Model(User{}).Where("email = ?", search).Take(&u).Error
	} else {
		err = DB.Model(User{}).Where("LOWER(REPLACE(TRIM(username), \" \", \"\")) = LOWER(REPLACE(TRIM(?), \" \", \"\"))", search).Take(&u).Error
	}

	if err != nil {
		return err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return err
	}

	return nil
}

func GetExistingUser(username string, email string) (User, error) {
	var err error

	u := User{}

	err = DB.Model(&User{}).Where("LOWER(REPLACE(TRIM(username), \" \", \"\")) = LOWER(REPLACE(TRIM(?), \" \", \"\"))", username).Or("email = ?", email).First(&u).Error

	return u, err
}

func (u *User) SaveUser() (*User, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) UpdateUser(fields ...string) (*User, error) {

	var err error
	err = DB.Select(fields).Save(&u).Error
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
