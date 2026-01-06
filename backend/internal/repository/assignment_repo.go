package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
)

// --- Assignment Repository ---

// AssignmentRepository defines the interface for assignment data operations.
type AssignmentRepository interface {
	CreateAssignment(ctx context.Context, assignment *models.Assignment) error
	GetAssignmentByID(ctx context.Context, id string) (*models.Assignment, error)
	GetAssignmentsByUserID(ctx context.Context, userID string) ([]models.Assignment, error)
	GetPendingAssignmentsByUserID(ctx context.Context, userID string) ([]models.Assignment, error)
	GetOverdueAssignmentsByUserID(ctx context.Context, userID string) ([]models.Assignment, error)
	UpdateAssignment(ctx context.Context, assignment *models.Assignment) error
	DeleteAssignment(ctx context.Context, id string) error
	UpdateAssignmentStatus(ctx context.Context, id string, status string) error
}

// PGAssignmentRepository implements AssignmentRepository for PostgreSQL.
type PGAssignmentRepository struct {
	db *pgxpool.Pool
}

// NewPGAssignmentRepository creates a new PostgreSQL assignment repository.
func NewPGAssignmentRepository(db *pgxpool.Pool) *PGAssignmentRepository {
	return &PGAssignmentRepository{db: db}
}

// CreateAssignment inserts a new assignment into the database.
func (r *PGAssignmentRepository) CreateAssignment(ctx context.Context, assignment *models.Assignment) error {
	query := `
		INSERT INTO assignments (
			id, user_id, subject_id, staff_id, title, description, instructions,
			assignment_type, assigned_date, due_date, submitted_at, status,
			max_marks, obtained_marks, feedback, priority, estimated_hours,
			actual_hours, reminder_enabled, reminder_before_hours, last_reminded_at,
			tags, is_recurring, recurrence_pattern, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16,
			$17, $18, $19, $20, $21, $22, $23, $24, $25, $26
		) RETURNING id, created_at, updated_at
	`
	assignment.ID = models.NewUUID()
	assignment.CreatedAt = time.Now()
	assignment.UpdatedAt = time.Now()

	_, err := r.db.Exec(ctx, query,
		assignment.ID, assignment.UserID, assignment.SubjectID, assignment.StaffID, assignment.Title, assignment.Description, assignment.Instructions,
		assignment.AssignmentType, assignment.AssignedDate, assignment.DueDate, assignment.SubmittedAt, assignment.Status,
		assignment.MaxMarks, assignment.ObtainedMarks, assignment.Feedback, assignment.Priority, assignment.EstimatedHours,
		assignment.ActualHours, assignment.ReminderEnabled, assignment.ReminderBeforeHours, assignment.LastRemindedAt,
		assignment.Tags, assignment.IsRecurring, assignment.RecurrencePattern, assignment.CreatedAt, assignment.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create assignment: %w", err)
	}
	return nil
}

// GetAssignmentByID retrieves an assignment by its ID.
func (r *PGAssignmentRepository) GetAssignmentByID(ctx context.Context, id string) (*models.Assignment, error) {
	assignment := &models.Assignment{}
	query := `
		SELECT
			id, user_id, subject_id, staff_id, title, description, instructions,
			assignment_type, assigned_date, due_date, submitted_at, status,
			max_marks, obtained_marks, feedback, priority, estimated_hours,
			actual_hours, reminder_enabled, reminder_before_hours, last_reminded_at,
			tags, is_recurring, recurrence_pattern, created_at, updated_at
		FROM assignments
		WHERE id = $1
	`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&assignment.ID, &assignment.UserID, &assignment.SubjectID, &assignment.StaffID, &assignment.Title, &assignment.Description, &assignment.Instructions,
		&assignment.AssignmentType, &assignment.AssignedDate, &assignment.DueDate, &assignment.SubmittedAt, &assignment.Status,
		&assignment.MaxMarks, &assignment.ObtainedMarks, &assignment.Feedback, &assignment.Priority, &assignment.EstimatedHours,
		&assignment.ActualHours, &assignment.ReminderEnabled, &assignment.ReminderBeforeHours, &assignment.LastRemindedAt,
		&assignment.Tags, &assignment.IsRecurring, &assignment.RecurrencePattern, &assignment.CreatedAt, &assignment.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get assignment by ID: %w", err)
	}
	return assignment, nil
}

