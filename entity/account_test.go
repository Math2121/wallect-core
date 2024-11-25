package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("tesdte", "teste@gmail.com")

	account := NewAccount(client)
	assert.NotNil(t, account)

}

func TestCreditBalance(t *testing.T) {
	client, _ := NewClient("teste", "teste@gmail.com")

	account := NewAccount(client)
	account.Credit(100.0)

	assert.Equal(t, 100.0, account.Balance)
}

func TestDeposit(t *testing.T) {
	client, _ := NewClient("teste", "teste@gmail.com")

	account := NewAccount(client)
	account.Credit(100.0)
	account.Debit(50.0)

	assert.Equal(t, 50.0, account.Balance)
}


