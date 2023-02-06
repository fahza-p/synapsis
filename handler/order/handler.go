package order

import (
	"github.com/fahza-p/synapsis/repository"
	"github.com/fahza-p/synapsis/service/order"
)

type Handler struct {
	service *order.Service
}

func NewHandler(
	orderRepo repository.OrderRepository,
	cartRepo repository.CartRepository,
	categoryRepo repository.CategoryRepository,
	productRepo repository.ProductRepository,
	userRepo repository.UserRepository,
) *Handler {
	handler := Handler{
		service: order.NewService(orderRepo, cartRepo, categoryRepo, productRepo, userRepo),
	}
	return &handler
}
