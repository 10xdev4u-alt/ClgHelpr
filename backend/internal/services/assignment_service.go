package services

import (
	"context"
	"fmt"
	"time"

	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/repository"
)

// AssignmentService defines the interface for assignment-related business logic.
type AssignmentService interface {
	CreateAssignment(ctx context.Context, userID string, input *models.AssignmentCreationInput) (*models.Assignment, error)
	GetAssignmentByID(ctx context.Context, id string) (*models.Assignment, error)
	GetAssignmentsByUserID(ctx context.Context, userID string) ([]models.Assignment, error)
	GetPendingAssignmentsByUserID(ctx context.Context, userID string) ([]models.Assignment, error)
	GetOverdueAssignmentsByUserID(ctx context.Context, userID string) ([]models.Assignment, error)
	UpdateAssignment(ctx context.Context, userID string, id string, input *models.AssignmentCreationInput) (*models.Assignment, error)
	UpdateAssignmentStatus(ctx context.Context, id string, status string) error
	DeleteAssignment(ctx context.Context, id string) error
}

// assignmentService implements AssignmentService.
type assignmentService struct {
	assignmentRepo repository.AssignmentRepository
}

// NewAssignmentService creates a new assignment service.
func NewAssignmentService(assignmentRepo repository.AssignmentRepository) AssignmentService {
	return &assignmentService{assignmentRepo: assignmentRepo}
}

// CreateAssignment creates a new assignment for a user.
func (s *assignmentService) CreateAssignment(ctx context.Context, userID string, input *models.AssignmentCreationInput) (*models.Assignment, error) {
	dueDate, err := time.Parse(time.RFC3339, input.DueDate)
	if err != nil {
		return nil, fmt.Errorf("invalid due date format: %w", err)
	}

	assignedDate := time.Now()
	if input.AssignedDate != nil {
		assignedDate, err = time.Parse("2006-01-02", *input.AssignedDate)
		if err != nil {
			return nil, fmt.Errorf("invalid assigned date format: %w", err)
		}
	}

	assignment := &models.Assignment{
		UserID:         userID,
		SubjectID:      sql.NullString{String: *input.SubjectID, Valid: input.SubjectID != nil},
		StaffID:        sql.NullString{String: *input.StaffID, Valid: input.StaffID != nil},
		Title:          input.Title,
		Description:    sql.NullString{String: *input.Description, Valid: input.Description != nil},
		Instructions:   sql.NullString{String: *input.Instructions, Valid: input.Instructions != nil},
		AssignmentType: input.AssignmentType,
		AssignedDate:   sql.NullTime{Time: assignedDate, Valid: true},
		DueDate:        dueDate,
		Status:         "pending",
		Priority:       "medium",
		ReminderEnabled: true,
		Tags:           input.Tags,
		IsRecurring:    false,
	}

	if input.Status != nil {
		assignment.Status = *input.Status
	}
	if input.Priority != nil {
		assignment.Priority = *input.Priority
	}
	if input.ReminderEnabled != nil {
		assignment.ReminderEnabled = *input.ReminderEnabled
	}
	if input.ReminderBeforeHours != nil {
		assignment.ReminderBeforeHours = sql.NullInt32{Int32: int32(*input.ReminderBeforeHours), Valid: true}
	}
	if input.MaxMarks != nil {
		assignment.MaxMarks = sql.NullFloat64{Float64: *input.MaxMarks, Valid: true}
	}
	if input.ObtainedMarks != nil {
		assignment.ObtainedMarks = sql.NullFloat64{Float64: *input.ObtainedMarks, Valid: true}
	}
	if input.EstimatedHours != nil {
		assignment.EstimatedHours = sql.NullFloat64{Float64: *input.EstimatedHours, Valid: true}
	}
	if input.ActualHours != nil {
		assignment.ActualHours = sql.NullFloat64{Float64: *input.ActualHours, Valid: true}
	}
	if input.IsRecurring != nil {
		assignment.IsRecurring = *input.IsRecurring
	}
	if input.RecurrencePattern != nil {
		assignment.RecurrencePattern = sql.NullString{String: *input.RecurrencePattern, Valid: true}
	}

	if err := s.assignmentRepo.CreateAssignment(ctx, assignment); err != nil {
		return nil, fmt.Errorf("failed to create assignment: %w", err)
	}
	return assignment, nil
}

// GetAssignmentByID retrieves a single assignment.
func (s *assignmentService) GetAssignmentByID(ctx context.Context, id string) (*models.Assignment, error) {
	return s.assignmentRepo.GetAssignmentByID(ctx, id)
}

// GetAssignmentsByUserID retrieves all assignments for a user.
func (s *assignmentService) GetAssignmentsByUserID(ctx context.Context, userID string) ([]models.Assignment, error) {
	return s.assignmentRepo.GetAssignmentsByUserID(ctx, userID)
}

// GetPendingAssignmentsByUserID retrieves pending assignments for a user.
func (s *assignmentService) GetPendingAssignmentsByUserID(ctx context.Context, userID string) ([]models.Assignment, error) {
	return s.assignmentRepo.GetPendingAssignmentsByUserID(ctx, userID)
}

