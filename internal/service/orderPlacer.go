package service

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type (
	OrderPlacer struct {
		producer   *kafka.Producer
		topic      string
		deliveryCh chan kafka.Event
	}
)

func NewOrderPlacer(p *kafka.Producer, topic string) *OrderPlacer {

	return &OrderPlacer{
		producer:   p,
		topic:      topic,
		deliveryCh: make(chan kafka.Event, 10000),
	}
}

func (op *OrderPlacer) PlaceOrder(orderType string, size int) error {

	var (
		format  = fmt.Sprintf("%s - %d", orderType, size)
		payload = []byte(format)
	)

	err := op.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &op.topic,
			Partition: kafka.PartitionAny},
		Value: payload,
	},
		op.deliveryCh)
	if err != nil {
		fmt.Printf("failed to produce message: %v\n", err)
	}

	<-op.deliveryCh
	fmt.Printf("placed order on the queue: %s\n", format)

	return nil
}
