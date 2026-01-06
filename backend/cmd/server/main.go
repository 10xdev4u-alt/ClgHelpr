package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/config"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/database"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/handlers"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/repository"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/services"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Connect to the database
	dbPool, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer dbPool.Close() // Close the database connection when the application exits

	app := fiber.New()

	// ... (imports)
	"github.com/princetheprogrammer/campus-pilot/backend/internal/middleware"

// ... (main function)

	api := app.Group("/api")

	// Public routes
	authRoutes := api.Group("/auth")
	authRoutes.Post("/register", authHandler.RegisterUser)
	authRoutes.Post("/login", authHandler.LoginUser)

	// Protected routes
	protected := api.Group("/", middleware.Protected(cfg.JWTSecret))
	protected.Get("/me", authHandler.GetUserProfile)

	// Base API route
	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to the Campus Pilot API!",
		})
	})

	log.Printf("Starting server on port %s", cfg.Port)
	app.Listen(":" + cfg.Port)
}
