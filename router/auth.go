package router

import (
	"github.com/fahza-p/synapsis/handler/auth"
	fiber "github.com/gofiber/fiber/v2"
)

func NewAuthRouter(app fiber.Router, handler *auth.Handler) {
	api := app.Group("/auth")

	api.Post("/signup", handler.Signup)
}
