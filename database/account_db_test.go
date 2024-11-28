package database_test

import (
	"database/sql"
	"testing"

	"github.com/Math2121/walletcore/database"
	"github.com/Math2121/walletcore/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDbTestSuite struct {
	suite.Suite
	accountDb *database.AccountDb
	client    *entity.Client
	db        *sql.DB
}

func (s *AccountDbTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date) ")
	db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance int, created_at date) ")

	s.accountDb = database.NewAccountDb(db)
	s.client, _ = entity.NewClient("teste", "teste@gmail.com")

}

func (s *AccountDbTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.Nil(s.db.Exec("DROP TABLE clients"))
	s.Nil(s.db.Exec("DROP TABLE accounts"))
}

func TestAccountDbTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDbTestSuite))

}

func (s *AccountDbTestSuite) TestCreateAccount(t *testing.T) {
	account := entity.NewAccount(s.client)
	err := s.accountDb.Create(account)
	s.Nil(err)
}

func (s *AccountDbTestSuite) TestFindByClientId(t *testing.T) {
	account := entity.NewAccount(s.client)
    err := s.accountDb.Create(account)
    s.Nil(err)

    accounts, err := s.accountDb.FindById(account.ID)
    s.Nil(err)
    s.Len(accounts, 1)
    s.Equal(account.ID, accounts.ID)
	s.Equal(account.Client.ID, accounts.Client.ID)
	s.Equal(account.Balance, accounts.Balance)
	s.Equal(account.CreatedAt, accounts.CreatedAt)

}


