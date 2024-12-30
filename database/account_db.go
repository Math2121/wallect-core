package database

import (
	"database/sql"

	"github.com/Math2121/walletcore/entity"
)

type AccountDb struct {
	Db *sql.DB
}

func NewAccountDb(db *sql.DB) *AccountDb {
	return &AccountDb{Db: db}
}

func (a *AccountDb) FindById(id string) (*entity.Account, error) {
	var account entity.Account
	var client entity.Client
	account.Client = &client

	stmt, err := a.Db.Prepare("Select a.id, a.client_id, a.balance, a.created_at, c.id, c.name, c.email, c.created_at from account a inner join client c on a.id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&account.ID, &account.Client.ID, &account.Balance, &account.CreatedAt, &client.ID, &client.Name, &client.Email, &client.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &account, nil
}


func (a *AccountDb) Create(account *entity.Account) error {
		

	stmt, err := a.Db.Prepare("INSERT INTO account (id, client_id, balance, created_at) VALUES (?,?,?,?)")
    if err!= nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(account.ID, account.Client.ID, account.Balance, account.CreatedAt)
    if err!= nil {
        return err
    }

    return nil
}

func (a *AccountDb) UpdateBalance(account *entity.Account) error {
	stmt, err := a.Db.Prepare("UPDATE account SET balance =? WHERE id =?")
    if err!= nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(account.Balance, account.ID)
    if err!= nil {
        return err
    }

    return nil

}