package createtransaction

import (
	"context"
	"walletcore/internal/entity"
	"walletcore/internal/gateway"
	"walletcore/pkg/events"
	"walletcore/pkg/unit_of_work"
)

type CreateTransactionInputDTO struct {
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionOutputDTO struct {
	ID            string  `json:"id"`
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type BalanceUpdatedOutputDTO struct {
	AccountIDFrom      string  `json:"account_id_from"`
	AccountIDTo        string  `json:"account_id_to"`
	BalanceAccountFrom float64 `json:"balance_account_id_from"`
	BalanceAccountTo   float64 `json:"balance_account_id_to"`
}

type CreateTransactionUseCase struct {
	UnitOfWork         unit_of_work.UnitOfWorkInterface
	EventDispatcher    events.EventDispatcherInterface
	TransactionCreated events.EventInterface
	BalanceUpdated     events.EventInterface
}

func NewCreateTransactionUseCase(
	UnitOfWork unit_of_work.UnitOfWorkInterface,
	eventDispatcher events.EventDispatcherInterface,
	transactionCreated events.EventInterface,
	balanceUpdated events.EventInterface,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		UnitOfWork:         UnitOfWork,
		EventDispatcher:    eventDispatcher,
		TransactionCreated: transactionCreated,
		BalanceUpdated:     balanceUpdated,
	}
}

func (uc *CreateTransactionUseCase) Execute(ctx context.Context, input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	output := &CreateTransactionOutputDTO{}
	balanceUpdatedOutput := &BalanceUpdatedOutputDTO{}
	err := uc.UnitOfWork.Do(ctx, func(_ *unit_of_work.UnitOfWork) error {
		accountRepository := uc.getAccountRepository(ctx)
		transactionRepository := uc.getTransactionRepository(ctx)

		accountFrom, err := accountRepository.FindByID(input.AccountIDFrom)
		if err != nil {
			return err
		}

		accountTo, err := accountRepository.FindByID(input.AccountIDTo)
		if err != nil {
			return err
		}

		transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
		if err != nil {
			return err
		}

		err = accountRepository.UpdateBalance(accountFrom)
		if err != nil {
			return err
		}

		err = accountRepository.UpdateBalance(accountTo)
		if err != nil {
			return err
		}

		err = transactionRepository.Create(transaction)
		if err != nil {
			return err
		}

		output.ID = transaction.ID
		output.AccountIDFrom = input.AccountIDFrom
		output.AccountIDTo = input.AccountIDTo
		output.Amount = input.Amount

		balanceUpdatedOutput.AccountIDFrom = input.AccountIDFrom
		balanceUpdatedOutput.AccountIDTo = input.AccountIDTo
		balanceUpdatedOutput.BalanceAccountFrom = accountFrom.Balance
		balanceUpdatedOutput.BalanceAccountTo = accountTo.Balance

		return nil
	})
	if err != nil {
		return nil, err
	}

	uc.TransactionCreated.SetPayload(output)
	uc.EventDispatcher.Dispatch(uc.TransactionCreated)

	uc.BalanceUpdated.SetPayload(balanceUpdatedOutput)
	uc.EventDispatcher.Dispatch(uc.BalanceUpdated)

	return output, nil
}

func (uc *CreateTransactionUseCase) getAccountRepository(ctx context.Context) gateway.AccountGateway {
	repository, err := uc.UnitOfWork.GetRepository(ctx, "AccountDB")
	if err != nil {
		panic(err)
	}
	return repository.(gateway.AccountGateway)
}

func (uc *CreateTransactionUseCase) getTransactionRepository(ctx context.Context) gateway.TransactionGateway {
	repository, err := uc.UnitOfWork.GetRepository(ctx, "TransactionDB")
	if err != nil {
		panic(err)
	}
	return repository.(gateway.TransactionGateway)
}
