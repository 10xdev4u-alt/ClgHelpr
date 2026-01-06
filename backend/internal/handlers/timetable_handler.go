package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/services"
)

// TimetableHandler handles HTTP requests related to timetables.
type TimetableHandler struct {
	timetableService services.TimetableService
	validator        *validator.Validate
}

// NewTimetableHandler creates a new TimetableHandler.
func NewTimetableHandler(timetableService services.TimetableService) *TimetableHandler {
	return &TimetableHandler{
		timetableService: timetableService,
		validator:        validator.New(),
	}
}

// CreateSubject handles creating a new subject.
// @Summary Create a new subject
// @Description Create a new academic subject.
// @Tags Timetable
// @Accept json
// @Produce json
// @Param subject body models.Subject true "Subject details"
// @Success 201 {object} models.Subject
// @Failure 400 {object} map[string]string
// @Router /timetable/subjects [post]
func (h *TimetableHandler) CreateSubject(c *fiber.Ctx) error {
	var subject models.Subject
	if err := c.BodyParser(&subject); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(subject); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.timetableService.CreateSubject(context.Background(), &subject); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create subject"})
	}
	return c.Status(fiber.StatusCreated).JSON(subject)
}

// GetAllSubjects handles retrieving all subjects.
// @Summary Get all subjects
// @Description Retrieve a list of all academic subjects.
// @Tags Timetable
// @Produce json
// @Success 200 {array} models.Subject
// @Failure 500 {object} map[string]string
// @Router /timetable/subjects [get]
func (h *TimetableHandler) GetAllSubjects(c *fiber.Ctx) error {
	subjects, err := h.timetableService.GetAllSubjects(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve subjects"})
	}
	return c.Status(fiber.StatusOK).JSON(subjects)
}

// CreateStaff handles creating a new staff member.
// @Summary Create a new staff member
// @Description Create a new staff entry.
// @Tags Timetable
// @Accept json
// @Produce json
// @Param staff body models.Staff true "Staff details"
// @Success 201 {object} models.Staff
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /timetable/staff [post]
func (h *TimetableHandler) CreateStaff(c *fiber.Ctx) error {
	var staff models.Staff
	if err := c.BodyParser(&staff); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(staff); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.timetableService.CreateStaff(context.Background(), &staff); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create staff"})
	}
	return c.Status(fiber.StatusCreated).JSON(staff)
}

// GetAllStaff handles retrieving all staff members.
// @Summary Get all staff members
// @Description Retrieve a list of all staff members.
// @Tags Timetable
// @Produce json
// @Success 200 {array} models.Staff
// @Failure 500 {object} map[string]string
// @Router /timetable/staff [get]
func (h *TimetableHandler) GetAllStaff(c *fiber.Ctx) error {
	staffMembers, err := h.timetableService.GetAllStaff(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve staff"})
	}
	return c.Status(fiber.StatusOK).JSON(staffMembers)
}

// CreateVenue handles creating a new venue.
// @Summary Create a new venue
// @Description Create a new venue entry.
// @Tags Timetable
// @Accept json
// @Produce json
// @Param venue body models.Venue true "Venue details"
// @Success 201 {object} models.Venue
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /timetable/venues [post]
func (h *TimetableHandler) CreateVenue(c *fiber.Ctx) error {
	var venue models.Venue
	if err := c.BodyParser(&venue); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(venue); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.timetableService.CreateVenue(context.Background(), &venue); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create venue"})
	}
	return c.Status(fiber.StatusCreated).JSON(venue)
}

// GetAllVenues handles retrieving all venues.
// @Summary Get all venues
// @Description Retrieve a list of all venues.
// @Tags Timetable
// @Produce json
// @Success 200 {array} models.Venue
// @Failure 500 {object} map[string]string
// @Router /timetable/venues [get]
func (h *TimetableHandler) GetAllVenues(c *fiber.Ctx) error {
	venues, err := h.timetableService.GetAllVenues(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve venues"})
	}
	return c.Status(fiber.StatusOK).JSON(venues)
}

// CreateTimetableSlot handles creating a new timetable slot.
// @Summary Create a new timetable slot
// @Description Create a new timetable slot for a user.
// @Tags Timetable
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param slot body models.TimetableSlot true "Timetable slot details"
// @Success 201 {object} models.TimetableSlot
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /timetable/slots [post]
func (h *TimetableHandler) CreateTimetableSlot(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var slot models.TimetableSlot
	if err := c.BodyParser(&slot); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(slot); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	slot.UserID = userID // Assign the authenticated user's ID
	if err := h.timetableService.CreateTimetableSlot(context.Background(), &slot); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create timetable slot"})
	}
	return c.Status(fiber.StatusCreated).JSON(slot)
}

// GetUserTimetableByDay handles retrieving a user's timetable for a specific day.
// @Summary Get user's timetable by day
// @Description Retrieve a user's recurring timetable slots for a specific day of the week.
// @Tags Timetable
// @Produce json
// @Security BearerAuth
// @Param dayOfWeek path int true "Day of the week (0=Sunday, 1=Monday...)"
// @Success 200 {array} models.TimetableSlot
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /timetable/day/:dayOfWeek [get]
func (h *TimetableHandler) GetUserTimetableByDay(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	dayOfWeekStr := c.Params("dayOfWeek")
	dayOfWeek, err := strconv.Atoi(dayOfWeekStr)
	if err != nil || dayOfWeek < 0 || dayOfWeek > 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid day of week. Must be 0-6."})
	}

	slots, err := h.timetableService.GetUserTimetableByDay(context.Background(), userID, int32(dayOfWeek))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve timetable slots"})
	}
	return c.Status(fiber.StatusOK).JSON(slots)
}

// GetUserTimetableByDateRange handles retrieving a user's timetable for a date range.
// @Summary Get user's timetable by date range
// @Description Retrieve a user's timetable slots (recurring and specific) within a given date range.
// @Tags Timetable
// @Produce json
// @Security BearerAuth
// @Param start query string true "Start date (YYYY-MM-DD)"
// @Param end query string true "End date (YYYY-MM-DD)"
// @Success 200 {array} models.TimetableSlot
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /timetable/range [get]
func (h *TimetableHandler) GetUserTimetableByDateRange(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	startStr := c.Query("start")
	endStr := c.Query("end")

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid start date format. Use YYYY-MM-DD."})
	}
	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid end date format. Use YYYY-MM-DD."})
	}

	// Adjust end date to include the whole day
	end = end.Add(24*time.Hour - time.Nanosecond)

	slots, err := h.timetableService.GetUserTimetableByDateRange(context.Background(), userID, start, end)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve timetable slots"})
	}
	return c.Status(fiber.StatusOK).JSON(slots)
}
