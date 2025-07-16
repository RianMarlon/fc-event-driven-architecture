package createtransaction

import (
	"context"
	"testing"
	"walletcore/internal/entity"
	event "walletcore/internal/events"
	"walletcore/internal/usecase/mocks"
	"walletcore/pkg/events"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *AccountGatewayMock) UpdateBalance(client *entity.Account) error {
	args := m.Called(client)
	return args.Error(1)
}

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("John Doe", "j@j.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("Jane Doe", "g@g.com")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	mockUnitOfWork := &mocks.UnitOfWorkMock{}
	mockUnitOfWork.On("Do", mock.Anything, mock.Anything).Return(nil)

	inputDto := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}
	dispatcher := events.NewEventDispatcher()
	transactionCreated := event.NewTransactionCreated()
	balanceUpdated := event.NewBalanceUpdated()

	ctx := context.Background()
	uc := NewCreateTransactionUseCase(mockUnitOfWork, dispatcher, transactionCreated, balanceUpdated)
	output, err := uc.Execute(ctx, inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output)

	mockUnitOfWork.AssertExpectations(t)
	mockUnitOfWork.AssertNumberOfCalls(t, "Do", 1)
}
