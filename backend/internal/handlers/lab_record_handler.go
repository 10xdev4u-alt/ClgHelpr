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

// LabRecordHandler handles HTTP requests related to lab records.
type LabRecordHandler struct {
	labRecordService services.LabRecordService
	validator        *validator.Validate
}

// NewLabRecordHandler creates a new LabRecordHandler.
func NewLabRecordHandler(labRecordService services.LabRecordService) *LabRecordHandler {
	return &LabRecordHandler{
		labRecordService: labRecordService,
		validator:        validator.New(),
	}
}

// CreateLabRecord handles creating a new lab record.
// @Summary Create a new lab record
// @Description Create a new lab record for the authenticated user.
// @Tags Lab Records
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param record body models.LabRecordCreationInput true "Lab record details"
// @Success 201 {object} models.LabRecord
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /lab-records [post]
func (h *LabRecordHandler) CreateLabRecord(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var input models.LabRecordCreationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	record, err := h.labRecordService.CreateLabRecord(context.Background(), userID, &input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create lab record: " + err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(record)
}

// GetLabRecords handles retrieving all lab records for the authenticated user.
// @Summary Get all lab records
// @Description Retrieve a list of all lab records for the authenticated user.
// @Tags Lab Records
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.LabRecord
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /lab-records [get]
func (h *LabRecordHandler) GetLabRecords(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	records, err := h.labRecordService.GetLabRecordsByUserID(context.Background(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve lab records: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(records)
}

// GetLabRecordByID handles retrieving a single lab record by ID.
// @Summary Get lab record by ID
// @Description Retrieve a single lab record by its ID for the authenticated user.
// @Tags Lab Records
// @Produce json
// @Security BearerAuth
// @Param id path string true "Lab Record ID"
// @Success 200 {object} models.LabRecord
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /lab-records/{id} [get]
func (h *LabRecordHandler) GetLabRecordByID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")
	record, err := h.labRecordService.GetLabRecordByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Lab record not found or not owned by user"})
	}
	if record.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Lab record does not belong to user"})
	}
	return c.Status(fiber.StatusOK).JSON(record)
}

// GetLabRecordsBySubjectID handles retrieving lab records by subject ID.
// @Summary Get lab records by subject ID
// @Description Retrieve a list of lab records for a specific subject for the authenticated user.
// @Tags Lab Records
// @Produce json
// @Security BearerAuth
// @Param subjectId path string true "Subject ID"
// @Success 200 {array} models.LabRecord
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /lab-records/subject/{subjectId} [get]
func (h *LabRecordHandler) GetLabRecordsBySubjectID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	subjectID := c.Params("subjectId")
	records, err := h.labRecordService.GetLabRecordsByUserIDAndSubjectID(context.Background(), userID, subjectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve lab records: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(records)
}

// UpdateLabRecord handles updating an existing lab record.
// @Summary Update a lab record
// @Description Update an existing lab record for the authenticated user.
// @Tags Lab Records
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Lab Record ID"
// @Param record body models.LabRecordCreationInput true "Updated lab record details"
// @Success 200 {object} models.LabRecord
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /lab-records/{id} [put]
func (h *LabRecordHandler) UpdateLabRecord(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")
	var input models.LabRecordCreationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	record, err := h.labRecordService.UpdateLabRecord(context.Background(), userID, id, &input)
	if err != nil {
		if err.Error() == "lab record does not belong to user" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
		}
		if err.Error() == "lab record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update lab record: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(record)
}

// UpdateLabRecordStatus handles updating the status of a lab record.
// @Summary Update lab record status
// @Description Update the status of a lab record for the authenticated user.
// @Tags Lab Records
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Lab Record ID"
// @Param status body string true "New status (e.g., 'submitted')"
// @Success 200 {object} map[string]string "Status updated"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /lab-records/{id}/status [patch]
func (h *LabRecordHandler) UpdateLabRecordStatus(c *fiber.Ctx) error {
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
		"pending": true, "practiced": true, "written": true, "printed": true,
		"submitted": true, "signed": true, "returned": true,
	}
	if !allowedStatus[body.Status] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status value"})
	}

	// Check ownership before updating
	record, err := h.labRecordService.GetLabRecordByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Lab record not found"})
	}
	if record.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Lab record does not belong to user"})
	}

	if err := h.labRecordService.UpdateLabRecordStatus(context.Background(), id, body.Status); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update lab record status: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Lab record status updated successfully"})
}

// DeleteLabRecord handles deleting a lab record.
// @Summary Delete a lab record
// @Description Delete a lab record for the authenticated user.
// @Tags Lab Records
// @Security BearerAuth
// @Param id path string true "Lab Record ID"
// @Success 204 "Lab record deleted"
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /lab-records/{id} [delete]
func (h *LabRecordHandler) DeleteLabRecord(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")

	// Check ownership before deleting
	record, err := h.labRecordService.GetLabRecordByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Lab record not found"})
	}
	if record.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Lab record does not belong to user"})
	}

	if err := h.labRecordService.DeleteLabRecord(context.Background(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete lab record: " + err.Error()})
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
}
