package product

import (
	"context"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/store"
	"github.com/fahza-p/synapsis/model"
)

func (s *Service) GetList(ctx context.Context, queryParams *store.QueryParams) ([]*model.ProductRes, int64, error) {
	logger := log.GetLogger(ctx, "Product.Service", "GetList")
	logger.Info("GetList")

	if len(queryParams.Filter) > 0 {
		if queryParams.Filter[0] == "category_name" {
			queryParams.Filter[0] = "category.name"
		}

		if queryParams.Filter[0] == "name" {
			queryParams.Filter[0] = "product.name"
		}
	}
	// Find Product
	return s.product.Get(ctx, queryParams)
}
