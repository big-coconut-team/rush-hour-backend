package utils

import (
    "github.com/confluentinc/confluent-kafka-go/kafka"
    "fmt"
    "os"
	"github.com/gin-gonic/gin"
    "log"
)

func StartOrder(c *gin.Context) {

    topic := "order"

    p, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost",
        // "client.id": "localhost",
        // "acks": "all"
    })
    
    if err != nil {
        fmt.Printf("Failed to create producer: %s\n", err)
        os.Exit(1)
    }

    value := "random value gogogo"

    delivery_chan := make(chan kafka.Event, 10000)
    err = p.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Value: []byte(value)},
        delivery_chan,
    )

    if err != nil {
        log.Panic(err)
    }

    // var resp string 

    go func() {
        for e := range p.Events() {
            switch ev := e.(type) {
            case *kafka.Message:
                if ev.TopicPartition.Error != nil {
                    // resp = fmt.Sprintf("Failed to deliver message: %v\n, order failed to create", ev.TopicPartition)
                    fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
                } else {
                    // resp = fmt.Sprintf("Successfully produced record to topic %s partition [%d] @ offset %v\n, order created",
                    // *ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
                    fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n, order created",
                    *ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
                }
            }
        }
    }()

    // return resp, err
}

func ListenOrder(c *gin.Context) {
    consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "weloveajkanat",
	})

    if err != nil {
        log.Panic(err)
    }
    // topics := "order"
    // topics := list.New()
    // topics.PushFront("order")
    // var topics [1]string
    // topics[0] = "order"
    topics := make([]string,1)
    topics = append(topics, "order")
    // topic := "order"
    // err = consumer.SubscribeTopics([]string{*topic}, nil)
    err = consumer.SubscribeTopics(topics, nil)

    run := true

    for run == true {
        ev := consumer.Poll(0)
        switch e := ev.(type) {
        case *kafka.Message:
            fmt.Printf("%% Message on %s:\n%s\n",
                e.TopicPartition, string(e.Value))
        case kafka.PartitionEOF:
            fmt.Printf("%% Reached %v\n", e)
        case kafka.Error:
            fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
            log.Panic(err)
            run = false
        default:
            fmt.Printf("Ignored %v\n", e)
        }
    }

    consumer.Close()
}
