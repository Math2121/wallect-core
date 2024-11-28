package database

import (
	"database/sql"

	"github.com/Math2121/walletcore/entity"
)

type TransactionDb struct {
	db *sql.DB
}

func NewTransactionDb(db *sql.DB) *TransactionDb {
	return &TransactionDb{db: db}
}

func (t *TransactionDb) Create(transaction *entity.Transaction) error {
	smt, err := t.db.Prepare(" INSERT INTO transaction  (id, account_id_from, account_id_to, amount, created_at) VALUES (?,?,?,?,?)")
	if err != nil {
		return err
	}

	defer smt.Close()

	_, err = smt.Exec(transaction.ID, transaction.AccountFrom, transaction.AccountTo, transaction.Amount, transaction.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

