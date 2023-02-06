package router

import (
	"github.com/fahza-p/synapsis/handler/order"
	"github.com/fahza-p/synapsis/middleware"
	fiber "github.com/gofiber/fiber/v2"
)

func NewOrderRouter(app fiber.Router, handler *order.Handler) {
	api := app.Group("/order")

	// Cumtomer API
	api.Post("", middleware.Protected(), handler.Create)
	api.Get("", middleware.Protected(), handler.GetList)
	// api.Get("/detail/:id", middleware.Protected(), handler.GetList)
}