// GetAssignmentsByUserID retrieves all assignments for a given user.
func (r *PGAssignmentRepository) GetAssignmentsByUserID(ctx context.Context, userID string) ([]models.Assignment, error) {
	var assignments []models.Assignment
	query := `
		SELECT
			id, user_id, subject_id, staff_id, title, description, instructions,
			assignment_type, assigned_date, due_date, submitted_at, status,
			max_marks, obtained_marks, feedback, priority, estimated_hours,
			actual_hours, reminder_enabled, reminder_before_hours, last_reminded_at,
			tags, is_recurring, recurrence_pattern, created_at, updated_at
		FROM assignments
		WHERE user_id = $1
		ORDER BY due_date ASC
	`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get assignments by user ID: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		assignment := models.Assignment{}
		err := rows.Scan(
			&assignment.ID, &assignment.UserID, &assignment.SubjectID, &assignment.StaffID, &assignment.Title, &assignment.Description, &assignment.Instructions,
			&assignment.AssignmentType, &assignment.AssignedDate, &assignment.DueDate, &assignment.SubmittedAt, &assignment.Status,
			&assignment.MaxMarks, &assignment.ObtainedMarks, &assignment.Feedback, &assignment.Priority, &assignment.EstimatedHours,
			&assignment.ActualHours, &assignment.ReminderEnabled, &assignment.ReminderBeforeHours, &assignment.LastRemindedAt,
			&assignment.Tags, &assignment.IsRecurring, &assignment.RecurrencePattern, &assignment.CreatedAt, &assignment.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan assignment row: %w", err)
		}
		assignments = append(assignments, assignment)
	}
	return assignments, nil
}

// GetPendingAssignmentsByUserID retrieves pending assignments for a given user.
func (r *PGAssignmentRepository) GetPendingAssignmentsByUserID(ctx context.Context, userID string) ([]models.Assignment, error) {
	var assignments []models.Assignment
	query := `
		SELECT
			id, user_id, subject_id, staff_id, title, description, instructions,
			assignment_type, assigned_date, due_date, submitted_at, status,
			max_marks, obtained_marks, feedback, priority, estimated_hours,
			actual_hours, reminder_enabled, reminder_before_hours, last_reminded_at,
			tags, is_recurring, recurrence_pattern, created_at, updated_at
		FROM assignments
		WHERE user_id = $1 AND status IN ('pending', 'in_progress') AND due_date >= NOW()
		ORDER BY due_date ASC
	`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get pending assignments: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		assignment := models.Assignment{}
		err := rows.Scan(
			&assignment.ID, &assignment.UserID, &assignment.SubjectID, &assignment.StaffID, &assignment.Title, &assignment.Description, &assignment.Instructions,
			&assignment.AssignmentType, &assignment.AssignedDate, &assignment.DueDate, &assignment.SubmittedAt, &assignment.Status,
			&assignment.MaxMarks, &assignment.ObtainedMarks, &assignment.Feedback, &assignment.Priority, &assignment.EstimatedHours,
			&assignment.ActualHours, &assignment.ReminderEnabled, &assignment.ReminderBeforeHours, &assignment.LastRemindedAt,
			&assignment.Tags, &assignment.IsRecurring, &assignment.RecurrencePattern, &assignment.CreatedAt, &assignment.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan pending assignment row: %w", err)
		}
		assignments = append(assignments, assignment)
	}
	return assignments, nil
}

