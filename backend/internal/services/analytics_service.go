package services

import (
	"context"
	"fmt"
	"time"

	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/repository"
)

// AnalyticsService defines the interface for analytics-related business logic.
type AnalyticsService interface {
	CreateActivityLog(ctx context.Context, userID string, input *models.ActivityLogCreationInput) (*models.ActivityLog, error)
	GetActivityLogsByUserID(ctx context.Context, userID string) ([]models.ActivityLog, error)
	GetDailyStatsByUserID(ctx context.Context, userID string) ([]models.DailyStats, error)
	GetDailyStatsByUserIDAndDate(ctx context.Context, userID string, date time.Time) (*models.DailyStats, error)
	UpsertDailyStats(ctx context.Context, userID string, input *models.DailyStatsUpdateInput) (*models.DailyStats, error)
}

// analyticsService implements AnalyticsService.
type analyticsService struct {
	activityLogRepo repository.ActivityLogRepository
	dailyStatsRepo  repository.DailyStatsRepository
}

// NewAnalyticsService creates a new analytics service.
func NewAnalyticsService(activityLogRepo repository.ActivityLogRepository, dailyStatsRepo repository.DailyStatsRepository) AnalyticsService {
	return &analyticsService{
		activityLogRepo: activityLogRepo,
		dailyStatsRepo:  dailyStatsRepo,
	}
}

// CreateActivityLog creates a new activity log for a user.
func (s *analyticsService) CreateActivityLog(ctx context.Context, userID string, input *models.ActivityLogCreationInput) (*models.ActivityLog, error) {
	activityLog := &models.ActivityLog{
		UserID:       userID,
		ActivityType: input.ActivityType,
		Metadata:     input.Metadata,
	}
	if input.Description != nil {
		activityLog.Description = sql.NullString{String: *input.Description, Valid: true}
	}
	if input.EntityType != nil {
		activityLog.EntityType = sql.NullString{String: *input.EntityType, Valid: true}
	}
	if input.EntityID != nil {
		activityLog.EntityID = sql.NullString{String: *input.EntityID, Valid: true}
	}
	if input.IPAddress != nil {
		activityLog.IPAddress = sql.NullString{String: *input.IPAddress, Valid: true}
	}
	if input.UserAgent != nil {
		activityLog.UserAgent = sql.NullString{String: *input.UserAgent, Valid: true}
	}

	if err := s.activityLogRepo.CreateActivityLog(ctx, activityLog); err != nil {
		return nil, fmt.Errorf("failed to create activity log: %w", err)
	}
	return activityLog, nil
}

// GetActivityLogsByUserID retrieves all activity logs for a user.
func (s *analyticsService) GetActivityLogsByUserID(ctx context.Context, userID string) ([]models.ActivityLog, error) {
	return s.activityLogRepo.GetActivityLogsByUserID(ctx, userID)
}

// GetDailyStatsByUserID retrieves all daily stats for a user.
func (s *analyticsService) GetDailyStatsByUserID(ctx context.Context, userID string) ([]models.DailyStats, error) {
	return s.dailyStatsRepo.GetDailyStatsByUserID(ctx, userID)
}

// GetDailyStatsByUserIDAndDate retrieves daily stats for a user on a specific date.
func (s *analyticsService) GetDailyStatsByUserIDAndDate(ctx context.Context, userID string, date time.Time) (*models.DailyStats, error) {
	return s.dailyStatsRepo.GetDailyStatsByUserIDAndDate(ctx, userID, date)
}

// UpsertDailyStats creates or updates daily stats for a user.
func (s *analyticsService) UpsertDailyStats(ctx context.Context, userID string, input *models.DailyStatsUpdateInput) (*models.DailyStats, error) {
	statDate, err := time.Parse("2006-01-02", input.StatDate)
	if err != nil {
		return nil, fmt.Errorf("invalid stat date format: %w", err)
	}

	stats := &models.DailyStats{
		UserID:            userID,
		StatDate:          statDate,
		StudyMinutes:      0,
		SessionsCompleted: 0,
		TopicsCovered:     0,
		AssignmentsCompleted: 0,
		AssignmentsAdded:     0,
		ClassesAttended:      0,
		TotalClasses:         0,
		XPEarned:             0,
	}

	// Try to get existing stats
	existingStats, err := s.dailyStatsRepo.GetDailyStatsByUserIDAndDate(ctx, userID, statDate)
	if err == nil { // Stats exist, update them
		stats = existingStats
	} else if err != nil && err.Error() != "failed to get daily stats by user ID and date: no rows in result set" {
		// Only propagate error if it's not a "not found" error
		return nil, fmt.Errorf("failed to check for existing daily stats: %w", err)
	}

	if input.StudyMinutes != nil {
		stats.StudyMinutes = *input.StudyMinutes
	}
	if input.SessionsCompleted != nil {
		stats.SessionsCompleted = *input.SessionsCompleted
	}
	if input.TopicsCovered != nil {
		stats.TopicsCovered = *input.TopicsCovered
	}
	if input.AssignmentsCompleted != nil {
		stats.AssignmentsCompleted = *input.AssignmentsCompleted
	}
	if input.AssignmentsAdded != nil {
		stats.AssignmentsAdded = *input.AssignmentsAdded
	}
	if input.ClassesAttended != nil {
		stats.ClassesAttended = *input.ClassesAttended
	}
	if input.TotalClasses != nil {
		stats.TotalClasses = *input.TotalClasses
	}
	if input.XPEarned != nil {
		stats.XPEarned = *input.XPEarned
	}

	if err := s.dailyStatsRepo.UpsertDailyStats(ctx, stats); err != nil {
		return nil, fmt.Errorf("failed to upsert daily stats: %w", err)
	}
	return stats, nil
}
