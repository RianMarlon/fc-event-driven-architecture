package usecase

import (
	"balance-ms/internal/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type BalanceGatewayMock struct {
	mock.Mock
}

func (m *BalanceGatewayMock) GetByAccountID(accountID string) *entity.Balance {
	called := m.Called(accountID)
	if called.Get(0) == nil {
		return nil
	}
	return called.Get(0).(*entity.Balance)
}

func (m *BalanceGatewayMock) Create(balance *entity.Balance) error {
	return m.Called(balance).Error(0)
}

func (m *BalanceGatewayMock) Update(balance *entity.Balance) error {
	return m.Called(balance).Error(0)
}

func TestUpdateAccountsBalanceWhenBalanceExists(t *testing.T) {
	balanceGatewayMock := BalanceGatewayMock{}

	balanceFrom := &entity.Balance{
		ID:        "1",
		AccountID: "1",
		Amount:    100.0,
	}
	balanceTo := &entity.Balance{
		ID:        "2",
		AccountID: "2",
		Amount:    200.0,
	}

	balanceGatewayMock.On("GetByAccountID", balanceFrom.AccountID).Return(balanceFrom)
	balanceGatewayMock.On("GetByAccountID", balanceTo.AccountID).Return(balanceTo)
	balanceGatewayMock.On("Update", mock.Anything).Return(nil)

	usecase := NewUpdateAccountsBalanceUseCase(&balanceGatewayMock)
	err := usecase.Execute(UpdateAccountsBalanceInput{
		AccountIDFrom:      "1",
		AccountIDTo:        "2",
		BalanceAccountFrom: 30.0,
		BalanceAccountTo:   270.0,
	})

	assert.Nil(t, err)
	balanceGatewayMock.AssertExpectations(t)

	balanceGatewayMock.AssertNumberOfCalls(t, "GetByAccountID", 2)
	balanceGatewayMock.AssertNumberOfCalls(t, "Create", 0)
	balanceGatewayMock.AssertNumberOfCalls(t, "Update", 2)

	balanceGatewayMock.AssertCalled(t, "GetByAccountID", balanceFrom.AccountID)
	balanceGatewayMock.AssertCalled(t, "GetByAccountID", balanceTo.AccountID)

	balanceGatewayMock.AssertCalled(t, "Update", matchBalance(balanceFrom.AccountID, 30.0))
	balanceGatewayMock.AssertCalled(t, "Update", matchBalance(balanceTo.AccountID, 270.0))
}

func TestUpdateAccountsBalanceWhenBalanceDoesNotExist(t *testing.T) {
	balanceGatewayMock := BalanceGatewayMock{}

	balanceGatewayMock.On("GetByAccountID", mock.Anything).Return(nil)
	balanceGatewayMock.On("Create", mock.Anything).Return(nil)

	usecase := NewUpdateAccountsBalanceUseCase(&balanceGatewayMock)
	err := usecase.Execute(UpdateAccountsBalanceInput{
		AccountIDFrom:      "1",
		AccountIDTo:        "2",
		BalanceAccountFrom: 30.0,
		BalanceAccountTo:   270.0,
	})

	assert.Nil(t, err)
	balanceGatewayMock.AssertExpectations(t)

	balanceGatewayMock.AssertNumberOfCalls(t, "GetByAccountID", 2)
	balanceGatewayMock.AssertNumberOfCalls(t, "Create", 2)
	balanceGatewayMock.AssertNumberOfCalls(t, "Update", 0)

	balanceGatewayMock.AssertCalled(t, "GetByAccountID", "1")
	balanceGatewayMock.AssertCalled(t, "GetByAccountID", "2")

	balanceGatewayMock.AssertCalled(t, "Create", matchBalance("1", 30.0))
	balanceGatewayMock.AssertCalled(t, "Create", matchBalance("2", 270.0))
}

func TestUpdateAccountsBalanceWhenBalanceFromDoesNotExist(t *testing.T) {
	balanceGatewayMock := BalanceGatewayMock{}

	balanceTo := &entity.Balance{
		ID:        "2",
		AccountID: "2",
		Amount:    200.0,
	}

	balanceGatewayMock.On("GetByAccountID", "1").Return(nil)
	balanceGatewayMock.On("GetByAccountID", balanceTo.AccountID).Return(balanceTo)
	balanceGatewayMock.On("Create", mock.Anything).Return(nil)
	balanceGatewayMock.On("Update", mock.Anything).Return(nil)

	usecase := NewUpdateAccountsBalanceUseCase(&balanceGatewayMock)
	err := usecase.Execute(UpdateAccountsBalanceInput{
		AccountIDFrom:      "1",
		AccountIDTo:        "2",
		BalanceAccountFrom: 70.0,
		BalanceAccountTo:   40.0,
	})

	assert.Nil(t, err)
	balanceGatewayMock.AssertExpectations(t)

	balanceGatewayMock.AssertNumberOfCalls(t, "GetByAccountID", 2)
	balanceGatewayMock.AssertNumberOfCalls(t, "Create", 1)
	balanceGatewayMock.AssertNumberOfCalls(t, "Update", 1)

	balanceGatewayMock.AssertCalled(t, "GetByAccountID", "1")
	balanceGatewayMock.AssertCalled(t, "GetByAccountID", balanceTo.AccountID)

	balanceGatewayMock.AssertCalled(t, "Create", matchBalance("1", 70.0))
	balanceGatewayMock.AssertCalled(t, "Update", matchBalance(balanceTo.AccountID, 40.0))
}

func TestUpdateAccountsBalanceWhenBalanceToDoesNotExist(t *testing.T) {
	balanceGatewayMock := BalanceGatewayMock{}

	balanceFrom := &entity.Balance{
		ID:        "1",
		AccountID: "1",
		Amount:    100.0,
	}

	balanceGatewayMock.On("GetByAccountID", balanceFrom.AccountID).Return(balanceFrom)
	balanceGatewayMock.On("GetByAccountID", "2").Return(nil)
	balanceGatewayMock.On("Create", mock.Anything).Return(nil)
	balanceGatewayMock.On("Update", mock.Anything).Return(nil)

	usecase := NewUpdateAccountsBalanceUseCase(&balanceGatewayMock)
	err := usecase.Execute(UpdateAccountsBalanceInput{
		AccountIDFrom:      "1",
		AccountIDTo:        "2",
		BalanceAccountFrom: 1000.0,
		BalanceAccountTo:   2400.0,
	})

	assert.Nil(t, err)
	balanceGatewayMock.AssertExpectations(t)

	balanceGatewayMock.AssertNumberOfCalls(t, "GetByAccountID", 2)
	balanceGatewayMock.AssertNumberOfCalls(t, "Create", 1)
	balanceGatewayMock.AssertNumberOfCalls(t, "Update", 1)

	balanceGatewayMock.AssertCalled(t, "GetByAccountID", balanceFrom.AccountID)
	balanceGatewayMock.AssertCalled(t, "GetByAccountID", "2")

	balanceGatewayMock.AssertCalled(t, "Create", matchBalance("2", 2400.0))
	balanceGatewayMock.AssertCalled(t, "Update", matchBalance(balanceFrom.AccountID, 1000.0))
}

func matchBalance(accountID string, amount float64) interface{} {
	return mock.MatchedBy(func(b *entity.Balance) bool {
		return b.AccountID == accountID && b.Amount == amount
	})
}
