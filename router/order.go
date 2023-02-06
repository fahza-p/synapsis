package router

import (
	"github.com/fahza-p/synapsis/handler/order"
	fiber "github.com/gofiber/fiber/v2"
)

func NewOrderRouter(app fiber.Router, handler *order.Handler) {
	// api := app.Group("/category")

	// // Cumtomer API
	// api.Get("/:id", middleware.Protected(), handler.FindById)

	// // Admin API
	// api.Post("/admin", middleware.Protected(), handler.Create)
	// api.Patch("/admin/:id", middleware.Protected(), handler.UpdatePatch)
	// api.Delete("/admin/:id", middleware.Protected(), handler.Remove)
}
