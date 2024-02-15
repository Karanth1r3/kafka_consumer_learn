package main

import (
	"fmt"
	"log"
	"time"

	kafkaconfig "github.com/Karanth1r3/kafka_learn/kafka"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	topic := "coordinates"
	kafkaCFG := kafkaconfig.QuickKafkaConsumerConfig()
	go func() {
		consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": kafkaCFG.BootstrapServers,
			"group.id":          kafkaCFG.GroupID,
			"auto.offset.reset": kafkaCFG.AutoOffsetReset,
		})
		if err != nil {
			log.Fatal(err)
		}

		err = consumer.Subscribe(topic, nil)
		if err != nil {
			log.Fatal(err)
		}

		for {
			ev := consumer.Poll(100)
			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("consumed message from the queue: %s\n", string(e.Value))
			case *kafka.Error:
				fmt.Printf("%s\n", e)
			}
		}
	}()
	// TODO - refactor (of course)
	time.Sleep(time.Second * 10)
}
