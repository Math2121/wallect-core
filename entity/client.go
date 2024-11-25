package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created"`
	Accounts  []*Account
}

func NewClient(name string, email string) (*Client, error) {

	client := &Client{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}
	err := client.Validate()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) Validate() error {
	if c.Name == "" {
		return errors.New("Name is required")
	}

	if c.Email == "" {
		return errors.New("Email is required")
	}

	return nil

}

func (c *Client) Update(name string, email string) error {
	c.Name = name
	c.Email = email

	errors := c.Validate()
	if errors != nil {
		return errors
	}
	return nil
}

func (c *Client) AddAccounts(account *Account) error {
	if account.Client.ID != c.ID {
		return errors.New("Account belongs to another client")
	}
	c.Accounts = append(c.Accounts, account)
	return nil
}
