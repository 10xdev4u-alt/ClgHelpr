package services

import (
	"context"
	"fmt"
	"time"

	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/repository"
)

// ExamService defines the interface for exam-related business logic.
type ExamService interface {
	CreateExam(ctx context.Context, userID string, input *models.ExamCreationInput) (*models.Exam, error)
	GetExamByID(ctx context.Context, id string) (*models.Exam, error)
	GetExamsByUserID(ctx context.Context, userID string) ([]models.Exam, error)
	GetUpcomingExamsByUserID(ctx context.Context, userID string) ([]models.Exam, error)
	UpdateExam(ctx context.Context, userID string, id string, input *models.ExamCreationInput) (*models.Exam, error)
	DeleteExam(ctx context.Context, id string) error
	UpdateExamPrepStatus(ctx context.Context, id string, status string) error

	CreateImportantQuestion(ctx context.Context, userID string, input *models.ImportantQuestionCreationInput) (*models.ImportantQuestion, error)
	GetImportantQuestionByID(ctx context.Context, id string) (*models.ImportantQuestion, error)
	GetImportantQuestionsByExamID(ctx context.Context, examID string) ([]models.ImportantQuestion, error)
	GetImportantQuestionsBySubjectID(ctx context.Context, subjectID string) ([]models.ImportantQuestion, error)
	UpdateImportantQuestion(ctx context.Context, userID string, id string, input *models.ImportantQuestionCreationInput) (*models.ImportantQuestion, error)
	DeleteImportantQuestion(ctx context.Context, id string) error
}

// examService implements ExamService.
type examService struct {
	examRepo              repository.ExamRepository
	importantQuestionRepo repository.ImportantQuestionRepository
}

// NewExamService creates a new exam service.
func NewExamService(examRepo repository.ExamRepository, importantQuestionRepo repository.ImportantQuestionRepository) ExamService {
	return &examService{
		examRepo:              examRepo,
		importantQuestionRepo: importantQuestionRepo,
	}
}

// CreateExam creates a new exam for a user.
func (s *examService) CreateExam(ctx context.Context, userID string, input *models.ExamCreationInput) (*models.Exam, error) {
	examDate, err := time.Parse("2006-01-02", input.ExamDate)
	if err != nil {
		return nil, fmt.Errorf("invalid exam date format: %w", err)
	}

	exam := &models.Exam{
		UserID:            userID,
		SubjectID:         sql.NullString{String: *input.SubjectID, Valid: input.SubjectID != nil},
		VenueID:           sql.NullString{String: *input.VenueID, Valid: input.VenueID != nil},
		Title:             input.Title,
		ExamType:          input.ExamType,
		ExamDate:          examDate,
		PrepStatus:        "not_started",
		ReminderEnabled:   true,
		SyllabusUnits:     input.SyllabusUnits,
		SyllabusTopics:    input.SyllabusTopics,
	}

	if input.StartTime != nil {
		startTime, err := time.Parse("15:04:05", *input.StartTime)
		if err != nil {
			return nil, fmt.Errorf("invalid start time format: %w", err)
		}
		exam.StartTime = sql.NullTime{Time: startTime, Valid: true}
	}
	if input.EndTime != nil {
		endTime, err := time.Parse("15:04:05", *input.EndTime)
		if err != nil {
			return nil, fmt.Errorf("invalid end time format: %w", err)
		}
		exam.EndTime = sql.NullTime{Time: endTime, Valid: true}
	}
	if input.DurationMinutes != nil {
		exam.DurationMinutes = sql.NullInt32{Int32: int32(*input.DurationMinutes), Valid: true}
	}
	if input.SyllabusNotes != nil {
		exam.SyllabusNotes = sql.NullString{String: *input.SyllabusNotes, Valid: true}
	}
	if input.MaxMarks != nil {
		exam.MaxMarks = sql.NullFloat64{Float64: *input.MaxMarks, Valid: true}
	}
	if input.ObtainedMarks != nil {
		exam.ObtainedMarks = sql.NullFloat64{Float64: *input.ObtainedMarks, Valid: true}
	}
	if input.Grade != nil {
		exam.Grade = sql.NullString{String: *input.Grade, Valid: true}
	}
	if input.PrepStatus != nil {
		exam.PrepStatus = *input.PrepStatus
	}
	if input.PrepNotes != nil {
		exam.PrepNotes = sql.NullString{String: *input.PrepNotes, Valid: true}
	}
	if input.StudyHoursLogged != nil {
		exam.StudyHoursLogged = sql.NullFloat64{Float64: *input.StudyHoursLogged, Valid: true}
	}
	if input.ReminderEnabled != nil {
		exam.ReminderEnabled = *input.ReminderEnabled
	}

	if err := s.examRepo.CreateExam(ctx, exam); err != nil {
		return nil, fmt.Errorf("failed to create exam: %w", err)
	}
	return exam, nil
}

