package product

import (
	"github.com/fahza-p/synapsis/repository"
	"github.com/fahza-p/synapsis/service/product"
)

type Handler struct {
	service *product.Service
}

func NewHandler(categoryRepo repository.CategoryRepository, productRepo repository.ProductRepository) *Handler {
	handler := Handler{
		service: product.NewService(categoryRepo, productRepo),
	}
	return &handler
}
