package usecase

import (
	"errors"

	"balance-ms/internal/domain/repository"
)

type GetBalanceByAccountInputDTO struct {
	AccountID string `json:"account_id"`
}

type GetBalanceByAccountOutputDTO struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}

type GetBalanceByAccountUsecase struct {
	BalanceRepository repository.BalanceRepository
}

func NewGetBalanceByAccountUsecase(balanceRepository repository.BalanceRepository) *GetBalanceByAccountUsecase {
	return &GetBalanceByAccountUsecase{
		BalanceRepository: balanceRepository,
	}
}

func (u *GetBalanceByAccountUsecase) Execute(input GetBalanceByAccountInputDTO) (*GetBalanceByAccountOutputDTO, error) {
	balance := u.BalanceRepository.GetByAccountID(input.AccountID)
	if balance == nil {
		return nil, errors.New("balance not found")
	}

	return &GetBalanceByAccountOutputDTO{
		ID:     balance.ID,
		Amount: balance.Amount,
	}, nil
}
