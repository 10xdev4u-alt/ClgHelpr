package handlers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/services"
)

// AssignmentHandler handles HTTP requests related to assignments.
type AssignmentHandler struct {
	assignmentService services.AssignmentService
	validator         *validator.Validate
}

// NewAssignmentHandler creates a new AssignmentHandler.
func NewAssignmentHandler(assignmentService services.AssignmentService) *AssignmentHandler {
	return &AssignmentHandler{
		assignmentService: assignmentService,
		validator:         validator.New(),
	}
}

// CreateAssignment handles creating a new assignment.
// @Summary Create a new assignment
// @Description Create a new assignment for the authenticated user.
// @Tags Assignments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param assignment body models.AssignmentCreationInput true "Assignment details"
// @Success 201 {object} models.Assignment
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /assignments [post]
func (h *AssignmentHandler) CreateAssignment(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var input models.AssignmentCreationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	assignment, err := h.assignmentService.CreateAssignment(context.Background(), userID, &input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create assignment: " + err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(assignment)
}

// GetAssignments handles retrieving all assignments for the authenticated user.
// @Summary Get all assignments
// @Description Retrieve a list of all assignments for the authenticated user.
// @Tags Assignments
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Assignment
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /assignments [get]
func (h *AssignmentHandler) GetAssignments(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	assignments, err := h.assignmentService.GetAssignmentsByUserID(context.Background(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve assignments: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(assignments)
}

// GetAssignmentByID handles retrieving a single assignment by ID.
// @Summary Get assignment by ID
// @Description Retrieve a single assignment by its ID for the authenticated user.
// @Tags Assignments
// @Produce json
// @Security BearerAuth
// @Param id path string true "Assignment ID"
// @Success 200 {object} models.Assignment
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /assignments/{id} [get]
func (h *AssignmentHandler) GetAssignmentByID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")
	assignment, err := h.assignmentService.GetAssignmentByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Assignment not found or not owned by user"})
	}
	if assignment.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Assignment does not belong to user"})
	}
	return c.Status(fiber.StatusOK).JSON(assignment)
}

// GetPendingAssignments handles retrieving pending assignments for the authenticated user.
// @Summary Get pending assignments
// @Description Retrieve a list of pending assignments for the authenticated user.
// @Tags Assignments
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Assignment
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /assignments/pending [get]
func (h *AssignmentHandler) GetPendingAssignments(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	assignments, err := h.assignmentService.GetPendingAssignmentsByUserID(context.Background(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve pending assignments: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(assignments)
}

// GetOverdueAssignments handles retrieving overdue assignments for the authenticated user.
// @Summary Get overdue assignments
// @Description Retrieve a list of overdue assignments for the authenticated user.
// @Tags Assignments
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Assignment
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /assignments/overdue [get]
func (h *AssignmentHandler) GetOverdueAssignments(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	assignments, err := h.assignmentService.GetOverdueAssignmentsByUserID(context.Background(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve overdue assignments: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(assignments)
}

// UpdateAssignment handles updating an existing assignment.
// @Summary Update an assignment
// @Description Update an existing assignment for the authenticated user.
// @Tags Assignments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Assignment ID"
// @Param assignment body models.AssignmentCreationInput true "Updated assignment details"
// @Success 200 {object} models.Assignment
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /assignments/{id} [put]
func (h *AssignmentHandler) UpdateAssignment(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")
	var input models.AssignmentCreationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	assignment, err := h.assignmentService.UpdateAssignment(context.Background(), userID, id, &input)
	if err != nil {
		if err.Error() == "assignment does not belong to user" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
		}
		if err.Error() == "assignment not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update assignment: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(assignment)
}

// UpdateAssignmentStatus handles updating the status of an assignment.
// @Summary Update assignment status
// @Description Update the status of an assignment for the authenticated user.
// @Tags Assignments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Assignment ID"
// @Param status body string true "New status (e.g., 'completed')"
// @Success 200 {object} map[string]string "Status updated"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /assignments/{id}/status [patch]
func (h *AssignmentHandler) UpdateAssignmentStatus(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")
	var body struct {
		Status string `json:"status" validate:"required"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := h.validator.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Basic check to ensure status is one of the allowed values
	allowedStatus := map[string]bool{
		"pending": true, "in_progress": true, "completed": true,
		"submitted": true, "graded": true, "overdue": true,
	}
	if !allowedStatus[body.Status] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status value"})
	}

	// Check ownership before updating
	assignment, err := h.assignmentService.GetAssignmentByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Assignment not found"})
	}
	if assignment.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Assignment does not belong to user"})
	}

	if err := h.assignmentService.UpdateAssignmentStatus(context.Background(), id, body.Status); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update assignment status: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Assignment status updated successfully"})
}

// DeleteAssignment handles deleting an assignment.
// @Summary Delete an assignment
// @Description Delete an assignment for the authenticated user.
// @Tags Assignments
// @Security BearerAuth
// @Param id path string true "Assignment ID"
// @Success 204 "Assignment deleted"
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /assignments/{id} [delete]
func (h *AssignmentHandler) DeleteAssignment(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")

	// Check ownership before deleting
	assignment, err := h.assignmentService.GetAssignmentByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Assignment not found"})
	}
	if assignment.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Assignment does not belong to user"})
	}

	if err := h.assignmentService.DeleteAssignment(context.Background(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete assignment: " + err.Error()})
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
}
