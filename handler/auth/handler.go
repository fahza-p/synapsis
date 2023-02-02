package auth

import (
	"github.com/fahza-p/synapsis/repository"
	"github.com/fahza-p/synapsis/service/auth"
)

type Handler struct {
	service *auth.Service
}

func NewHandler(authRepo repository.AuthRepository) *Handler {
	handler := Handler{
		service: auth.NewService(authRepo),
	}
	return &handler
}
