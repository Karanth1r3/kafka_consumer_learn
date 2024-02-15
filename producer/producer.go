package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// Producer config
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "kafka_client",
		"acks":              "all",
	}

	//Creating producer
	kp, err := kafka.NewProducer(config)
	if err != nil {
		log.Fatal(err)
	}
	defer kp.Close()

	// Creating topic
	topic := "coordinates"
	deliverCh := make(chan kafka.Event, 10000)

	// Writing message to the topic
	for i := 0; i < 10; i++ {
		value := fmt.Sprintf("message %d", i)
		err := kp.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic,
				Partition: kafka.PartitionAny},
			Value: ([]byte(value)),
		}, deliverCh)
		if err != nil {
			fmt.Printf("failed to produce message %d: %v\n", i, err)
		} else {
			fmt.Printf("Produced message: %d: %s\n", i, value)
		}
		<-deliverCh
		time.Sleep(time.Second * 4)
	}
	//kp.Flush(15 * 1000)
}
