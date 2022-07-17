package p_controllers

import (
	// "encoding/json"
	"net/http"
	"net/http/httptest"
	"product_svc/p_models"
	"time"

	"context"
	"fmt"

	// "io"
	"encoding/json"
	"log"
	"os"
	"strconv"

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

	_, err = p.SaveProd(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success adding product"})

}

func DownloadPhoto(c *gin.Context) {
	// dt := time.Now()
	path, err := os.Getwd()
	prods_list, prod_details := p_models.GetProdByTime(c)
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

	c.JSON(http.StatusOK, gin.H{"details": prod_details})
}

type UpdateStockInput struct {
	ProdID  int `json:"prod_id" binding:"required"`
	NumItem int `json:"num_item" binding:"required"`
}

func GetStockUpdate(c *gin.Context) {

	var input UpdateStockInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := p_models.Product{}

	_, err = p.UpdateStock(input.ProdID, input.NumItem, c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "stock updated"})
}


func UpdateManyStock(input []byte) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	// "{1:2,4:3,7:25}"
	var tempData map[string]interface{}
	// fmt.Printf("INPUT:\n%s\n", input)
	err := json.Unmarshal([]byte(input), &tempData)
	if err != nil {
		log.Panic(err)
	}
	for key, element := range tempData {
		// fmt.Println("Key:", key, "=>", "Element:", element)
		prod := p_models.Product{}
		intKey, err := strconv.Atoi(key)
		if err != nil {
			log.Panic(err)
		}
		_, err = prod.UpdateStock(intKey, int(element.(float64)), c)
		if err != nil {
			fmt.Println(err.Error())
			log.Panic(err)
		}
	}
}
