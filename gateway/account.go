package gateway

import "github.com/Math2121/walletcore/entity"

type AccountGateway interface {
	Create(account *entity.Account) error
	FindById(id string) (*entity.Account, error)
	UpdateBalance(account *entity.Account) error
}
