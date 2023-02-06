package router

import (
	"github.com/fahza-p/synapsis/handler/product"
	"github.com/fahza-p/synapsis/middleware"
	fiber "github.com/gofiber/fiber/v2"
)

func NewProductRouter(app fiber.Router, handler *product.Handler) {
	api := app.Group("/product")

	// Cumtomer API
	api.Get("", middleware.Protected(), handler.GetList)

	// Admin API
	api.Post("/admin", middleware.Protected(), handler.Create)
	api.Delete("/admin/:id", middleware.Protected(), handler.Remove)
}
