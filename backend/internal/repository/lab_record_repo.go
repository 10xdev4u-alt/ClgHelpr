package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
)

// --- LabRecord Repository ---

// LabRecordRepository defines the interface for lab record data operations.
type LabRecordRepository interface {
	CreateLabRecord(ctx context.Context, record *models.LabRecord) error
	GetLabRecordByID(ctx context.Context, id string) (*models.LabRecord, error)
	GetLabRecordsByUserID(ctx context.Context, userID string) ([]models.LabRecord, error)
	GetLabRecordsByUserIDAndSubjectID(ctx context.Context, userID, subjectID string) ([]models.LabRecord, error)
	UpdateLabRecord(ctx context.Context, record *models.LabRecord) error
	DeleteLabRecord(ctx context.Context, id string) error
	UpdateLabRecordStatus(ctx context.Context, id string, status string) error
}

// PGLabRecordRepository implements LabRecordRepository for PostgreSQL.
type PGLabRecordRepository struct {
	db *pgxpool.Pool
}

// NewPGLabRecordRepository creates a new PostgreSQL lab record repository.
func NewPGLabRecordRepository(db *pgxpool.Pool) *PGLabRecordRepository {
	return &PGLabRecordRepository{db: db}
}

// CreateLabRecord inserts a new lab record into the database.
func (r *PGLabRecordRepository) CreateLabRecord(ctx context.Context, record *models.LabRecord) error {
	query := `
		INSERT INTO lab_records (
			id, user_id, subject_id, experiment_number, title, lab_date, record_written_date,
			submitted_date, status, aim, algorithm, code, output, observations, result,
			viva_questions, print_required, pages_to_print, printed_at, marks,
			staff_remarks, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16,
			$17, $18, $19, $20, $21, $22, $23
		) RETURNING id, created_at, updated_at
	`
	record.ID = models.NewUUID()
	record.CreatedAt = time.Now()
	record.UpdatedAt = time.Now()

	_, err := r.db.Exec(ctx, query,
		record.ID, record.UserID, record.SubjectID, record.ExperimentNumber, record.Title, record.LabDate, record.RecordWrittenDate,
		record.SubmittedDate, record.Status, record.Aim, record.Algorithm, record.Code, record.Output, record.Observations, record.Result,
		record.VivaQuestions, record.PrintRequired, record.PagesToPrint, record.PrintedAt, record.Marks,
		record.StaffRemarks, record.CreatedAt, record.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create lab record: %w", err)
	}
	return nil
}

// GetLabRecordByID retrieves a lab record by its ID.
func (r *PGLabRecordRepository) GetLabRecordByID(ctx context.Context, id string) (*models.LabRecord, error) {
	record := &models.LabRecord{}
	query := `
		SELECT
			id, user_id, subject_id, experiment_number, title, lab_date, record_written_date,
			submitted_date, status, aim, algorithm, code, output, observations, result,
			viva_questions, print_required, pages_to_print, printed_at, marks,
			staff_remarks, created_at, updated_at
		FROM lab_records
		WHERE id = $1
	`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&record.ID, &record.UserID, &record.SubjectID, &record.ExperimentNumber, &record.Title, &record.LabDate, &record.RecordWrittenDate,
		&record.SubmittedDate, &record.Status, &record.Aim, &record.Algorithm, &record.Code, &record.Output, &record.Observations, &record.Result,
		&record.VivaQuestions, &record.PrintRequired, &record.PagesToPrint, &record.PrintedAt, &record.Marks,
		&record.StaffRemarks, &record.CreatedAt, &record.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get lab record by ID: %w", err)
	}
	return record, nil
}

