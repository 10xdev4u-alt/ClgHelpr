package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
)

// --- Exam Repository ---

// ExamRepository defines the interface for exam data operations.
type ExamRepository interface {
	CreateExam(ctx context.Context, exam *models.Exam) error
	GetExamByID(ctx context.Context, id string) (*models.Exam, error)
	GetExamsByUserID(ctx context.Context, userID string) ([]models.Exam, error)
	GetUpcomingExamsByUserID(ctx context.Context, userID string) ([]models.Exam, error)
	UpdateExam(ctx context.Context, exam *models.Exam) error
	DeleteExam(ctx context.Context, id string) error
	UpdateExamPrepStatus(ctx context.Context, id string, status string) error
}

// PGExamRepository implements ExamRepository for PostgreSQL.
type PGExamRepository struct {
	db *pgxpool.Pool
}

// NewPGExamRepository creates a new PostgreSQL exam repository.
func NewPGExamRepository(db *pgxpool.Pool) *PGExamRepository {
	return &PGExamRepository{db: db}
}

// CreateExam inserts a new exam into the database.
func (r *PGExamRepository) CreateExam(ctx context.Context, exam *models.Exam) error {
	query := `
		INSERT INTO exams (
			id, user_id, subject_id, title, exam_type, exam_date, start_time, end_time,
			duration_minutes, venue_id, syllabus_units, syllabus_topics, syllabus_notes,
			max_marks, obtained_marks, grade, prep_status, prep_notes, study_hours_logged,
			reminder_enabled, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16,
			$17, $18, $19, $20, $21, $22
		) RETURNING id, created_at, updated_at
	`
	exam.ID = models.NewUUID()
	exam.CreatedAt = time.Now()
	exam.UpdatedAt = time.Now()

	_, err := r.db.Exec(ctx, query,
		exam.ID, exam.UserID, exam.SubjectID, exam.Title, exam.ExamType, exam.ExamDate, exam.StartTime, exam.EndTime,
		exam.DurationMinutes, exam.VenueID, exam.SyllabusUnits, exam.SyllabusTopics, exam.SyllabusNotes,
		exam.MaxMarks, exam.ObtainedMarks, exam.Grade, exam.PrepStatus, exam.PrepNotes, exam.StudyHoursLogged,
		exam.ReminderEnabled, exam.CreatedAt, exam.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create exam: %w", err)
	}
	return nil
}

// GetExamByID retrieves an exam by its ID.
func (r *PGExamRepository) GetExamByID(ctx context.Context, id string) (*models.Exam, error) {
	exam := &models.Exam{}
	query := `
		SELECT
			id, user_id, subject_id, title, exam_type, exam_date, start_time, end_time,
			duration_minutes, venue_id, syllabus_units, syllabus_topics, syllabus_notes,
			max_marks, obtained_marks, grade, prep_status, prep_notes, study_hours_logged,
			reminder_enabled, created_at, updated_at
		FROM exams
		WHERE id = $1
	`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&exam.ID, &exam.UserID, &exam.SubjectID, &exam.Title, &exam.ExamType, &exam.ExamDate, &exam.StartTime, &exam.EndTime,
		&exam.DurationMinutes, &exam.VenueID, &exam.SyllabusUnits, &exam.SyllabusTopics, &exam.SyllabusNotes,
		&exam.MaxMarks, &exam.ObtainedMarks, &exam.Grade, &exam.PrepStatus, &exam.PrepNotes, &exam.StudyHoursLogged,
		&exam.ReminderEnabled, &exam.CreatedAt, &exam.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get exam by ID: %w", err)
	}
	return exam, nil
}

// GetExamsByUserID retrieves all exams for a given user.
func (r *PGExamRepository) GetExamsByUserID(ctx context.Context, userID string) ([]models.Exam, error) {
	var exams []models.Exam
	query := `
		SELECT
			id, user_id, subject_id, title, exam_type, exam_date, start_time, end_time,
			duration_minutes, venue_id, syllabus_units, syllabus_topics, syllabus_notes,
			max_marks, obtained_marks, grade, prep_status, prep_notes, study_hours_logged,
			reminder_enabled, created_at, updated_at
		FROM exams
		WHERE user_id = $1
		ORDER BY exam_date ASC, start_time ASC
	`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get exams by user ID: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		exam := models.Exam{}
		err := rows.Scan(
			&exam.ID, &exam.UserID, &exam.SubjectID, &exam.Title, &exam.ExamType, &exam.ExamDate, &exam.StartTime, &exam.EndTime,
			&exam.DurationMinutes, &exam.VenueID, &exam.SyllabusUnits, &exam.SyllabusTopics, &exam.SyllabusNotes,
			&exam.MaxMarks, &exam.ObtainedMarks, &exam.Grade, &exam.PrepStatus, &exam.PrepNotes, &exam.StudyHoursLogged,
			&exam.ReminderEnabled, &exam.CreatedAt, &exam.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan exam row: %w", err)
		}
		exams = append(exams, exam)
	}
	return exams, nil
}