// GetExamByID retrieves a single exam.
func (s *examService) GetExamByID(ctx context.Context, id string) (*models.Exam, error) {
	return s.examRepo.GetExamByID(ctx, id)
}

// GetExamsByUserID retrieves all exams for a user.
func (s *examService) GetExamsByUserID(ctx context.Context, userID string) ([]models.Exam, error) {
	return s.examRepo.GetExamsByUserID(ctx, userID)
}

// GetUpcomingExamsByUserID retrieves upcoming exams for a user.
func (s *examService) GetUpcomingExamsByUserID(ctx context.Context, userID string) ([]models.Exam, error) {
	return s.examRepo.GetUpcomingExamsByUserID(ctx, userID)
}

// UpdateExam updates an existing exam.
func (s *examService) UpdateExam(ctx context.Context, userID string, id string, input *models.ExamCreationInput) (*models.Exam, error) {
	existingExam, err := s.examRepo.GetExamByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("exam not found: %w", err)
	}
	if existingExam.UserID != userID {
		return nil, fmt.Errorf("exam does not belong to user")
	}

	if input.ExamDate != "" {
		examDate, err := time.Parse("2006-01-02", input.ExamDate)
		if err != nil {
			return nil, fmt.Errorf("invalid exam date format: %w", err)
		}
		existingExam.ExamDate = examDate
	}
	if input.StartTime != nil {
		startTime, err := time.Parse("15:04:05", *input.StartTime)
		if err != nil {
			return nil, fmt.Errorf("invalid start time format: %w", err)
		}
		existingExam.StartTime = sql.NullTime{Time: startTime, Valid: true}
	} else {
		existingExam.StartTime = sql.NullTime{Valid: false}
	}
	if input.EndTime != nil {
		endTime, err := time.Parse("15:04:05", *input.EndTime)
		if err != nil {
			return nil, fmt.Errorf("invalid end time format: %w", err)
		}
		existingExam.EndTime = sql.NullTime{Time: endTime, Valid: true}
	} else {
		existingExam.EndTime = sql.NullTime{Valid: false}
	}
	if input.SubjectID != nil {
		existingExam.SubjectID = sql.NullString{String: *input.SubjectID, Valid: true}
	} else {
		existingExam.SubjectID = sql.NullString{Valid: false}
	}
	if input.VenueID != nil {
		existingExam.VenueID = sql.NullString{String: *input.VenueID, Valid: true}
	} else {
		existingExam.VenueID = sql.NullString{Valid: false}
	}
	if input.Title != "" {
		existingExam.Title = input.Title
	}
	if input.ExamType != "" {
		existingExam.ExamType = input.ExamType
	}
	if input.DurationMinutes != nil {
		existingExam.DurationMinutes = sql.NullInt32{Int32: int32(*input.DurationMinutes), Valid: true}
	} else {
		existingExam.DurationMinutes = sql.NullInt32{Valid: false}
	}
	if input.SyllabusUnits != nil {
		existingExam.SyllabusUnits = input.SyllabusUnits
	}
	if input.SyllabusTopics != nil {
		existingExam.SyllabusTopics = input.SyllabusTopics
	}
	if input.SyllabusNotes != nil {
		existingExam.SyllabusNotes = sql.NullString{String: *input.SyllabusNotes, Valid: true}
	} else {
		existingExam.SyllabusNotes = sql.NullString{Valid: false}
	}
	if input.MaxMarks != nil {
		existingExam.MaxMarks = sql.NullFloat64{Float64: *input.MaxMarks, Valid: true}
	} else {
		existingExam.MaxMarks = sql.NullFloat64{Valid: false}
	}
	if input.ObtainedMarks != nil {
		existingExam.ObtainedMarks = sql.NullFloat64{Float64: *input.ObtainedMarks, Valid: true}
	} else {
		existingExam.ObtainedMarks = sql.NullFloat64{Valid: false}
	}
	if input.Grade != nil {
		existingExam.Grade = sql.NullString{String: *input.Grade, Valid: true}
	} else {
		existingExam.Grade = sql.NullString{Valid: false}
	}
	if input.PrepStatus != nil {
		existingExam.PrepStatus = *input.PrepStatus
	}
	if input.PrepNotes != nil {
		existingExam.PrepNotes = sql.NullString{String: *input.PrepNotes, Valid: true}
	} else {
		existingExam.PrepNotes = sql.NullString{Valid: false}
	}
	if input.StudyHoursLogged != nil {
		existingExam.StudyHoursLogged = sql.NullFloat64{Float64: *input.StudyHoursLogged, Valid: true}
	} else {
		existingExam.StudyHoursLogged = sql.NullFloat64{Valid: false}
	}
	if input.ReminderEnabled != nil {
		existingExam.ReminderEnabled = *input.ReminderEnabled
	}

	if err := s.examRepo.UpdateExam(ctx, existingExam); err != nil {
		return nil, fmt.Errorf("failed to update exam: %w", err)
	}
	return existingExam, nil
}

