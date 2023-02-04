package product

import "github.com/fahza-p/synapsis/repository"

type Service struct {
	product  repository.ProductRepository
	category repository.CategoryRepository
}

func NewService(categoryRepo repository.CategoryRepository, productRepo repository.ProductRepository) *Service {
	return &Service{
		product:  productRepo,
		category: categoryRepo,
	}
}
