package controllers

import (
	"fmt"
	"net/http"

	// "log"
	"bytes"
	"controller_svc/utils"
	"encoding/json"
	"io/ioutil"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
)

var orderid int

func SendMSG(topic string, data []byte) error {
	delivery_chan := make(chan kafka.Event, 10000)
	err := utils.Getp_client().Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data},
		delivery_chan,
	)
	return err
}

type PlaceOrderInput struct {
	ProdIDs map[string]int `json:"prod_dict" binding:"required"`
}

func PlaceOrder(c *gin.Context) {
	var input PlaceOrderInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := utils.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	prod_dict, err := json.Marshal(input.ProdIDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data := fmt.Sprintf(
		`{
			"prod_dict": %s
		}`, prod_dict)

	resp, err := http.Post("http://localhost:8001/api/user/calc", "application/json", bytes.NewBuffer([]byte(data)))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r_body, err := ioutil.ReadAll(resp.Body)
	// buff_r_body := bytes.NewBuffer(r_body)

	var tempData map[string]interface{}

	err = json.Unmarshal(r_body, &tempData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	total_price := tempData["total"]

	data = fmt.Sprintf(
		`{
			"send_from": "controller",
			"action" : "CreateOrder",
			"data": {
				"made_by_id": %d,
				"total_price": %d,
				"prod_dict": %s
			}
		}`, id, int(total_price.(float64)), prod_dict)
	// total price sent from controllers
	err = SendMSG("orchest", []byte(data))
	if err != nil {
		// log.Panic(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("TOTAL PRICE: %d\n PROD DICT: %s\n", total_price, prod_dict)
}

type PayInput struct {
	ProdIDs   map[string]int `json:"prod_dict" binding:"required"`
	PaymentID int            `json:"payment_id" binding:"required"`
}

func Pay(c *gin.Context) {
	var input PayInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	prod_dict, err := json.Marshal(input.ProdIDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pay_id := input.PaymentID

	data := fmt.Sprintf(
		`{
			"send_from": "order",
			"action" : "MakePayment",
			"data": {
				"prod_dict": %s,
				"payment_id": %d
			}
		}`, prod_dict, pay_id)
	orderid = pay_id
	// prod_dict from frontend cart

	// total price sent from controllers
	err = SendMSG("orchest", []byte(data))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func GetOrderId(c *gin.Context) {
	resp, err := http.Get("http://localhost:8099/api/get_oid")
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

	var tempData map[string]interface{}

	err = json.Unmarshal(body, &tempData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	oid := tempData["order_id"]

	c.JSON(http.StatusOK, gin.H{"order_id": oid})
}
