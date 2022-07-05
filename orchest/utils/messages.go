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
        "bootstrap.servers": "localhost:9092",
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
