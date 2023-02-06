package cmd

import (
	"context"
	"os"

	"github.com/fahza-p/synapsis/handler/auth"
	"github.com/fahza-p/synapsis/handler/cart"
	"github.com/fahza-p/synapsis/handler/category"
	"github.com/fahza-p/synapsis/handler/order"
	"github.com/fahza-p/synapsis/handler/product"
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

	userRepo, err := repository.NewUserRepository(store)
	if err != nil {
		logger.WithError(err).Fatal("Unable to initialize user repository")
	}

	categoryRepo, err := repository.NewCategoryRepository(store)
	if err != nil {
		logger.WithError(err).Fatal("Unable to initialize category repository")
	}

	productRepo, err := repository.NewProductRepository(store)
	if err != nil {
		logger.WithError(err).Fatal("Unable to initialize product repository")
	}

	cartRepo, err := repository.NewCartRepository(store)
	if err != nil {
		logger.WithError(err).Fatal("Unable to initialize cart repository")
	}

	orderRepo, err := repository.NewOrderRepository(store)
	if err != nil {
		logger.WithError(err).Fatal("Unable to initialize order repository")
	}

	/* Initialize Handler */
	authHandler := auth.NewHandler(authRepo)
	categoryHandler := category.NewHandler(categoryRepo, productRepo, userRepo)
	productHandler := product.NewHandler(categoryRepo, productRepo)
	cartHandler := cart.NewHandler(cartRepo, categoryRepo, productRepo)
	orderHandler := order.NewHandler(orderRepo, cartRepo, categoryRepo, productRepo, userRepo)

	app := fiber.New(fiber.Config{
		AppName:               "synapsis",
		CaseSensitive:         true,
		DisableStartupMessage: true,
		StrictRouting:         true,
	})

	api := app.Group("/api")
	router.NewAuthRouter(api, authHandler)
	router.NewCategoryRouter(api, categoryHandler)
	router.NewProductRouter(api, productHandler)
	router.NewCartRouter(api, cartHandler)
	router.NewOrderRouter(api, orderHandler)

	app.Listen(":" + os.Getenv("PORT"))
}
