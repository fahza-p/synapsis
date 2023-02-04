package router

import (
	"github.com/fahza-p/synapsis/handler/cart"
	"github.com/fahza-p/synapsis/middleware"
	fiber "github.com/gofiber/fiber/v2"
)

func NewCartRouter(app fiber.Router, handler *cart.Handler) {
	api := app.Group("/cart")

	// Cumtomer API
	api.Get("", middleware.Protected(), handler.GetCartMe)

	// Cart Items
	apiItems := api.Group("/item")

	apiItems.Get("", middleware.Protected(), handler.GetCartItemMe)
	apiItems.Post("/add", middleware.Protected(), handler.AddProduct)
	apiItems.Delete("/remove/:id", middleware.Protected(), handler.RemoveProduct)
	// Admin API
	// api.Post("/admin", middleware.Protected(), handler.Create)
	// api.Delete("/admin/:id", middleware.Protected(), handler.Remove)
}
