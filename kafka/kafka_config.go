package kafkaconfig

// TODO - replace with better config when the whole idea is up
type (
	KafkaProducerConfig struct {
		BootstrapServers string
		ClientID         string
		Acks             string
	}

	KafkaConsumerConfig struct {
		BootstrapServers string
		GroupID          string
		AutoOffsetReset  string
	}
)

// TODO - make something better
func QuickKafkaProducerConfig() *KafkaProducerConfig {
	return &KafkaProducerConfig{
		BootstrapServers: "localhost:9092",
		ClientID:         "learnGroup",
		Acks:             "all",
	}
}

// TODO - make read from file i guess
func QuickKafkaConsumerConfig() *KafkaConsumerConfig {
	return &KafkaConsumerConfig{
		BootstrapServers: "localhost:9092",
		GroupID:          "learnGroup",
		AutoOffsetReset:  "smallest",
	}
}
