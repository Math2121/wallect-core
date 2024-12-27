package web

import (
	"encoding/json"
	"net/http"

	createaccount "github.com/Math2121/walletcore/usecase/account/create_account"
)

type WebAccountHandler struct {
	CreateAccountUseCase createaccount.CreateAccountUseCase
}

func NewWebAccountHandler(createAccount createaccount.CreateAccountUseCase) *WebAccountHandler {
	return &WebAccountHandler{CreateAccountUseCase: createAccount}
}

func (h *WebAccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {

	var dto createaccount.CreateAccountInputDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, error := h.CreateAccountUseCase.Execute(dto)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
