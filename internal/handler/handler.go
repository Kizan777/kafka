package handler

import (
	"context"

	"kafka_producer/internal/logs"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandleMessage(ctx context.Context, message []byte, topic kafka.TopicPartition, cn int) error {
	l := logs.FromContext(ctx)
	l.Infof("Consumer #%d, Message from kafka with offset %d '%s' on partition %d", cn, topic.Offset, string(message), topic.Partition)
	return nil
}
