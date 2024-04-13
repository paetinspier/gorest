package router

import (
	"psn/gorest/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", services.HandleHealthCheck)

	// Users
	users := app.Group("/users")
	users.Get("/", services.GetAllUsers)
	users.Post("/create", services.CreateUser)
}
