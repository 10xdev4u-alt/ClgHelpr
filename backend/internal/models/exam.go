package models

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// Exam represents an academic exam or test.
type Exam struct {
	ID                string         `json:"id"`
	UserID            string         `json:"userId"`
	SubjectID         sql.NullString `json:"subjectId"`
	
	Title             string         `json:"title"`
	ExamType          string         `json:"examType"` // 'cat1', 'fat', 'model', etc.
	
	ExamDate          time.Time      `json:"examDate"`
	StartTime         sql.NullTime   `json:"startTime"`
	EndTime           sql.NullTime   `json:"endTime"`
	DurationMinutes   sql.NullInt32  `json:"durationMinutes"`
	VenueID           sql.NullString `json:"venueId"`
	
	SyllabusUnits     pgtype.FlatTextArray `json:"syllabusUnits"` // TEXT[]
	SyllabusTopics    pgtype.FlatTextArray `json:"syllabusTopics"` // TEXT[]
	SyllabusNotes     sql.NullString       `json:"syllabusNotes"`
	
	MaxMarks          sql.NullFloat64      `json:"maxMarks"`
	ObtainedMarks     sql.NullFloat64      `json:"obtainedMarks"`
	Grade             sql.NullString       `json:"grade"`
	
	PrepStatus        string         `json:"prepStatus"` // 'not_started', 'in_progress', 'revision', 'ready'
	PrepNotes         sql.NullString `json:"prepNotes"`
	StudyHoursLogged  sql.NullFloat64 `json:"studyHoursLogged"`
	
	ReminderEnabled   bool           `json:"reminderEnabled"`
	
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
}

// ImportantQuestion represents an important question for an exam.
type ImportantQuestion struct {
	ID                string         `json:"id"`
	UserID            string         `json:"userId"`
	SubjectID         sql.NullString `json:"subjectId"`
	ExamID            sql.NullString `json:"examId"`
	
	QuestionText      string         `json:"questionText"`
	AnswerText        sql.NullString `json:"answerText"`
	Source            sql.NullString `json:"source"`
	
	Unit              sql.NullString `json:"unit"`
	Topic             sql.NullString `json:"topic"`
	Marks             sql.NullInt32  `json:"marks"`
	FrequencyCount    sql.NullInt32  `json:"frequencyCount"`
	
	IsPracticed       bool           `json:"isPracticed"`
	LastPracticedAt   sql.NullTime   `json:"lastPracticedAt"`
	ConfidenceLevel   sql.NullInt32  `json:"confidenceLevel"`
	
	Tags              pgtype.FlatTextArray `json:"tags"` // TEXT[]
	CreatedAt         time.Time      `json:"createdAt"`
}

// ExamCreationInput defines the expected input for creating an exam.
type ExamCreationInput struct {
	SubjectID         *string   `json:"subjectId"` // Pointer for optional foreign key
	VenueID           *string   `json:"venueId"` // Pointer for optional foreign key

	Title             string    `json:"title" validate:"required"`
	ExamType          string    `json:"examType" validate:"required"` // 'cat1', 'fat', etc.
	
	ExamDate          string    `json:"examDate" validate:"required"` // Date string (YYYY-MM-DD)
	StartTime         *string   `json:"startTime"` // Time string (HH:MM:SS)
	EndTime           *string   `json:"endTime"`   // Time string (HH:MM:SS)
	DurationMinutes   *int      `json:"durationMinutes"`
	
	SyllabusUnits     []string  `json:"syllabusUnits"`
	SyllabusTopics    []string  `json:"syllabusTopics"`
	SyllabusNotes     *string   `json:"syllabusNotes"`
	
	MaxMarks          *float64  `json:"maxMarks"`
	ObtainedMarks     *float64  `json:"obtainedMarks"`
	Grade             *string   `json:"grade"`
	
	PrepStatus        *string   `json:"prepStatus"` // defaults to 'not_started'
	PrepNotes         *string   `json:"prepNotes"`
	StudyHoursLogged  *float64  `json:"studyHoursLogged"`
	
	ReminderEnabled   *bool     `json:"reminderEnabled"` // defaults to true
}

// ImportantQuestionCreationInput defines input for creating an important question.
type ImportantQuestionCreationInput struct {
	SubjectID         *string   `json:"subjectId"`
	ExamID            *string   `json:"examId"`
	
	QuestionText      string    `json:"questionText" validate:"required"`
	AnswerText        *string   `json:"answerText"`
	Source            *string   `json:"source"`
	
	Unit              *string   `json:"unit"`
	Topic             *string   `json:"topic"`
	Marks             *int      `json:"marks"`
	FrequencyCount    *int      `json:"frequencyCount"`
	
	IsPracticed       *bool     `json:"isPracticed"`
	LastPracticedAt   *string   `json:"lastPracticedAt"`
	ConfidenceLevel   *int      `json:"confidenceLevel"`
	
	Tags              []string  `json:"tags"`
}