// GetUpcomingExamsByUserID retrieves upcoming exams for a given user.
func (r *PGExamRepository) GetUpcomingExamsByUserID(ctx context.Context, userID string) ([]models.Exam, error) {
	var exams []models.Exam
	query := `
		SELECT
			id, user_id, subject_id, title, exam_type, exam_date, start_time, end_time,
			duration_minutes, venue_id, syllabus_units, syllabus_topics, syllabus_notes,
			max_marks, obtained_marks, grade, prep_status, prep_notes, study_hours_logged,
			reminder_enabled, created_at, updated_at
		FROM exams
		WHERE user_id = $1 AND exam_date >= CURRENT_DATE
		ORDER BY exam_date ASC, start_time ASC
	`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get upcoming exams: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		exam := models.Exam{}
		err := rows.Scan(
			&exam.ID, &exam.UserID, &exam.SubjectID, &exam.Title, &exam.ExamType, &exam.ExamDate, &exam.StartTime, &exam.EndTime,
			&exam.DurationMinutes, &exam.VenueID, &exam.SyllabusUnits, &exam.SyllabusTopics, &exam.SyllabusNotes,
			&exam.MaxMarks, &exam.ObtainedMarks, &exam.Grade, &exam.PrepStatus, &exam.PrepNotes, &exam.StudyHoursLogged,
			&exam.ReminderEnabled, &exam.CreatedAt, &exam.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan upcoming exam row: %w", err)
		}
		exams = append(exams, exam)
	}
	return exams, nil
}

// UpdateExam updates an existing exam in the database.
func (r *PGExamRepository) UpdateExam(ctx context.Context, exam *models.Exam) error {
	query := `
		UPDATE exams SET
			subject_id = $1, title = $2, exam_type = $3, exam_date = $4, start_time = $5, end_time = $6,
			duration_minutes = $7, venue_id = $8, syllabus_units = $9, syllabus_topics = $10, syllabus_notes = $11,
			max_marks = $12, obtained_marks = $13, grade = $14, prep_status = $15, prep_notes = $16,
			study_hours_logged = $17, reminder_enabled = $18, updated_at = $19
		WHERE id = $20 AND user_id = $21
	`
	exam.UpdatedAt = time.Now()

	cmdTag, err := r.db.Exec(ctx, query,
		exam.SubjectID, exam.Title, exam.ExamType, exam.ExamDate, exam.StartTime, exam.EndTime,
		exam.DurationMinutes, exam.VenueID, exam.SyllabusUnits, exam.SyllabusTopics, exam.SyllabusNotes,
		exam.MaxMarks, exam.ObtainedMarks, exam.Grade, exam.PrepStatus, exam.PrepNotes,
		exam.StudyHoursLogged, exam.ReminderEnabled, exam.UpdatedAt,
		exam.ID, exam.UserID,
	)
	if err != nil {
		return fmt.Errorf("failed to update exam: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("exam with ID %s not found or not owned by user", exam.ID)
	}
	return nil
}

// UpdateExamPrepStatus updates the preparation status of an exam.
func (r *PGExamRepository) UpdateExamPrepStatus(ctx context.Context, id string, status string) error {
	query := `
		UPDATE exams SET
			prep_status = $1, updated_at = $2
		WHERE id = $3
	`
	_, err := r.db.Exec(ctx, query, status, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to update exam prep status: %w", err)
	}
	return nil
}

// DeleteExam deletes an exam from the database.
func (r *PGExamRepository) DeleteExam(ctx context.Context, id string) error {
	query := `DELETE FROM exams WHERE id = $1`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete exam: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("exam with ID %s not found", id)
	}
	return nil
}

// --- ImportantQuestion Repository ---

// ImportantQuestionRepository defines the interface for important question data operations.
type ImportantQuestionRepository interface {
	CreateImportantQuestion(ctx context.Context, question *models.ImportantQuestion) error
	GetImportantQuestionByID(ctx context.Context, id string) (*models.ImportantQuestion, error)
	GetImportantQuestionsByExamID(ctx context.Context, examID string) ([]models.ImportantQuestion, error)
	GetImportantQuestionsBySubjectID(ctx context.Context, subjectID string) ([]models.ImportantQuestion, error)
	UpdateImportantQuestion(ctx context.Context, question *models.ImportantQuestion) error
	DeleteImportantQuestion(ctx context.Context, id string) error
}

// PGImportantQuestionRepository implements ImportantQuestionRepository for PostgreSQL.
type PGImportantQuestionRepository struct {
	db *pgxpool.Pool
}

// NewPGImportantQuestionRepository creates a new PostgreSQL important question repository.
func NewPGImportantQuestionRepository(db *pgxpool.Pool) *PGImportantQuestionRepository {
	return &PGImportantQuestionRepository{db: db}
}

// CreateImportantQuestion inserts a new important question into the database.
func (r *PGImportantQuestionRepository) CreateImportantQuestion(ctx context.Context, question *models.ImportantQuestion) error {
	query := `
		INSERT INTO important_questions (
			id, user_id, subject_id, exam_id, question_text, answer_text, source,
			unit, topic, marks, frequency_count, is_practiced, last_practiced_at,
			confidence_level, tags, created_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
		) RETURNING id, created_at
	`
	question.ID = models.NewUUID()
	question.CreatedAt = time.Now()

	_, err := r.db.Exec(ctx, query,
		question.ID, question.UserID, question.SubjectID, question.ExamID, question.QuestionText, question.AnswerText, question.Source,
		question.Unit, question.Topic, question.Marks, question.FrequencyCount, question.IsPracticed, question.LastPracticedAt,
		question.ConfidenceLevel, question.Tags, question.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create important question: %w", err)
	}
	return nil
}

