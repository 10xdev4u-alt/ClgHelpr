package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
)

// --- ActivityLog Repository ---

// ActivityLogRepository defines the interface for activity log data operations.
type ActivityLogRepository interface {
	CreateActivityLog(ctx context.Context, log *models.ActivityLog) error
	GetActivityLogsByUserID(ctx context.Context, userID string) ([]models.ActivityLog, error)
	GetActivityLogsByUserIDAndType(ctx context.Context, userID, activityType string) ([]models.ActivityLog, error)
}

// PGActivityLogRepository implements ActivityLogRepository for PostgreSQL.
type PGActivityLogRepository struct {
	db *pgxpool.Pool
}

// NewPGActivityLogRepository creates a new PostgreSQL activity log repository.
func NewPGActivityLogRepository(db *pgxpool.Pool) *PGActivityLogRepository {
	return &PGActivityLogRepository{db: db}
}

// CreateActivityLog inserts a new activity log into the database.
func (r *PGActivityLogRepository) CreateActivityLog(ctx context.Context, log *models.ActivityLog) error {
	query := `
		INSERT INTO activity_logs (
			id, user_id, activity_type, description, entity_type, entity_id,
			metadata, ip_address, user_agent, created_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
		) RETURNING id, created_at
	`
	log.ID = models.NewUUID()
	log.CreatedAt = time.Now()

	_, err := r.db.Exec(ctx, query,
		log.ID, log.UserID, log.ActivityType, log.Description, log.EntityType, log.EntityID,
		log.Metadata, log.IPAddress, log.UserAgent, log.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create activity log: %w", err)
	}
	return nil
}

// GetActivityLogsByUserID retrieves all activity logs for a given user.
func (r *PGActivityLogRepository) GetActivityLogsByUserID(ctx context.Context, userID string) ([]models.ActivityLog, error) {
	var logs []models.ActivityLog
	query := `
		SELECT
			id, user_id, activity_type, description, entity_type, entity_id,
			metadata, ip_address, user_agent, created_at
		FROM activity_logs
		WHERE user_id = $1
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get activity logs by user ID: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		log := models.ActivityLog{}
		err := rows.Scan(
			&log.ID, &log.UserID, &log.ActivityType, &log.Description, &log.EntityType, &log.EntityID,
			&log.Metadata, &log.IPAddress, &log.UserAgent, &log.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan activity log row: %w", err)
		}
		logs = append(logs, log)
	}
	return logs, nil
}

// GetActivityLogsByUserIDAndType retrieves activity logs for a given user and activity type.
func (r *PGActivityLogRepository) GetActivityLogsByUserIDAndType(ctx context.Context, userID, activityType string) ([]models.ActivityLog, error) {
	var logs []models.ActivityLog
	query := `
		SELECT
			id, user_id, activity_type, description, entity_type, entity_id,
			metadata, ip_address, user_agent, created_at
		FROM activity_logs
		WHERE user_id = $1 AND activity_type = $2
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, query, userID, activityType)
	if err != nil {
		return nil, fmt.Errorf("failed to get activity logs by user ID and type: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		log := models.ActivityLog{}
		err := rows.Scan(
			&log.ID, &log.UserID, &log.ActivityType, &log.Description, &log.EntityType, &log.EntityID,
			&log.Metadata, &log.IPAddress, &log.UserAgent, &log.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan activity log row: %w", err)
		}
		logs = append(logs, log)
	}
	return logs, nil
}

// --- DailyStats Repository ---

// DailyStatsRepository defines the interface for daily stats data operations.
type DailyStatsRepository interface {
	GetDailyStatsByUserIDAndDate(ctx context.Context, userID string, date time.Time) (*models.DailyStats, error)
	GetDailyStatsByUserID(ctx context.Context, userID string) ([]models.DailyStats, error)
	UpsertDailyStats(ctx context.Context, stats *models.DailyStats) error
}

// PGDailyStatsRepository implements DailyStatsRepository for PostgreSQL.
type PGDailyStatsRepository struct {
	db *pgxpool.Pool
}

