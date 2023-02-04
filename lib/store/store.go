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

type Transaction interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type TxFn func(Transaction) error

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

func (m *MysqlStore) exec(ctx context.Context, statment string, args ...interface{}) error {
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

func (m *MysqlStore) Insert(ctx context.Context, table string, values interface{}, args ...interface{}) (int64, error) {
	var (
		cols, sep []string
		val       []interface{}
		docs      map[string]interface{}
	)

	data, err := json.Marshal(values)
	if err != nil {
		return 0, err
	}

	if err := json.Unmarshal(data, &docs); err != nil {
		return 0, err
	}

	for k, v := range docs {
		cols = append(cols, k)
		sep = append(sep, "?")
		val = append(val, v)
	}

	args = append(val, args...)

	statment := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%v)", table, strings.Join(cols, ","), strings.Join(sep, ","))

	res, err := m.store.ExecContext(ctx, statment, args...)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (m *MysqlStore) Delete(ctx context.Context, statment string, args ...interface{}) error {
	return m.exec(ctx, statment, args...)
}

func (m *MysqlStore) Update(ctx context.Context, table string, query string, values interface{}, args ...interface{}) error {
	var (
		cols []string
		val  []interface{}
		docs map[string]interface{}
	)

	data, err := json.Marshal(values)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &docs); err != nil {
		return err
	}

	for k, v := range docs {
		cols = append(cols, k+"=?")
		val = append(val, v)
	}

	args = append(val, args...)

	statment := fmt.Sprintf("UPDATE %s SET %s WHERE %s", table, strings.Join(cols, ","), query)

	return m.exec(ctx, statment, args...)
}

func (m *MysqlStore) Increment(ctx context.Context, table string, field string, query string, args ...interface{}) error {
	statment := fmt.Sprintf("UPDATE %s SET %s = %s + 1 WHERE %s", table, field, field, query)
	return m.exec(ctx, statment, args...)
}

func (m *MysqlStore) Decrement(ctx context.Context, table string, field string, query string, args ...interface{}) error {
	statment := fmt.Sprintf("UPDATE %s SET %s = %s - 1 WHERE %s", table, field, field, query)
	return m.exec(ctx, statment, args...)
}

func (m *MysqlStore) ExecTx(ctx context.Context, fn TxFn) error {
	tx, err := m.store.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = fn(tx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
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
			row[cols[i]] = v
		}

		rows = append(rows, row)
	}

	return
}
