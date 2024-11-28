package database_test

import (
	"database/sql"
	"testing"

	"github.com/Math2121/walletcore/database"
	"github.com/Math2121/walletcore/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDbSuite struct {
	suite.Suite
	db            *sql.DB
	client        *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDb *database.TransactionDb
}

func (s *TransactionDbSuite) SetupSuite() {

	db, err := sql.Open("sqlite3", "./wallet_test.db")
	s.Nil(err)
	s.db = db
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date) ")
	db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance int, created_at date) ")
	db.Exec("Create table transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), balance int, created_at date)")
	s.client = &entity.Client{Name: "Alice", Email: "alice@gmail.com"}
	s.client2 = &entity.Client{Name: "Bob", Email: "bob@gmail.com"}

	s.accountFrom = &entity.Account{Client: s.client, Balance: 1000}
	s.accountFrom.Balance = 1000

	s.accountTo = &entity.Account{Client: s.client2, Balance: 500}
	s.accountTo.Balance = 1000

	s.transactionDb = database.NewTransactionDb(db)

}

func (s *TransactionDbSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDbSuite(t *testing.T) {
	suite.Run(t, new(TransactionDbSuite))
}

func (s *TransactionDbSuite) TestCreateTransaction(t *testing.T) {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 200)
	s.Nil(err)
	err = s.transactionDb.Create(transaction)
	s.Nil(err)

}
