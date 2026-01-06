package services

import (
	"context"
	"fmt"
	"time"

	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/repository"
)

// LabRecordService defines the interface for lab record-related business logic.
type LabRecordService interface {
	CreateLabRecord(ctx context.Context, userID string, input *models.LabRecordCreationInput) (*models.LabRecord, error)
	GetLabRecordByID(ctx context.Context, id string) (*models.LabRecord, error)
	GetLabRecordsByUserID(ctx context.Context, userID string) ([]models.LabRecord, error)
	GetLabRecordsByUserIDAndSubjectID(ctx context.Context, userID, subjectID string) ([]models.LabRecord, error)
	UpdateLabRecord(ctx context.Context, userID string, id string, input *models.LabRecordCreationInput) (*models.LabRecord, error)
	DeleteLabRecord(ctx context.Context, id string) error
	UpdateLabRecordStatus(ctx context.Context, id string, status string) error
}

// labRecordService implements LabRecordService.
type labRecordService struct {
	labRecordRepo repository.LabRecordRepository
}

// NewLabRecordService creates a new lab record service.
func NewLabRecordService(labRecordRepo repository.LabRecordRepository) LabRecordService {
	return &labRecordService{labRecordRepo: labRecordRepo}
}

// CreateLabRecord creates a new lab record for a user.
func (s *labRecordService) CreateLabRecord(ctx context.Context, userID string, input *models.LabRecordCreationInput) (*models.LabRecord, error) {
	labRecord := &models.LabRecord{
		UserID:            userID,
		SubjectID:         sql.NullString{String: *input.SubjectID, Valid: input.SubjectID != nil},
		ExperimentNumber:  input.ExperimentNumber,
		Title:             input.Title,
		PrintRequired:     true, // Default
		Status:            "pending", // Default
		VivaQuestions:     input.VivaQuestions,
	}

	if input.LabDate != nil {
		labDate, err := time.Parse("2006-01-02", *input.LabDate)
		if err != nil {
			return nil, fmt.Errorf("invalid lab date format: %w", err)
		}
		labRecord.LabDate = sql.NullTime{Time: labDate, Valid: true}
	}
	if input.RecordWrittenDate != nil {
		recordWrittenDate, err := time.Parse("2006-01-02", *input.RecordWrittenDate)
		if err != nil {
			return nil, fmt.Errorf("invalid record written date format: %w", err)
		}
		labRecord.RecordWrittenDate = sql.NullTime{Time: recordWrittenDate, Valid: true}
	}
	if input.SubmittedDate != nil {
		submittedDate, err := time.Parse("2006-01-02", *input.SubmittedDate)
		if err != nil {
			return nil, fmt.Errorf("invalid submitted date format: %w", err)
		}
		labRecord.SubmittedDate = sql.NullTime{Time: submittedDate, Valid: true}
	}
	if input.Status != nil {
		labRecord.Status = *input.Status
	}
	if input.Aim != nil {
		labRecord.Aim = sql.NullString{String: *input.Aim, Valid: true}
	}
	if input.Algorithm != nil {
		labRecord.Algorithm = sql.NullString{String: *input.Algorithm, Valid: true}
	}
	if input.Code != nil {
		labRecord.Code = sql.NullString{String: *input.Code, Valid: true}
	}
	if input.Output != nil {
		labRecord.Output = sql.NullString{String: *input.Output, Valid: true}
	}
	if input.Observations != nil {
		labRecord.Observations = sql.NullString{String: *input.Observations, Valid: true}
	}
	if input.Result != nil {
		labRecord.Result = sql.NullString{String: *input.Result, Valid: true}
	}
	if input.PrintRequired != nil {
		labRecord.PrintRequired = *input.PrintRequired
	}
	if input.PagesToPrint != nil {
		labRecord.PagesToPrint = sql.NullInt32{Int32: int32(*input.PagesToPrint), Valid: true}
	}
	if input.PrintedAt != nil {
		printedAt, err := time.Parse(time.RFC3339, *input.PrintedAt)
		if err != nil {
			return nil, fmt.Errorf("invalid printed at format: %w", err)
		}
		labRecord.PrintedAt = sql.NullTime{Time: printedAt, Valid: true}
	}
	if input.Marks != nil {
		labRecord.Marks = sql.NullFloat64{Float64: *input.Marks, Valid: true}
	}
	if input.StaffRemarks != nil {
		labRecord.StaffRemarks = sql.NullString{String: *input.StaffRemarks, Valid: true}
	}

	if err := s.labRecordRepo.CreateLabRecord(ctx, labRecord); err != nil {
		return nil, fmt.Errorf("failed to create lab record: %w", err)
	}
	return labRecord, nil
}

// GetLabRecordByID retrieves a single lab record.
func (s *labRecordService) GetLabRecordByID(ctx context.Context, id string) (*models.LabRecord, error) {
	return s.labRecordRepo.GetLabRecordByID(ctx, id)
}

// GetLabRecordsByUserID retrieves all lab records for a user.
func (s *labRecordService) GetLabRecordsByUserID(ctx context.Context, userID string) ([]models.LabRecord, error) {
	return s.labRecordRepo.GetLabRecordsByUserID(ctx, userID)
}

// GetLabRecordsByUserIDAndSubjectID retrieves all lab records for a user and subject.
func (s *labRecordService) GetLabRecordsByUserIDAndSubjectID(ctx context.Context, userID, subjectID string) ([]models.LabRecord, error) {
	return s.labRecordRepo.GetLabRecordsByUserIDAndSubjectID(ctx, userID, subjectID)
}

