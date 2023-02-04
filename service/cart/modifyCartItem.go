package cart

import (
	"context"
	"errors"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/model"
)

func (s *Service) AddProduct(ctx context.Context, model *model.CartItemData, authData map[string]interface{}) error {
	logger := log.GetLogger(ctx, "Cart.Service", "AddProduct")
	logger.Info("AddProduct")

	// Check Product IsExisted Data
	isProductExists, err := s.product.IsExists(ctx, "id", model.ProductId)
	if err != nil {
		logger.Errorf("can't get product data with slug: '%s'", model.ProductId)
		return err
	}

	if !isProductExists {
		return errors.New("product not found")
	}

	// Get Cart
	cartData, err := s.cart.FindOne(ctx, "user_id", authData["id"])
	if err != nil {
		return err
	}

	// Check Product Is Exist On Cart
	query := map[string]interface{}{
		"cart_id":    cartData.Id,
		"product_id": model.ProductId,
	}
	isCartHaveProduct, err := s.cart.IsExistsProduct(ctx, query)
	if err != nil {
		logger.Errorf("can't get product data with id: '%s'", model.ProductId)
		return err
	}

	if !isCartHaveProduct {
		model.SetCartItemAddProductData(cartData.Id, authData["email"].(string))

		// Add Item
		return s.cart.AddItem(ctx, model)
	}

	// Update Item Qty
	return s.cart.UpdateItemQty(ctx, cartData.Id, model.ProductId, model.Qty)
}

func (s *Service) RemoveProduct(ctx context.Context, productId int64, authData map[string]interface{}) error {
	logger := log.GetLogger(ctx, "Cart.Service", "RemoveProduct")
	logger.Info("RemoveProduct")

	// Get Cart
	cartData, err := s.cart.FindOne(ctx, "user_id", authData["id"])
	if err != nil {
		return err
	}

	// Delete Category
	return s.cart.RemoveItem(ctx, cartData.Id, productId)
}
