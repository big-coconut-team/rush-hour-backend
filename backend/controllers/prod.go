package controllers

import (
	"net/http"
	"scalable-final-proj/backend/models"
	"scalable-final-proj/backend/utils"
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
}

func AddProduct(c *gin.Context) {

	var input ProductInput

		
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := models.Product{}

	p.ProdName = input.ProdName
	p.Details = input.Details
	p.StartTime = input.StartTime
	p.EndTime = input.EndTime
	p.InitialPrice = input.InitialPrice
	p.DiscountedPrice = input.DiscountedPrice
	p.Stock = input.Stock
	p.NumSold = 0
	p.UserID = id

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

	// uid, err := utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for object := range minioClient.ListObjects(context.Background(), "product", minio.ListObjectsOptions{Recursive: true}) {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		list = append(list, object.Key)
		err = minioClient.FGetObject(context.Background(), "product", object.Key, path+"/product-list/"+object.Key, minio.GetObjectOptions{})
		if err != nil {
			println(list[0])
			fmt.Println(err)
			return
		}

		// localFile, err := os.Create("my-testfile")
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// defer localFile.Close()

		// stat, err := reader.Stat()
		// if err != nil {
		// 	log.Fatalln(err)
		// }

		// if _, err := io.CopyN(localFile, reader, stat.Size); err != nil {
		// 	log.Fatalln(err)
		// }
	}
	c.JSON(http.StatusOK, gin.H{"message": list})
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func GetProdByTime(c *gin.Context) {
	dt := time.Now()
	var input ProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productdb, err := models.DB.DB().Query("SELECT * FROM products WHERE start_time <= ? AND end_time >= ?", dt)

	if err != nil {
		panic(err.Error())
	}

	prod := ProductInput{}
	for productdb.Next() {
		var prodID, userID int
		var start, end time.Time
		err = productdb.Scan(&prodID, &userID, &start, &end)
		if err != nil {
			panic(err.Error())
		}
		prod.StartTime = start
		prod.EndTime = end
	}
	defer models.DB.Close()

}
