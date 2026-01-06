package models

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// LabRecord represents an academic lab record.
type LabRecord struct {
	ID                string         `json:"id"`
	UserID            string         `json:"userId"`
	SubjectID         sql.NullString `json:"subjectId"`
	
	ExperimentNumber  int32          `json:"experimentNumber"`
	Title             string         `json:"title"`
	
	LabDate           sql.NullTime   `json:"labDate"`
	RecordWrittenDate sql.NullTime   `json:"recordWrittenDate"`
	SubmittedDate     sql.NullTime   `json:"submittedDate"`
	
	Status            string         `json:"status"` // 'pending', 'practiced', 'written', 'printed', 'submitted', 'signed', 'returned'
	
	Aim               sql.NullString `json:"aim"`
	Algorithm         sql.NullString `json:"algorithm"`
	Code              sql.NullString `json:"code"`
	Output            sql.NullString `json:"output"`
	Observations      sql.NullString `json:"observations"`
	Result            sql.NullString `json:"result"`
	VivaQuestions     pgtype.FlatTextArray `json:"vivaQuestions"` // TEXT[]
	
	PrintRequired     bool           `json:"printRequired"`
	PagesToPrint      sql.NullInt32  `json:"pagesToPrint"`
	PrintedAt         sql.NullTime   `json:"printedAt"`
	
	Marks             sql.NullFloat64 `json:"marks"`
	StaffRemarks      sql.NullString `json:"staffRemarks"`
	
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
}

// LabRecordAttachment represents an attachment related to a lab record.
type LabRecordAttachment struct {
	ID              string         `json:"id"`
	LabRecordID     string         `json:"labRecordId"`
	
	FileName        string         `json:"fileName"`
	FileType        sql.NullString `json:"fileType"`
	FileURL         string         `json:"fileUrl"`
	StorageKey      sql.NullString `json:"storageKey"`
	
	AttachmentType  string         `json:"attachmentType"` // 'code_file', 'output_screenshot', 'record_pdf', 'signed_record', 'other'
	
	UploadedAt      time.Time      `json:"uploadedAt"`
}

// LabRecordCreationInput defines the expected input for creating a lab record.
type LabRecordCreationInput struct {
	SubjectID         *string   `json:"subjectId"` // Pointer for optional foreign key
	
	ExperimentNumber  int32     `json:"experimentNumber" validate:"required"`
	Title             string    `json:"title" validate:"required"`
	
	LabDate           *string   `json:"labDate"` // Date string (YYYY-MM-DD)
	RecordWrittenDate *string   `json:"recordWrittenDate"` // Date string (YYYY-MM-DD)
	SubmittedDate     *string   `json:"submittedDate"`     // Date string (YYYY-MM-DD)
	
	Status            *string   `json:"status"` // defaults to 'pending'
	
	Aim               *string   `json:"aim"`
	Algorithm         *string   `json:"algorithm"`
	Code              *string   `json:"code"`
	Output            *string   `json:"output"`
	Observations      *string   `json:"observations"`
	Result            *string   `json:"result"`
	VivaQuestions     []string  `json:"vivaQuestions"`
	
	PrintRequired     *bool     `json:"printRequired"` // defaults to true
	PagesToPrint      *int      `json:"pagesToPrint"`
	PrintedAt         *string   `json:"printedAt"`
	
	Marks             *float64  `json:"marks"`
	StaffRemarks      *string   `json:"staffRemarks"`
}
