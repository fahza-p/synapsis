package router

import (
	"github.com/fahza-p/synapsis/handler/category"
	"github.com/fahza-p/synapsis/middleware"
	fiber "github.com/gofiber/fiber/v2"
)

func NewCategoryRouter(app fiber.Router, handler *category.Handler) {
	api := app.Group("/category")

	// Cumtomer API
	api.Get("", middleware.Protected(), handler.GetList)
	api.Get("/:id", middleware.Protected(), handler.FindById)

	// Admin API
	api.Post("/admin", middleware.Protected(), handler.Create)
	api.Patch("/admin/:id", middleware.Protected(), handler.UpdatePatch)
	api.Delete("/admin/:id", middleware.Protected(), handler.Remove)
}
