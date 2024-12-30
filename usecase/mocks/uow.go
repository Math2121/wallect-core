package mocks

import (
	"context"

	"github.com/Math2121/walletcore/pkg/uow"
	"github.com/stretchr/testify/mock"
)

type UowMocks struct {
	mock.Mock
}


func (m *UowMocks) Register(name string, fc uow.RepositoryFactory) {
    m.Called(name, fc)
}

func (m *UowMocks) GetRepository(ctx context.Context, name string) (interface{}, error) {
    args := m.Called(ctx, name)
    return args.Get(0), args.Error(1)
}

func (m *UowMocks) Do(ctx context.Context, fn func(uow *uow.Uow) error) error {
    args := m.Called(ctx, fn)
    return args.Error(0)

}

func (m *UowMocks) CommitOrRollback() error {
    args := m.Called()
    return args.Error(0)
}

func (m *UowMocks) Rollback() error {
    args := m.Called()
    return args.Error(0)
}

func (m *UowMocks) UnRegister(name string) {
    m.Called(name)
}