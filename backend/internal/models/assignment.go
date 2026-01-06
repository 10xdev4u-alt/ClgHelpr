package models

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// Assignment represents an assignment or task.
type Assignment struct {
	ID                 string         `json:"id"`
	UserID             string         `json:"userId"`
	SubjectID          sql.NullString `json:"subjectId"`
	StaffID            sql.NullString `json:"staffId"`

	Title              string         `json:"title"`
	Description        sql.NullString `json:"description"`
	Instructions       sql.NullString `json:"instructions"`

	AssignmentType     string         `json:"assignmentType"` // e.g., 'assignment', 'lab_record', 'project'

	AssignedDate       sql.NullTime   `json:"assignedDate"`
	DueDate            time.Time      `json:"dueDate"`
	SubmittedAt        sql.NullTime   `json:"submittedAt"`

	Status             string         `json:"status"` // 'pending', 'in_progress', 'completed', 'submitted', 'graded', 'overdue'

	MaxMarks           sql.NullFloat64 `json:"maxMarks"`
	ObtainedMarks      sql.NullFloat64 `json:"obtainedMarks"`
	Feedback           sql.NullString  `json:"feedback"`

	Priority           string         `json:"priority"` // 'low', 'medium', 'high', 'urgent'
	EstimatedHours     sql.NullFloat64 `json:"estimatedHours"`
	ActualHours        sql.NullFloat64 `json:"actualHours"`

	ReminderEnabled    bool           `json:"reminderEnabled"`
	ReminderBeforeHours sql.NullInt32 `json:"reminderBeforeHours"`
	LastRemindedAt     sql.NullTime   `json:"lastRemindedAt"`

	Tags               pgtype.FlatTextArray `json:"tags"` // TEXT[]
	IsRecurring        bool                 `json:"isRecurring"`
	RecurrencePattern  sql.NullString       `json:"recurrencePattern"`

	CreatedAt          time.Time      `json:"createdAt"`
	UpdatedAt          time.Time      `json:"updatedAt"`
}

// AssignmentAttachment represents an attachment related to an assignment.
type AssignmentAttachment struct {
	ID              string         `json:"id"`
	AssignmentID    string         `json:"assignmentId"`

	FileName        string         `json:"fileName"`
	FileType        sql.NullString `json:"fileType"`
	FileSize        sql.NullInt64  `json:"fileSize"`
	FileURL         string         `json:"fileUrl"`
	StorageKey      sql.NullString `json:"storageKey"`

	AttachmentType  string         `json:"attachmentType"` // 'reference', 'submission', 'feedback'

	UploadedAt      time.Time      `json:"uploadedAt"`
}

// AssignmentCreationInput defines the expected input for creating an assignment.
type AssignmentCreationInput struct {
	SubjectID          *string  `json:"subjectId"` // Pointers for optional foreign keys
	StaffID            *string  `json:"staffId"`

	Title              string   `json:"title" validate:"required"`
	Description        *string  `json:"description"`
	Instructions       *string  `json:"instructions"`

	AssignmentType     string   `json:"assignmentType" validate:"required"`

	AssignedDate       *string  `json:"assignedDate"` // Date string (YYYY-MM-DD)
	DueDate            string   `json:"dueDate" validate:"required"` // ISO timestamp or date string

	Status             *string  `json:"status"` // defaults to 'pending'

	MaxMarks           *float64 `json:"maxMarks"`
	ObtainedMarks      *float64 `json:"obtainedMarks"`
	Feedback           *string  `json:"feedback"`

	Priority           *string  `json:"priority"` // defaults to 'medium'
	EstimatedHours     *float64 `json:"estimatedHours"`
	ActualHours        *float64 `json:"actualHours"`

	ReminderEnabled    *bool    `json:"reminderEnabled"` // defaults to true
	ReminderBeforeHours *int    `json:"reminderBeforeHours"`

	Tags               []string `json:"tags"`
	IsRecurring        *bool    `json:"isRecurring"`
	RecurrencePattern  *string  `json:"recurrencePattern"`
}