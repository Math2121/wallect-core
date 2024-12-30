package web

import (
	"encoding/json"
	"net/http"

	createtransaction "github.com/Math2121/walletcore/usecase/transaction/create_transaction"
)

type WebTransactionHandler struct {
	CreateTransactionUseCase createtransaction.CreateTransactionUseCase
}

func NewWebTransactionHandler(createtransaction createtransaction.CreateTransactionUseCase) *WebTransactionHandler {
	return &WebTransactionHandler{CreateTransactionUseCase: createtransaction}
}

func (h *WebTransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {

	var dto createtransaction.CreateTransactionInputDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	output, error := h.CreateTransactionUseCase.Execute(ctx, dto)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(error.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
