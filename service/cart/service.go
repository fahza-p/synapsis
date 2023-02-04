package cart

import (
	"github.com/fahza-p/synapsis/repository"
)

type Service struct {
	cart     repository.CartRepository
	product  repository.ProductRepository
	category repository.CategoryRepository
}

func NewService(cartRepo repository.CartRepository, categoryRepo repository.CategoryRepository, productRepo repository.ProductRepository) *Service {
	return &Service{
		cart:     cartRepo,
		product:  productRepo,
		category: categoryRepo,
	}
}