// UpdateExamPrepStatus updates the preparation status of an exam.
func (s *examService) UpdateExamPrepStatus(ctx context.Context, id string, status string) error {
	return s.examRepo.UpdateExamPrepStatus(ctx, id, status)
}

// DeleteExam deletes an exam.
func (s *examService) DeleteExam(ctx context.Context, id string) error {
	return s.examRepo.DeleteExam(ctx, id)
}

// CreateImportantQuestion creates a new important question for a user.
func (s *examService) CreateImportantQuestion(ctx context.Context, userID string, input *models.ImportantQuestionCreationInput) (*models.ImportantQuestion, error) {
	question := &models.ImportantQuestion{
		UserID:       userID,
		SubjectID:    sql.NullString{String: *input.SubjectID, Valid: input.SubjectID != nil},
		ExamID:       sql.NullString{String: *input.ExamID, Valid: input.ExamID != nil},
		QuestionText: input.QuestionText,
		IsPracticed:  false,
		Tags:         input.Tags,
	}

	if input.AnswerText != nil {
		question.AnswerText = sql.NullString{String: *input.AnswerText, Valid: true}
	}
	if input.Source != nil {
		question.Source = sql.NullString{String: *input.Source, Valid: true}
	}
	if input.Unit != nil {
		question.Unit = sql.NullString{String: *input.Unit, Valid: true}
	}
	if input.Topic != nil {
		question.Topic = sql.NullString{String: *input.Topic, Valid: true}
	}
	if input.Marks != nil {
		question.Marks = sql.NullInt32{Int32: int32(*input.Marks), Valid: true}
	}
	if input.FrequencyCount != nil {
		question.FrequencyCount = sql.NullInt32{Int32: int32(*input.FrequencyCount), Valid: true}
	}
	if input.IsPracticed != nil {
		question.IsPracticed = *input.IsPracticed
	}
	if input.LastPracticedAt != nil {
		lastPracticedAt, err := time.Parse(time.RFC3339, *input.LastPracticedAt)
		if err != nil {
			return nil, fmt.Errorf("invalid last practiced at format: %w", err)
		}
		question.LastPracticedAt = sql.NullTime{Time: lastPracticedAt, Valid: true}
	}
	if input.ConfidenceLevel != nil {
		question.ConfidenceLevel = sql.NullInt32{Int32: int32(*input.ConfidenceLevel), Valid: true}
	}

	if err := s.importantQuestionRepo.CreateImportantQuestion(ctx, question); err != nil {
		return nil, fmt.Errorf("failed to create important question: %w", err)
	}
	return question, nil
}

// GetImportantQuestionByID retrieves a single important question.
func (s *examService) GetImportantQuestionByID(ctx context.Context, id string) (*models.ImportantQuestion, error) {
	return s.importantQuestionRepo.GetImportantQuestionByID(ctx, id)
}

