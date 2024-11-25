package gateway

import "github.com/Math2121/walletcore/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Create(client *entity.Client) error

}