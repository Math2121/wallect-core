package createaccount_test

import (
	"testing"

	"github.com/Math2121/walletcore/entity"
	createaccount "github.com/Math2121/walletcore/usecase/account/create_account"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) Create(client *entity.Client) error {
	args := m.Called(client)
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

func TestCreateAccountUseCase(t *testing.T) {
	client, _ := entity.NewClient("teste", "teste@gmail.com")
	clientMock := &ClientGatewayMock{}
	clientMock.On("Get", client.ID).Return(client, nil)

	accountMock := &AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	useCase := createaccount.NewCreateAccountUseCase(accountMock, clientMock)
	input := createaccount.CreateAccountInputDto{
		ClientID: client.ID,
	}

	output, err := useCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output.AccountID)
	accountMock.AssertExpectations(t)
	clientMock.AssertExpectations(t)

}
