package utils

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var p_client *kafka.Producer

func Initp_client() {


	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka-broker:9092",
	})

	if err != nil {
		log.Panic(err)
	}

	p_client = p
}

func Getp_client() *kafka.Producer {
	return p_client
}
