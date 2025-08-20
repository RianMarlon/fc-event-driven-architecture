package handler

import (
	usecase "balance-ms/internal/usecase/update_accounts_balance"
	"balance-ms/pkg/events"
	"fmt"
	"sync"
)

type BalanceUpdatedKafkaHandler struct {
	UpdateAccountsBalanceUseCase usecase.UpdateAccountsBalanceUseCase
}

func NewBalanceUpdatedKafkaHandler(updateAccountsBalanceUseCase usecase.UpdateAccountsBalanceUseCase) *BalanceUpdatedKafkaHandler {
	return &BalanceUpdatedKafkaHandler{UpdateAccountsBalanceUseCase: updateAccountsBalanceUseCase}
}

func (h *BalanceUpdatedKafkaHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	h.UpdateAccountsBalanceUseCase.Execute(event.GetPayload().(usecase.UpdateAccountsBalanceInput))
	fmt.Println("BalanceUpdatedKafkaHandler called")
	wg.Done()
}
