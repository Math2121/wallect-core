package createclient

import (
	"time"

	"github.com/Math2121/walletcore/entity"
	"github.com/Math2121/walletcore/gateway"
)

type CreateClientInputDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateClientOutputDto struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateClientUseCase struct {
	ClientGateway gateway.ClientGateway
}

func NewCreateClientUseCase(clientGateway gateway.ClientGateway) *CreateClientUseCase {
	return &CreateClientUseCase{ClientGateway: clientGateway}
}

func (u *CreateClientUseCase) Execute(input CreateClientInputDto) (*CreateClientOutputDto, error) {
	client, err := entity.NewClient(input.Name, input.Email)
	if err != nil {
		return nil, err
	}

	err = u.ClientGateway.Create(client)
	if err != nil {
		return nil, err
	}
	return &CreateClientOutputDto{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
	}, nil
}