// GetOverdueAssignmentsByUserID retrieves overdue assignments for a user.
func (s *assignmentService) GetOverdueAssignmentsByUserID(ctx context.Context, userID string) ([]models.Assignment, error) {
	return s.assignmentRepo.GetOverdueAssignmentsByUserID(ctx, userID)
}

// UpdateAssignment updates an existing assignment.
func (s *assignmentService) UpdateAssignment(ctx context.Context, userID string, id string, input *models.AssignmentCreationInput) (*models.Assignment, error) {
	existingAssignment, err := s.assignmentRepo.GetAssignmentByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("assignment not found: %w", err)
	}
	if existingAssignment.UserID != userID {
		return nil, fmt.Errorf("assignment does not belong to user")
	}

	if input.DueDate != "" {
		dueDate, err := time.Parse(time.RFC3339, input.DueDate)
		if err != nil {
			return nil, fmt.Errorf("invalid due date format: %w", err)
		}
		existingAssignment.DueDate = dueDate
	}

	if input.AssignedDate != nil {
		assignedDate, err := time.Parse("2006-01-02", *input.AssignedDate)
		if err != nil {
			return nil, fmt.Errorf("invalid assigned date format: %w", err)
		}
		existingAssignment.AssignedDate = sql.NullTime{Time: assignedDate, Valid: true}
	} else {
		existingAssignment.AssignedDate = sql.NullTime{Valid: false}
	}

	if input.SubjectID != nil {
		existingAssignment.SubjectID = sql.NullString{String: *input.SubjectID, Valid: true}
	} else {
		existingAssignment.SubjectID = sql.NullString{Valid: false}
	}
	if input.StaffID != nil {
		existingAssignment.StaffID = sql.NullString{String: *input.StaffID, Valid: true}
	} else {
		existingAssignment.StaffID = sql.NullString{Valid: false}
	}
	if input.Title != "" {
		existingAssignment.Title = input.Title
	}
	if input.Description != nil {
		existingAssignment.Description = sql.NullString{String: *input.Description, Valid: true}
	} else {
		existingAssignment.Description = sql.NullString{Valid: false}
	}
	if input.Instructions != nil {
		existingAssignment.Instructions = sql.NullString{String: *input.Instructions, Valid: true}
	} else {
		existingAssignment.Instructions = sql.NullString{Valid: false}
	}
	if input.AssignmentType != "" {
		existingAssignment.AssignmentType = input.AssignmentType
	}
	if input.Status != nil {
		existingAssignment.Status = *input.Status
	}
	if input.MaxMarks != nil {
		existingAssignment.MaxMarks = sql.NullFloat64{Float64: *input.MaxMarks, Valid: true}
	} else {
		existingAssignment.MaxMarks = sql.NullFloat64{Valid: false}
	}
	if input.ObtainedMarks != nil {
		existingAssignment.ObtainedMarks = sql.NullFloat64{Float64: *input.ObtainedMarks, Valid: true}
	} else {
		existingAssignment.ObtainedMarks = sql.NullFloat64{Valid: false}
	}
	if input.Feedback != nil {
		existingAssignment.Feedback = sql.NullString{String: *input.Feedback, Valid: true}
	} else {
		existingAssignment.Feedback = sql.NullString{Valid: false}
	}
	if input.Priority != nil {
		existingAssignment.Priority = *input.Priority
	}
	if input.EstimatedHours != nil {
		existingAssignment.EstimatedHours = sql.NullFloat64{Float64: *input.EstimatedHours, Valid: true}
	} else {
		existingAssignment.EstimatedHours = sql.NullFloat64{Valid: false}
	}
	if input.ActualHours != nil {
		existingAssignment.ActualHours = sql.NullFloat64{Float64: *input.ActualHours, Valid: true}
	} else {
		existingAssignment.ActualHours = sql.NullFloat64{Valid: false}
	}
	if input.ReminderEnabled != nil {
		existingAssignment.ReminderEnabled = *input.ReminderEnabled
	}
	if input.ReminderBeforeHours != nil {
		existingAssignment.ReminderBeforeHours = sql.NullInt32{Int32: int32(*input.ReminderBeforeHours), Valid: true}
	} else {
		existingAssignment.ReminderBeforeHours = sql.NullInt32{Valid: false}
	}
	if input.Tags != nil {
		existingAssignment.Tags = input.Tags
	}
	if input.IsRecurring != nil {
		existingAssignment.IsRecurring = *input.IsRecurring
	}
	if input.RecurrencePattern != nil {
		existingAssignment.RecurrencePattern = sql.NullString{String: *input.RecurrencePattern, Valid: true}
	} else {
		existingAssignment.RecurrencePattern = sql.NullString{Valid: false}
	}

	if err := s.assignmentRepo.UpdateAssignment(ctx, existingAssignment); err != nil {
		return nil, fmt.Errorf("failed to update assignment: %w", err)
	}
	return existingAssignment, nil
}

// UpdateAssignmentStatus updates the status of an assignment.
func (s *assignmentService) UpdateAssignmentStatus(ctx context.Context, id string, status string) error {
	return s.assignmentRepo.UpdateAssignmentStatus(ctx, id, status)
}

// DeleteAssignment deletes an assignment.
func (s *assignmentService) DeleteAssignment(ctx context.Context, id string) error {
	return s.assignmentRepo.DeleteAssignment(ctx, id)
}
