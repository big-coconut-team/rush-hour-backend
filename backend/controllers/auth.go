package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"scalable-final-proj/backend/utils"

	"github.com/gin-gonic/gin"
)

var USER_SERVICE_ADDR = "localhost" // change this to localhost if testing locally
var USER_SERVICE_PORT = "8000"

type UpdateInfoInput struct {
	NewPassword string `json:"newpassword"`
	NewEmail    string `json:"newemail"`
}

type UpdateInfoOutput struct {
	UserID      int    `json:"uid"`
	NewPassword string `json:"newpassword"`
}

func UpdateInfo(c *gin.Context) {
	var input UpdateInfoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := utils.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload, errr := json.Marshal(&UpdateInfoOutput{UserID: id, NewPassword: input.NewPassword, NewEmail: input.NewEmail})
	if errr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res := bytes.NewBuffer(payload)

	resp, err := http.Post("http://"+USER_SERVICE_ADDR+":"+USER_SERVICE_PORT+"/changepassword", "application/json", res)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sb := string(body)

	c.JSON(resp.StatusCode, sb)

}

type ReturnedVerfiedMessage struct {
	Message string `json:"message"`
	UserID  int    `json:"uid"`
}

func Login(c *gin.Context) {

	payload, err := ioutil.ReadAll(c.Request.Body)
	res := bytes.NewBuffer(payload)

	resp, err := http.Post("http://"+USER_SERVICE_ADDR+":"+USER_SERVICE_PORT+"/verifypassword", "application/json", res)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sb := string(body)

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadRequest, sb)
		return
	}

	var payloadStruct ReturnedVerfiedMessage
	if err3 := json.Unmarshal([]byte(sb), &payloadStruct); err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err3 := utils.GenerateToken(payloadStruct.UserID)

	if err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": payloadStruct.Message, "token": token})

}

func Register(c *gin.Context) {
	payload, err := ioutil.ReadAll(c.Request.Body)
	res := bytes.NewBuffer(payload)

	resp, err := http.Post("http://"+USER_SERVICE_ADDR+":"+USER_SERVICE_PORT+"/signup", "application/json", res)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sb := string(body)

	c.JSON(resp.StatusCode, sb)
}
