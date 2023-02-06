package order

import (
	"context"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/store"
	"github.com/fahza-p/synapsis/model"
)

func (s *Service) GetList(ctx context.Context, queryParams *store.QueryParams, userId int64) ([]*model.OrderData, int64, error) {
	logger := log.GetLogger(ctx, "Order.Service", "GetList")
	logger.Info("GetList")

	// Find Order
	return s.order.Get(ctx, queryParams, userId)
}
