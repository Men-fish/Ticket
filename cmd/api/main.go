package main

import (
	"fmt"

	"github.com/Men-fish/ticket-v1/config"
	"github.com/Men-fish/ticket-v1/db"
	"github.com/Men-fish/ticket-v1/handlers"
	"github.com/Men-fish/ticket-v1/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "Ticket-Booking",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(db)

	// Routing
	server := app.Group("/api")

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
