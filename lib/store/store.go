package store

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlStore struct {
	store *sql.DB
}

func NewStore() (*MysqlStore, error) {
	client, err := NewMysqlConnection()
	if err != nil {
		return nil, err
	}

	return NewMysqlStore(client)
}

func NewMysqlStore(db *sql.DB) (*MysqlStore, error) {
	return &MysqlStore{store: db}, nil
}

func (m *MysqlStore) Ping(ctx context.Context) error {
	return m.store.PingContext(ctx)
}