// UpdateLabRecord updates an existing lab record.
func (s *labRecordService) UpdateLabRecord(ctx context.Context, userID string, id string, input *models.LabRecordCreationInput) (*models.LabRecord, error) {
	existingRecord, err := s.labRecordRepo.GetLabRecordByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("lab record not found: %w", err)
	}
	if existingRecord.UserID != userID {
		return nil, fmt.Errorf("lab record does not belong to user")
	}

	if input.LabDate != nil {
		labDate, err := time.Parse("2006-01-02", *input.LabDate)
		if err != nil {
			return nil, fmt.Errorf("invalid lab date format: %w", err)
		}
		existingRecord.LabDate = sql.NullTime{Time: labDate, Valid: true}
	} else {
		existingRecord.LabDate = sql.NullTime{Valid: false}
	}
	if input.RecordWrittenDate != nil {
		recordWrittenDate, err := time.Parse("2006-01-02", *input.RecordWrittenDate)
		if err != nil {
			return nil, fmt.Errorf("invalid record written date format: %w", err)
		}
		existingRecord.RecordWrittenDate = sql.NullTime{Time: recordWrittenDate, Valid: true}
	} else {
		existingRecord.RecordWrittenDate = sql.NullTime{Valid: false}
	}
	if input.SubmittedDate != nil {
		submittedDate, err := time.Parse("2006-01-02", *input.SubmittedDate)
		if err != nil {
			return nil, fmt.Errorf("invalid submitted date format: %w", err)
		}
		existingRecord.SubmittedDate = sql.NullTime{Time: submittedDate, Valid: true}
	} else {
		existingRecord.SubmittedDate = sql.NullTime{Valid: false}
	}
	if input.SubjectID != nil {
		existingRecord.SubjectID = sql.NullString{String: *input.SubjectID, Valid: true}
	} else {
		existingRecord.SubjectID = sql.NullString{Valid: false}
	}
	if input.ExperimentNumber != 0 { // Cannot be 0, so check directly
		existingRecord.ExperimentNumber = input.ExperimentNumber
	}
	if input.Title != "" {
		existingRecord.Title = input.Title
	}
	if input.Status != nil {
		existingRecord.Status = *input.Status
	}
	if input.Aim != nil {
		existingRecord.Aim = sql.NullString{String: *input.Aim, Valid: true}
	} else {
		existingRecord.Aim = sql.NullString{Valid: false}
	}
	if input.Algorithm != nil {
		existingRecord.Algorithm = sql.NullString{String: *input.Algorithm, Valid: true}
	} else {
		existingRecord.Algorithm = sql.NullString{Valid: false}
	}
	if input.Code != nil {
		existingRecord.Code = sql.NullString{String: *input.Code, Valid: true}
	} else {
		existingRecord.Code = sql.NullString{Valid: false}
	}
	if input.Output != nil {
		existingRecord.Output = sql.NullString{String: *input.Output, Valid: true}
	} else {
		existingRecord.Output = sql.NullString{Valid: false}
	}
	if input.Observations != nil {
		existingRecord.Observations = sql.NullString{String: *input.Observations, Valid: true}
	} else {
		existingRecord.Observations = sql.NullString{Valid: false}
	}
	if input.Result != nil {
		existingRecord.Result = sql.NullString{String: *input.Result, Valid: true}
	} else {
		existingRecord.Result = sql.NullString{Valid: false}
	}
	if input.VivaQuestions != nil {
		existingRecord.VivaQuestions = input.VivaQuestions
	}
	if input.PrintRequired != nil {
		existingRecord.PrintRequired = *input.PrintRequired
	}
	if input.PagesToPrint != nil {
		existingRecord.PagesToPrint = sql.NullInt32{Int32: int32(*input.PagesToPrint), Valid: true}
	} else {
		existingRecord.PagesToPrint = sql.NullInt32{Valid: false}
	}
	if input.PrintedAt != nil {
		printedAt, err := time.Parse(time.RFC3339, *input.PrintedAt)
		if err != nil {
			return nil, fmt.Errorf("invalid printed at format: %w", err)
		}
		existingRecord.PrintedAt = sql.NullTime{Time: printedAt, Valid: true}
	} else {
		existingRecord.PrintedAt = sql.NullTime{Valid: false}
	}
	if input.Marks != nil {
		existingRecord.Marks = sql.NullFloat64{Float64: *input.Marks, Valid: true}
	} else {
		existingRecord.Marks = sql.NullFloat64{Valid: false}
	}
	if input.StaffRemarks != nil {
		existingRecord.StaffRemarks = sql.NullString{String: *input.StaffRemarks, Valid: true}
	} else {
		existingRecord.StaffRemarks = sql.NullString{Valid: false}
	}

	if err := s.labRecordRepo.UpdateLabRecord(ctx, existingRecord); err != nil {
		return nil, fmt.Errorf("failed to update lab record: %w", err)
	}
	return existingRecord, nil
}

// DeleteLabRecord deletes a lab record.
func (s *labRecordService) DeleteLabRecord(ctx context.Context, id string) error {
	return s.labRecordRepo.DeleteLabRecord(ctx, id)
}

// UpdateLabRecordStatus updates the status of a lab record.
func (s *labRecordService) UpdateLabRecordStatus(ctx context.Context, id string, status string) error {
	return s.labRecordRepo.UpdateLabRecordStatus(ctx, id, status)
}
