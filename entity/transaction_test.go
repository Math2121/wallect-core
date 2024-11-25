package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {

	client, _ := NewClient("teste", "teste@gmail.com")
	clinet2, err := NewClient("teste2", "teste2@gmail.com")

	account1 := NewAccount(client)
	account2 := NewAccount(clinet2)

	account1.Credit(200)
	account2.Credit(100)

	transaction, err := NewTransaction(account1, account2, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, float64(100), account1.Balance)
	assert.Equal(t, float64(200), account2.Balance)

}

func TestCreateTransactionWithInsuficientBalance(t *testing.T) {
	client, _ := NewClient("teste", "teste@gmail.com")
	clinet2, err := NewClient("teste2", "teste2@gmail.com")

	account1 := NewAccount(client)
	account2 := NewAccount(clinet2)

	account1.Credit(100)

	transaction, err := NewTransaction(account1, account2, 200)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, float64(100), account1.Balance)
	assert.Equal(t, float64(0), account2.Balance)
	assert.Equal(t, "Insufficient balance", err.Error())
}
