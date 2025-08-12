package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBalance(t *testing.T) {
	balance, err := NewBalance("123", 100.0)
	assert.Nil(t, err)
	assert.NotNil(t, balance)
	assert.NotEmpty(t, balance.ID)
	assert.Equal(t, "123", balance.AccountID)
	assert.Equal(t, 100.0, balance.Amount)
	assert.NotEmpty(t, balance.CreatedAt)
	assert.NotEmpty(t, balance.UpdatedAt)
}

func TestNewBalanceNegativeAmount(t *testing.T) {
	balance, err := NewBalance("123", -100.0)
	assert.Nil(t, balance)
	assert.Error(t, err)
	assert.Equal(t, "amount cannot be negative", err.Error())
}

func TestBalanceCredit(t *testing.T) {
	balance, err := NewBalance("123", 100.0)
	assert.Nil(t, err)
	assert.NotNil(t, balance)
	assert.Equal(t, 100.0, balance.Amount)
	err = balance.Credit(50.0)
	assert.Nil(t, err)
	assert.Equal(t, 150.0, balance.Amount)
}

func TestBalanceDebit(t *testing.T) {
	balance, err := NewBalance("123", 100.0)
	assert.Nil(t, err)
	assert.NotNil(t, balance)
	assert.Equal(t, 100.0, balance.Amount)
	err = balance.Debit(50.0)
	assert.Nil(t, err)
	assert.Equal(t, 50.0, balance.Amount)
}

func TestBalanceDebitInsufficientBalance(t *testing.T) {
	balance, err := NewBalance("123", 100.0)
	assert.Nil(t, err)
	assert.NotNil(t, balance)
	assert.Equal(t, 100.0, balance.Amount)
	err = balance.Debit(150.0)
	assert.Error(t, err)
	assert.Equal(t, "insufficient balance", err.Error())
}
