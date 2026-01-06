package models

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// ActivityLog represents a user activity entry.
type ActivityLog struct {
	ID           string         `json:"id"`
	UserID       string         `json:"userId"`
	
	ActivityType string         `json:"activityType"` // 'assignment_created', 'exam_updated', etc.
	Description  sql.NullString `json:"description"`
	
	EntityType   sql.NullString `json:"entityType"` // 'assignment', 'exam', 'study_session', etc.
	EntityID     sql.NullString `json:"entityId"`
	
	Metadata     pgtype.JSONB   `json:"metadata"` // Additional JSON data
	
	IPAddress    sql.NullString `json:"ipAddress"` // INET type might need specific pgtype handling, using string for simplicity
	UserAgent    sql.NullString `json:"userAgent"`
	
	CreatedAt    time.Time      `json:"createdAt"`
}

// DailyStats represents pre-aggregated daily statistics for a user.
type DailyStats struct {
	ID                   string    `json:"id"`
	UserID               string    `json:"userId"`
	StatDate             time.Time `json:"statDate"` // DATE type
	
	StudyMinutes         int32     `json:"studyMinutes"`
	SessionsCompleted    int32     `json:"sessionsCompleted"`
	TopicsCovered        int32     `json:"topicsCovered"`
	
	AssignmentsCompleted int32     `json:"assignmentsCompleted"`
	AssignmentsAdded     int32     `json:"assignmentsAdded"`
	
	ClassesAttended      int32     `json:"classesAttended"`
	TotalClasses         int32     `json:"totalClasses"`
	
	XPEarned             int32     `json:"xpEarned"`
}

// ActivityLogCreationInput defines the expected input for creating an activity log.
type ActivityLogCreationInput struct {
	ActivityType string      `json:"activityType" validate:"required"`
	Description  *string     `json:"description"`
	EntityType   *string     `json:"entityType"`
	EntityID     *string     `json:"entityId"`
	Metadata     interface{} `json:"metadata"` // Can be any JSON object
	IPAddress    *string     `json:"ipAddress"`
	UserAgent    *string     `json:"userAgent"`
}

// DailyStatsCreationInput defines the expected input for creating/updating daily stats.
type DailyStatsUpdateInput struct {
	StatDate             string `json:"statDate" validate:"required"` // YYYY-MM-DD
	StudyMinutes         *int32 `json:"studyMinutes"`
	SessionsCompleted    *int32 `json:"sessionsCompleted"`
	TopicsCovered        *int32 `json:"topicsCovered"`
	AssignmentsCompleted *int32 `json:"assignmentsCompleted"`
	AssignmentsAdded     *int32 `json:"assignmentsAdded"`
	ClassesAttended      *int32 `json:"classesAttended"`
	TotalClasses         *int32 `json:"totalClasses"`
	XPEarned             *int32 `json:"xpEarned"`
}
