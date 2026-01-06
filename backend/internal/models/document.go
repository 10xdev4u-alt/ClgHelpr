package models

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// Document represents a document stored in the vault.
type Document struct {
	ID              string         `json:"id"`
	UserID          string         `json:"userId"`
	SubjectID       sql.NullString `json:"subjectId"`
	
	Title           string         `json:"title"`
	Description     sql.NullString `json:"description"`
	
	DocumentType    string         `json:"documentType"` // 'notes', 'textbook', 'slides', etc.
	
	FileName        sql.NullString `json:"fileName"`
	FileType        sql.NullString `json:"fileType"`
	FileSize        sql.NullInt64  `json:"fileSize"`
	FileURL         string         `json:"fileUrl"`
	StorageKey      sql.NullString `json:"storageKey"`
	
	Folder          sql.NullString `json:"folder"`
	Tags            pgtype.FlatTextArray `json:"tags"` // TEXT[]
	
	IsPublic        bool           `json:"isPublic"`
	SharedWith      pgtype.FlatTextArray `json:"sharedWith"` // UUID[]
	
	ViewCount       int32          `json:"viewCount"`
	DownloadCount   int32          `json:"downloadCount"`
	LastAccessedAt  sql.NullTime   `json:"lastAccessedAt"`
	
	CreatedAt       time.Time      `json:"createdAt"`
	UpdatedAt       time.Time      `json:"updatedAt"`
}

// DocumentCreationInput defines the expected input for creating a document.
type DocumentCreationInput struct {
	SubjectID       *string  `json:"subjectId"`
	
	Title           string   `json:"title" validate:"required"`
	Description     *string  `json:"description"`
	
	DocumentType    string   `json:"documentType" validate:"required"`
	
	FileName        *string  `json:"fileName"`
	FileType        *string  `json:"fileType"`
	FileSize        *int64   `json:"fileSize"`
	FileURL         string   `json:"fileUrl" validate:"required"`
	StorageKey      *string  `json:"storageKey"`
	
	Folder          *string  `json:"folder"`
	Tags            []string `json:"tags"`
	
	IsPublic        *bool    `json:"isPublic"`
	SharedWith      []string `json:"sharedWith"`
}
