package order

import (
	"context"
	"errors"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/model"
)

func (s *Service) Create(ctx context.Context, reqData *model.OrderReq, orderModel *model.OrderData, authData map[string]interface{}) error {
	logger := log.GetLogger(ctx, "Order.Service", "Create")
	logger.Info("Create")

	var (
		totalItem, totalProduct int32
		totalPrice, totalPaid   float64
	)

	// Find Cart
	cartData, err := s.cart.FindCart(ctx, authData["id"])
	if err != nil {
		return err
	}

	// Get Cart Item With Cart ID & Product ID
	cartItem, err := s.cart.GetCartItem(ctx, cartData.Id, reqData.ProductId)
	if err != nil {
		return err
	}

	if len(cartItem) == 0 {
		return errors.New("cart is empty")
	}

	for _, v := range cartItem {
		totalProduct += 1
		totalItem += v.Qty
		totalPrice += v.TotalPrice
		totalPaid += v.TotalPrice
	}

	// Build Order Create Data
	itemsModel, statusModel := orderModel.SetOrderCreateData(totalProduct, totalItem, totalPrice, totalPaid, int64(authData["id"].(float64)), authData["email"].(string), cartItem)

	// Create Order
	return s.order.Create(ctx, cartData.Id, orderModel, itemsModel, statusModel)
}
