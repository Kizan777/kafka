package kafka

import (
	"context"
	"strings"

	"kafka_producer/internal/logs"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	noTimeout = -1
)

type Handler interface {
	HandleMessage(ctx context.Context, message []byte, topic kafka.TopicPartition, cn int) error
}

type Consumer struct {
	consumer       *kafka.Consumer
	handler        Handler
	stop           bool
	consumerNumber int
}

func NewConsumer(handler Handler, address []string, topic, consumerGroup string, consumerNumber, consumerSessionTimeout, autoCommitIntervalMs int) (*Consumer, error) {
	cfg := &kafka.ConfigMap{
		"bootstrap.servers":        strings.Join(address, ","),
		"group.id":                 consumerGroup,
		"session.timeout.ms":       consumerSessionTimeout,
		"enable.auto.offset.store": false,
		"enable.auto.commit":       true,
		"auto.commit.interval.ms":  autoCommitIntervalMs,
		"auto.offset.reset":        "earliest",
	}

	c, err := kafka.NewConsumer(cfg)
	if err != nil {
		return nil, err
	}

	if err = c.Subscribe(topic, nil); err != nil {
		return nil, err
	}
	return &Consumer{
		consumer:       c,
		handler:        handler,
		consumerNumber: consumerNumber,
	}, nil
}

func (c *Consumer) Start(ctx context.Context) {
	l := logs.FromContext(ctx)
	for {
		if c.stop {
			break
		}
		kafkaMsg, err := c.consumer.ReadMessage(noTimeout)
		if err != nil {
			l.Error(err)
		}
		if kafkaMsg == nil {
			continue
		}
		if err = c.handler.HandleMessage(ctx, kafkaMsg.Value, kafkaMsg.TopicPartition, c.consumerNumber); err != nil {
			l.Error(err)
			continue
		}
		if _, err = c.consumer.StoreMessage(kafkaMsg); err != nil {
			l.Error(err)
			continue
		}
	}
}

func (c *Consumer) Stop(ctx context.Context) error {
	l := logs.FromContext(ctx)
	c.stop = true
	if _, err := c.consumer.Commit(); err != nil {
		return err
	}
	l.Infof("Commited offset")
	return c.consumer.Close()
}
