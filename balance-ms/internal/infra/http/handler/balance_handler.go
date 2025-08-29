package http_handler

import (
	usecase "balance-ms/internal/app/usecase/get_balance_by_account"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type HttpBalanceHandler struct {
	GetBalanceByAccountUseCase usecase.GetBalanceByAccountUsecase
}

func NewHttpBalanceHandler(getBalanceByAccountUseCase usecase.GetBalanceByAccountUsecase) *HttpBalanceHandler {
	return &HttpBalanceHandler{GetBalanceByAccountUseCase: getBalanceByAccountUseCase}
}

func (h *HttpBalanceHandler) GetBalanceByAccount(w http.ResponseWriter, r *http.Request) {
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
