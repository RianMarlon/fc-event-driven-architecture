package main

import (
	"context"
	"database/sql"
	"log"
	"walletcore/internal/database"
	event "walletcore/internal/events"
	"walletcore/internal/events/handler"
	createaccount "walletcore/internal/usecase/create_account"
	createclient "walletcore/internal/usecase/create_client"
	createtransaction "walletcore/internal/usecase/create_transaction"
	"walletcore/internal/web"
	"walletcore/internal/web/webserver"
	"walletcore/pkg/events"
	"walletcore/pkg/kafka"
	"walletcore/pkg/unit_of_work"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/wallet?parseTime=true")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9094",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	transactionCreatedEvent := event.NewTransactionCreated()

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	unitOfWork := unit_of_work.NewUnitOfWork(ctx, db)

	unitOfWork.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})

	unitOfWork.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(unitOfWork, eventDispatcher, transactionCreatedEvent)

	webserver := webserver.NewWebServer(":3000")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)
	webserver.Start()
}
