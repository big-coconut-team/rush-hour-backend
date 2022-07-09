package utils

import (
	"fmt"
	"log"
	"os"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"encoding/json"
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

	// var topics []string
	// topics = append(topics, "order", "payment")
	// err = consumer.SubscribeTopics(topics, nil)

	err = consumer.Subscribe("orchest", nil)

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

			// determine which svc to go to
			var tempData map[string]interface{}

			err = json.Unmarshal(e.Value, &tempData)
			if err != nil {
				log.Panic(err)
			}

			from := tempData["send_from"]

			data, err := json.Marshal(tempData["data"])

			if err != nil {
				log.Panic(err)
			}

			switch from {

			case "controller":
				// go to order
				SendMSG("order", data)
			case "order":
				// go to payment
			case "payment":
				// go update prod
			}
			

			// messageChan := fmt.Sprintf("%s",e.TopicPartition)

			// // order => payment
			// if strings.Contains(messageChan, "order") {
			// 	fmt.Println("order")
			// 	send data to order_queue and trigger create_order?
			// }
			// // payment => stock
			// if strings.Contains(messageChan, "payment") {
			// 	fmt.Println("payment")
			// }	

			// if strings.Contains(messageChan, "orchest") {
			// 	fmt.Println("orchest")
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
