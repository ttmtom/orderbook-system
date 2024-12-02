package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"orderbook/config"
)

type Kafka struct {
	Consumer *kafka.Consumer
	Producer *kafka.Producer
}

func NewProducer(config config.KafkaConfig) (*kafka.Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.Brokers,
	})
	if err != nil {
		panic(err)
	}

	return producer, nil
}

func NewConsumer(config config.KafkaConfig) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	return consumer, nil
}
