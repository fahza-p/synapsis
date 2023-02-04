package product

import (
	"context"
	"errors"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/model"
)

func (s *Service) Create(ctx context.Context, reqModel *model.ProductCreateReq, model *model.Product, userEmail string) error {
	logger := log.GetLogger(ctx, "Product.Service", "Create")
	logger.Info("Create")

	// Check Category
	isCategoryExists, err := s.category.IsExists(ctx, "id", reqModel.CategoryId)
	if err != nil {
		logger.Errorf("can't get category data with id: '%s'", reqModel.CategoryId)
		return err
	}

	if !isCategoryExists {
		return errors.New("category not found")
	}

	// Check Product Is Exists
	isProductExists, err := s.product.IsExists(ctx, "sku", reqModel.Sku)
	if err != nil {
		logger.Errorf("can't get product data with sku: '%s'", reqModel.Sku)
		return err
	}

	if isProductExists {
		return errors.New("data is already exist")
	}

	// Build Create Data
	model.SetProductCreateData(reqModel, userEmail)

	// Create
	return s.product.Create(ctx, model)
}
