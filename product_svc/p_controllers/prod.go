package p_controllers

import (
	"encoding/json"
	"net/http"
	"product_svc/p_models"
	"strconv"
	"time"

	"context"
	"fmt"

	// "io"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type ProductInput struct {
	ProdName        string    `json:"prod_name" binding:"required"`
	Details         string    `json:"details" binding:"required"`
	StartTime       time.Time `json:"start_time" binding:"required"`
	EndTime         time.Time `json:"end_time" binding:"required"`
	InitialPrice    int       `json:"initial_price" binding:"required"`
	DiscountedPrice int       `json:"discounted_price" binding:"required"`
	Stock           int       `json:"stock" binding:"required"`
	UID             int       `json:"uid" binding:"required"`
}

func AddProduct(c *gin.Context) {

	var input ProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// id, err := utils.ExtractTokenID(c)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	p := p_models.Product{}

	p.ProdName = input.ProdName
	p.Details = input.Details
	p.StartTime = input.StartTime
	p.EndTime = input.EndTime
	p.InitialPrice = input.InitialPrice
	p.DiscountedPrice = input.DiscountedPrice
	p.Stock = input.Stock
	p.NumSold = 0
	p.UserID = input.UID

	_, err = p.SaveProd()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success adding product"})

}

func DownloadPhoto(c *gin.Context) {
	// dt := time.Now()
	path, err := os.Getwd()
	prods_list, prod_details := GetProdByTime()
	if err != nil {
		log.Println(err)
	}
	var list []string
	endpoint := "localhost:7000"
	accessKeyID := "minio"
	secretAccessKey := "minio123"
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, object := range prods_list {
		err = minioClient.FGetObject(context.Background(), "product", object, path+"/product-list/"+object, minio.GetObjectOptions{})

		if err != nil {
			println(list[0])
			fmt.Println(err)
			return
		}

	}

	// for object := range minioClient.ListObjects(context.Background(), "product", minio.ListObjectsOptions{Recursive: true}) {
	// 	if object.Err != nil {
	// 		fmt.Println(object.Err)
	// 		return
	// 	}
	// 	list = append(list, object.Key)
	// 	if stringInSlice(object.Key, prods_list) {
	// 		err = minioClient.FGetObject(context.Background(), "product", object.Key, path+"/product-list/"+object.Key, minio.GetObjectOptions{})
	// 	}
	// 	if err != nil {
	// 		println(list[0])
	// 		fmt.Println(err)
	// 		return
	// 	}

	// }
	c.JSON(http.StatusOK, gin.H{"details": prod_details})
}

// func inTimeSpan(start, end, check time.Time) bool {
// 	return check.After(start) && check.Before(end)
// }

// func stringInSlice(a string, list []string) bool {
// 	for _, b := range list {
// 		if b == a {
// 			return true
// 		}
// 	}
// 	return false
// }

func GetProdByTime() ([]string, []p_models.Product) {

	dt := time.Now()
	var l []string
	var list []p_models.Product

	productdb, err := p_models.DB.DB().Query("SELECT * FROM products WHERE start_time <= ? AND end_time >= ?", dt)

	if err != nil {
		panic(err.Error())
	}

	for productdb.Next() {
		var product p_models.Product
		var prod_id, user_id, initial_price, discounted_price, stock, num_sold int
		var path, prod_name, details string
		var start_time, end_time time.Time
		err = productdb.Scan(&prod_id, &prod_name, &details, &start_time, &end_time, &initial_price, &discounted_price, &stock, &num_sold, &user_id)
		if err != nil {
			panic(err.Error())
		}

		path = "/" + strconv.Itoa(user_id) + "/" + strconv.Itoa(prod_id) + "/"
		l = append(l, path)

		product_json := map[string]interface{}{
			"prod_id":          prod_id,
			"prod_name":        prod_name,
			"details":          details,
			"start_time":       start_time,
			"end_time":         end_time,
			"initial_price":    initial_price,
			"discounted_price": discounted_price,
			"stock":            stock,
			"num_sold":         num_sold,
			"user_id":          user_id,
		}
		json.Unmarshal([]byte(product_json), &product)
		list = append(list, product)

	}
	defer p_models.DB.Close()

	return l, list
}