// GetLabRecordsByUserID retrieves all lab records for a given user.
func (r *PGLabRecordRepository) GetLabRecordsByUserID(ctx context.Context, userID string) ([]models.LabRecord, error) {
	var records []models.LabRecord
	query := `
		SELECT
			id, user_id, subject_id, experiment_number, title, lab_date, record_written_date,
			submitted_date, status, aim, algorithm, code, output, observations, result,
			viva_questions, print_required, pages_to_print, printed_at, marks,
			staff_remarks, created_at, updated_at
		FROM lab_records
		WHERE user_id = $1
		ORDER BY lab_date DESC, experiment_number ASC
	`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get lab records by user ID: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		record := models.LabRecord{}
		err := rows.Scan(
			&record.ID, &record.UserID, &record.SubjectID, &record.ExperimentNumber, &record.Title, &record.LabDate, &record.RecordWrittenDate,
			&record.SubmittedDate, &record.Status, &record.Aim, &record.Algorithm, &record.Code, &record.Output, &record.Observations, &record.Result,
			&record.VivaQuestions, &record.PrintRequired, &record.PagesToPrint, &record.PrintedAt, &record.Marks,
			&record.StaffRemarks, &record.CreatedAt, &record.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan lab record row: %w", err)
		}
		records = append(records, record)
	}
	return records, nil
}

// GetLabRecordsByUserIDAndSubjectID retrieves all lab records for a given user and subject.
func (r *PGLabRecordRepository) GetLabRecordsByUserIDAndSubjectID(ctx context.Context, userID, subjectID string) ([]models.LabRecord, error) {
	var records []models.LabRecord
	query := `
		SELECT
			id, user_id, subject_id, experiment_number, title, lab_date, record_written_date,
			submitted_date, status, aim, algorithm, code, output, observations, result,
			viva_questions, print_required, pages_to_print, printed_at, marks,
			staff_remarks, created_at, updated_at
		FROM lab_records
		WHERE user_id = $1 AND subject_id = $2
		ORDER BY experiment_number ASC
	`
	rows, err := r.db.Query(ctx, query, userID, subjectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get lab records by user ID and subject ID: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		record := models.LabRecord{}
		err := rows.Scan(
			&record.ID, &record.UserID, &record.SubjectID, &record.ExperimentNumber, &record.Title, &record.LabDate, &record.RecordWrittenDate,
			&record.SubmittedDate, &record.Status, &record.Aim, &record.Algorithm, &record.Code, &record.Output, &record.Observations, &record.Result,
			&record.VivaQuestions, &record.PrintRequired, &record.PagesToPrint, &record.PrintedAt, &record.Marks,
			&record.StaffRemarks, &record.CreatedAt, &record.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan lab record row: %w", err)
		}
		records = append(records, record)
	}
	return records, nil
}

// UpdateLabRecord updates an existing lab record in the database.
func (r *PGLabRecordRepository) UpdateLabRecord(ctx context.Context, record *models.LabRecord) error {
	query := `
		UPDATE lab_records SET
			subject_id = $1, experiment_number = $2, title = $3, lab_date = $4,
			record_written_date = $5, submitted_date = $6, status = $7, aim = $8,
			algorithm = $9, code = $10, output = $11, observations = $12, result = $13,
			viva_questions = $14, print_required = $15, pages_to_print = $16, printed_at = $17,
			marks = $18, staff_remarks = $19, updated_at = $20
		WHERE id = $21 AND user_id = $22
	`
	record.UpdatedAt = time.Now()

	cmdTag, err := r.db.Exec(ctx, query,
		record.SubjectID, record.ExperimentNumber, record.Title, record.LabDate,
		record.RecordWrittenDate, record.SubmittedDate, record.Status, record.Aim,
		record.Algorithm, record.Code, record.Output, record.Observations, record.Result,
		record.VivaQuestions, record.PrintRequired, record.PagesToPrint, record.PrintedAt,
		record.Marks, record.StaffRemarks, record.UpdatedAt,
		record.ID, record.UserID,
	)
	if err != nil {
		return fmt.Errorf("failed to update lab record: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("lab record with ID %s not found or not owned by user", record.ID)
	}
	return nil
}

// UpdateLabRecordStatus updates the status of a lab record.
func (r *PGLabRecordRepository) UpdateLabRecordStatus(ctx context.Context, id string, status string) error {
	query := `
		UPDATE lab_records SET
			status = $1, updated_at = $2
		WHERE id = $3
	`
	_, err := r.db.Exec(ctx, query, status, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to update lab record status: %w", err)
	}
	return nil
}

// DeleteLabRecord deletes a lab record from the database.
func (r *PGLabRecordRepository) DeleteLabRecord(ctx context.Context, id string) error {
	query := `DELETE FROM lab_records WHERE id = $1`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete lab record: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("lab record with ID %s not found", id)
	}
	return nil
}

