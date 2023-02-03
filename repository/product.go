package repository

import (
	"context"
	"fmt"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/store"
)

type ProductRepository interface {
	IsExists(ctx context.Context, key string, val interface{}) (bool, error)
}

type ProductStore struct {
	db    *store.MysqlStore
	table string
}

func NewProductRepository(store *store.MysqlStore) (ProductRepository, error) {
	logger := log.GetLogger(context.Background(), "Product.Repository", "NewProductRepository")
	logger.Info("Add New Product Repository")

	return &ProductStore{db: store, table: "product"}, nil
}

func (s *ProductStore) IsExists(ctx context.Context, key string, val interface{}) (bool, error) {
	logger := log.GetLogger(ctx, "Product.Repository", "IsExists")
	logger.Info("Repository IsExists Product")

	statment := fmt.Sprintf("SELECT COUNT(id) AS total FROM %s WHERE %s=?", s.table, key)

	total, err := s.db.Count(ctx, statment, val)
	if err != nil {
		return false, err
	}

	if total == 0 {
		return false, nil
	}

	return true, nil
}
