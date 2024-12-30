package createtransaction

import (
	"context"

	"github.com/Math2121/walletcore/entity"
	"github.com/Math2121/walletcore/gateway"
	"github.com/Math2121/walletcore/pkg/eventos/pkg/events"
	"github.com/Math2121/walletcore/pkg/uow"
)

type CreateTransactionInputDto struct {
	AccountIDFrom string  `json:"account_id"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutputDto struct {
	ID            string  `json:"id"`
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionUseCase struct {
	Uow                uow.UowInterface
	EventDispatcher    events.EventDispatcherInterface
	TransactionCreated events.EventInterface
}

func NewCreateTransactionUseCase(Uow uow.UowInterface,
	eventDispatcher events.EventDispatcherInterface,
	transactionCreated events.EventInterface,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		Uow:                Uow,
		EventDispatcher:    eventDispatcher,
		TransactionCreated: transactionCreated,
	}
}

func (u *CreateTransactionUseCase) Execute(ctx context.Context, input CreateTransactionInputDto) (*CreateTransactionOutputDto, error) {
	output := &CreateTransactionOutputDto{}
	err := u.Uow.Do(ctx, func(_ *uow.Uow) error {

		accountRepo := u.getAccountRepository(ctx)
		transactionRepo := u.getTransactionRepository(ctx)

		accountFrom, err := accountRepo.FindById(input.AccountIDFrom)
		if err != nil {
			return err
		}
		accountTo, err := accountRepo.FindById(input.AccountIDTo)
		if err != nil {
			return err
		}

		transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
		if err != nil {
			return err
		}

		err = accountRepo.UpdateBalance(accountFrom)
		if err != nil {
			return err
		}

		err = accountRepo.UpdateBalance(accountTo)
		if err != nil {
			return err
		}

		err = transactionRepo.Create(transaction)
		if err != nil {
			return err
		}
		output.ID = transaction.ID
		output.AccountIDFrom = input.AccountIDFrom
		output.AccountIDTo = input.AccountIDTo
		output.Amount = input.Amount
		return nil

	})
	
	if err!= nil {
        return nil, err
    }

	u.TransactionCreated.SetPayload(output)
	u.EventDispatcher.Dispatch(u.TransactionCreated)

	return output, nil

}

func (u *CreateTransactionUseCase) getAccountRepository(ctx context.Context) gateway.AccountGateway {
	repo, err := u.Uow.GetRepository(ctx, "AccountDb")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.AccountGateway)

}

func (u *CreateTransactionUseCase) getTransactionRepository(ctx context.Context) gateway.TransactionGateway {
	repo, err := u.Uow.GetRepository(ctx, "TransactionDb")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.TransactionGateway)
}
