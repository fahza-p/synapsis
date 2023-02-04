package cart

import (
	"context"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/model"
)

func (s *Service) GetCartMe(ctx context.Context, authData map[string]interface{}) (*model.Cart, error) {
	logger := log.GetLogger(ctx, "Cart.Service", "GetCartMe")
	logger.Info("GetCartMe")

	// Find Cart
	cartData, err := s.cart.FindCart(ctx, authData["id"])
	if err != nil {
		return nil, err
	}

	return cartData, nil
}

func (s *Service) GetCartItemMe(ctx context.Context, authData map[string]interface{}) ([]*model.CartItems, error) {
	logger := log.GetLogger(ctx, "Cart.Service", "GetCartMe")
	logger.Info("GetCartMe")

	// Get Cart
	cartData, err := s.cart.FindOne(ctx, "user_id", authData["id"])
	if err != nil {
		return nil, err
	}

	// Find Cart
	cartItemData, err := s.cart.GetCartItemByCartId(ctx, cartData.Id)
	if err != nil {
		return nil, err
	}

	return cartItemData, nil
}
