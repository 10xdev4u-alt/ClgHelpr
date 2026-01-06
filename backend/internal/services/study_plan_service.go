package services

import (
	"context"
	"fmt"
	"time"

	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/repository"
)

// StudyPlanService defines the interface for study plan-related business logic.
type StudyPlanService interface {
	CreateStudyPlan(ctx context.Context, userID string, input *models.StudyPlanCreationInput) (*models.StudyPlan, error)
	GetStudyPlanByID(ctx context.Context, id string) (*models.StudyPlan, error)
	GetStudyPlansByUserID(ctx context.Context, userID string) ([]models.StudyPlan, error)
	GetStudyPlansByUserIDAndDate(ctx context.Context, userID string, date time.Time) ([]models.StudyPlan, error)
	UpdateStudyPlan(ctx context.Context, userID string, id string, input *models.StudyPlanCreationInput) (*models.StudyPlan, error)
	DeleteStudyPlan(ctx context.Context, id string) error

	CreateStudySession(ctx context.Context, userID string, input *models.StudySessionCreationInput) (*models.StudySession, error)
	GetStudySessionByID(ctx context.Context, id string) (*models.StudySession, error)
	GetStudySessionsByUserID(ctx context.Context, userID string) ([]models.StudySession, error)
	GetStudySessionsByStudyPlanID(ctx context.Context, studyPlanID string) ([]models.StudySession, error)
	UpdateStudySession(ctx context.Context, userID string, id string, input *models.StudySessionCreationInput) (*models.StudySession, error)
	DeleteStudySession(ctx context.Context, id string) error
}

// studyPlanService implements StudyPlanService.
type studyPlanService struct {
	studyPlanRepo repository.StudyPlanRepository
	sessionRepo   repository.StudySessionRepository
}

// NewStudyPlanService creates a new study plan service.
func NewStudyPlanService(studyPlanRepo repository.StudyPlanRepository, sessionRepo repository.StudySessionRepository) StudyPlanService {
	return &studyPlanService{
		studyPlanRepo: studyPlanRepo,
		sessionRepo:   sessionRepo,
	}
}

// CreateStudyPlan creates a new study plan for a user.
func (s *studyPlanService) CreateStudyPlan(ctx context.Context, userID string, input *models.StudyPlanCreationInput) (*models.StudyPlan, error) {
	planDate, err := time.Parse("2006-01-02", input.PlanDate)
	if err != nil {
		return nil, fmt.Errorf("invalid plan date format: %w", err)
	}

	studyPlan := &models.StudyPlan{
		UserID:   userID,
		Title:    input.Title,
		PlanDate: planDate,
		PlanType: input.PlanType,
		Status:   "planned", // Default
	}

	if input.Notes != nil {
		studyPlan.Notes = sql.NullString{String: *input.Notes, Valid: true}
	}
	if input.Status != nil {
		studyPlan.Status = *input.Status
	}

	if err := s.studyPlanRepo.CreateStudyPlan(ctx, studyPlan); err != nil {
		return nil, fmt.Errorf("failed to create study plan: %w", err)
	}
	return studyPlan, nil
}

// GetStudyPlanByID retrieves a single study plan.
func (s *studyPlanService) GetStudyPlanByID(ctx context.Context, id string) (*models.StudyPlan, error) {
	return s.studyPlanRepo.GetStudyPlanByID(ctx, id)
}

// GetStudyPlansByUserID retrieves all study plans for a user.
func (s *studyPlanService) GetStudyPlansByUserID(ctx context.Context, userID string) ([]models.StudyPlan, error) {
	return s.studyPlanRepo.GetStudyPlansByUserID(ctx, userID)
}

// GetStudyPlansByUserIDAndDate retrieves all study plans for a user on a specific date.
func (s *studyPlanService) GetStudyPlansByUserIDAndDate(ctx context.Context, userID string, date time.Time) ([]models.StudyPlan, error) {
	return s.studyPlanRepo.GetStudyPlansByUserIDAndDate(ctx, userID, date)
}

// UpdateStudyPlan updates an existing study plan.
func (s *studyPlanService) UpdateStudyPlan(ctx context.Context, userID string, id string, input *models.StudyPlanCreationInput) (*models.StudyPlan, error) {
	existingPlan, err := s.studyPlanRepo.GetStudyPlanByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("study plan not found: %w", err)
	}
	if existingPlan.UserID != userID {
		return nil, fmt.Errorf("study plan does not belong to user")
	}

	if input.PlanDate != "" {
		planDate, err := time.Parse("2006-01-02", input.PlanDate)
		if err != nil {
			return nil, fmt.Errorf("invalid plan date format: %w", err)
		}
		existingPlan.PlanDate = planDate
	}
	if input.Title != "" {
		existingPlan.Title = input.Title
	}
	if input.PlanType != "" {
		existingPlan.PlanType = input.PlanType
	}
	if input.Notes != nil {
		existingPlan.Notes = sql.NullString{String: *input.Notes, Valid: true}
	} else {
		existingPlan.Notes = sql.NullString{Valid: false}
	}
	if input.Status != nil {
		existingPlan.Status = *input.Status
	}

	if err := s.studyPlanRepo.UpdateStudyPlan(ctx, existingPlan); err != nil {
		return nil, fmt.Errorf("failed to update study plan: %w", err)
	}
	return existingPlan, nil
}