// GetImportantQuestionsByExamID retrieves all important questions for an exam.
func (s *examService) GetImportantQuestionsByExamID(ctx context.Context, examID string) ([]models.ImportantQuestion, error) {
	return s.importantQuestionRepo.GetImportantQuestionsByExamID(ctx, examID)
}

// GetImportantQuestionsBySubjectID retrieves all important questions for a subject.
func (s *examService) GetImportantQuestionsBySubjectID(ctx context.Context, subjectID string) ([]models.ImportantQuestion, error) {
	return s.importantQuestionRepo.GetImportantQuestionsBySubjectID(ctx, subjectID)
}

// UpdateImportantQuestion updates an existing important question.
func (s *examService) UpdateImportantQuestion(ctx context.Context, userID string, id string, input *models.ImportantQuestionCreationInput) (*models.ImportantQuestion, error) {
	existingQuestion, err := s.importantQuestionRepo.GetImportantQuestionByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("important question not found: %w", err)
	}
	if existingQuestion.UserID != userID {
		return nil, fmt.Errorf("important question does not belong to user")
	}

	if input.SubjectID != nil {
		existingQuestion.SubjectID = sql.NullString{String: *input.SubjectID, Valid: true}
	} else {
		existingQuestion.SubjectID = sql.NullString{Valid: false}
	}
	if input.ExamID != nil {
		existingQuestion.ExamID = sql.NullString{String: *input.ExamID, Valid: true}
	} else {
		existingQuestion.ExamID = sql.NullString{Valid: false}
	}
	if input.QuestionText != "" {
		existingQuestion.QuestionText = input.QuestionText
	}
	if input.AnswerText != nil {
		existingQuestion.AnswerText = sql.NullString{String: *input.AnswerText, Valid: true}
	} else {
		existingQuestion.AnswerText = sql.NullString{Valid: false}
	}
	if input.Source != nil {
		existingQuestion.Source = sql.NullString{String: *input.Source, Valid: true}
	} else {
		existingQuestion.Source = sql.NullString{Valid: false}
	}
	if input.Unit != nil {
		existingQuestion.Unit = sql.NullString{String: *input.Unit, Valid: true}
	} else {
		existingQuestion.Unit = sql.NullString{Valid: false}
	}
	if input.Topic != nil {
		existingQuestion.Topic = sql.NullString{String: *input.Topic, Valid: true}
	} else {
		existingQuestion.Topic = sql.NullString{Valid: false}
	}
	if input.Marks != nil {
		existingQuestion.Marks = sql.NullInt32{Int32: int32(*input.Marks), Valid: true}
	} else {
		existingQuestion.Marks = sql.NullInt32{Valid: false}
	}
	if input.FrequencyCount != nil {
		existingQuestion.FrequencyCount = sql.NullInt32{Int32: int32(*input.FrequencyCount), Valid: true}
	} else {
		existingQuestion.FrequencyCount = sql.NullInt32{Valid: false}
	}
	if input.IsPracticed != nil {
		existingQuestion.IsPracticed = *input.IsPracticed
	}
	if input.LastPracticedAt != nil {
		lastPracticedAt, err := time.Parse(time.RFC3339, *input.LastPracticedAt)
		if err != nil {
			return nil, fmt.Errorf("invalid last practiced at format: %w", err)
		}
		existingQuestion.LastPracticedAt = sql.NullTime{Time: lastPracticedAt, Valid: true}
	} else {
		existingQuestion.LastPracticedAt = sql.NullTime{Valid: false}
	}
	if input.ConfidenceLevel != nil {
		existingQuestion.ConfidenceLevel = sql.NullInt32{Int32: int32(*input.ConfidenceLevel), Valid: true}
	} else {
		existingQuestion.ConfidenceLevel = sql.NullInt32{Valid: false}
	}
	if input.Tags != nil {
		existingQuestion.Tags = input.Tags
	}

	if err := s.importantQuestionRepo.UpdateImportantQuestion(ctx, existingQuestion); err != nil {
		return nil, fmt.Errorf("failed to update important question: %w", err)
	}
	return existingQuestion, nil
}

// DeleteImportantQuestion deletes an important question.
func (s *examService) DeleteImportantQuestion(ctx context.Context, id string) error {
	return s.importantQuestionRepo.DeleteImportantQuestion(ctx, id)
}
