package kafka

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log/slog"
	"orderbook/config"
)

type Manager struct {
	producer    *kafka.Producer
	consumerMap map[string]*ConsumerGroup
	config      *config.KafkaConfig
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

	consumerMap := make(map[string]*ConsumerGroup)

	return &Manager{producer, consumerMap, config.KafkaConfig}
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

func (m *Manager) CloseAll() {
	for _, consumer := range m.consumerMap {
		consumer.StopPolling()
	}
}

func (m *Manager) StartPolling() {
	for _, consumer := range m.consumerMap {
		consumer.StartPolling()
	}
}

func (m *Manager) SetUpGroupConsumer(
	group string,
	topicMap map[string]func(event any),
) *ConsumerGroup {
	c := NewConsumerGroup(group, topicMap, m.config)

	m.consumerMap[group] = c

	return c
}
