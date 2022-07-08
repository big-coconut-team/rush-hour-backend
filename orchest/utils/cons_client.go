package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ListenOrder() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               "localhost:9092",
		"group.id":                        "order-group",
		"go.application.rebalance.enable": true,
	})
	if err != nil {
		log.Panic(err)
	}

	var topics []string
	topics = append(topics, "order", "payment")
	err = consumer.SubscribeTopics(topics, nil)

	if err != nil {
		log.Panic(err)
	}

	run := true

	for run == true {
		ev := consumer.Poll(0)
		switch e := ev.(type) {
		case *kafka.Message:
			//  Message on order[0]@4:
			fmt.Printf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))
			
			messageChan := fmt.Sprintf("%s",e.TopicPartition)

			// order => payment
			if strings.Contains(messageChan, "order") {
				fmt.Println("order")
			}
			// payment => stock
			if strings.Contains(messageChan, "payment") {
				fmt.Println("payment")
			}	
			// else {
			// 	fmt.Println("topic doesn't exist")
			// }	

		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", e)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			log.Panic(err)
			run = false
		}
	}

	consumer.Close()
}
