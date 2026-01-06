package services

import (
	"context"
	"fmt"
	"time"

	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/repository"
)

// DocumentService defines the interface for document-related business logic.
type DocumentService interface {
	CreateDocument(ctx context.Context, userID string, input *models.DocumentCreationInput) (*models.Document, error)
	GetDocumentByID(ctx context.Context, id string) (*models.Document, error)
	GetDocumentsByUserID(ctx context.Context, userID string) ([]models.Document, error)
	GetDocumentsByUserIDAndSubjectID(ctx context.Context, userID, subjectID string) ([]models.Document, error)
	UpdateDocument(ctx context.Context, userID string, id string, input *models.DocumentCreationInput) (*models.Document, error)
	DeleteDocument(ctx context.Context, id string) error
	IncrementViewCount(ctx context.Context, id string) error
	IncrementDownloadCount(ctx context.Context, id string) error
}

// documentService implements DocumentService.
type documentService struct {
	documentRepo repository.DocumentRepository
}

// NewDocumentService creates a new document service.
func NewDocumentService(documentRepo repository.DocumentRepository) DocumentService {
	return &documentService{documentRepo: documentRepo}
}

// CreateDocument creates a new document for a user.
func (s *documentService) CreateDocument(ctx context.Context, userID string, input *models.DocumentCreationInput) (*models.Document, error) {
	document := &models.Document{
		UserID:       userID,
		SubjectID:    sql.NullString{String: *input.SubjectID, Valid: input.SubjectID != nil},
		Title:        input.Title,
		Description:  sql.NullString{String: *input.Description, Valid: input.Description != nil},
		DocumentType: input.DocumentType,
		FileName:     sql.NullString{String: *input.FileName, Valid: input.FileName != nil},
		FileType:     sql.NullString{String: *input.FileType, Valid: input.FileType != nil},
		FileSize:     sql.NullInt64{Int64: *input.FileSize, Valid: input.FileSize != nil},
		FileURL:      input.FileURL,
		StorageKey:   sql.NullString{String: *input.StorageKey, Valid: input.StorageKey != nil},
		Folder:       sql.NullString{String: *input.Folder, Valid: input.Folder != nil},
		Tags:         input.Tags,
		IsPublic:     false, // Default
		SharedWith:   input.SharedWith,
	}

	if input.IsPublic != nil {
		document.IsPublic = *input.IsPublic
	}

	if err := s.documentRepo.CreateDocument(ctx, document); err != nil {
		return nil, fmt.Errorf("failed to create document: %w", err)
	}
	return document, nil
}

// GetDocumentByID retrieves a single document.
func (s *documentService) GetDocumentByID(ctx context.Context, id string) (*models.Document, error) {
	return s.documentRepo.GetDocumentByID(ctx, id)
}

// GetDocumentsByUserID retrieves all documents for a user.
func (s *documentService) GetDocumentsByUserID(ctx context.Context, userID string) ([]models.Document, error) {
	return s.documentRepo.GetDocumentsByUserID(ctx, userID)
}

// GetDocumentsByUserIDAndSubjectID retrieves all documents for a user and subject.
func (s *documentService) GetDocumentsByUserIDAndSubjectID(ctx context.Context, userID, subjectID string) ([]models.Document, error) {
	return s.documentRepo.GetDocumentsByUserIDAndSubjectID(ctx, userID, subjectID)
}

// UpdateDocument updates an existing document.
func (s *documentService) UpdateDocument(ctx context.Context, userID string, id string, input *models.DocumentCreationInput) (*models.Document, error) {
	existingDocument, err := s.documentRepo.GetDocumentByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("document not found: %w", err)
	}
	if existingDocument.UserID != userID {
		return nil, fmt.Errorf("document does not belong to user")
	}

	if input.SubjectID != nil {
		existingDocument.SubjectID = sql.NullString{String: *input.SubjectID, Valid: true}
	} else {
		existingDocument.SubjectID = sql.NullString{Valid: false}
	}
	if input.Title != "" {
		existingDocument.Title = input.Title
	}
	if input.Description != nil {
		existingDocument.Description = sql.NullString{String: *input.Description, Valid: true}
	} else {
		existingDocument.Description = sql.NullString{Valid: false}
	}
	if input.DocumentType != "" {
		existingDocument.DocumentType = input.DocumentType
	}
	if input.FileName != nil {
		existingDocument.FileName = sql.NullString{String: *input.FileName, Valid: true}
	} else {
		existingDocument.FileName = sql.NullString{Valid: false}
	}
	if input.FileType != nil {
		existingDocument.FileType = sql.NullString{String: *input.FileType, Valid: true}
	} else {
		existingDocument.FileType = sql.NullString{Valid: false}
	}
	if input.FileSize != nil {
		existingDocument.FileSize = sql.NullInt64{Int64: *input.FileSize, Valid: true}
	} else {
		existingDocument.FileSize = sql.NullInt64{Valid: false}
	}
	if input.FileURL != "" {
		existingDocument.FileURL = input.FileURL
	}
	if input.StorageKey != nil {
		existingDocument.StorageKey = sql.NullString{String: *input.StorageKey, Valid: true}
	} else {
		existingDocument.StorageKey = sql.NullString{Valid: false}
	}
	if input.Folder != nil {
		existingDocument.Folder = sql.NullString{String: *input.Folder, Valid: true}
	} else {
		existingDocument.Folder = sql.NullString{Valid: false}
	}
	if input.Tags != nil {
		existingDocument.Tags = input.Tags
	}
	if input.IsPublic != nil {
		existingDocument.IsPublic = *input.IsPublic
	}
	if input.SharedWith != nil {
		existingDocument.SharedWith = input.SharedWith
	}

	if err := s.documentRepo.UpdateDocument(ctx, existingDocument); err != nil {
		return nil, fmt.Errorf("failed to update document: %w", err)
	}
	return existingDocument, nil
}

// DeleteDocument deletes a document.
func (s *documentService) DeleteDocument(ctx context.Context, id string) error {
	return s.documentRepo.DeleteDocument(ctx, id)
}

// IncrementViewCount increments the view_count for a document.
func (s *documentService) IncrementViewCount(ctx context.Context, id string) error {
	return s.documentRepo.IncrementViewCount(ctx, id)
}

// IncrementDownloadCount increments the download_count for a document.
func (s *documentService) IncrementDownloadCount(ctx context.Context, id string) error {
	return s.documentRepo.IncrementDownloadCount(ctx, id)
}
