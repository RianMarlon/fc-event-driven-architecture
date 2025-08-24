package usecase

import (
	"balance-ms/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type BalanceGatewayMock struct {
	mock.Mock
}

func (m *BalanceGatewayMock) GetByAccountID(accountID string) *entity.Balance {
	args := m.Called(accountID)
	if balance, ok := args.Get(0).(*entity.Balance); ok {
		return balance
	}
	return nil
}

func (m *BalanceGatewayMock) Create(balance *entity.Balance) error {
	args := m.Called(balance)
	return args.Error(0)
}

func (m *BalanceGatewayMock) Update(balance *entity.Balance) error {
	args := m.Called(balance)
	return args.Error(0)
}

func TestGetBalanceByAccountUsecase_Execute(t *testing.T) {
	balanceGatewayMock := BalanceGatewayMock{}
	balanceGatewayMock.On("GetByAccountID", mock.Anything).Return(&entity.Balance{
		ID:        "1",
		AccountID: "1",
		Amount:    100.0,
	})

	usecase := NewGetBalanceByAccountUsecase(&balanceGatewayMock)
	input := GetBalanceByAccountInputDTO{
		AccountID: "1",
	}
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, "1", output.ID)
	assert.Equal(t, 100.0, output.Amount)
	balanceGatewayMock.AssertExpectations(t)
	balanceGatewayMock.AssertNumberOfCalls(t, "GetByAccountID", 1)
}

func TestGetBalanceByAccountUsecase_Execute_BalanceNotFound(t *testing.T) {
	balanceGatewayMock := BalanceGatewayMock{}
	balanceGatewayMock.On("GetByAccountID", mock.Anything).Return(nil)

	usecase := NewGetBalanceByAccountUsecase(&balanceGatewayMock)
	input := GetBalanceByAccountInputDTO{
		AccountID: "1",
	}
	output, err := usecase.Execute(input)
	assert.Nil(t, output)
	assert.Error(t, err)
	assert.Equal(t, "balance not found", err.Error())
	balanceGatewayMock.AssertExpectations(t)
	balanceGatewayMock.AssertNumberOfCalls(t, "GetByAccountID", 1)
}
