package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/services"
)

// StudyPlanHandler handles HTTP requests related to study plans and sessions.
type StudyPlanHandler struct {
	studyPlanService services.StudyPlanService
	validator        *validator.Validate
}

// NewStudyPlanHandler creates a new StudyPlanHandler.
func NewStudyPlanHandler(studyPlanService services.StudyPlanService) *StudyPlanHandler {
	return &StudyPlanHandler{
		studyPlanService: studyPlanService,
		validator:        validator.New(),
	}
}

// CreateStudyPlan handles creating a new study plan.
// @Summary Create a new study plan
// @Description Create a new study plan for the authenticated user.
// @Tags Study Plans
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param plan body models.StudyPlanCreationInput true "Study plan details"
// @Success 201 {object} models.StudyPlan
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /study-plans [post]
func (h *StudyPlanHandler) CreateStudyPlan(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var input models.StudyPlanCreationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	plan, err := h.studyPlanService.CreateStudyPlan(context.Background(), userID, &input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create study plan: " + err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(plan)
}

// GetStudyPlans handles retrieving all study plans for the authenticated user.
// @Summary Get all study plans
// @Description Retrieve a list of all study plans for the authenticated user.
// @Tags Study Plans
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.StudyPlan
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /study-plans [get]
func (h *StudyPlanHandler) GetStudyPlans(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	plans, err := h.studyPlanService.GetStudyPlansByUserID(context.Background(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve study plans: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(plans)
}

// GetStudyPlanByID handles retrieving a single study plan by ID.
// @Summary Get study plan by ID
// @Description Retrieve a single study plan by its ID for the authenticated user.
// @Tags Study Plans
// @Produce json
// @Security BearerAuth
// @Param id path string true "Study Plan ID"
// @Success 200 {object} models.StudyPlan
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /study-plans/{id} [get]
func (h *StudyPlanHandler) GetStudyPlanByID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")
	plan, err := h.studyPlanService.GetStudyPlanByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Study plan not found or not owned by user"})
	}
	if plan.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Study plan does not belong to user"})
	}
	return c.Status(fiber.StatusOK).JSON(plan)
}

// GetStudyPlansByDate handles retrieving study plans for a specific date for the authenticated user.
// @Summary Get study plans by date
// @Description Retrieve a list of study plans for a specific date for the authenticated user.
// @Tags Study Plans
// @Produce json
// @Security BearerAuth
// @Param date query string true "Date (YYYY-MM-DD)"
// @Success 200 {array} models.StudyPlan
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /study-plans/date [get]
func (h *StudyPlanHandler) GetStudyPlansByDate(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	dateStr := c.Query("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date format. Use YYYY-MM-DD."})
	}

	plans, err := h.studyPlanService.GetStudyPlansByUserIDAndDate(context.Background(), userID, date)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve study plans: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(plans)
}

// UpdateStudyPlan handles updating an existing study plan.
// @Summary Update a study plan
// @Description Update an existing study plan for the authenticated user.
// @Tags Study Plans
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Study Plan ID"
// @Param plan body models.StudyPlanCreationInput true "Updated study plan details"
// @Success 200 {object} models.StudyPlan
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /study-plans/{id} [put]
func (h *StudyPlanHandler) UpdateStudyPlan(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")
	var input models.StudyPlanCreationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	plan, err := h.studyPlanService.UpdateStudyPlan(context.Background(), userID, id, &input)
	if err != nil {
		if err.Error() == "study plan does not belong to user" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
		}
		if err.Error() == "study plan not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update study plan: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(plan)
}

