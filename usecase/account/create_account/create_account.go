package createaccount

import (
	"github.com/Math2121/walletcore/entity"
	"github.com/Math2121/walletcore/gateway"
)

type CreateAccountInputDto struct {
	ClientID string `json:"client_id"`
}

type CreateAccountOutputDto struct {
	AccountID string `json:"account_id"`
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	CLientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(a gateway.AccountGateway, c gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{AccountGateway: a, CLientGateway: c}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDto) (*CreateAccountOutputDto, error) {
	client, err := uc.CLientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(client)
	err = uc.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}
	return &CreateAccountOutputDto{AccountID: account.ID}, nil

}
