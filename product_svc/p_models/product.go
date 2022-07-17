package p_models

import (
	// "github.com/jinzhu/gorm"
	"errors"
	//"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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

func (p *Product) SaveProd(c *gin.Context) (*Product, error) {

	var err error
	err = DB.WithContext(c.Request.Context()).Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func GetProductByUID(uid uint, c *gin.Context) (Product, error) {

	var p Product

	if err := DB.WithContext(c.Request.Context()).First(&p, uid).Error; err != nil {
		return p, errors.New("User not found!")
	}

	return p, nil

}

func GetProdByTime(c *gin.Context) ([]string, []Product) {

	dt := time.Now()
	var l []string
	var list []Product

	productdb, err := DB.WithContext(c.Request.Context()).Find("start_time <= ? AND end_time >= ?", dt).Rows()
	defer productdb.Close()

	if err != nil {
		panic(err.Error())
	}

	for productdb.Next() {
		// var product p_models.Product
		var prod_id, user_id, initial_price, discounted_price, stock, num_sold int
		var path, prod_name, details string
		var start_time, end_time time.Time
		err = productdb.Scan(&prod_id, &prod_name, &details, &start_time, &end_time, &initial_price, &discounted_price, &stock, &num_sold, &user_id)
		if err != nil {
			panic(err.Error())
		}

		path = "/" + strconv.Itoa(user_id) + "/" + strconv.Itoa(prod_id) + "/"
		l = append(l, path)

		// product_json := map[string]interface{}{
		// 	"prod_id":          prod_id,
		// 	"prod_name":        prod_name,
		// 	"details":          details,
		// 	"start_time":       start_time,
		// 	"end_time":         end_time,
		// 	"initial_price":    initial_price,
		// 	"discounted_price": discounted_price,
		// 	"stock":            stock,
		// 	"num_sold":         num_sold,
		// 	"user_id":          user_id,
		// }

		// product_json := fmt.Sprintf(`{"prod_id":%s,"prod_name":%s,"details":%s,"start_time":%s,"end_time":%s,"initial_price":%s,"discounted_price":%s,"stock":%s,"num_sold":%s,"user_id":%s}`
		// , prod_id, prod_name, details, start_time, end_time, initial_price, discounted_price, stock, num_sold, user_id)

		// json.Unmarshal([]byte(product_json), &product)

		p := Product{}

		p.ProdID = prod_id
		p.ProdName = prod_name
		p.Details = details
		p.StartTime = start_time
		p.EndTime = end_time
		p.InitialPrice = initial_price
		p.DiscountedPrice = discounted_price
		p.Stock = stock
		p.NumSold = num_sold
		p.UserID = user_id

		list = append(list, p)

	}

	return l, list
}

func (p *Product) UpdateStock(prodID int, numItems int, c *gin.Context) (*Product, error) {

	productdb, err := DB.WithContext(c.Request.Context()).Select("stock", "num_sold").Find("prod_id = ?", prodID).Rows()
	defer productdb.Close()

	if err != nil {
		panic(err.Error())
	}
	//fmt.Println("QUERY: %v\n", productdb)

	for productdb.Next() {
		var stock, num_sold int

		err = productdb.Scan(&stock, &num_sold)
		if err != nil {
			panic(err.Error())
		}
		if stock >= numItems {
			var err error
			newStock := stock - numItems
			newNumSold := num_sold + numItems
			err = DB.WithContext(c.Request.Context()).Model(&p).Where("prod_id = ?", prodID).Select("stock", "num_sold").Updates(Product{Stock: newStock, NumSold: newNumSold}).Error

			if err != nil {
				return &Product{}, err
			}
		} else {

			return p, errors.New("error: cannot update stock")
		}
	}

	return p, nil
}
