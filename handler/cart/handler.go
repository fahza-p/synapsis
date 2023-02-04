package cart

import (
	"github.com/fahza-p/synapsis/repository"
	"github.com/fahza-p/synapsis/service/cart"
)

type Handler struct {
	service *cart.Service
}

func NewHandler(cartRepo repository.CartRepository, categoryRepo repository.CategoryRepository, productRepo repository.ProductRepository) *Handler {
	handler := Handler{
		service: cart.NewService(cartRepo, categoryRepo, productRepo),
	}
	return &handler
}
