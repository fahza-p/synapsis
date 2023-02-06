package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/store"
	"github.com/fahza-p/synapsis/model"
)

type ProductRepository interface {
	IsExists(ctx context.Context, key string, val interface{}) (bool, error)
	Get(ctx context.Context, queryParams *store.QueryParams) ([]*model.ProductRes, int64, error)
	Create(ctx context.Context, req *model.Product) error
	Delete(ctx context.Context, key string, val interface{}) error
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

func (s *ProductStore) Create(ctx context.Context, req *model.Product) error {
	logger := log.GetLogger(ctx, "Product.Repository", "Create")
	logger.Info("Repository Create Product")

	id, err := s.db.Insert(ctx, s.table, req)
	if err != nil {
		return err
	}

	req.Id = id

	return nil
}

func (s *ProductStore) Delete(ctx context.Context, key string, val interface{}) error {
	logger := log.GetLogger(ctx, "Product.Repository", "Delete")
	logger.Info("Repository Delete Product")

	return s.db.ExecTx(ctx, func(tx store.Transaction) error {
		// Delet Product
		statment := fmt.Sprintf("DELETE FROM %s WHERE %s= ?", s.table, key)
		res, err := tx.ExecContext(ctx, statment, val)
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

		// Delete Cart Item
		statment = fmt.Sprintf("DELETE FROM %s WHERE %s= ?", "cart_item", "product_id")
		_, err = tx.ExecContext(ctx, statment, val)
		if err != nil {
			return err
		}

		return nil
	})
}

func (s *ProductStore) Get(ctx context.Context, queryParams *store.QueryParams) ([]*model.ProductRes, int64, error) {
	logger := log.GetLogger(ctx, "Product.Repository", "Get")
	logger.Info("Repository Get Product")

	var models []*model.ProductRes
	limit, offset, sort, filter, keywords := queryParams.BuildPagination(model.ProductFilter)

	statment := fmt.Sprintf(`
	SELECT 
		product.id,
		product.category_id,
		category.name AS category_name,
		product.sku,
		product.name,
		product.image,
		product.price,
		product.stock,
		product.created_at,
		product.updated_at,
		product.created_by,
		product.updated_by
	FROM product
	INNER JOIN category ON product.category_id = category.id
	WHERE %s AND (%s)
	%s
	%s %s
	`, filter, keywords, sort, limit, offset)

	if err := s.db.Query(ctx, &models, statment, true); err != nil {
		return nil, 0, err
	}

	countStetment := fmt.Sprintf(`
	SELECT 
		COUNT(product.id) AS total
	FROM product
	INNER JOIN category ON product.category_id = category.id
	WHERE %s AND (%s)
	`, filter, keywords)

	totalData, err := s.db.Count(ctx, countStetment)
	if err != nil {
		return nil, 0, err
	}

	return models, totalData, nil
}
