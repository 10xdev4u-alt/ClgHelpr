package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
)

// DocumentRepository defines the interface for document data operations.
type DocumentRepository interface {
	CreateDocument(ctx context.Context, document *models.Document) error
	GetDocumentByID(ctx context.Context, id string) (*models.Document, error)
	GetDocumentsByUserID(ctx context.Context, userID string) ([]models.Document, error)
	GetDocumentsByUserIDAndSubjectID(ctx context.Context, userID, subjectID string) ([]models.Document, error)
	UpdateDocument(ctx context.Context, document *models.Document) error
	DeleteDocument(ctx context.Context, id string) error
	IncrementViewCount(ctx context.Context, id string) error
	IncrementDownloadCount(ctx context.Context, id string) error
}

// PGDocumentRepository implements DocumentRepository for PostgreSQL.
type PGDocumentRepository struct {
	db *pgxpool.Pool
}

// NewPGDocumentRepository creates a new PostgreSQL document repository.
func NewPGDocumentRepository(db *pgxpool.Pool) *PGDocumentRepository {
	return &PGDocumentRepository{db: db}
}

// CreateDocument inserts a new document into the database.
func (r *PGDocumentRepository) CreateDocument(ctx context.Context, document *models.Document) error {
	query := `
		INSERT INTO documents (
			id, user_id, subject_id, title, description, document_type, file_name,
			file_type, file_size, file_url, storage_key, folder, tags, is_public,
			shared_with, view_count, download_count, last_accessed_at, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16,
			$17, $18, $19, $20
		) RETURNING id, created_at, updated_at
	`
	document.ID = models.NewUUID()
	document.CreatedAt = time.Now()
	document.UpdatedAt = time.Now()

	_, err := r.db.Exec(ctx, query,
		document.ID, document.UserID, document.SubjectID, document.Title, document.Description, document.DocumentType, document.FileName,
		document.FileType, document.FileSize, document.FileURL, document.StorageKey, document.Folder, document.Tags, document.IsPublic,
		document.SharedWith, document.ViewCount, document.DownloadCount, document.LastAccessedAt, document.CreatedAt, document.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create document: %w", err)
	}
	return nil
}

// GetDocumentByID retrieves a document by its ID.
func (r *PGDocumentRepository) GetDocumentByID(ctx context.Context, id string) (*models.Document, error) {
	document := &models.Document{}
	query := `
		SELECT
			id, user_id, subject_id, title, description, document_type, file_name,
			file_type, file_size, file_url, storage_key, folder, tags, is_public,
			shared_with, view_count, download_count, last_accessed_at, created_at, updated_at
		FROM documents
		WHERE id = $1
	`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&document.ID, &document.UserID, &document.SubjectID, &document.Title, &document.Description, &document.DocumentType, &document.FileName,
		&document.FileType, &document.FileSize, &document.FileURL, &document.StorageKey, &document.Folder, &document.Tags, &document.IsPublic,
		&document.SharedWith, &document.ViewCount, &document.DownloadCount, &document.LastAccessedAt, &document.CreatedAt, &document.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get document by ID: %w", err)
	}
	return document, nil
}

// GetDocumentsByUserID retrieves all documents for a given user.
func (r *PGDocumentRepository) GetDocumentsByUserID(ctx context.Context, userID string) ([]models.Document, error) {
	var documents []models.Document
	query := `
		SELECT
			id, user_id, subject_id, title, description, document_type, file_name,
			file_type, file_size, file_url, storage_key, folder, tags, is_public,
			shared_with, view_count, download_count, last_accessed_at, created_at, updated_at
		FROM documents
		WHERE user_id = $1
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get documents by user ID: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		document := models.Document{}
		err := rows.Scan(
			&document.ID, &document.UserID, &document.SubjectID, &document.Title, &document.Description, &document.DocumentType, &document.FileName,
			&document.FileType, &document.FileSize, &document.FileURL, &document.StorageKey, &document.Folder, &document.Tags, &document.IsPublic,
			&document.SharedWith, &document.ViewCount, &document.DownloadCount, &document.LastAccessedAt, &document.CreatedAt, &document.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan document row: %w", err)
		}
		documents = append(documents, document)
	}
	return documents, nil
}

// GetDocumentsByUserIDAndSubjectID retrieves all documents for a given user and subject.
func (r *PGDocumentRepository) GetDocumentsByUserIDAndSubjectID(ctx context.Context, userID, subjectID string) ([]models.Document, error) {
	var documents []models.Document
	query := `
		SELECT
			id, user_id, subject_id, title, description, document_type, file_name,
			file_type, file_size, file_url, storage_key, folder, tags, is_public,
			shared_with, view_count, download_count, last_accessed_at, created_at, updated_at
		FROM documents
		WHERE user_id = $1 AND subject_id = $2
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, query, userID, subjectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get documents by user ID and subject ID: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		document := models.Document{}
		err := rows.Scan(
			&document.ID, &document.UserID, &document.SubjectID, &document.Title, &document.Description, &document.DocumentType, &document.FileName,
			&document.FileType, &document.FileSize, &document.FileURL, &document.StorageKey, &document.Folder, &document.Tags, &document.IsPublic,
			&document.SharedWith, &document.ViewCount, &document.DownloadCount, &document.LastAccessedAt, &document.CreatedAt, &document.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan document row: %w", err)
		}
		documents = append(documents, document)
	}
	return documents, nil
}

// UpdateDocument updates an existing document in the database.
func (r *PGDocumentRepository) UpdateDocument(ctx context.Context, document *models.Document) error {
	query := `
		UPDATE documents SET
			subject_id = $1, title = $2, description = $3, document_type = $4, file_name = $5,
			file_type = $6, file_size = $7, file_url = $8, storage_key = $9, folder = $10,
			tags = $11, is_public = $12, shared_with = $13, updated_at = $14
		WHERE id = $15 AND user_id = $16
	`
	document.UpdatedAt = time.Now()

	cmdTag, err := r.db.Exec(ctx, query,
		document.SubjectID, document.Title, document.Description, document.DocumentType, document.FileName,
		document.FileType, document.FileSize, document.FileURL, document.StorageKey, document.Folder,
		document.Tags, document.IsPublic, document.SharedWith, document.UpdatedAt,
		document.ID, document.UserID,
	)
	if err != nil {
		return fmt.Errorf("failed to update document: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("document with ID %s not found or not owned by user", document.ID)
	}
	return nil
}

// DeleteDocument deletes a document from the database.
func (r *PGDocumentRepository) DeleteDocument(ctx context.Context, id string) error {
	query := `DELETE FROM documents WHERE id = $1`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete document: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("document with ID %s not found", id)
	}
	return nil
}

// IncrementViewCount increments the view_count for a document.
func (r *PGDocumentRepository) IncrementViewCount(ctx context.Context, id string) error {
	query := `UPDATE documents SET view_count = view_count + 1 WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

// IncrementDownloadCount increments the download_count for a document.
func (r *PGDocumentRepository) IncrementDownloadCount(ctx context.Context, id string) error {
	query := `UPDATE documents SET download_count = download_count + 1 WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
