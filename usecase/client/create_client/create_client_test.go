package createclient_test

import (
	"testing"

	"github.com/Math2121/walletcore/entity"
	createclient "github.com/Math2121/walletcore/usecase/client/create_client"
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
func TestCreateClientUseCase(t *testing.T){

	clientGatewayMock := &ClientGatewayMock{}
	clientGatewayMock.On("Create", mock.Anything).Return(nil)

	useCase := createclient.NewCreateClientUseCase(clientGatewayMock)
	output, err := useCase.Execute(createclient.CreateClientInputDto{
		Name:      "teste",
        Email:     "teste@gmail.com",
	})
	assert.Nil(t, err)
	assert.Equal(t, "teste", output.Name)
	clientGatewayMock.AssertExpectations(t)

}