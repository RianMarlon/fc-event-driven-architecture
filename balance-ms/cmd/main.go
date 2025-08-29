package main

import (
	getBalanceByAccountUsecase "balance-ms/internal/app/usecase/get_balance_by_account"
	updateAccountsBalanceUsecase "balance-ms/internal/app/usecase/update_accounts_balance"
	event "balance-ms/internal/domain/event"
	database "balance-ms/internal/infra/db/postgres"
	httpserver "balance-ms/internal/infra/http"
	web "balance-ms/internal/infra/http/handler"
	kafka_handler "balance-ms/internal/infra/messaging/kafka/handler"
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

	balanceGateway := database.NewBalanceRepositoryDB(db)
	updateAccountsBalanceUseCase := updateAccountsBalanceUsecase.NewUpdateAccountsBalanceUseCase(balanceGateway)
	getBalanceByAccountUseCase := getBalanceByAccountUsecase.NewGetBalanceByAccountUsecase(balanceGateway)

	server := httpserver.NewWebServer(":3003")
	balanceHandler := web.NewHttpBalanceHandler(*getBalanceByAccountUseCase)
	server.AddHandler(httpserver.Handler{
		Path:        "/balances/{account_id}",
		Method:      "GET",
		HandlerFunc: balanceHandler.GetBalanceByAccount,
	})
	go server.Start()

	eventDispatcher.Register("BalanceUpdated", kafka_handler.NewBalanceUpdatedKafkaHandler(*updateAccountsBalanceUseCase))
	go kafkaConsumer.Consume(msgChan)

	for msg := range msgChan {
		updateBalanceEvent := event.NewBalanceUpdated()
		updateBalanceEvent.SetPayload(msg.Value)
		eventDispatcher.Dispatch(updateBalanceEvent)
	}
}
