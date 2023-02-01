package cmd

import (
	fiber "github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New(fiber.Config{
		AppName:               "synapsis",
		CaseSensitive:         true,
		DisableStartupMessage: true,
		StrictRouting:         true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
