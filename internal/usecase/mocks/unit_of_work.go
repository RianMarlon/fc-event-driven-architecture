package mocks

import (
	"context"
	"walletcore/pkg/unit_of_work"

	"github.com/stretchr/testify/mock"
)

type UnitOfWorkMock struct {
	mock.Mock
}

func (m *UnitOfWorkMock) Register(name string, fc unit_of_work.RepositoryFactory) {
	m.Called(name, fc)
}

func (m *UnitOfWorkMock) GetRepository(ctx context.Context, name string) (interface{}, error) {
	args := m.Called(name)
	return args.Get(0), args.Error(1)
}

func (m *UnitOfWorkMock) Do(ctx context.Context, fn func(unit_of_work *unit_of_work.UnitOfWork) error) error {
	args := m.Called(fn)
	return args.Error(0)
}

func (m *UnitOfWorkMock) CommitOrRollback() error {
	args := m.Called()
	return args.Error(0)
}

func (m *UnitOfWorkMock) Rollback() error {
	args := m.Called()
	return args.Error(0)
}

func (m *UnitOfWorkMock) UnRegister(name string) {
	m.Called(name)
}
