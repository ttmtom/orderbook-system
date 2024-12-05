package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log/slog"
	"orderbook/config"
	"time"
)

type ConsumerGroup struct {
	consumer        *kafka.Consumer
	eventMap        map[string]func(event []byte) error
	Group           string
	PollingInterval int
	Retry           int
}

var run = false

func NewConsumerGroup(
	group string,
	topicMap map[string]func(event []byte) error,
	config *config.KafkaConfig,
	pollingInterval int,
	retry int,
) *ConsumerGroup {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        config.Brokers,
		"group.id":                 group,
		"session.timeout.ms":       6000,
		"auto.offset.reset":        "earliest",
		"enable.auto.offset.store": true,
	})

	if err != nil {
		slog.Error("fail on create consumer", "group", group, "err", err)
		panic(err.Error())
	}

	topics := make([]string, 0, len(topicMap))
	for key := range topicMap {
		topics = append(topics, key)
	}

	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		slog.Error("fail on subscribe", "topics", topics, "err", err)
		panic(err.Error())
	}

	return &ConsumerGroup{
		c,
		topicMap,
		group,
		pollingInterval,
		retry,
	}
}

func (cg *ConsumerGroup) processMessage(topic string, event []byte) error {
	handler, ok := cg.eventMap[topic]
	if !ok {
		slog.Error("Event handler not exist", "topic", topic)
		return nil
	}

	slog.Info("event handling")

	var err error
	for attempt := 0; attempt < cg.Retry; attempt++ {
		err = handler(event)
		if err == nil {
			break
		} else {
			time.Sleep(time.Duration(cg.PollingInterval))
		}
	}

	return err
}

func (cg *ConsumerGroup) StartPolling() {
	run = true
	go func() {
		for run {
			ev := cg.consumer.Poll(cg.PollingInterval)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				slog.Info("Event received", "group", cg.Group, "topic", *e.TopicPartition.Topic, "event", e.Value, "header", e.Headers)
				if e.TopicPartition.Topic == nil {
					slog.Error("Topic is nil")
					break
				}
				go func() {
					err := cg.processMessage(*e.TopicPartition.Topic, e.Value)
					if err != nil {
						slog.Info("Event done")
						cg.consumer.CommitMessage(e)
					}
				}()
			case kafka.Error:
				slog.Error("Kafka Error", "group", cg.Group, "code", e.Code(), "error", e.Error())
				if e.Code() == kafka.ErrAllBrokersDown {
					cg.StopPolling()
				}
			default:
				slog.Info("Ignored Event", "group", cg.Group, "event", e)
			}
		}
	}()
}

func (cg *ConsumerGroup) StopPolling() {
	run = false
	cg.consumer.Close()
	slog.Info("Consumer stop", "group", cg.Group)
}
