package p_models

import (
	// "github.com/jinzhu/gorm"
	"errors"
	"time"
)

type Product struct {
	// gorm.Model
	ProdID          int       `gorm:"primary_key;size:11;not null;" json:"prod_id"`
	ProdName        string    `gorm:"size:30;not null;" json:"prod_name"`
	Details         string    `gorm:"size:1000;not null;" json:"details"`
	StartTime       time.Time `gorm:"not null;" json:"start_time"`
	EndTime         time.Time `gorm:"not null;" json:"end_time"`
	InitialPrice    int       `gorm:"size:10;not null;" json:"initial_price"`
	DiscountedPrice int       `gorm:"size:10;not null;" json:"discounted_price"`
	Stock           int       `gorm:"size:11;not null;" json:"stock"`
	NumSold         int       `gorm:"size:11;not null;" json:"num_sold"`
	UserID          int       `gorm:"size:11;not null;" json:"uid"`
}

func (p *Product) SaveProd() (*Product, error) {

	var err error
	err = DB.Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func GetProductByUID(uid uint) (Product, error) {

	var p Product

	if err := DB.First(&p, uid).Error; err != nil {
		return p, errors.New("User not found!")
	}

	return p, nil

}
