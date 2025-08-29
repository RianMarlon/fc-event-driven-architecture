package kafka_handler

import (
	usecase "balance-ms/internal/app/usecase/update_accounts_balance"
	"balance-ms/pkg/events"
	"encoding/json"
	"fmt"
	"sync"
)

type EventPayload struct {
	Payload usecase.UpdateAccountsBalanceInput
}

type BalanceUpdatedKafkaHandler struct {
	UpdateAccountsBalanceUseCase usecase.UpdateAccountsBalanceUseCase
}

func NewBalanceUpdatedKafkaHandler(updateAccountsBalanceUseCase usecase.UpdateAccountsBalanceUseCase) *BalanceUpdatedKafkaHandler {
	return &BalanceUpdatedKafkaHandler{UpdateAccountsBalanceUseCase: updateAccountsBalanceUseCase}
}

func (h *BalanceUpdatedKafkaHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	var input EventPayload
	err := json.Unmarshal(event.GetPayload().([]byte), &input)
	if err != nil {
		fmt.Println("Error unmarshalling event", err)
	}
	fmt.Println("BalanceUpdatedKafkaHandler called", input.Payload)
	h.UpdateAccountsBalanceUseCase.Execute(input.Payload)
}
