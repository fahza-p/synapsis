package cmd

import (
	"context"

	"github.com/fahza-p/synapsis/handler/auth"
	"github.com/fahza-p/synapsis/lib/log"
	"github.com/fahza-p/synapsis/lib/store"
	"github.com/fahza-p/synapsis/repository"
	"github.com/fahza-p/synapsis/router"
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
	store, err := store.NewStore()
	if err != nil {
		logger.WithError(err).Fatal("Unable to connect database")
	}

	/* Initialize Repository */
	authRepo, err := repository.NewAuthRepository(store)
	if err != nil {
		logger.WithError(err).Fatal("Unable to initialize auth repository")
	}

	/* Initialize Handler */
	authHandler := auth.NewHandler(authRepo)

	app := fiber.New(fiber.Config{
		AppName:               "synapsis",
		CaseSensitive:         true,
		DisableStartupMessage: true,
		StrictRouting:         true,
	})

	api := app.Group("/api")
	router.NewAuthRouter(api, authHandler)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
