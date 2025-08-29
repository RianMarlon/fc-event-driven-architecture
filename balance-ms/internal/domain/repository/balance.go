package repository

import entity "balance-ms/internal/domain/entity"

type BalanceRepository interface {
	GetByAccountID(accountID string) *entity.Balance
	Create(balance *entity.Balance) error
	Update(balance *entity.Balance) error
}
