package utils

import (
	"fmt"
	"log"
	"os"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"encoding/json"
)

func ListenOrder() {

	kafka_add := fmt.Sprintf("%s:9092", os.Getenv("KAFKA_SERVICE_ADDRESS"))

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               kafka_add,
		//"bootstrap.servers":               "10.109.2.246:9092",
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
			action := tempData["action"]

			data, err := json.Marshal(tempData["data"])

			if err != nil {
				log.Panic(err)
			}

			switch from {

			case "controller":
				// go to order
				switch action {
				case "CreateOrder":
					err =SendMSG("order", data)
					if err != nil {
						log.Panic(err)
					}
				}
			case "order":
				// go to payment
				res := fmt.Sprintf(
					`{
						"action" : "%s",
						"data": %s
					}`, action, data)
				switch action{
				case "CreatePayment":
					err = SendMSG("payment", []byte(res))
					if err != nil {
						log.Panic(err)
					}
				case "MakePayment":
					err = SendMSG("payment", []byte(res))
					if err != nil {
						log.Panic(err)
					}
				}
			case "payment":
				// go update prod
				switch action {
				case "UpdateStock":
					err =SendMSG("product", data)
					if err != nil {
						log.Panic(err)
					}
				}
			// notify user
			// case "product":
			// 	switch action {
			// 	case "NotifyUser":
					
			// 	}
			}

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
