package web

import (
	usecase "balance-ms/internal/usecase/get_balance_by_account"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type WebBalanceHandler struct {
	GetBalanceByAccountUseCase usecase.GetBalanceByAccountUsecase
}

func NewWebBalanceHandler(getBalanceByAccountUseCase usecase.GetBalanceByAccountUsecase) *WebBalanceHandler {
	return &WebBalanceHandler{GetBalanceByAccountUseCase: getBalanceByAccountUseCase}
}

func (h *WebBalanceHandler) GetBalanceByAccount(w http.ResponseWriter, r *http.Request) {
	accountID := chi.URLParam(r, "account_id")
	balance, err := h.GetBalanceByAccountUseCase.Execute(usecase.GetBalanceByAccountInputDTO{AccountID: accountID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