// --- LabRecordAttachment Repository ---

// LabRecordAttachmentRepository defines the interface for lab record attachment data operations.
type LabRecordAttachmentRepository interface {
	CreateAttachment(ctx context.Context, attachment *models.LabRecordAttachment) error
	GetAttachmentsByLabRecordID(ctx context.Context, labRecordID string) ([]models.LabRecordAttachment, error)
	GetAttachmentByID(ctx context.Context, id string) (*models.LabRecordAttachment, error)
	DeleteAttachment(ctx context.Context, id string) error
}

// PGLabRecordAttachmentRepository implements LabRecordAttachmentRepository for PostgreSQL.
type PGLabRecordAttachmentRepository struct {
	db *pgxpool.Pool
}

// NewPGLabRecordAttachmentRepository creates a new PostgreSQL lab record attachment repository.
func NewPGLabRecordAttachmentRepository(db *pgxpool.Pool) *PGLabRecordAttachmentRepository {
	return &PGLabRecordAttachmentRepository{db: db}
}

// CreateAttachment inserts a new lab record attachment into the database.
func (r *PGLabRecordAttachmentRepository) CreateAttachment(ctx context.Context, attachment *models.LabRecordAttachment) error {
	query := `
		INSERT INTO lab_record_attachments (
			id, lab_record_id, file_name, file_type, file_url, storage_key, attachment_type, uploaded_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8
		) RETURNING id, uploaded_at
	`
	attachment.ID = models.NewUUID()
	attachment.UploadedAt = time.Now()

	_, err := r.db.Exec(ctx, query,
		attachment.ID, attachment.LabRecordID, attachment.FileName, attachment.FileType, attachment.FileURL,
		attachment.StorageKey, attachment.AttachmentType, attachment.UploadedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create lab record attachment: %w", err)
	}
	return nil
}

// GetAttachmentsByLabRecordID retrieves all attachments for a given lab record.
func (r *PGLabRecordAttachmentRepository) GetAttachmentsByLabRecordID(ctx context.Context, labRecordID string) ([]models.LabRecordAttachment, error) {
	var attachments []models.LabRecordAttachment
	query := `
		SELECT
			id, lab_record_id, file_name, file_type, file_url, storage_key, attachment_type, uploaded_at
		FROM lab_record_attachments
		WHERE lab_record_id = $1
		ORDER BY uploaded_at ASC
	`
	rows, err := r.db.Query(ctx, query, labRecordID)
	if err != nil {
		return nil, fmt.Errorf("failed to get attachments by lab record ID: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		attachment := models.LabRecordAttachment{}
		err := rows.Scan(
			&attachment.ID, &attachment.LabRecordID, &attachment.FileName, &attachment.FileType, &attachment.FileURL,
			&attachment.StorageKey, &attachment.AttachmentType, &attachment.UploadedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan attachment row: %w", err)
		}
		attachments = append(attachments, attachment)
	}
	return attachments, nil
}

// GetAttachmentByID retrieves an attachment by its ID.
func (r *PGLabRecordAttachmentRepository) GetAttachmentByID(ctx context.Context, id string) (*models.LabRecordAttachment, error) {
	attachment := &models.LabRecordAttachment{}
	query := `
		SELECT
			id, lab_record_id, file_name, file_type, file_url, storage_key, attachment_type, uploaded_at
		FROM lab_record_attachments
		WHERE id = $1
	`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&attachment.ID, &attachment.LabRecordID, &attachment.FileName, &attachment.FileType, &attachment.FileURL,
		&attachment.StorageKey, &attachment.AttachmentType, &attachment.UploadedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get attachment by ID: %w", err)
	}
	return attachment, nil
}

// DeleteAttachment deletes a lab record attachment from the database.
func (r *PGLabRecordAttachmentRepository) DeleteAttachment(ctx context.Context, id string) error {
	query := `DELETE FROM lab_record_attachments WHERE id = $1`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete lab record attachment: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("lab record attachment with ID %s not found", id)
	}
	return nil
}
