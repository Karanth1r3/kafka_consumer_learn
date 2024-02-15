package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// Producer config
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}

	// Creating topic
	topic := "coordinates"

	//Creating producer
	kp, err := kafka.NewProducer(config)
	if err != nil {
		log.Fatal(err)
	}
	defer kp.Close()
	// Writing message to the topic
	for i := 0; i < 10; i++ {
		value := fmt.Sprintf("message %d", i)
		err := kp.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic,
				Partition: kafka.PartitionAny},
			Value: ([]byte(value)),
		}, nil)
		if err != nil {
			fmt.Printf("failed to produce message %d: %v\n", i, err)
		} else {
			fmt.Printf("Produced message: %d: %s\n", i, value)
		}
	}

	kp.Flush(15 * 1000)

}