// DeleteStudyPlan deletes a study plan.
func (s *studyPlanService) DeleteStudyPlan(ctx context.Context, id string) error {
	return s.studyPlanRepo.DeleteStudyPlan(ctx, id)
}

// CreateStudySession creates a new study session for a user.
func (s *studyPlanService) CreateStudySession(ctx context.Context, userID string, input *models.StudySessionCreationInput) (*models.StudySession, error) {
	studySession := &models.StudySession{
		UserID:             userID,
		StudyPlanID:        sql.NullString{String: *input.StudyPlanID, Valid: input.StudyPlanID != nil},
		SubjectID:          sql.NullString{String: *input.SubjectID, Valid: input.SubjectID != nil},
		SessionType:        input.SessionType,
		TopicsToCover:      input.TopicsToCover,
		TopicsCovered:      input.TopicsCovered,
		Status:             "planned",
		CompletionPercentage: 0,
	}

	if input.PlannedStartTime != nil {
		plannedStartTime, err := time.Parse(time.RFC3339, *input.PlannedStartTime)
		if err != nil {
			return nil, fmt.Errorf("invalid planned start time format: %w", err)
		}
		studySession.PlannedStartTime = sql.NullTime{Time: plannedStartTime, Valid: true}
	}
	if input.PlannedEndTime != nil {
		plannedEndTime, err := time.Parse(time.RFC3339, *input.PlannedEndTime)
		if err != nil {
			return nil, fmt.Errorf("invalid planned end time format: %w", err)
		}
		studySession.PlannedEndTime = sql.NullTime{Time: plannedEndTime, Valid: true}
	}
	if input.PlannedDurationMinutes != nil {
		studySession.PlannedDurationMinutes = sql.NullInt32{Int32: int32(*input.PlannedDurationMinutes), Valid: true}
	}
	if input.ActualStartTime != nil {
		actualStartTime, err := time.Parse(time.RFC3339, *input.ActualStartTime)
		if err != nil {
			return nil, fmt.Errorf("invalid actual start time format: %w", err)
		}
		studySession.ActualStartTime = sql.NullTime{Time: actualStartTime, Valid: true}
	}
	if input.ActualEndTime != nil {
		actualEndTime, err := time.Parse(time.RFC3339, *input.ActualEndTime)
		if err != nil {
			return nil, fmt.Errorf("invalid actual end time format: %w", err)
		}
		studySession.ActualEndTime = sql.NullTime{Time: actualEndTime, Valid: true}
	}
	if input.ActualDurationMinutes != nil {
		studySession.ActualDurationMinutes = sql.NullInt32{Int32: int32(*input.ActualDurationMinutes), Valid: true}
	}
	if input.Status != nil {
		studySession.Status = *input.Status
	}
	if input.CompletionPercentage != nil {
		studySession.CompletionPercentage = int32(*input.CompletionPercentage)
	}
	if input.ProductivityRating != nil {
		studySession.ProductivityRating = sql.NullInt32{Int32: int32(*input.ProductivityRating), Valid: true}
	}
	if input.Notes != nil {
		studySession.Notes = sql.NullString{String: *input.Notes, Valid: true}
	}
	if input.Blockers != nil {
		studySession.Blockers = sql.NullString{String: *input.Blockers, Valid: true}
	}

	if err := s.sessionRepo.CreateStudySession(ctx, studySession); err != nil {
		return nil, fmt.Errorf("failed to create study session: %w", err)
	}
	return studySession, nil
}

// GetStudySessionByID retrieves a single study session.
func (s *studyPlanService) GetStudySessionByID(ctx context.Context, id string) (*models.StudySession, error) {
	return s.sessionRepo.GetStudySessionByID(ctx, id)
}

// GetStudySessionsByUserID retrieves all study sessions for a user.
func (s *studyPlanService) GetStudySessionsByUserID(ctx context.Context, userID string) ([]models.StudySession, error) {
	return s.sessionRepo.GetStudySessionsByUserID(ctx, userID)
}

// GetStudySessionsByStudyPlanID retrieves all study sessions for a study plan.
func (s *studyPlanService) GetStudySessionsByStudyPlanID(ctx context.Context, studyPlanID string) ([]models.StudySession, error) {
	return s.sessionRepo.GetStudySessionsByStudyPlanID(ctx, studyPlanID)
}

