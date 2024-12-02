package config

import (
	"os"
)

type KafkaConfig struct {
	Brokers string
}

func LoadKafkaConfig() *KafkaConfig {
	return &KafkaConfig{
		os.Getenv("KAFKA_BROKERS"),
	}
}
