package controllers

import (
	"github.com/gin-gonic/gin"
	"bytes"
	"net/http"
	"io/ioutil"
)

func AddProduct(c *gin.Context) {
	data,err := ioutil.ReadAll(c.Request.Body)
	responseBody := bytes.NewBuffer(data)	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := http.Post("localhost:8001/add_product", "application/json", responseBody)
	//Handle Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()	

	c.JSON(http.StatusOK, gin.H{"message": "Success adding product"})
}

func DownloadPhoto(c *gin.Context){
	resp, err := http.Get("localhost:8001/list_product")
	//Handle Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
   	}

	c.JSON(http.StatusOK, gin.H{"message": body})
}