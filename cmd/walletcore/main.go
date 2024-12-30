package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Math2121/walletcore/database"
	"github.com/Math2121/walletcore/event"
	"github.com/Math2121/walletcore/pkg/eventos/pkg/events"
	"github.com/Math2121/walletcore/pkg/uow"
	createaccount "github.com/Math2121/walletcore/usecase/account/create_account"
	createclient "github.com/Math2121/walletcore/usecase/client/create_client"
	createtransaction "github.com/Math2121/walletcore/usecase/transaction/create_transaction"
	"github.com/Math2121/walletcore/web"
	"github.com/Math2121/walletcore/web/webserver"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDb(db)

	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDb", func(tx *sql.Tx) interface{} {
		return database.NewAccountDb(db)
	})
	uow.Register("TransactionDb", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDb(db)
	})

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent)

	webserver := webserver.NewWebServer(":3000")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/client", clientHandler.CreateClient)
	webserver.AddHandler("/account", accountHandler.CreateAccount)
	webserver.AddHandler("/transaction", transactionHandler.CreateTransaction)

	webserver.Start()

}
