package controllers

import (
	"fmt"
	"log"
	"os"
	// "strings"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"encoding/json"
)

func RunQueue() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               "localhost:9092",
		"group.id":                        "order-group",
		"go.application.rebalance.enable": true,
	})
	if err != nil {
		log.Panic(err)
	}

	err = consumer.Subscribe("payment", nil)

	if err != nil {
		log.Panic(err)
	}

	run := true

	for run == true {
		ev := consumer.Poll(0)
		switch e := ev.(type) {
		case *kafka.Message:

			fmt.Printf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))

			// TODO data sent from order to orc might not line up with what we send to pay svc

			// CreateNewPayment([]byte(e.Value))
			var tempData map[string]interface{}

			err = json.Unmarshal(e.Value, &tempData)
			if err != nil {
				log.Panic(err)
			}

			action := tempData["action"]

			// fmt.Printf("DO THIS ACTION: %s\n", action)
			switch action {
			case "CreatePayment":
				data,err := json.Marshal(tempData["data"])
				if err != nil {
					log.Panic(err)
				}
				CreateNewPayment([]byte(data))
			case "MakePayment":
				data,err := json.Marshal(tempData["data"])
				if err != nil {
					log.Panic(err)
				}
				MakePayment([]byte(data))
			}

			res := fmt.Sprintf(
				`{
					"send_from": "payment",
					"action": "UpdateStock",
					"data": %s
				}`, e.Value)

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