// GetOverdueAssignmentsByUserID retrieves overdue assignments for a given user.
func (r *PGAssignmentRepository) GetOverdueAssignmentsByUserID(ctx context.Context, userID string) ([]models.Assignment, error) {
	var assignments []models.Assignment
	query := `
		SELECT
			id, user_id, subject_id, staff_id, title, description, instructions,
			assignment_type, assigned_date, due_date, submitted_at, status,
			max_marks, obtained_marks, feedback, priority, estimated_hours,
			actual_hours, reminder_enabled, reminder_before_hours, last_reminded_at,
			tags, is_recurring, recurrence_pattern, created_at, updated_at
		FROM assignments
		WHERE user_id = $1 AND status IN ('pending', 'in_progress') AND due_date < NOW()
		ORDER BY due_date ASC
	`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get overdue assignments: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		assignment := models.Assignment{}
		err := rows.Scan(
			&assignment.ID, &assignment.UserID, &assignment.SubjectID, &assignment.StaffID, &assignment.Title, &assignment.Description, &assignment.Instructions,
			&assignment.AssignmentType, &assignment.AssignedDate, &assignment.DueDate, &assignment.SubmittedAt, &assignment.Status,
			&assignment.MaxMarks, &assignment.ObtainedMarks, &assignment.Feedback, &assignment.Priority, &assignment.EstimatedHours,
			&assignment.ActualHours, &assignment.ReminderEnabled, &assignment.ReminderBeforeHours, &assignment.LastRemindedAt,
			&assignment.Tags, &assignment.IsRecurring, &assignment.RecurrencePattern, &assignment.CreatedAt, &assignment.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan overdue assignment row: %w", err)
		}
		assignments = append(assignments, assignment)
	}
	return assignments, nil
}

// UpdateAssignment updates an existing assignment in the database.
func (r *PGAssignmentRepository) UpdateAssignment(ctx context.Context, assignment *models.Assignment) error {
	query := `
		UPDATE assignments SET
			subject_id = $1, staff_id = $2, title = $3, description = $4, instructions = $5,
			assignment_type = $6, assigned_date = $7, due_date = $8, submitted_at = $9, status = $10,
			max_marks = $11, obtained_marks = $12, feedback = $13, priority = $14, estimated_hours = $15,
			actual_hours = $16, reminder_enabled = $17, reminder_before_hours = $18, last_reminded_at = $19,
			tags = $20, is_recurring = $21, recurrence_pattern = $22, updated_at = $23
		WHERE id = $24 AND user_id = $25
	`
	assignment.UpdatedAt = time.Now()

	cmdTag, err := r.db.Exec(ctx, query,
		assignment.SubjectID, assignment.StaffID, assignment.Title, assignment.Description, assignment.Instructions,
		assignment.AssignmentType, assignment.AssignedDate, assignment.DueDate, assignment.SubmittedAt, assignment.Status,
		assignment.MaxMarks, assignment.ObtainedMarks, assignment.Feedback, assignment.Priority, assignment.EstimatedHours,
		assignment.ActualHours, assignment.ReminderEnabled, assignment.ReminderBeforeHours, assignment.LastRemindedAt,
		assignment.Tags, assignment.IsRecurring, assignment.RecurrencePattern, assignment.UpdatedAt,
		assignment.ID, assignment.UserID,
	)
	if err != nil {
		return fmt.Errorf("failed to update assignment: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("assignment with ID %s not found or not owned by user", assignment.ID)
	}
	return nil
}

// UpdateAssignmentStatus updates the status of an assignment.
func (r *PGAssignmentRepository) UpdateAssignmentStatus(ctx context.Context, id string, status string) error {
	query := `
		UPDATE assignments SET
			status = $1, updated_at = $2
		WHERE id = $3
	`
	_, err := r.db.Exec(ctx, query, status, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to update assignment status: %w", err)
	}
	return nil
}

// DeleteAssignment deletes an assignment from the database.
func (r *PGAssignmentRepository) DeleteAssignment(ctx context.Context, id string) error {
	query := `DELETE FROM assignments WHERE id = $1`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete assignment: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("assignment with ID %s not found", id)
	}
	return nil
}

// --- Assignment Attachment Repository ---

