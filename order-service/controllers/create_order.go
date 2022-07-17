package controllers


import (
	// "context"
	// "net/http"
	// "net/http"
	// "github.com/gin-gonic/gin"
	"encoding/json"
	// "bytes"

	// "bytes"
	// "fmt"
	// "github.com/confluentinc/confluent-kafka-go/kafka"
	// "scalable-final-proj/order-service/utils"
	"scalable-final-proj/order-service/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"	
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func CreateOrder(input []byte) {
	var o models.Order
	err := json.Unmarshal([]byte(input), &o)
	if err != nil {
		log.Panic(err)
	}
	_, err = o.SaveOrder()
	if err != nil {
		log.Panic(err)
	}
}