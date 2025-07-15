package kafka

import (
	"fmt"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const flushTimeout = 5000

type Producer struct {
	producer *kafka.Producer
}

func NewProducer(address []string) (*Producer, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": strings.Join(address, ","),
	}
	producer, err := kafka.NewProducer(config)
	if err != nil {
		return nil, fmt.Errorf("error creating kafka kafka_producer_consumer: %v", err)
	}
	return &Producer{producer: producer}, nil
}

func (p *Producer) Produce(message, topic, key string, time time.Time) error {
	kafkaMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value:     []byte(message),
		Key:       []byte(key),
		Timestamp: time,
	}
	kafkaChan := make(chan kafka.Event)
	if err := p.producer.Produce(kafkaMsg, kafkaChan); err != nil {
		return fmt.Errorf("error sending message to kafka: %v", err)
	}
	event := <-kafkaChan
	switch ev := event.(type) {
	case *kafka.Message:
		return nil
	case kafka.Error:
		return ev
	default:
		return fmt.Errorf("unknown event type: %T", ev)
	}
}

func (p *Producer) Close() {
	p.producer.Flush(flushTimeout)
	p.producer.Close()
}
