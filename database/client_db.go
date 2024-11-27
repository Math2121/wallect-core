package database

import (
	"database/sql"

	"github.com/Math2121/walletcore/entity"
)

type ClientDB struct {
	Db *sql.DB
}

func NewClientDB(db *sql.DB) *ClientDB {
    return &ClientDB{Db: db}
}
func (c *ClientDB) Get(id string) (*entity.Client, error) {
	client := &entity.Client{}
    query := "SELECT id, name, email, created_at FROM clients WHERE id =?"
    row := c.Db.QueryRow(query, id)

    err := row.Scan(&client.ID, &client.Name, &client.Email, &client.CreatedAt)
    if err == sql.ErrNoRows {
        return nil, nil
    } else if err!= nil {
        return nil, err
    }
	


    return client, nil
}

func (c *ClientDB) Create(client *entity.Client) error {
	query := "INSERT INTO clients (id, name, email, created_at) VALUES (?,?,?,?)"
    _, err := c.Db.Exec(query, client.ID, client.Name, client.Email, client.CreatedAt)
    if err!= nil {
        return err
    }
    return nil
}