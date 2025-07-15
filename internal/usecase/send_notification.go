package usecase

import (
	"context"
	"fmt"
	"time"

	"kafka_producer/internal/pkg/key_generator"

	"github.com/pkg/errors"
)

func (s Sender) SendTestNotification(ctx context.Context, topic string, keysNum int) error {
	keys := key_generator.GenerateUUIDString(keysNum)
	for i := 1; i <= 150; i++ {
		msg := fmt.Sprintf("Kafka message #%d", i)
		key := keys[i%keysNum]
		if err := s.producer.Produce(msg, topic, key, time.Now()); err != nil {
			return errors.Errorf("Failed to produce message #%d: %s", i, err)
		}
	}
	return nil
}
