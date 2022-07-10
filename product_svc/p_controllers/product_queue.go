package p_controllers

import (
	"fmt"
	"log"
	"os"
	// "strings"
	"github.com/confluentinc/confluent-kafka-go/kafka"

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

	err = consumer.Subscribe("product", nil)

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

			UpdateManyStock([]byte(e.Value))

			res := fmt.Sprintf(
				`{
					"send_from": "product",
					"action": "None",
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