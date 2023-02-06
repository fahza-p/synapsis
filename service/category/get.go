package category

import (
	"context"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/store"
	"github.com/fahza-p/synapsis/model"
)

func (s *Service) FindById(ctx context.Context, id string) (*model.Category, error) {
	logger := log.GetLogger(ctx, "Category.Service", "FindById")
	logger.Info("FindById")

	// Find Category
	categoryData, err := s.category.FindOne(ctx, "id", id)
	if err != nil {
		return nil, err
	}

	return categoryData, nil
}

func (s *Service) GetList(ctx context.Context, queryParams *store.QueryParams) ([]*model.Category, int64, error) {
	logger := log.GetLogger(ctx, "Category.Service", "GetList")
	logger.Info("GetList")

	// Find Category
	return s.category.Get(ctx, queryParams)
}
