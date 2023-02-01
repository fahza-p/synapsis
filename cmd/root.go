package cmd

import (
	"context"

	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/store"
	fiber "github.com/gofiber/fiber/v2"
)

func Run() {
	/* Initialize Context */
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	/* Initialize Logger */
	log.Configure("json", "debug")
	logger := log.GetLogger(ctx, "cmd", "Run")

	/* Initialize Database */
	_, err := store.NewStore()
	if err != nil {
		logger.WithError(err).Fatal("Unable to connect database")
	}

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
