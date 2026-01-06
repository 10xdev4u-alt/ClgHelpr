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

				examRepo := repository.NewPGExamRepository(dbPool)

				importantQuestionRepo := repository.NewPGImportantQuestionRepository(dbPool)

				labRecordRepo := repository.NewPGLabRecordRepository(dbPool)

				documentRepo := repository.NewPGDocumentRepository(dbPool)

				studyPlanRepo := repository.NewPGStudyPlanRepository(dbPool)

				studySessionRepo := repository.NewPGStudySessionRepository(dbPool)

			

				authService := services.NewAuthService(userRepo, cfg.JWTSecret)

				timetableService := services.NewTimetableService(subjectRepo, staffRepo, venueRepo, slotRepo)

				assignmentService := services.NewAssignmentService(assignmentRepo)

				examService := services.NewExamService(examRepo, importantQuestionRepo)

				labRecordService := services.NewLabRecordService(labRecordRepo)

				documentService := services.NewDocumentService(documentRepo)

				studyPlanService := services.NewStudyPlanService(studyPlanRepo, studySessionRepo)

			

				authHandler := handlers.NewAuthHandler(authService)

				timetableHandler := handlers.NewTimetableHandler(timetableService)

				assignmentHandler := handlers.NewAssignmentHandler(assignmentService)

				examHandler := handlers.NewExamHandler(examService)

				labRecordHandler := handlers.NewLabRecordHandler(labRecordService)

				documentHandler := handlers.NewDocumentHandler(documentService)

				studyPlanHandler := handlers.NewStudyPlanHandler(studyPlanService)

			

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

			

				// Exam Protected Routes

				examProtectedRoutes := protected.Group("/exams")

				examProtectedRoutes.Post("/", examHandler.CreateExam)

				examProtectedRoutes.Get("/", examHandler.GetExams)

				examProtectedRoutes.Get("/upcoming", examHandler.GetUpcomingExams)

				examProtectedRoutes.Get("/:id", examHandler.GetExamByID)

				examProtectedRoutes.Put("/:id", examHandler.UpdateExam)

				examProtectedRoutes.Patch("/:id/prep-status", examHandler.UpdateExamPrepStatus)

				examProtectedRoutes.Delete("/:id", examHandler.DeleteExam)

				

				importantQuestionProtectedRoutes := protected.Group("/important-questions")

				importantQuestionProtectedRoutes.Post("/", examHandler.CreateImportantQuestion)

				importantQuestionProtectedRoutes.Get("/exam/:examId", examHandler.GetImportantQuestionsByExamID)

				importantQuestionProtectedRoutes.Get("/subject/:subjectId", examHandler.GetImportantQuestionsBySubjectID)

				importantQuestionProtectedRoutes.Get("/:id", examHandler.GetImportantQuestionByID)

				importantQuestionProtectedRoutes.Put("/:id", examHandler.UpdateImportantQuestion)

				importantQuestionProtectedRoutes.Delete("/:id", examHandler.DeleteImportantQuestion)

			

				// Lab Record Protected Routes

				labRecordProtectedRoutes := protected.Group("/lab-records")

				labRecordProtectedRoutes.Post("/", labRecordHandler.CreateLabRecord)

				labRecordProtectedRoutes.Get("/", labRecordHandler.GetLabRecords)

				labRecordProtectedRoutes.Get("/subject/:subjectId", labRecordHandler.GetLabRecordsBySubjectID)

				labRecordProtectedRoutes.Get("/:id", labRecordHandler.GetLabRecordByID)

				labRecordProtectedRoutes.Put("/:id", labRecordHandler.UpdateLabRecord)

				labRecordProtectedRoutes.Patch("/:id/status", labRecordHandler.UpdateLabRecordStatus)

				labRecordProtectedRoutes.Delete("/:id", labRecordHandler.DeleteLabRecord)

			

				// Study Plan Protected Routes

				studyPlanProtectedRoutes := protected.Group("/study-plans")

				studyPlanProtectedRoutes.Post("/", studyPlanHandler.CreateStudyPlan)

				studyPlanProtectedRoutes.Get("/", studyPlanHandler.GetStudyPlans)

				studyPlanProtectedRoutes.Get("/date", studyPlanHandler.GetStudyPlansByDate)

				studyPlanProtectedRoutes.Get("/:id", studyPlanHandler.GetStudyPlanByID)

				studyPlanProtectedRoutes.Put("/:id", studyPlanHandler.UpdateStudyPlan)

				studyPlanProtectedRoutes.Delete("/:id", studyPlanHandler.DeleteStudyPlan)

			

				studySessionProtectedRoutes := protected.Group("/study-sessions")

				studySessionProtectedRoutes.Post("/", studyPlanHandler.CreateStudySession)

				studySessionProtectedRoutes.Get("/", studyPlanHandler.GetStudySessions)

				studySessionProtectedRoutes.Get("/plan/:studyPlanId", studyPlanHandler.GetStudySessionsByStudyPlanID)

				studySessionProtectedRoutes.Get("/:id", studyPlanHandler.GetStudySessionByID)

				studySessionProtectedRoutes.Put("/:id", studyPlanHandler.UpdateStudySession)

				studySessionProtectedRoutes.Delete("/:id", studyPlanHandler.DeleteStudySession)

			

			

				log.Printf("Starting server on port %s", cfg.Port)

				log.Fatal(app.Listen(":" + cfg.Port))

			

		

	
}