// GetImportantQuestionByID retrieves an important question by its ID.
func (r *PGImportantQuestionRepository) GetImportantQuestionByID(ctx context.Context, id string) (*models.ImportantQuestion, error) {
	question := &models.ImportantQuestion{}
	query := `
		SELECT
			id, user_id, subject_id, exam_id, question_text, answer_text, source,
			unit, topic, marks, frequency_count, is_practiced, last_practiced_at,
			confidence_level, tags, created_at
		FROM important_questions
		WHERE id = $1
	`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&question.ID, &question.UserID, &question.SubjectID, &question.ExamID, &question.QuestionText, &question.AnswerText, &question.Source,
		&question.Unit, &question.Topic, &question.Marks, &question.FrequencyCount, &question.IsPracticed, &question.LastPracticedAt,
		&question.ConfidenceLevel, &question.Tags, &question.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get important question by ID: %w", err)
	}
	return question, nil
}

// GetImportantQuestionsByExamID retrieves all important questions for a given exam.
func (r *PGImportantQuestionRepository) GetImportantQuestionsByExamID(ctx context.Context, examID string) ([]models.ImportantQuestion, error) {
	var questions []models.ImportantQuestion
	query := `
		SELECT
			id, user_id, subject_id, exam_id, question_text, answer_text, source,
			unit, topic, marks, frequency_count, is_practiced, last_practiced_at,
			confidence_level, tags, created_at
		FROM important_questions
		WHERE exam_id = $1
		ORDER BY created_at ASC
	`
	rows, err := r.db.Query(ctx, query, examID)
	if err != nil {
		return nil, fmt.Errorf("failed to get important questions by exam ID: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		question := models.ImportantQuestion{}
		err := rows.Scan(
			&question.ID, &question.UserID, &question.SubjectID, &question.ExamID, &question.QuestionText, &question.AnswerText, &question.Source,
			&question.Unit, &question.Topic, &question.Marks, &question.FrequencyCount, &question.IsPracticed, &question.LastPracticedAt,
			&question.ConfidenceLevel, &question.Tags, &question.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan important question row: %w", err)
		}
		questions = append(questions, question)
	}
	return questions, nil
}

// GetImportantQuestionsBySubjectID retrieves all important questions for a given subject.
func (r *PGImportantQuestionRepository) GetImportantQuestionsBySubjectID(ctx context.Context, subjectID string) ([]models.ImportantQuestion, error) {
	var questions []models.ImportantQuestion
	query := `
		SELECT
			id, user_id, subject_id, exam_id, question_text, answer_text, source,
			unit, topic, marks, frequency_count, is_practiced, last_practiced_at,
			confidence_level, tags, created_at
		FROM important_questions
		WHERE subject_id = $1
		ORDER BY created_at ASC
	`
	rows, err := r.db.Query(ctx, query, subjectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get important questions by subject ID: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		question := models.ImportantQuestion{}
		err := rows.Scan(
			&question.ID, &question.UserID, &question.SubjectID, &question.ExamID, &question.QuestionText, &question.AnswerText, &question.Source,
			&question.Unit, &question.Topic, &question.Marks, &question.FrequencyCount, &question.IsPracticed, &question.LastPracticedAt,
			&question.ConfidenceLevel, &question.Tags, &question.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan important question row: %w", err)
		}
		questions = append(questions, question)
	}
	return questions, nil
}

// UpdateImportantQuestion updates an existing important question in the database.
func (r *PGImportantQuestionRepository) UpdateImportantQuestion(ctx context.Context, question *models.ImportantQuestion) error {
	query := `
		UPDATE important_questions SET
			subject_id = $1, exam_id = $2, question_text = $3, answer_text = $4, source = $5,
			unit = $6, topic = $7, marks = $8, frequency_count = $9, is_practiced = $10,
			last_practiced_at = $11, confidence_level = $12, tags = $13
		WHERE id = $14 AND user_id = $15
	`
	_, err := r.db.Exec(ctx, query,
		question.SubjectID, question.ExamID, question.QuestionText, question.AnswerText, question.Source,
		question.Unit, question.Topic, question.Marks, question.FrequencyCount, question.IsPracticed,
		question.LastPracticedAt, question.ConfidenceLevel, question.Tags,
		question.ID, question.UserID,
	)
	if err != nil {
		return fmt.Errorf("failed to update important question: %w", err)
	}
	return nil
}

// DeleteImportantQuestion deletes an important question from the database.
func (r *PGImportantQuestionRepository) DeleteImportantQuestion(ctx context.Context, id string) error {
	query := `DELETE FROM important_questions WHERE id = $1`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete important question: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("important question with ID %s not found", id)
	}
	return nil
}
