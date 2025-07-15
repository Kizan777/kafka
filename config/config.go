package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type AppConfig struct {
	Kafka Kafka
}

func NewAppConfig() AppConfig {
	var appConfig AppConfig
	err := env.Parse(&appConfig)
	if err != nil {
		log.Fatal("Error parsing config", err)
	}
	if appConfig.Kafka.NumberOfKeys == 0 {
		log.Fatal("Error parsing config: Kafka must have at least one key")
	}
	return appConfig
}

type Kafka struct {
	KafkaPort              string `env:"KAFKA_PORT"`
	Topic1                 string `env:"TOPIC1"`
	NumberOfKeys           int    `env:"NUMBER_OF_KEYS"`
	ConsumerGroup          string `env:"CONSUMER_GROUP"`
	ConsumerNumber         int    `env:"CONSUMER_NUMBER"`
	ConsumerSessionTimeout int    `env:"CONSUMER_SESSION_TIMEOUT"`
	AutoCommitIntervalMs   int    `env:"AUTO_COMMIT_INTERVAL_MS"`
}
