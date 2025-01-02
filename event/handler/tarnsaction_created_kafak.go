package handler

import (
	"sync"

	"github.com/Math2121/walletcore/pkg/eventos/pkg/events"
	"github.com/Math2121/walletcore/pkg/kafka"
)

type TransactionCreatedKafkaHandler struct {
	Kafka *kafka.Producer

}

func NewTransactionCreatedKafkaHandler(kafka *kafka.Producer) *TransactionCreatedKafkaHandler {
    return &TransactionCreatedKafkaHandler{Kafka: kafka}
}

func (h *TransactionCreatedKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup)  {
	defer wg.Done()

	h.Kafka.Publish(message, nil, "transactions")
	

}