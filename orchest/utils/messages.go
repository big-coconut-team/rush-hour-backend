package utils

import (
	"log"
	"net/http"
	"encoding/json"
    // "fmt"
	// "strconv"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
)

type COInput struct {
    OrderID         int     `json:"order_id" binding:"required"`
	MadeByUserID	int		`json:"made_by_id" binding:"required"`
	ProdIDs			string	`json:"prod_list" binding:"required"`
}

func StartOrder(c *gin.Context) {

    var input COInput

    err := c.ShouldBindJSON(&input)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	topic := "order"
	value, err := json.Marshal(&input)

	if err != nil {
		log.Panic(err)
	}

	delivery_chan := make(chan kafka.Event, 10000)
	err = Getp_client().Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value)},
		delivery_chan,
	)

	if err != nil {
		log.Panic(err)
	}
}

type CPInput struct {
    	// gorm.Model
	PaymentID 		int 	`json:"payment_id" binding:"required"`
	MadeByUserID   	int 	`json:"made_by_id" binding:"required"`
	Amount			int		`json:"amount" binding:"required"`
	Paid			bool	`json:"paid" binding:"required"`
}

func SendMSGPayment(c *gin.Context) {

    var input CPInput

    err := c.ShouldBindJSON(&input)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	topic := "payment"
	value, err := json.Marshal(&input)

	if err != nil {
		log.Panic(err)
	}

	delivery_chan := make(chan kafka.Event, 10000)
	err = Getp_client().Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value)},
		delivery_chan,
	)

	if err != nil {
		log.Panic(err)
	}
}
