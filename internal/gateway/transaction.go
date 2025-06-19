package gateway

import "walletcore/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
