package usecase

import (
	"balance-ms/internal/gateway"
	"errors"
)

type GetBalanceByAccountInputDTO struct {
	AccountID string `json:"account_id"`
}

type GetBalanceByAccountOutputDTO struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}

type GetBalanceByAccountUsecase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewGetBalanceByAccountUsecase(balanceGateway gateway.BalanceGateway) *GetBalanceByAccountUsecase {
	return &GetBalanceByAccountUsecase{
		BalanceGateway: balanceGateway,
	}
}

func (u *GetBalanceByAccountUsecase) Execute(input GetBalanceByAccountInputDTO) (*GetBalanceByAccountOutputDTO, error) {
	balance := u.BalanceGateway.GetByAccountID(input.AccountID)
	if balance == nil {
		return nil, errors.New("balance not found")
	}

	return &GetBalanceByAccountOutputDTO{
		ID:     balance.ID,
		Amount: balance.Amount,
	}, nil
}
