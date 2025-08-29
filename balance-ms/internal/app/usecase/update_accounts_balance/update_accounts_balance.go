package usecase

import (
	"balance-ms/internal/domain/entity"
	"balance-ms/internal/domain/repository"
)

type UpdateAccountsBalanceUseCase struct {
	BalanceRepository repository.BalanceRepository
}

type UpdateAccountsBalanceInput struct {
	AccountIDFrom      string  `json:"account_id_from"`
	AccountIDTo        string  `json:"account_id_to"`
	BalanceAccountFrom float64 `json:"balance_account_id_from"`
	BalanceAccountTo   float64 `json:"balance_account_id_to"`
}

type CreateBalanceInput struct {
	AccountID string
	Balance   float64
}

func NewUpdateAccountsBalanceUseCase(balanceRepository repository.BalanceRepository) *UpdateAccountsBalanceUseCase {
	return &UpdateAccountsBalanceUseCase{BalanceRepository: balanceRepository}
}

func (u *UpdateAccountsBalanceUseCase) Execute(input UpdateAccountsBalanceInput) error {
	err := u.createOrUpdateBalance(CreateBalanceInput{
		AccountID: input.AccountIDFrom,
		Balance:   input.BalanceAccountFrom,
	})
	if err != nil {
		return err
	}

	err = u.createOrUpdateBalance(CreateBalanceInput{
		AccountID: input.AccountIDTo,
		Balance:   input.BalanceAccountTo,
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *UpdateAccountsBalanceUseCase) createOrUpdateBalance(input CreateBalanceInput) error {
	balanceExists := u.BalanceRepository.GetByAccountID(input.AccountID)
	if balanceExists == nil {
		balance, err := entity.NewBalance(input.AccountID, input.Balance)
		if err != nil {
			return err
		}
		err = u.BalanceRepository.Create(balance)
		if err != nil {
			return err
		}
		return nil
	}
	balanceExists.UpdateBalance(input.Balance)
	err := u.BalanceRepository.Update(balanceExists)
	if err != nil {
		return err
	}
	return nil
}
