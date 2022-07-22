package controllers

import (
	"bytes"
	"controller_svc/utils"
<<<<<<< HEAD
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	// "fmt"
=======
>>>>>>> 7f55db27b2f4abadba5fdbb62d32a456d7ded1f7
)

func AddProduct(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)

	var tempData map[string]interface{}

	err = json.Unmarshal(data, &tempData)

	id, err := utils.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tempData["uid"] = id
	data, err = json.Marshal(tempData)
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

<<<<<<< HEAD
func DownloadPhoto(c *gin.Context) {
	resp, err := http.Get("http://localhost:8001/api/user/list_product")
=======
func DownloadPhoto(c *gin.Context){
	resp, err := http.Get("http://rush-hour-product:8001/api/user/list_product")
>>>>>>> 7f55db27b2f4abadba5fdbb62d32a456d7ded1f7
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
