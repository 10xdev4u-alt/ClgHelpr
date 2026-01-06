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

// AnalyticsHandler handles HTTP requests related to analytics.
type AnalyticsHandler struct {
	analyticsService services.AnalyticsService
	validator        *validator.Validate
}

// NewAnalyticsHandler creates a new AnalyticsHandler.
func NewAnalyticsHandler(analyticsService services.AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{
		analyticsService: analyticsService,
		validator:        validator.New(),
	}
}

// CreateActivityLog handles creating a new activity log.
// @Summary Create a new activity log
// @Description Create a new activity log for the authenticated user.
// @Tags Analytics
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param log body models.ActivityLogCreationInput true "Activity log details"
// @Success 201 {object} models.ActivityLog
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /analytics/activity-logs [post]
func (h *AnalyticsHandler) CreateActivityLog(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var input models.ActivityLogCreationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logEntry, err := h.analyticsService.CreateActivityLog(context.Background(), userID, &input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create activity log: " + err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(logEntry)
}

// GetActivityLogs handles retrieving all activity logs for the authenticated user.
// @Summary Get all activity logs
// @Description Retrieve a list of all activity logs for the authenticated user.
// @Tags Analytics
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.ActivityLog
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /analytics/activity-logs [get]
func (h *AnalyticsHandler) GetActivityLogs(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	logs, err := h.analyticsService.GetActivityLogsByUserID(context.Background(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve activity logs: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(logs)
}

// GetDailyStats handles retrieving all daily stats for the authenticated user.
// @Summary Get all daily stats
// @Description Retrieve a list of all daily statistics for the authenticated user.
// @Tags Analytics
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.DailyStats
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /analytics/daily-stats [get]
func (h *AnalyticsHandler) GetDailyStats(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	stats, err := h.analyticsService.GetDailyStatsByUserID(context.Background(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve daily stats: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(stats)
}

// GetDailyStatsByDate handles retrieving daily stats for a specific date for the authenticated user.
// @Summary Get daily stats by date
// @Description Retrieve daily statistics for a specific date for the authenticated user.
// @Tags Analytics
// @Produce json
// @Security BearerAuth
// @Param date query string true "Date (YYYY-MM-DD)"
// @Success 200 {object} models.DailyStats
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string "Daily stats not found for date"
// @Failure 500 {object} map[string]string
// @Router /analytics/daily-stats/date [get]
func (h *AnalyticsHandler) GetDailyStatsByDate(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	dateStr := c.Query("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date format. Use YYYY-MM-DD."})
	}

	stats, err := h.analyticsService.GetDailyStatsByUserIDAndDate(context.Background(), userID, date)
	if err != nil {
		if err.Error() == "failed to get daily stats by user ID and date: no rows in result set" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Daily stats not found for date"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve daily stats: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(stats)
}

// UpsertDailyStats handles creating or updating daily stats.
// @Summary Create or update daily stats
// @Description Create or update daily statistics for the authenticated user for a specific date.
// @Tags Analytics
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param stats body models.DailyStatsUpdateInput true "Daily stats update details"
// @Success 200 {object} models.DailyStats
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /analytics/daily-stats [put]
func (h *AnalyticsHandler) UpsertDailyStats(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var input models.DailyStatsUpdateInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	stats, err := h.analyticsService.UpsertDailyStats(context.Background(), userID, &input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to upsert daily stats: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(stats)
}
