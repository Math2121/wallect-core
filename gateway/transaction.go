package gateway

import "github.com/Math2121/walletcore/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}