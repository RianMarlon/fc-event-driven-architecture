package main

import (
	"balance-ms/internal/database"
	"balance-ms/internal/event"
	"balance-ms/internal/event/handler"
	usecase "balance-ms/internal/usecase/update_accounts_balance"
	"balance-ms/pkg/events"
	"balance-ms/pkg/kafka"
	"database/sql"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://root:root@balance-ms-database:5432/balance?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "balance",
	}
	kafkaConsumer := kafka.NewConsumer(&configMap, []string{"balances"})
	msgChan := make(chan *ckafka.Message)

	eventDispatcher := events.NewEventDispatcher()

	balanceGateway := database.NewBalanceDB(db)
	updateAccountsBalanceUseCase := usecase.NewUpdateAccountsBalanceUseCase(balanceGateway)
	eventDispatcher.Register("BalanceUpdated", handler.NewBalanceUpdatedKafkaHandler(*updateAccountsBalanceUseCase))

	go kafkaConsumer.Consume(msgChan)

	for msg := range msgChan {
		updateBalanceEvent := event.NewBalanceUpdated()
		updateBalanceEvent.SetPayload(msg.Value)
		eventDispatcher.Dispatch(updateBalanceEvent)
	}
}
