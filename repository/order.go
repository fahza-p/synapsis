package repository

import (
	"context"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/store"
)

type OrderRepository interface {
}

type OrderStore struct {
	db    *store.MysqlStore
	table string
}

func NewOrderRepository(store *store.MysqlStore) (OrderRepository, error) {
	logger := log.GetLogger(context.Background(), "Order.Repository", "NewOrderRepository")
	logger.Info("Add New Order Repository")

	return &OrderStore{db: store, table: "order"}, nil
}
