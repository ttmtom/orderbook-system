package kafka

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log/slog"
	"orderbook/config"
)

type Manager struct {
	producer    *kafka.Producer
	consumerMap map[string]*kafka.Consumer
}

func NewKafkaManager(
	config *config.Config,
) *Manager {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":        config.KafkaConfig.Brokers,
		"allow.auto.create.topics": true, // @TODO update it for prod
	})
	if err != nil {
		slog.Info("Init User module error", "err", err)
		panic(err)
	}

	consumerMap := make(map[string]*kafka.Consumer)

	return &Manager{producer, consumerMap}
}

func (m *Manager) PublishEvent(topic string, event any) error {
	slog.Info("Publish event", "topic", topic, "event", event)
	eventData, err := json.Marshal(event)
	if err != nil {
		return err
	}
	err = m.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: eventData,
	}, nil)

	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) SubscriptEvent(topic string, handler func(event any)) {

}
