package createtransaction

import (
	"github.com/Math2121/walletcore/entity"
	"github.com/Math2121/walletcore/gateway"
	"github.com/Math2121/walletcore/pkg/eventos/pkg/events"
)

type CreateTransactionInputDto struct {
	AccountIDFrom string  `json:"account_id"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutputDto struct {
	ID string `json:"id"`
}

type CreateTransactionUseCase struct {
	transactionGateway gateway.TransactionGateway
	accountGateway     gateway.AccountGateway
	EventDispatcher   events.EventDispatcherInterface
    TransactionCreated events.EventInterface
}

func NewCreateTransactionUseCase(transactionGateway gateway.TransactionGateway, 
	accountGateway gateway.AccountGateway,
	eventDispatcher events.EventDispatcherInterface,
	transactionCreated events.EventInterface,
	) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		transactionGateway: transactionGateway,
		accountGateway:     accountGateway,
		EventDispatcher:     eventDispatcher,
        TransactionCreated: transactionCreated,

	}
}

func (u *CreateTransactionUseCase) Execute(input CreateTransactionInputDto) (*CreateTransactionOutputDto, error) {
	accountFrom, err := u.accountGateway.FindById(input.AccountIDFrom)
	if err != nil {
		return nil, err
	}
	accountTo, err := u.accountGateway.FindById(input.AccountIDTo)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}
	err = u.transactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}
	output := &CreateTransactionOutputDto{ID: transaction.ID}
	u.TransactionCreated.SetPayload(output)
	u.EventDispatcher.Dispatch(u.TransactionCreated)

	return output, nil

}
