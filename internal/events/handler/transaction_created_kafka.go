package handler

import (
	"fmt"
	"walletcore/pkg/events"
	"walletcore/pkg/kafka"
)

type TransactionCreatedKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewTransactionCreatedKafkaHandler(kafka *kafka.Producer) *TransactionCreatedKafkaHandler {
	return &TransactionCreatedKafkaHandler{
		Kafka: kafka,
	}
}

func (h *TransactionCreatedKafkaHandler) Handle(message events.EventInterface) {
	h.Kafka.Publish(message, nil, "transactions")
	fmt.Println("TransactionCreatedKafkaHandler: ", message.GetPayload())
}
