package usecase

import k "kafka_producer/internal/kafka"

type Sender struct {
	producer k.Producer
}

func NewSender(producer k.Producer) *Sender {
	return &Sender{
		producer: producer,
	}
}
