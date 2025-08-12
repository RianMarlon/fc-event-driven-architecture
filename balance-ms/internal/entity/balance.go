package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Balance struct {
	ID        string
	AccountID string
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBalance(accountID string, amount float64) (*Balance, error) {
	if amount < 0 {
		return nil, errors.New("amount cannot be negative")
	}
	return &Balance{
		ID:        uuid.New().String(),
		AccountID: accountID,
		Amount:    amount,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (b *Balance) Credit(amount float64) error {
	b.Amount += amount
	b.UpdatedAt = time.Now()
	return nil
}

func (b *Balance) Debit(amount float64) error {
	if b.Amount < amount {
		return errors.New("insufficient balance")
	}
	b.Amount -= amount
	b.UpdatedAt = time.Now()
	return nil
}
