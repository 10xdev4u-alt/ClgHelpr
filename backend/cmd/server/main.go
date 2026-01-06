package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/config"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/database"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/handlers"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/middleware"
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
	defer dbPool.Close()

	app := fiber.New()

	api := app.Group("/api")

	// Initialize dependencies
	userRepo := repository.NewPGUserRepository(dbPool)
	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	authHandler := handlers.NewAuthHandler(authService)

	// --- Public Routes ---
	authRoutes := api.Group("/auth")
	authRoutes.Post("/register", authHandler.RegisterUser)
	authRoutes.Post("/login", authHandler.LoginUser)

	// Base welcome route
	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome to the Campus Pilot API!"})
	})

	// --- Protected Routes ---
	protected := api.Group("/") // This group is now protected by the middleware
	protected.Use(middleware.Protected(cfg.JWTSecret))
	protected.Get("/me", authHandler.GetUserProfile)

	log.Printf("Starting server on port %s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}