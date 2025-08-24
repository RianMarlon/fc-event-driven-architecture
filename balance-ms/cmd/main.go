package main

import (
	"balance-ms/internal/database"
	"balance-ms/internal/event"
	"balance-ms/internal/event/handler"
	getBalanceByAccountUsecase "balance-ms/internal/usecase/get_balance_by_account"
	updateAccountsBalanceUsecase "balance-ms/internal/usecase/update_accounts_balance"
	"balance-ms/internal/web"
	"balance-ms/internal/web/webserver"
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
	updateAccountsBalanceUseCase := updateAccountsBalanceUsecase.NewUpdateAccountsBalanceUseCase(balanceGateway)
	getBalanceByAccountUseCase := getBalanceByAccountUsecase.NewGetBalanceByAccountUsecase(balanceGateway)

	server := webserver.NewWebServer(":3003")
	balanceHandler := web.NewWebBalanceHandler(*getBalanceByAccountUseCase)
	server.AddHandler(webserver.Handler{
		Path:        "/balances/{account_id}",
		Method:      "GET",
		HandlerFunc: balanceHandler.GetBalanceByAccount,
	})
	go server.Start()

	eventDispatcher.Register("BalanceUpdated", handler.NewBalanceUpdatedKafkaHandler(*updateAccountsBalanceUseCase))
	go kafkaConsumer.Consume(msgChan)

	for msg := range msgChan {
		updateBalanceEvent := event.NewBalanceUpdated()
		updateBalanceEvent.SetPayload(msg.Value)
		eventDispatcher.Dispatch(updateBalanceEvent)
	}
}
