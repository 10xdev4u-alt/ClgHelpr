package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/config"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/database"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Connect to the database
	_, err = database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	app := fiber.New()

	api := app.Group("/api")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to the Campus Pilot API!",
		})
	})

	log.Printf("Starting server on port %s", cfg.Port)
	app.Listen(":" + cfg.Port)
}
