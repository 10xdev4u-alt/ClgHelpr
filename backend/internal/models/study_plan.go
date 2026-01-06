package models

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// StudyPlan represents a user's study plan.
type StudyPlan struct {
	ID        string         `json:"id"`
	UserID    string         `json:"userId"`
	
	Title     string         `json:"title"`
	PlanDate  time.Time      `json:"planDate"`
	
	PlanType  string         `json:"planType"` // 'daily', 'weekly', 'weekend', 'exam_prep', 'revision', 'custom'
	
	Status    string         `json:"status"` // 'planned', 'in_progress', 'completed', 'partial', 'skipped'
	
	Notes     sql.NullString `json:"notes"`
	
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}

// StudySession represents a single study session within a plan.
type StudySession struct {
	ID                     string         `json:"id"`
	UserID                 string         `json:"userId"`
	StudyPlanID            sql.NullString `json:"studyPlanId"`
	SubjectID              sql.NullString `json:"subjectId"`
	
	PlannedStartTime       sql.NullTime   `json:"plannedStartTime"`
	PlannedEndTime         sql.NullTime   `json:"plannedEndTime"`
	PlannedDurationMinutes sql.NullInt32  `json:"plannedDurationMinutes"`
	
	ActualStartTime        sql.NullTime   `json:"actualStartTime"`
	ActualEndTime          sql.NullTime   `json:"actualEndTime"`
	ActualDurationMinutes  sql.NullInt32  `json:"actualDurationMinutes"`
	
	SessionType            string         `json:"sessionType"` // 'study', 'revision', 'practice', 'assignment', 'lab_prep', 'exam_prep'
	TopicsToCover          pgtype.FlatTextArray `json:"topicsToCover"`
	TopicsCovered          pgtype.FlatTextArray `json:"topicsCovered"`
	
	Status                 string         `json:"status"` // 'planned', 'in_progress', 'completed', 'partial', 'skipped'
	CompletionPercentage   int32          `json:"completionPercentage"`
	
	ProductivityRating     sql.NullInt32  `json:"productivityRating"`
	Notes                  sql.NullString `json:"notes"`
	Blockers               sql.NullString `json:"blockers"`
	
	CreatedAt              time.Time      `json:"createdAt"`
	UpdatedAt              time.Time      `json:"updatedAt"`
}

// StudyPlanCreationInput defines the expected input for creating a study plan.
type StudyPlanCreationInput struct {
	Title    string   `json:"title" validate:"required"`
	PlanDate string   `json:"planDate" validate:"required"` // YYYY-MM-DD
	PlanType string   `json:"planType" validate:"required"`
	Notes    *string  `json:"notes"`
	Status   *string  `json:"status"`
}

// StudySessionCreationInput defines the expected input for creating a study session.
type StudySessionCreationInput struct {
	StudyPlanID            *string  `json:"studyPlanId"`
	SubjectID              *string  `json:"subjectId"`

	PlannedStartTime       *string  `json:"plannedStartTime"` // ISO string
	PlannedEndTime         *string  `json:"plannedEndTime"`   // ISO string
	PlannedDurationMinutes *int     `json:"plannedDurationMinutes"`

	ActualStartTime        *string  `json:"actualStartTime"` // ISO string
	ActualEndTime          *string  `json:"actualEndTime"`   // ISO string
	ActualDurationMinutes  *int     `json:"actualDurationMinutes"`

	SessionType            string   `json:"sessionType" validate:"required"`
	TopicsToCover          []string `json:"topicsToCover"`
	TopicsCovered          []string `json:"topicsCovered"`

	Status                 *string  `json:"status"` // defaults to 'planned'
	CompletionPercentage   *int     `json:"completionPercentage"` // defaults to 0

	ProductivityRating     *int     `json:"productivityRating"`
	Notes                  *string  `json:"notes"`
	Blockers               *string  `json:"blockers"`
}