// NewPGDailyStatsRepository creates a new PostgreSQL daily stats repository.
func NewPGDailyStatsRepository(db *pgxpool.Pool) *PGDailyStatsRepository {
	return &PGDailyStatsRepository{db: db}
}

// GetDailyStatsByUserIDAndDate retrieves daily stats for a given user and date.
func (r *PGDailyStatsRepository) GetDailyStatsByUserIDAndDate(ctx context.Context, userID string, date time.Time) (*models.DailyStats, error) {
	stats := &models.DailyStats{}
	query := `
		SELECT
			id, user_id, stat_date, study_minutes, sessions_completed, topics_covered,
			assignments_completed, assignments_added, classes_attended, total_classes, xp_earned
		FROM daily_stats
		WHERE user_id = $1 AND stat_date = $2
	`
	err := r.db.QueryRow(ctx, query, userID, date).Scan(
		&stats.ID, &stats.UserID, &stats.StatDate, &stats.StudyMinutes, &stats.SessionsCompleted, &stats.TopicsCovered,
		&stats.AssignmentsCompleted, &stats.AssignmentsAdded, &stats.ClassesAttended, &stats.TotalClasses, &stats.XPEarned,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get daily stats by user ID and date: %w", err)
	}
	return stats, nil
}

// GetDailyStatsByUserID retrieves all daily stats for a given user.
func (r *PGDailyStatsRepository) GetDailyStatsByUserID(ctx context.Context, userID string) ([]models.DailyStats, error) {
	var statsList []models.DailyStats
	query := `
		SELECT
			id, user_id, stat_date, study_minutes, sessions_completed, topics_covered,
			assignments_completed, assignments_added, classes_attended, total_classes, xp_earned
		FROM daily_stats
		WHERE user_id = $1
		ORDER BY stat_date DESC
	`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get daily stats by user ID: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		stats := models.DailyStats{}
		err := rows.Scan(
			&stats.ID, &stats.UserID, &stats.StatDate, &stats.StudyMinutes, &stats.SessionsCompleted, &stats.TopicsCovered,
			&stats.AssignmentsCompleted, &stats.AssignmentsAdded, &stats.ClassesAttended, &stats.TotalClasses, &stats.XPEarned,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan daily stats row: %w", err)
		}
		statsList = append(statsList, stats)
	}
	return statsList, nil
}

// UpsertDailyStats inserts or updates daily stats for a user on a specific date.
func (r *PGDailyStatsRepository) UpsertDailyStats(ctx context.Context, stats *models.DailyStats) error {
	query := `
		INSERT INTO daily_stats (
			id, user_id, stat_date, study_minutes, sessions_completed, topics_covered,
			assignments_completed, assignments_added, classes_attended, total_classes, xp_earned
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
		) ON CONFLICT (user_id, stat_date) DO UPDATE SET
			study_minutes = EXCLUDED.study_minutes,
			sessions_completed = EXCLUDED.sessions_completed,
			topics_covered = EXCLUDED.topics_covered,
			assignments_completed = EXCLUDED.assignments_completed,
			assignments_added = EXCLUDED.assignments_added,
			classes_attended = EXCLUDED.classes_attended,
			total_classes = EXCLUDED.total_classes,
			xp_earned = EXCLUDED.xp_earned,
			id = EXCLUDED.id -- Update ID if it was generated
		RETURNING id
	`
	// If stats.ID is empty, generate a new UUID
	if stats.ID == "" {
		stats.ID = models.NewUUID()
	}

	_, err := r.db.Exec(ctx, query,
		stats.ID, stats.UserID, stats.StatDate, stats.StudyMinutes, stats.SessionsCompleted, stats.TopicsCovered,
		stats.AssignmentsCompleted, stats.AssignmentsAdded, stats.ClassesAttended, stats.TotalClasses, stats.XPEarned,
	)
	if err != nil {
		return fmt.Errorf("failed to upsert daily stats: %w", err)
	}
	return nil
}
