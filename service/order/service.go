package order

import "github.com/fahza-p/synapsis/repository"

type Service struct {
	order    repository.OrderRepository
	cart     repository.CartRepository
	category repository.CategoryRepository
	product  repository.ProductRepository
	user     repository.UserRepository
}

func NewService(
	orderRepo repository.OrderRepository,
	cartRepo repository.CartRepository,
	categoryRepo repository.CategoryRepository,
	productRepo repository.ProductRepository,
	userRepo repository.UserRepository,
) *Service {
	return &Service{
		order:    orderRepo,
		cart:     cartRepo,
		category: categoryRepo,
		product:  productRepo,
		user:     userRepo,
	}
}
