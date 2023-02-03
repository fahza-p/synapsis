package category

import (
	"github.com/fahza-p/synapsis/repository"
	"github.com/fahza-p/synapsis/service/category"
)

type Handler struct {
	service *category.Service
}

func NewHandler(categoryRepo repository.CategoryRepository, productRepo repository.ProductRepository, userRepo repository.UserRepository) *Handler {
	handler := Handler{
		service: category.NewService(categoryRepo, productRepo, userRepo),
	}
	return &handler
}
