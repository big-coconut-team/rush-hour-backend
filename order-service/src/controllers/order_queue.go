package controllers

import (
	"fmt"
	"log"
	// "os"
	// "strings"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"

)

func RunQueue() {

	// kafka_add := fmt.Sprintf("%s:9092", os.Getenv("KAFKA_SERVICE_ADDRESS"))

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		// "bootstrap.servers":               kafka_add,
		"bootstrap.servers":               "10.109.2.246:9092",
		"group.id":                        "order-group",
		"go.application.rebalance.enable": true,
	})
	if err != nil {
		log.Panic(err)
	}

	err = consumer.Subscribe("order", nil)

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
			// controllers.CreateNewOrder(e.Value)

			oid := CreateOrder([]byte(e.Value))

			var tempData map[string]interface{}

			err = json.Unmarshal(e.Value, tempData)
			if err != nil {
				log.Panic(err)
			}
			tempData["payment_id"] = oid
			newData, err := json.Marshal(tempData)
			if err != nil {
				log.Panic(err)
			}

			res := fmt.Sprintf(
				`{
					"send_from": "order",
					"action": "CreatePayment",
					"data": %s,
				}`, newData)
			
			SendMSG("orchest", []byte(res))

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
