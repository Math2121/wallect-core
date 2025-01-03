package web

import (
	"encoding/json"
	"net/http"

	createclient "github.com/Math2121/walletcore/usecase/client/create_client"
)

type WebClientHandler struct {
	CreateClientUseCase createclient.CreateClientUseCase
}

func NewWebClientHandler(createClientUseCase createclient.CreateClientUseCase) *WebClientHandler {
	return &WebClientHandler{createClientUseCase}
}

func (h *WebClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {

	var dto createclient.CreateClientInputDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, error := h.CreateClientUseCase.Execute(dto)
	if error!= nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err!= nil {
        w.WriteHeader(http.StatusInternalServerError)
		return
    }
	w.WriteHeader(http.StatusCreated)
}