// UpdateStudySession updates an existing study session.
func (s *studyPlanService) UpdateStudySession(ctx context.Context, userID string, id string, input *models.StudySessionCreationInput) (*models.StudySession, error) {
	existingSession, err := s.sessionRepo.GetStudySessionByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("study session not found: %w", err)
	}
	if existingSession.UserID != userID {
		return nil, fmt.Errorf("study session does not belong to user")
	}

	if input.StudyPlanID != nil {
		existingSession.StudyPlanID = sql.NullString{String: *input.StudyPlanID, Valid: true}
	} else {
		existingSession.StudyPlanID = sql.NullString{Valid: false}
	}
	if input.SubjectID != nil {
		existingSession.SubjectID = sql.NullString{String: *input.SubjectID, Valid: true}
	} else {
		existingSession.SubjectID = sql.NullString{Valid: false}
	}
	if input.PlannedStartTime != nil {
		plannedStartTime, err := time.Parse(time.RFC3339, *input.PlannedStartTime)
		if err != nil {
			return nil, fmt.Errorf("invalid planned start time format: %w", err)
		}
		existingSession.PlannedStartTime = sql.NullTime{Time: plannedStartTime, Valid: true}
	} else {
		existingSession.PlannedStartTime = sql.NullTime{Valid: false}
	}
	if input.PlannedEndTime != nil {
		plannedEndTime, err := time.Parse(time.RFC3339, *input.PlannedEndTime)
		if err != nil {
			return nil, fmt.Errorf("invalid planned end time format: %w", err)
		}
		existingSession.PlannedEndTime = sql.NullTime{Time: plannedEndTime, Valid: true}
	} else {
		existingSession.PlannedEndTime = sql.NullTime{Valid: false}
	}
	if input.PlannedDurationMinutes != nil {
		existingSession.PlannedDurationMinutes = sql.NullInt32{Int32: int32(*input.PlannedDurationMinutes), Valid: true}
	} else {
		existingSession.PlannedDurationMinutes = sql.NullInt32{Valid: false}
	}
	if input.ActualStartTime != nil {
		actualStartTime, err := time.Parse(time.RFC3339, *input.ActualStartTime)
		if err != nil {
			return nil, fmt.Errorf("invalid actual start time format: %w", err)
		}
		existingSession.ActualStartTime = sql.NullTime{Time: actualStartTime, Valid: true}
	} else {
		existingSession.ActualStartTime = sql.NullTime{Valid: false}
	}
	if input.ActualEndTime != nil {
		actualEndTime, err := time.Parse(time.RFC3339, *input.ActualEndTime)
		if err != nil {
			return nil, fmt.Errorf("invalid actual end time format: %w", err)
		}
		existingSession.ActualEndTime = sql.NullTime{Time: actualEndTime, Valid: true}
	} else {
		existingSession.ActualEndTime = sql.NullTime{Valid: false}
	}
	if input.ActualDurationMinutes != nil {
		existingSession.ActualDurationMinutes = sql.NullInt32{Int32: int32(*input.ActualDurationMinutes), Valid: true}
	} else {
		existingSession.ActualDurationMinutes = sql.NullInt32{Valid: false}
	}
	if input.SessionType != "" {
		existingSession.SessionType = input.SessionType
	}
	if input.TopicsToCover != nil {
		existingSession.TopicsToCover = input.TopicsToCover
	}
	if input.TopicsCovered != nil {
		existingSession.TopicsCovered = input.TopicsCovered
	}
	if input.Status != nil {
		existingSession.Status = *input.Status
	}
	if input.CompletionPercentage != nil {
		existingSession.CompletionPercentage = int32(*input.CompletionPercentage)
	}
	if input.ProductivityRating != nil {
		existingSession.ProductivityRating = sql.NullInt32{Int32: int32(*input.ProductivityRating), Valid: true}
	} else {
		existingSession.ProductivityRating = sql.NullInt32{Valid: false}
	}
	if input.Notes != nil {
		existingSession.Notes = sql.NullString{String: *input.Notes, Valid: true}
	} else {
		existingSession.Notes = sql.NullString{Valid: false}
	}
	if input.Blockers != nil {
		existingSession.Blockers = sql.NullString{String: *input.Blockers, Valid: true}
	} else {
		existingSession.Blockers = sql.NullString{Valid: false}
	}

	if err := s.sessionRepo.UpdateStudySession(ctx, existingSession); err != nil {
		return nil, fmt.Errorf("failed to update study session: %w", err)
	}
	return existingSession, nil
}

// DeleteStudySession deletes a study session.
func (s *studyPlanService) DeleteStudySession(ctx context.Context, id string) error {
	return s.sessionRepo.DeleteStudySession(ctx, id)
}
