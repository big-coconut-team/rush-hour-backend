package utils

import (
	"log"
	// "fmt"
	// "os"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var p_client *kafka.Producer

func Initp_client() {

	// kafka_add := fmt.Sprintf("%s:9092", os.Getenv("KAFKA_SERVICE_ADDRESS"))

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka-service:9092",
		// "client.id": "localhost",
		// "acks": "all"
	})

	if err != nil {
		log.Panic(err)
	}

	p_client = p
}

func Getp_client() *kafka.Producer {
	return p_client
}
