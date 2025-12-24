package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mardini1414/simple-procurement-system-be/internal/config"
	"github.com/mardini1414/simple-procurement-system-be/internal/database"
	"github.com/mardini1414/simple-procurement-system-be/internal/handler"
	"github.com/mardini1414/simple-procurement-system-be/internal/middleware"
	"github.com/mardini1414/simple-procurement-system-be/internal/repository"
	"github.com/mardini1414/simple-procurement-system-be/internal/service"
)

func main() {
	cfg := config.LoadConfig()
	db := database.LoadDB(cfg)

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorMidleware,
	})

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))
	api := app.Group("/api")

	api.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Hello": "world",
		})
	})

	userRepo := repository.NewUserRepository(db)
	suplRepo := repository.NewSupplierRepository(db)
	itemRepo := repository.NewItemRepository(db)
	purcRepo := repository.NewPurchasingRepository(db)

	authService := service.NewAuthService(userRepo, cfg)
	suplService := service.NewSupplierService(suplRepo)
	itemService := service.NewItemService(itemRepo)
	purcService := service.NewPurchasingService(purcRepo, itemRepo, db, cfg)

	authHandler := handler.NewAuthHandler(authService)
	suplHandler := handler.NewSupplierHandler(suplService)
	itemHandler := handler.NewItemHandler(itemService)
	purcHandler := handler.NewPurchasingHandler(purcService)

	api.Post("/auth/login", authHandler.Login)
	api.Post("/auth/register", authHandler.Register)

	authenticated := api.Group("", middleware.AuthMiddleware(cfg))

	authenticated.Post("/suppliers", suplHandler.Create)
	authenticated.Put("/suppliers/:id", suplHandler.Update)
	authenticated.Get("/suppliers", suplHandler.GetAll)
	authenticated.Delete("/suppliers/:id", suplHandler.Delete)

	authenticated.Post("/items", itemHandler.Create)
	authenticated.Put("/items/:id", itemHandler.Update)
	authenticated.Get("/items", itemHandler.GetAll)
	authenticated.Delete("/items/:id", itemHandler.Delete)

	authenticated.Post("/purchasings", purcHandler.Create)
	authenticated.Get("/purchasings", purcHandler.FindAll)
	authenticated.Get("/purchasings/:id", purcHandler.FindById)

	port := fmt.Sprintf(":%s", cfg.Port)
	app.Listen(port)
}
