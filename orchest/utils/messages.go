package utils

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
)

func StartOrder(c *gin.Context) {

	topic := "order"
	value := "random value gogogo"

	delivery_chan := make(chan kafka.Event, 10000)
	err := Getp_client().Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value)},
		delivery_chan,
	)

	if err != nil {
		log.Panic(err)
	}

	// var resp string

	// return resp, err
}
