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
	subjectRepo := repository.NewPGSubjectRepository(dbPool)
	staffRepo := repository.NewPGStaffRepository(dbPool)
	venueRepo := repository.NewPGVenueRepository(dbPool)
	slotRepo := repository.NewPGTimetableSlotRepository(dbPool)
	assignmentRepo := repository.NewPGAssignmentRepository(dbPool)

	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	timetableService := services.NewTimetableService(subjectRepo, staffRepo, venueRepo, slotRepo)
	assignmentService := services.NewAssignmentService(assignmentRepo)

	authHandler := handlers.NewAuthHandler(authService)
	timetableHandler := handlers.NewTimetableHandler(timetableService)
	assignmentHandler := handlers.NewAssignmentHandler(assignmentService)

	// --- Public Routes ---
	authRoutes := api.Group("/auth")
	authRoutes.Post("/register", authHandler.RegisterUser)
	authRoutes.Post("/login", authHandler.LoginUser)

	// Timetable Public Routes (e.g., for dropdowns before login)
	api.Get("/subjects", timetableHandler.GetAllSubjects)
	api.Get("/staff", timetableHandler.GetAllStaff)
	api.Get("/venues", timetableHandler.GetAllVenues)

	// Base welcome route
	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome to the Campus Pilot API!"})
	})

	// --- Protected Routes ---
	protected := api.Group("/") // This group is now protected by the middleware
	protected.Use(middleware.Protected(cfg.JWTSecret))
	protected.Get("/me", authHandler.GetUserProfile)

	// Timetable Protected Routes
	timetableProtectedRoutes := protected.Group("/timetable")
	timetableProtectedRoutes.Post("/subjects", timetableHandler.CreateSubject)
	timetableProtectedRoutes.Post("/staff", timetableHandler.CreateStaff)
	timetableProtectedRoutes.Post("/venues", timetableHandler.CreateVenue)
	timetableProtectedRoutes.Post("/slots", timetableHandler.CreateTimetableSlot)
	timetableProtectedRoutes.Get("/day/:dayOfWeek", timetableHandler.GetUserTimetableByDay)
	timetableProtectedRoutes.Get("/range", timetableHandler.GetUserTimetableByDateRange)
	timetableProtectedRoutes.Get("/export-ics", timetableHandler.ExportICSCalendar)

	// Assignment Protected Routes
	assignmentProtectedRoutes := protected.Group("/assignments")
	assignmentProtectedRoutes.Post("/", assignmentHandler.CreateAssignment)
	assignmentProtectedRoutes.Get("/", assignmentHandler.GetAssignments)
	assignmentProtectedRoutes.Get("/pending", assignmentHandler.GetPendingAssignments)
	assignmentProtectedRoutes.Get("/overdue", assignmentHandler.GetOverdueAssignments)
	assignmentProtectedRoutes.Get("/:id", assignmentHandler.GetAssignmentByID)
	assignmentProtectedRoutes.Put("/:id", assignmentHandler.UpdateAssignment)
	assignmentProtectedRoutes.Patch("/:id/status", assignmentHandler.UpdateAssignmentStatus)
	assignmentProtectedRoutes.Delete("/:id", assignmentHandler.DeleteAssignment)


	log.Printf("Starting server on port %s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}