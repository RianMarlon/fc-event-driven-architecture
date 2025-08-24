package usecase

import (
	"balance-ms/internal/entity"
	"balance-ms/internal/gateway"
)

type UpdateAccountsBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
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

func NewUpdateAccountsBalanceUseCase(balanceGateway gateway.BalanceGateway) *UpdateAccountsBalanceUseCase {
	return &UpdateAccountsBalanceUseCase{BalanceGateway: balanceGateway}
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
	balanceExists := u.BalanceGateway.GetByAccountID(input.AccountID)
	if balanceExists == nil {
		balance, err := entity.NewBalance(input.AccountID, input.Balance)
		if err != nil {
			return err
		}
		err = u.BalanceGateway.Create(balance)
		if err != nil {
			return err
		}
		return nil
	}
	balanceExists.UpdateBalance(input.Balance)
	err := u.BalanceGateway.Update(balanceExists)
	if err != nil {
		return err
	}
	return nil
}
