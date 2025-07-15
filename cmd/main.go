package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"kafka_producer/config"
	"kafka_producer/internal/handler"
	k "kafka_producer/internal/kafka"
	"kafka_producer/internal/logs"
	"kafka_producer/internal/usecase"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var defaultLevel = zap.NewAtomicLevelAt(zap.InfoLevel)

func main() {
	ctx := context.Background()
	l := logs.NewLogger(defaultLevel, os.Stdout)
	ctx = logs.WithLogger(ctx, l)
	err := godotenv.Load()
	if err != nil {
		l.Fatal("error loading .env file")
	}
	conf := config.NewAppConfig()

	address := []string{conf.Kafka.KafkaPort}
	p, err := k.NewProducer(address)
	if err != nil {
		l.Fatal(err)
	}
	sender := usecase.NewSender(*p)
	keysNum := conf.Kafka.NumberOfKeys
	topic := conf.Kafka.Topic1

	err = sender.SendTestNotification(ctx, topic, keysNum)
	if err != nil {
		l.Fatal(err)
	}

	consumerGroup := conf.Kafka.ConsumerGroup
	consumerNumber := conf.Kafka.ConsumerNumber
	consumerSessionTimeout := conf.Kafka.ConsumerSessionTimeout
	autoCommitIntervalMs := conf.Kafka.AutoCommitIntervalMs
	h := handler.NewHandler()
	c1, err := k.NewConsumer(h, address, topic,
		consumerGroup, consumerNumber, consumerSessionTimeout, autoCommitIntervalMs)
	if err != nil {
		l.Fatal(err)
	}

	go c1.Start(ctx)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	p.Close()
	l.Fatal(c1.Stop(ctx))

}
