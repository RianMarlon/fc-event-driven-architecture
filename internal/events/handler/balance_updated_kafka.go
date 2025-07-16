package handler

import (
	"fmt"
	"walletcore/pkg/events"
	"walletcore/pkg/kafka"
)

type BalanceUpdatedKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewBalanceUpdatedKafkaHandler(kafka *kafka.Producer) *BalanceUpdatedKafkaHandler {
	return &BalanceUpdatedKafkaHandler{Kafka: kafka}
}

func (h *BalanceUpdatedKafkaHandler) Handle(event events.EventInterface) {
	h.Kafka.Publish(event, nil, "balances")
	fmt.Println("BalanceUpdatedKafkaHandler called")
}