// AssignmentAttachmentRepository defines the interface for assignment attachment data operations.
type AssignmentAttachmentRepository interface {
	CreateAttachment(ctx context.Context, attachment *models.AssignmentAttachment) error
	GetAttachmentsByAssignmentID(ctx context.Context, assignmentID string) ([]models.AssignmentAttachment, error)
	GetAttachmentByID(ctx context.Context, id string) (*models.AssignmentAttachment, error)
	DeleteAttachment(ctx context.Context, id string) error
}

// PGAssignmentAttachmentRepository implements AssignmentAttachmentRepository for PostgreSQL.
type PGAssignmentAttachmentRepository struct {
	db *pgxpool.Pool
}

// NewPGAssignmentAttachmentRepository creates a new PostgreSQL assignment attachment repository.
func NewPGAssignmentAttachmentRepository(db *pgxpool.Pool) *PGAssignmentAttachmentRepository {
	return &PGAssignmentAttachmentRepository{db: db}
}

// CreateAttachment inserts a new assignment attachment into the database.
func (r *PGAssignmentAttachmentRepository) CreateAttachment(ctx context.Context, attachment *models.AssignmentAttachment) error {
	query := `
		INSERT INTO assignment_attachments (
			id, assignment_id, file_name, file_type, file_size, file_url, storage_key, attachment_type, uploaded_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		) RETURNING id, uploaded_at
	`
	attachment.ID = models.NewUUID()
	attachment.UploadedAt = time.Now()

	_, err := r.db.Exec(ctx, query,
		attachment.ID, attachment.AssignmentID, attachment.FileName, attachment.FileType, attachment.FileSize,
		attachment.FileURL, attachment.StorageKey, attachment.AttachmentType, attachment.UploadedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create assignment attachment: %w", err)
	}
	return nil
}

// GetAttachmentsByAssignmentID retrieves all attachments for a given assignment.
func (r *PGAssignmentAttachmentRepository) GetAttachmentsByAssignmentID(ctx context.Context, assignmentID string) ([]models.AssignmentAttachment, error) {
	var attachments []models.AssignmentAttachment
	query := `
		SELECT
			id, assignment_id, file_name, file_type, file_size, file_url, storage_key, attachment_type, uploaded_at
		FROM assignment_attachments
		WHERE assignment_id = $1
		ORDER BY uploaded_at ASC
	`
	rows, err := r.db.Query(ctx, query, assignmentID)
	if err != nil {
		return nil, fmt.Errorf("failed to get attachments by assignment ID: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		attachment := models.AssignmentAttachment{}
		err := rows.Scan(
			&attachment.ID, &attachment.AssignmentID, &attachment.FileName, &attachment.FileType, &attachment.FileSize,
			&attachment.FileURL, &attachment.StorageKey, &attachment.AttachmentType, &attachment.UploadedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan attachment row: %w", err)
		}
		attachments = append(attachments, attachment)
	}
	return attachments, nil
}

// GetAttachmentByID retrieves an attachment by its ID.
func (r *PGAssignmentAttachmentRepository) GetAttachmentByID(ctx context.Context, id string) (*models.AssignmentAttachment, error) {
	attachment := &models.AssignmentAttachment{}
	query := `
		SELECT
			id, assignment_id, file_name, file_type, file_size, file_url, storage_key, attachment_type, uploaded_at
		FROM assignment_attachments
		WHERE id = $1
	`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&attachment.ID, &attachment.AssignmentID, &attachment.FileName, &attachment.FileType, &attachment.FileSize,
		&attachment.FileURL, &attachment.StorageKey, &attachment.AttachmentType, &attachment.UploadedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get attachment by ID: %w", err)
	}
	return attachment, nil
}

// DeleteAttachment deletes an assignment attachment from the database.
func (r *PGAssignmentAttachmentRepository) DeleteAttachment(ctx context.Context, id string) error {
	query := `DELETE FROM assignment_attachments WHERE id = $1`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete assignment attachment: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("assignment attachment with ID %s not found", id)
	}
	return nil
}
