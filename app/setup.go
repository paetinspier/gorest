package app

import (
	"errors"
	"log"
	"os"
	"psn/gorest/config"
	"psn/gorest/database"
	"psn/gorest/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupAndRunApp() error {
	// load env
	err := config.LoadENV()
	if err != nil {
		return err
	}

	// start database
	err = database.StartPostgresDB()
	if err != nil {
		return err
	}

	// defer closing database
	defer database.ClosePostgresDB()

	// create new app
	app := fiber.New()

	// attach middleware
	// app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	// setup routes
	router.SetupRoutes(app)

	// attach swagger
	config.AddSwaggerRoutes(app)

	// get the port and start
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		return errors.New("missing DB_DATABASE environment variable")
	}

	log.Fatal(app.Listen(":" + port))

	return nil
}
