package database_test

import (
	"database/sql"
	"testing"

	"github.com/Math2121/walletcore/database"
	"github.com/Math2121/walletcore/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"

)

type ClientDbTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDb *database.ClientDB
}

func (s *ClientDbTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date) ")
	s.clientDb = database.NewClientDB(db)
}

func (s *ClientDbTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}
func TestClientDbTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDbTestSuite))
}

func (s *ClientDbTestSuite) TestGet() {
	client, _ := entity.NewClient("Teste", "j@gmail.com")

	s.clientDb.Create(client)

	clientDb, err := s.clientDb.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, clientDb.ID)
	s.Equal(client.Name, clientDb.Name)
	s.Equal(client.Email, clientDb.Email)
}