// DeleteStudyPlan handles deleting a study plan.
// @Summary Delete a study plan
// @Description Delete a study plan for the authenticated user.
// @Tags Study Plans
// @Security BearerAuth
// @Param id path string true "Study Plan ID"
// @Success 204 "Study plan deleted"
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /study-plans/{id} [delete]
func (h *StudyPlanHandler) DeleteStudyPlan(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")

	// Check ownership before deleting
	plan, err := h.studyPlanService.GetStudyPlanByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Study plan not found"})
	}
	if plan.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Study plan does not belong to user"})
	}

	if err := h.studyPlanService.DeleteStudyPlan(context.Background(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete study plan: " + err.Error()})
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
}

// CreateStudySession handles creating a new study session.
// @Summary Create a new study session
// @Description Create a new study session for the authenticated user.
// @Tags Study Sessions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param session body models.StudySessionCreationInput true "Study session details"
// @Success 201 {object} models.StudySession
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /study-sessions [post]
func (h *StudyPlanHandler) CreateStudySession(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var input models.StudySessionCreationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	session, err := h.studyPlanService.CreateStudySession(context.Background(), userID, &input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create study session: " + err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(session)
}

// GetStudySessions handles retrieving all study sessions for the authenticated user.
// @Summary Get all study sessions
// @Description Retrieve a list of all study sessions for the authenticated user.
// @Tags Study Sessions
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.StudySession
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /study-sessions [get]
func (h *StudyPlanHandler) GetStudySessions(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	sessions, err := h.studyPlanService.GetStudySessionsByUserID(context.Background(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve study sessions: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(sessions)
}

// GetStudySessionByID handles retrieving a single study session by ID.
// @Summary Get study session by ID
// @Description Retrieve a single study session by its ID for the authenticated user.
// @Tags Study Sessions
// @Produce json
// @Security BearerAuth
// @Param id path string true "Study Session ID"
// @Success 200 {object} models.StudySession
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /study-sessions/{id} [get]
func (h *StudyPlanHandler) GetStudySessionByID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")
	session, err := h.studyPlanService.GetStudySessionByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Study session not found or not owned by user"})
	}
	if session.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Study session does not belong to user"})
	}
	return c.Status(fiber.StatusOK).JSON(session)
}

// GetStudySessionsByStudyPlanID handles retrieving study sessions for a specific study plan ID.
// @Summary Get study sessions by study plan ID
// @Description Retrieve a list of study sessions for a specific study plan for the authenticated user.
// @Tags Study Sessions
// @Produce json
// @Security BearerAuth
// @Param studyPlanId path string true "Study Plan ID"
// @Success 200 {array} models.StudySession
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /study-sessions/plan/{studyPlanId} [get]
func (h *StudyPlanHandler) GetStudySessionsByStudyPlanID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	studyPlanID := c.Params("studyPlanId")
	sessions, err := h.studyPlanService.GetStudySessionsByStudyPlanID(context.Background(), studyPlanID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve study sessions: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(sessions)
}

// UpdateStudySession handles updating an existing study session.
// @Summary Update a study session
// @Description Update an existing study session for the authenticated user.
// @Tags Study Sessions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Study Session ID"
// @Param session body models.StudySessionCreationInput true "Updated study session details"
// @Success 200 {object} models.StudySession
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /study-sessions/{id} [put]
func (h *StudyPlanHandler) UpdateStudySession(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")
	var input models.StudySessionCreationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	session, err := h.studyPlanService.UpdateStudySession(context.Background(), userID, id, &input)
	if err != nil {
		if err.Error() == "study session does not belong to user" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
		}
		if err.Error() == "study session not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update study session: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(session)
}

// DeleteStudySession handles deleting a study session.
// @Summary Delete a study session
// @Description Delete a study session for the authenticated user.
// @Tags Study Sessions
// @Security BearerAuth
// @Param id path string true "Study Session ID"
// @Success 204 "Study session deleted"
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /study-sessions/{id} [delete]
func (h *StudyPlanHandler) DeleteStudySession(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")

	// Check ownership before deleting
	session, err := h.studyPlanService.GetStudySessionByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Study session not found"})
	}
	if session.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Study session does not belong to user"})
	}

	if err := h.studyPlanService.DeleteStudySession(context.Background(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete study session: " + err.Error()})
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
}
