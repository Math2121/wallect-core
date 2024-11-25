package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {

	client, err := NewClient("teste", "teste@gmail.com")

	assert.Nil(t, err)
	assert.Equal(t, "teste", client.Name)
	assert.Equal(t, "teste@gmail.com", client.Email)
}


func TestCreateNewClientWhenArgsAreInvalid(t *testing.T){
	
	_, err := NewClient("", "")

	assert.NotNil(t, err)
}

func TestUpdateNewClient(t *testing.T){
	
    client := &Client{Name: "teste", Email: "teste@gmail.com"}

    err := client.Update("novo nome", "novo email@gmail.com")

    assert.Nil(t, err)
    assert.Equal(t, "novo nome", client.Name)
    assert.Equal(t, "novo email@gmail.com", client.Email)
}

func TestUpdateNewClientWhenArgsAreInvalid(t *testing.T){
	
    client := &Client{Name: "teste", Email: "teste@gmail.com"}

    err := client.Update("", "")

    assert.NotNil(t, err)
}

func TestAddAccountClient(t *testing.T){
	
    client := &Client{Name: "teste", Email: "teste@gmail.com"}

    account := &Account{Client: client}

    err := client.AddAccounts(account)

    assert.Nil(t, err)
    assert.Equal(t, 1, len(client.Accounts))
    assert.Equal(t, account, client.Accounts[0])
}