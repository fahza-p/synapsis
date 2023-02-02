package auth

import (
	"github.com/fahza-p/synapsis/repository"
)

type Service struct {
	auth repository.AuthRepository
}

func NewService(authRepo repository.AuthRepository) *Service {
	return &Service{
		auth: authRepo,
	}
}
