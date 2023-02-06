package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/store"
	"github.com/fahza-p/synapsis/model"
)

type CartRepository interface {
	FindOne(ctx context.Context, key string, val interface{}) (*model.CartData, error)
	FindCart(ctx context.Context, val interface{}) (*model.Cart, error)
	GetCartItemByCartId(ctx context.Context, val interface{}, queryParams *store.QueryParams) ([]*model.CartItems, int64, error)
	IsExistsProduct(ctx context.Context, query map[string]interface{}) (bool, error)
	AddItem(ctx context.Context, req *model.CartItemData) error
	UpdateItemQty(ctx context.Context, cartId int64, productId int64, qty int32) error
	RemoveItem(ctx context.Context, cartId int64, productId int64) error
}

type CartStore struct {
	db    *store.MysqlStore
	table string
}

func NewCartRepository(store *store.MysqlStore) (CartRepository, error) {
	logger := log.GetLogger(context.Background(), "Cart.Repository", "NewCartRepository")
	logger.Info("Add New Cart Repository")

	return &CartStore{db: store, table: "cart"}, nil
}

func (s *CartStore) FindCart(ctx context.Context, val interface{}) (*model.Cart, error) {
	logger := log.GetLogger(ctx, "Cart.Repository", "FindCart")
	logger.Info("Repository FindCart Cart")

	var model model.Cart

	statment := `
	SELECT 
		cart.id,
		cart.user_id,
		CAST(COUNT(cart_item.id) AS int) AS total_product,
		CAST(SUM(cart_item.qty) AS int)  AS total_items,
		CAST(SUM((cart_item.qty * product.price)) AS float) AS total_price,
		cart.created_at,
		cart.updated_at,
		cart.created_by,
		cart.updated_by
	FROM cart
	INNER JOIN cart_item ON cart_item.cart_id = cart.id 
	INNER JOIN product ON product.id = cart_item.product_id 
	WHERE cart.user_id = ?
	GROUP BY cart.id  
	`

	if err := s.db.QueryRow(ctx, &model, statment, val); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return &model, errors.New("document not found")
		}

		return &model, err
	}

	return &model, nil
}

func (s *CartStore) GetCartItemByCartId(ctx context.Context, val interface{}, queryParams *store.QueryParams) ([]*model.CartItems, int64, error) {
	logger := log.GetLogger(ctx, "Cart.Repository", "GetCartItemByCartId")
	logger.Info("Repository GetCartItemByCartId Cart")

	var models []*model.CartItems
	limit, offset, sort, filter, keywords := queryParams.BuildPagination(model.CartItemFilter)

	statment := fmt.Sprintf(`
	SELECT 
		cart_item.id,
		cart_item.cart_id,
		cart_item.product_id,
		product.sku,
		product.name,
		product.price,
		product.stock,
		cart_item.qty,
		CAST((cart_item.qty * product.price) AS float) AS total_price,
		cart_item.created_at,
		cart_item.updated_at,
		cart_item.created_by,
		cart_item.updated_by
	FROM cart_item
	INNER JOIN product ON product.id = cart_item.product_id 
	WHERE 
		cart_item.cart_id = ? AND %s AND (%s)
	%s
	%s %s
	`, filter, keywords, sort, limit, offset)
	if err := s.db.Query(ctx, &models, statment, false, val); err != nil {
		return nil, 0, err
	}

	countStetment := fmt.Sprintf(`
	SELECT 
		COUNT(cart_item.id) AS total
	FROM cart_item
	INNER JOIN product ON product.id = cart_item.product_id 
	WHERE 
		cart_item.cart_id = ? AND %s AND (%s)
	`, filter, keywords)

	totalData, err := s.db.Count(ctx, countStetment, val)
	if err != nil {
		return nil, 0, err
	}

	return models, totalData, nil
}

func (s *CartStore) FindOne(ctx context.Context, key string, val interface{}) (*model.CartData, error) {
	logger := log.GetLogger(ctx, "Cart.Repository", "FindOne")
	logger.Info("Repository FindOne Cart")

	var model model.CartData

	statment := fmt.Sprintf("SELECT * FROM %s WHERE %s= ?", s.table, key)
	if err := s.db.QueryRow(ctx, &model, statment, val); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return &model, errors.New("document not found")
		}

		return &model, err
	}

	return &model, nil
}

func (s *CartStore) IsExistsProduct(ctx context.Context, query map[string]interface{}) (bool, error) {
	logger := log.GetLogger(ctx, "Cart.Repository", "IsExistsProduct")
	logger.Info("Repository IsExistsProduct Cart")

	var (
		key    []string
		params []interface{}
	)

	for k, v := range query {
		key = append(key, k+"=?")
		params = append(params, v)
	}

	statment := fmt.Sprintf("SELECT COUNT(id) AS total FROM %s WHERE %s", "cart_item", strings.Join(key, " AND "))

	total, err := s.db.Count(ctx, statment, params...)
	if err != nil {
		return false, err
	}

	if total == 0 {
		return false, nil
	}

	return true, nil
}

func (s *CartStore) AddItem(ctx context.Context, req *model.CartItemData) error {
	logger := log.GetLogger(ctx, "Cart.Repository", "AddItem")
	logger.Info("Repository AddItem Cart")

	id, err := s.db.Insert(ctx, "cart_item", req)
	if err != nil {
		return err
	}

	req.Id = id

	return nil
}

func (s *CartStore) UpdateItemQty(ctx context.Context, cartId int64, productId int64, qty int32) error {
	logger := log.GetLogger(ctx, "Cart.Repository", "UpdateItemQty")
	logger.Info("Repository UpdateItemQty Cart")

	query := "cart_id=? AND product_id=?"
	updateField := map[string]interface{}{"qty": qty}
	return s.db.Update(ctx, "cart_item", query, updateField, cartId, productId)
}

func (s *CartStore) IncrementItemQty(ctx context.Context, cartId int64, productId int64) error {
	logger := log.GetLogger(ctx, "Cart.Repository", "IncrementItemQty")
	logger.Info("Repository IncrementItemQty Cart")

	query := "cart_id=? AND product_id=?"
	return s.db.Increment(ctx, "cart_item", "qty", query, cartId, productId)
}

func (s *CartStore) RemoveItem(ctx context.Context, cartId int64, productId int64) error {
	logger := log.GetLogger(ctx, "Cart.Repository", "RemoveItem")
	logger.Info("Repository RemoveItem Cart")

	statment := fmt.Sprintf("DELETE FROM %s WHERE cart_id=? AND product_id=?", "cart_item")

	return s.db.Delete(ctx, statment, cartId, productId)
}
