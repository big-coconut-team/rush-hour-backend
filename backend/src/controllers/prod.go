package controllers

import (
	"github.com/gin-gonic/gin"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"controller_svc/utils"
)

func AddProduct(c *gin.Context) {
	data,err := ioutil.ReadAll(c.Request.Body)

	var tempData map[string]interface{}

	err = json.Unmarshal(data, &tempData)
	
	id, err := utils.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}	

	tempData["uid"] = id
	data,err = json.Marshal(tempData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseBody := bytes.NewBuffer(data)	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// fmt.Printf("add product: %s",data)

	resp, err := http.Post("http://rush-hour-product:8001/api/user/add_product", "application/json", responseBody)
	//Handle Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()	

	c.JSON(http.StatusOK, gin.H{"message": "Success adding product"})
}

func DownloadPhoto(c *gin.Context){
	resp, err := http.Get("http://rush-hour-product:8001/api/user/list_product")
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