package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/store"
	"github.com/fahza-p/synapsis/model"
)

type OrderRepository interface {
	Create(ctx context.Context, cartId int64, orderModel *model.OrderData, itemsModel []*model.OrderItems, statusModel *model.StatusChangelog) error
	Get(ctx context.Context, queryParams *store.QueryParams, userId int64) ([]*model.OrderData, int64, error)
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

func (s *OrderStore) Create(ctx context.Context, cartId int64, orderModel *model.OrderData, itemsModel []*model.OrderItems, statusModel *model.StatusChangelog) error {
	logger := log.GetLogger(ctx, "Order.Repository", "Create")
	logger.Info("Repository Create Order")

	return s.db.ExecTx(ctx, func(tx store.Transaction) error {
		// Order
		var (
			orderInput map[string]interface{}
			orderCols  []string
			orderVals  []interface{}
			orderSep   []string
		)

		orderData, err := json.Marshal(orderModel)
		if err != nil {
			return err
		}
		json.Unmarshal(orderData, &orderInput)

		for k, v := range orderInput {
			orderCols = append(orderCols, k)
			orderSep = append(orderSep, "?")
			orderVals = append(orderVals, v)
		}

		statment := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%v)", "order", strings.Join(orderCols, ","), strings.Join(orderSep, ","))

		res, err := tx.ExecContext(ctx, statment, orderVals...)
		if err != nil {
			return err
		}
		orderId, err := res.LastInsertId()
		if err != nil {
			return err
		}
		orderModel.Id = orderId

		// Order Item
		for _, v := range itemsModel {
			var (
				orderItemInput map[string]interface{}
				orderItemCols  []string
				orderItemSep   []string
				orderItemVals  []interface{}
			)

			v.OrderId = orderId

			orderItemData, err := json.Marshal(v)
			if err != nil {
				return err
			}
			json.Unmarshal(orderItemData, &orderItemInput)

			for k, v := range orderItemInput {
				orderItemCols = append(orderItemCols, k)
				orderItemSep = append(orderItemSep, "?")
				orderItemVals = append(orderItemVals, v)
			}

			itemStatment := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%v)", "order_item", strings.Join(orderItemCols, ","), strings.Join(orderItemSep, ","))

			_, err = tx.ExecContext(ctx, itemStatment, orderItemVals...)
			if err != nil {
				return err
			}

			// Decrement Stock
			stockStetment := fmt.Sprintf("UPDATE %s SET %s = %s - %d WHERE id = %d", "product", "stock", "stock", v.Qty, v.ProductId)
			_, err = tx.ExecContext(ctx, stockStetment)
			if err != nil {
				return err
			}

			// Remove From Cart
			removeCartItemStatment := "DELETE FROM cart_item WHERE product_id= ? AND cart_id= ?"
			_, err = tx.ExecContext(ctx, removeCartItemStatment, v.ProductId, cartId)
			if err != nil {
				return err
			}
		}

		// Status
		var (
			orderStatusInput map[string]interface{}
			orderStatusCols  []string
			orderStatusSep   []string
			orderStatusVals  []interface{}
		)

		statusModel.OrderId = orderId
		orderStatusData, err := json.Marshal(statusModel)
		if err != nil {
			return err
		}
		json.Unmarshal(orderStatusData, &orderStatusInput)

		for k, v := range orderStatusInput {
			orderStatusCols = append(orderStatusCols, "`"+k+"`")
			orderStatusSep = append(orderStatusSep, "?")
			orderStatusVals = append(orderStatusVals, v)
		}

		statusStatment := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (%v)", "status_changelog", strings.Join(orderStatusCols, ","), strings.Join(orderStatusSep, ","))
		_, err = tx.ExecContext(ctx, statusStatment, orderStatusVals...)
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *OrderStore) Get(ctx context.Context, queryParams *store.QueryParams, userId int64) ([]*model.OrderData, int64, error) {
	logger := log.GetLogger(ctx, "Order.Repository", "Get")
	logger.Info("Repository Get Order")

	var models []*model.OrderData
	limit, offset, sort, filter, keywords := queryParams.BuildPagination(model.OrderFilter)

	statment := fmt.Sprintf(`
	SELECT *
	FROM %s
	WHERE user_id = %d AND %s AND (%s)
	%s
	%s %s
	`, "`order`", userId, filter, keywords, sort, limit, offset)

	if err := s.db.Query(ctx, &models, statment, true); err != nil {
		return nil, 0, err
	}

	countStetment := fmt.Sprintf(`
	SELECT 
		COUNT(order.id) AS total
	FROM %s
	WHERE user_id = %d AND %s AND (%s)
	`, "`order`", userId, filter, keywords)

	totalData, err := s.db.Count(ctx, countStetment)
	if err != nil {
		return nil, 0, err
	}

	return models, totalData, nil
}
