package p_controllers

import (
	// "fmt"
	// "github.com/gin-gonic/gin"
	"product_svc/p_utils"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func SendMSG(topic string, data []byte) (error){
	delivery_chan := make(chan kafka.Event, 10000)
	err := p_utils.Getp_client().Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          data},
		delivery_chan,
	)
	return err
}