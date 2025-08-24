package gateway

import "balance-ms/internal/entity"

type BalanceGateway interface {
	GetByAccountID(accountID string) *entity.Balance
	Create(balance *entity.Balance) error
	Update(balance *entity.Balance) error
}
