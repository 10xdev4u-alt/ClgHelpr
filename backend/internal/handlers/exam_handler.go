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

// ExamHandler handles HTTP requests related to exams.
type ExamHandler struct {
	examService services.ExamService
	validator   *validator.Validate
}

// NewExamHandler creates a new ExamHandler.
func NewExamHandler(examService services.ExamService) *ExamHandler {
	return &ExamHandler{
		examService: examService,
		validator:   validator.New(),
	}
}

// CreateExam handles creating a new exam.
// @Summary Create a new exam
// @Description Create a new exam for the authenticated user.
// @Tags Exams
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param exam body models.ExamCreationInput true "Exam details"
// @Success 201 {object} models.Exam
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /exams [post]
func (h *ExamHandler) CreateExam(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var input models.ExamCreationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	exam, err := h.examService.CreateExam(context.Background(), userID, &input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create exam: " + err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(exam)
}

// GetExams handles retrieving all exams for the authenticated user.
// @Summary Get all exams
// @Description Retrieve a list of all exams for the authenticated user.
// @Tags Exams
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Exam
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /exams [get]
func (h *ExamHandler) GetExams(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	exams, err := h.examService.GetExamsByUserID(context.Background(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve exams: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(exams)
}

// GetExamByID handles retrieving a single exam by ID.
// @Summary Get exam by ID
// @Description Retrieve a single exam by its ID for the authenticated user.
// @Tags Exams
// @Produce json
// @Security BearerAuth
// @Param id path string true "Exam ID"
// @Success 200 {object} models.Exam
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /exams/{id} [get]
func (h *ExamHandler) GetExamByID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")
	exam, err := h.examService.GetExamByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Exam not found or not owned by user"})
	}
	if exam.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Exam does not belong to user"})
	}
	return c.Status(fiber.StatusOK).JSON(exam)
}

// GetUpcomingExams handles retrieving upcoming exams for the authenticated user.
// @Summary Get upcoming exams
// @Description Retrieve a list of upcoming exams for the authenticated user.
// @Tags Exams
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Exam
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /exams/upcoming [get]
func (h *ExamHandler) GetUpcomingExams(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	exams, err := h.examService.GetUpcomingExamsByUserID(context.Background(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve upcoming exams: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(exams)
}

// UpdateExam handles updating an existing exam.
// @Summary Update an exam
// @Description Update an existing exam for the authenticated user.
// @Tags Exams
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Exam ID"
// @Param exam body models.ExamCreationInput true "Updated exam details"
// @Success 200 {object} models.Exam
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /exams/{id} [put]
func (h *ExamHandler) UpdateExam(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")
	var input models.ExamCreationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	exam, err := h.examService.UpdateExam(context.Background(), userID, id, &input)
	if err != nil {
		if err.Error() == "exam does not belong to user" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
		}
		if err.Error() == "exam not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update exam: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(exam)
}

// UpdateExamPrepStatus handles updating the prep status of an exam.
// @Summary Update exam prep status
// @Description Update the preparation status of an exam for the authenticated user.
// @Tags Exams
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Exam ID"
// @Param status body string true "New prep status (e.g., 'revision')"
// @Success 200 {object} map[string]string "Status updated"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /exams/{id}/prep-status [patch]
func (h *ExamHandler) UpdateExamPrepStatus(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")
	var body struct {
		PrepStatus string `json:"prepStatus" validate:"required"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := h.validator.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Basic check to ensure status is one of the allowed values
	allowedStatus := map[string]bool{
		"not_started": true, "in_progress": true, "revision": true, "ready": true,
	}
	if !allowedStatus[body.PrepStatus] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid prep status value"})
	}

	// Check ownership before updating
	exam, err := h.examService.GetExamByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Exam not found"})
	}
	if exam.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Exam does not belong to user"})
	}

	if err := h.examService.UpdateExamPrepStatus(context.Background(), id, body.PrepStatus); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update exam prep status: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Exam prep status updated successfully"})
}

// DeleteExam handles deleting an exam.
// @Summary Delete an exam
// @Description Delete an exam for the authenticated user.
// @Tags Exams
// @Security BearerAuth
// @Param id path string true "Exam ID"
// @Success 204 "Exam deleted"
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /exams/{id} [delete]
func (h *ExamHandler) DeleteExam(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")

	// Check ownership before deleting
	exam, err := h.examService.GetExamByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Exam not found"})
	}
	if exam.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Exam does not belong to user"})
	}

	if err := h.examService.DeleteExam(context.Background(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete exam: " + err.Error()})
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
}

// CreateImportantQuestion handles creating a new important question.
// @Summary Create a new important question
// @Description Create a new important question for the authenticated user.
// @Tags Important Questions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param question body models.ImportantQuestionCreationInput true "Important question details"
// @Success 201 {object} models.ImportantQuestion
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /important-questions [post]
func (h *ExamHandler) CreateImportantQuestion(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var input models.ImportantQuestionCreationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	question, err := h.examService.CreateImportantQuestion(context.Background(), userID, &input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create important question: " + err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(question)
}

// GetImportantQuestionsByExamID handles retrieving important questions by exam ID.
// @Summary Get important questions by exam ID
// @Description Retrieve a list of important questions for a specific exam for the authenticated user.
// @Tags Important Questions
// @Produce json
// @Security BearerAuth
// @Param examId path string true "Exam ID"
// @Success 200 {array} models.ImportantQuestion
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /important-questions/exam/{examId} [get]
func (h *ExamHandler) GetImportantQuestionsByExamID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	examID := c.Params("examId")
	questions, err := h.examService.GetImportantQuestionsByExamID(context.Background(), examID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve important questions: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(questions)
}

// GetImportantQuestionsBySubjectID handles retrieving important questions by subject ID.
// @Summary Get important questions by subject ID
// @Description Retrieve a list of important questions for a specific subject for the authenticated user.
// @Tags Important Questions
// @Produce json
// @Security BearerAuth
// @Param subjectId path string true "Subject ID"
// @Success 200 {array} models.ImportantQuestion
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /important-questions/subject/{subjectId} [get]
func (h *ExamHandler) GetImportantQuestionsBySubjectID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	subjectID := c.Params("subjectId")
	questions, err := h.examService.GetImportantQuestionsBySubjectID(context.Background(), subjectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve important questions: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(questions)
}

// UpdateImportantQuestion handles updating an existing important question.
// @Summary Update an important question
// @Description Update an existing important question for the authenticated user.
// @Tags Important Questions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Important Question ID"
// @Param question body models.ImportantQuestionCreationInput true "Updated important question details"
// @Success 200 {object} models.ImportantQuestion
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /important-questions/{id} [put]
func (h *ExamHandler) UpdateImportantQuestion(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")
	var input models.ImportantQuestionCreationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	question, err := h.examService.UpdateImportantQuestion(context.Background(), userID, id, &input)
	if err != nil {
		if err.Error() == "important question does not belong to user" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
		}
		if err.Error() == "important question not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update important question: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(question)
}

// DeleteImportantQuestion handles deleting an important question.
// @Summary Delete an important question
// @Description Delete an important question for the authenticated user.
// @Tags Important Questions
// @Security BearerAuth
// @Param id path string true "Important Question ID"
// @Success 204 "Important question deleted"
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /important-questions/{id} [delete]
func (h *ExamHandler) DeleteImportantQuestion(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")

	// Check ownership before deleting
	question, err := h.examService.GetImportantQuestionByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Important question not found"})
	}
	if question.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Important question does not belong to user"})
	}

	if err := h.examService.DeleteImportantQuestion(context.Background(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete important question: " + err.Error()})
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
}
