package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/fahza-p/synapsis/lib/utils"
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

func (m *MysqlStore) Exect(ctx context.Context, statment string, args ...interface{}) error {
	res, err := m.store.ExecContext(ctx, statment, args...)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errors.New("document not found")
	}

	return nil
}

func (m *MysqlStore) Query(ctx context.Context, docs interface{}, statment string, args ...interface{}) error {
	rows, err := m.store.QueryContext(ctx, statment, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	out := scan(rows)

	if err := rows.Close(); err != nil {
		return err
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return utils.DecodeJSON(out, docs)
}

func (m *MysqlStore) QueryRow(ctx context.Context, docs interface{}, statment string, args ...interface{}) error {
	cols := utils.GetAllJsonTagName(docs)
	scans := make([]interface{}, len(cols))
	out := make(map[string]interface{})

	for i := range cols {
		scans[i] = &scans[i]
	}

	if err := m.store.QueryRowContext(ctx, statment, args...).Scan(scans...); err != nil {
		return err
	}

	for i, v := range scans {
		out[cols[i]] = v
	}

	return utils.DecodeJSON(out, docs)
}

func (m *MysqlStore) Count(ctx context.Context, statment string, args ...interface{}) (int64, error) {
	var total int64

	if err := m.store.QueryRowContext(ctx, statment, args...).Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}

func (m *MysqlStore) Insert(ctx context.Context, table string, val interface{}, args ...interface{}) (int64, error) {
	var (
		fields, values []string
		docs           map[string]interface{}
	)

	data, err := json.Marshal(val)
	if err != nil {
		return 0, err
	}

	json.Unmarshal(data, &docs)

	for k, v := range docs {
		fields = append(fields, k)
		values = append(values, "?")
		args = append(args, v)
	}

	statment := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%v)", table, strings.Join(fields, ","), strings.Join(values, ","))

	res, err := m.store.ExecContext(ctx, statment, args...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func scan(list *sql.Rows) (rows []map[string]interface{}) {
	cols, _ := list.Columns()
	for list.Next() {
		scans := make([]interface{}, len(cols))
		row := make(map[string]interface{})

		for i := range scans {
			scans[i] = &scans[i]
		}

		list.Scan(scans...)

		for i, v := range scans {
			var value = ""
			if v != nil {
				value = fmt.Sprintf("%s", v)
			}
			row[cols[i]] = value
		}

		rows = append(rows, row)
	}

	return
}
