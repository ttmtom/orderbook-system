package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log/slog"
	"orderbook/config"
	"os"
)

type ConsumerGroup struct {
	consumer *kafka.Consumer
	eventMap map[string]func(event any)
}

var run = false

func NewConsumerGroup(
	group string,
	topicMap map[string]func(event any),
	config *config.KafkaConfig,
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

	return &ConsumerGroup{c, topicMap}
}

func (cg *ConsumerGroup) StartPolling() {
	run = true
	go func() {
		for run {
			ev := cg.consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n",
					e.TopicPartition, string(e.Value))
				if e.Headers != nil {
					fmt.Printf("%% Headers: %v\n", e.Headers)
				}
				break
			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				// But in this example we choose to terminate
				// the application if all brokers are down.
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
				break
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}()
}

func (cg *ConsumerGroup) StopPolling() {
	run = false
	cg.consumer.Close()
}
