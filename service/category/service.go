package category

import "github.com/fahza-p/synapsis/repository"

type Service struct {
	category repository.CategoryRepository
	product  repository.ProductRepository
	user     repository.UserRepository
}

func NewService(categoryRepo repository.CategoryRepository, productRepo repository.ProductRepository, userRepo repository.UserRepository) *Service {
	return &Service{
		category: categoryRepo,
		product:  productRepo,
		user:     userRepo,
	}
}
