package kafka

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log/slog"
	"orderbook/config"
	"orderbook/internal/core/port"
)

type EventRepository struct {
	producer    *kafka.Producer
	consumerMap map[string]port.ConsumerGroup
	config      *config.KafkaConfig
}

func New(
	config *config.Config,
) port.EventRepository {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":        config.KafkaConfig.Brokers,
		"allow.auto.create.topics": true, // @TODO update it for prod
	})
	if err != nil {
		slog.Info("Init User module error", "err", err)
		panic(err)
	}

	consumerMap := make(map[string]port.ConsumerGroup)

	return &EventRepository{producer, consumerMap, config.KafkaConfig}
}

func (m *EventRepository) PublishEvent(topic string, event any) error {
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

func (m *EventRepository) CloseAll() {
	for _, consumer := range m.consumerMap {
		consumer.StopPolling()
	}
}

func (m *EventRepository) StartPolling() {
	for _, consumer := range m.consumerMap {
		consumer.StartPolling()
	}
}

func (m *EventRepository) SetUpGroupConsumer(
	group string,
	topicMap map[string]func(event []byte) error,
	pollingInterval int,
	retry int,
) port.ConsumerGroup {
	c := NewConsumerGroup(group, topicMap, m.config, pollingInterval, retry)

	m.consumerMap[group] = c

	return c
}

var _ port.EventRepository = (*EventRepository)(nil)
