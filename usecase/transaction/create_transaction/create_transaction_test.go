package createtransaction_test

import (
	"context"
	"testing"

	"github.com/Math2121/walletcore/entity"
	"github.com/Math2121/walletcore/event"
	"github.com/Math2121/walletcore/pkg/eventos/pkg/events"
	"github.com/Math2121/walletcore/usecase/mocks"
	createtransaction "github.com/Math2121/walletcore/usecase/transaction/create_transaction"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Create(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindById(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateTransactionUseCase(t *testing.T) {
	client, _ := entity.NewClient("client1", "teste@test.com")
	account := entity.NewAccount(client)
	account.Credit(1000)

	client2, _ := entity.NewClient("client2", "teste@test.com")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	mockUow := &mocks.UowMocks{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)



	inputDto := createtransaction.CreateTransactionInputDto{
		Amount:        100,
		AccountIDFrom: account.ID,
		AccountIDTo:   account2.ID,
	}

	dispatcher := events.NewEventDispatcher()
	event := event.NewTransactionCreated()
	ctx := context.Background()

	uc := createtransaction.NewCreateTransactionUseCase(mockUow, dispatcher, event)
	output, err := uc.Execute(ctx, inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)

}
