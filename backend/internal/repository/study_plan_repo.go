package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
)

// --- StudyPlan Repository ---

// StudyPlanRepository defines the interface for study plan data operations.
type StudyPlanRepository interface {
	CreateStudyPlan(ctx context.Context, plan *models.StudyPlan) error
	GetStudyPlanByID(ctx context.Context, id string) (*models.StudyPlan, error)
	GetStudyPlansByUserID(ctx context.Context, userID string) ([]models.StudyPlan, error)
	GetStudyPlansByUserIDAndDate(ctx context.Context, userID string, date time.Time) ([]models.StudyPlan, error)
	UpdateStudyPlan(ctx context.Context, plan *models.StudyPlan) error
	DeleteStudyPlan(ctx context.Context, id string) error
}

// PGStudyPlanRepository implements StudyPlanRepository for PostgreSQL.
type PGStudyPlanRepository struct {
	db *pgxpool.Pool
}

// NewPGStudyPlanRepository creates a new PostgreSQL study plan repository.
func NewPGStudyPlanRepository(db *pgxpool.Pool) *PGStudyPlanRepository {
	return &PGStudyPlanRepository{db: db}
}

// CreateStudyPlan inserts a new study plan into the database.
func (r *PGStudyPlanRepository) CreateStudyPlan(ctx context.Context, plan *models.StudyPlan) error {
	query := `
		INSERT INTO study_plans (
			id, user_id, title, plan_date, plan_type, status, notes, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		) RETURNING id, created_at, updated_at
	`
	plan.ID = models.NewUUID()
	plan.CreatedAt = time.Now()
	plan.UpdatedAt = time.Now()

	_, err := r.db.Exec(ctx, query,
		plan.ID, plan.UserID, plan.Title, plan.PlanDate, plan.PlanType, plan.Status, plan.Notes, plan.CreatedAt, plan.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create study plan: %w", err)
	}
	return nil
}

// GetStudyPlanByID retrieves a study plan by its ID.
func (r *PGStudyPlanRepository) GetStudyPlanByID(ctx context.Context, id string) (*models.StudyPlan, error) {
	plan := &models.StudyPlan{}
	query := `
		SELECT
			id, user_id, title, plan_date, plan_type, status, notes, created_at, updated_at
		FROM study_plans
		WHERE id = $1
	`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&plan.ID, &plan.UserID, &plan.Title, &plan.PlanDate, &plan.PlanType, &plan.Status, &plan.Notes, &plan.CreatedAt, &plan.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get study plan by ID: %w", err)
	}
	return plan, nil
}

// GetStudyPlansByUserID retrieves all study plans for a given user.
func (r *PGStudyPlanRepository) GetStudyPlansByUserID(ctx context.Context, userID string) ([]models.StudyPlan, error) {
	var plans []models.StudyPlan
	query := `
		SELECT
			id, user_id, title, plan_date, plan_type, status, notes, created_at, updated_at
		FROM study_plans
		WHERE user_id = $1
		ORDER BY plan_date DESC, created_at DESC
	`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get study plans by user ID: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		plan := models.StudyPlan{}
		err := rows.Scan(
			&plan.ID, &plan.UserID, &plan.Title, &plan.PlanDate, &plan.PlanType, &plan.Status, &plan.Notes, &plan.CreatedAt, &plan.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan study plan row: %w", err)
		}
		plans = append(plans, plan)
	}
	return plans, nil
}

// GetStudyPlansByUserIDAndDate retrieves all study plans for a given user on a specific date.
func (r *PGStudyPlanRepository) GetStudyPlansByUserIDAndDate(ctx context.Context, userID string, date time.Time) ([]models.StudyPlan, error) {
	var plans []models.StudyPlan
	query := `
		SELECT
			id, user_id, title, plan_date, plan_type, status, notes, created_at, updated_at
		FROM study_plans
		WHERE user_id = $1 AND plan_date = $2
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, query, userID, date)
	if err != nil {
		return nil, fmt.Errorf("failed to get study plans by user ID and date: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		plan := models.StudyPlan{}
		err := rows.Scan(
			&plan.ID, &plan.UserID, &plan.Title, &plan.PlanDate, &plan.PlanType, &plan.Status, &plan.Notes, &plan.CreatedAt, &plan.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan study plan row: %w", err)
		}
		plans = append(plans, plan)
	}
	return plans, nil
}

// UpdateStudyPlan updates an existing study plan in the database.
func (r *PGStudyPlanRepository) UpdateStudyPlan(ctx context.Context, plan *models.StudyPlan) error {
	query := `
		UPDATE study_plans SET
			title = $1, plan_date = $2, plan_type = $3, status = $4, notes = $5, updated_at = $6
		WHERE id = $7 AND user_id = $8
	`
	plan.UpdatedAt = time.Now()

	cmdTag, err := r.db.Exec(ctx, query,
		plan.Title, plan.PlanDate, plan.PlanType, plan.Status, plan.Notes, plan.UpdatedAt,
		plan.ID, plan.UserID,
	)
	if err != nil {
		return fmt.Errorf("failed to update study plan: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("study plan with ID %s not found or not owned by user", plan.ID)
	}
	return nil
}

// DeleteStudyPlan deletes a study plan from the database.
func (r *PGStudyPlanRepository) DeleteStudyPlan(ctx context.Context, id string) error {
	query := `DELETE FROM study_plans WHERE id = $1`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete study plan: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("study plan with ID %s not found", id)
	}
	return nil
}

// --- StudySession Repository ---

// StudySessionRepository defines the interface for study session data operations.
type StudySessionRepository interface {
	CreateStudySession(ctx context.Context, session *models.StudySession) error
	GetStudySessionByID(ctx context.Context, id string) (*models.StudySession, error)
	GetStudySessionsByUserID(ctx context.Context, userID string) ([]models.StudySession, error)
	GetStudySessionsByStudyPlanID(ctx context.Context, studyPlanID string) ([]models.StudySession, error)
	UpdateStudySession(ctx context.Context, session *models.StudySession) error
	DeleteStudySession(ctx context.Context, id string) error
}

// PGStudySessionRepository implements StudySessionRepository for PostgreSQL.
type PGStudySessionRepository struct {
	db *pgxpool.Pool
}

// NewPGStudySessionRepository creates a new PostgreSQL study session repository.
func NewPGStudySessionRepository(db *pgxpool.Pool) *PGStudySessionRepository {
	return &PGStudySessionRepository{db: db}
}

// CreateStudySession inserts a new study session into the database.
func (r *PGStudySessionRepository) CreateStudySession(ctx context.Context, session *models.StudySession) error {
	query := `
		INSERT INTO study_sessions (
			id, user_id, study_plan_id, subject_id, planned_start_time, planned_end_time,
			planned_duration_minutes, actual_start_time, actual_end_time, actual_duration_minutes,
			session_type, topics_to_cover, topics_covered, status, completion_percentage,
			productivity_rating, notes, blockers, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20
		) RETURNING id, created_at, updated_at
	`
	session.ID = models.NewUUID()
	session.CreatedAt = time.Now()
	session.UpdatedAt = time.Now()

	_, err := r.db.Exec(ctx, query,
		session.ID, session.UserID, session.StudyPlanID, session.SubjectID, session.PlannedStartTime, session.PlannedEndTime,
		session.PlannedDurationMinutes, session.ActualStartTime, session.ActualEndTime, session.ActualDurationMinutes,
		session.SessionType, session.TopicsToCover, session.TopicsCovered, session.Status, session.CompletionPercentage,
		session.ProductivityRating, session.Notes, session.Blockers, session.CreatedAt, session.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create study session: %w", err)
	}
	return nil
}

// GetStudySessionByID retrieves a study session by its ID.
func (r *PGStudySessionRepository) GetStudySessionByID(ctx context.Context, id string) (*models.StudySession, error) {
	session := &models.StudySession{}
	query := `
		SELECT
			id, user_id, study_plan_id, subject_id, planned_start_time, planned_end_time,
			planned_duration_minutes, actual_start_time, actual_end_time, actual_duration_minutes,
			session_type, topics_to_cover, topics_covered, status, completion_percentage,
			productivity_rating, notes, blockers, created_at, updated_at
		FROM study_sessions
		WHERE id = $1
	`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&session.ID, &session.UserID, &session.StudyPlanID, &session.SubjectID, &session.PlannedStartTime, &session.PlannedEndTime,
		&session.PlannedDurationMinutes, &session.ActualStartTime, &session.ActualEndTime, &session.ActualDurationMinutes,
		&session.SessionType, &session.TopicsToCover, &session.TopicsCovered, &session.Status, &session.CompletionPercentage,
		&session.ProductivityRating, &session.Notes, &session.Blockers, &session.CreatedAt, &session.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get study session by ID: %w", err)
	}
	return session, nil
}

// GetStudySessionsByUserID retrieves all study sessions for a given user.
func (r *PGStudySessionRepository) GetStudySessionsByUserID(ctx context.Context, userID string) ([]models.StudySession, error) {
	var sessions []models.StudySession
	query := `
		SELECT
			id, user_id, study_plan_id, subject_id, planned_start_time, planned_end_time,
			planned_duration_minutes, actual_start_time, actual_end_time, actual_duration_minutes,
			session_type, topics_to_cover, topics_covered, status, completion_percentage,
			productivity_rating, notes, blockers, created_at, updated_at
		FROM study_sessions
		WHERE user_id = $1
		ORDER BY planned_start_time DESC, created_at DESC
	`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get study sessions by user ID: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		session := models.StudySession{}
		err := rows.Scan(
			&session.ID, &session.UserID, &session.StudyPlanID, &session.SubjectID, &session.PlannedStartTime, &session.PlannedEndTime,
			&session.PlannedDurationMinutes, &session.ActualStartTime, &session.ActualEndTime, &session.ActualDurationMinutes,
			&session.SessionType, &session.TopicsToCover, &session.TopicsCovered, &session.Status, &session.CompletionPercentage,
			&session.ProductivityRating, &session.Notes, &session.Blockers, &session.CreatedAt, &session.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan study session row: %w", err)
		}
		sessions = append(sessions, session)
	}
	return sessions, nil
}

// GetStudySessionsByStudyPlanID retrieves all study sessions for a given study plan.
func (r *PGStudySessionRepository) GetStudySessionsByStudyPlanID(ctx context.Context, studyPlanID string) ([]models.StudySession, error) {
	var sessions []models.StudySession
	query := `
		SELECT
			id, user_id, study_plan_id, subject_id, planned_start_time, planned_end_time,
			planned_duration_minutes, actual_start_time, actual_end_time, actual_duration_minutes,
			session_type, topics_to_cover, topics_covered, status, completion_percentage,
			productivity_rating, notes, blockers, created_at, updated_at
		FROM study_sessions
		WHERE study_plan_id = $1
		ORDER BY planned_start_time ASC, created_at ASC
	`
	rows, err := r.db.Query(ctx, query, studyPlanID)
	if err != nil {
		return nil, fmt.Errorf("failed to get study sessions by study plan ID: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		session := models.StudySession{}
		err := rows.Scan(
			&session.ID, &session.UserID, &session.StudyPlanID, &session.SubjectID, &session.PlannedStartTime, &session.PlannedEndTime,
			&session.PlannedDurationMinutes, &session.ActualStartTime, &session.ActualEndTime, &session.ActualDurationMinutes,
			&session.SessionType, &session.TopicsToCover, &session.TopicsCovered, &session.Status, &session.CompletionPercentage,
			&session.ProductivityRating, &session.Notes, &session.Blockers, &session.CreatedAt, &session.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan study session row: %w", err)
		}
		sessions = append(sessions, session)
	}
	return sessions, nil
}

// UpdateStudySession updates an existing study session in the database.
func (r *PGStudySessionRepository) UpdateStudySession(ctx context.Context, session *models.StudySession) error {
	query := `
		UPDATE study_sessions SET
			study_plan_id = $1, subject_id = $2, planned_start_time = $3, planned_end_time = $4,
			planned_duration_minutes = $5, actual_start_time = $6, actual_end_time = $7,
			actual_duration_minutes = $8, session_type = $9, topics_to_cover = $10,
			topics_covered = $11, status = $12, completion_percentage = $13,
			productivity_rating = $14, notes = $15, blockers = $16, updated_at = $17
		WHERE id = $18 AND user_id = $19
	`
	session.UpdatedAt = time.Now()

	cmdTag, err := r.db.Exec(ctx, query,
		session.StudyPlanID, session.SubjectID, session.PlannedStartTime, session.PlannedEndTime,
		session.PlannedDurationMinutes, session.ActualStartTime, session.ActualEndTime,
		session.ActualDurationMinutes, session.SessionType, session.TopicsToCover,
		session.TopicsCovered, session.Status, session.CompletionPercentage,
		session.ProductivityRating, session.Notes, session.Blockers, session.UpdatedAt,
		session.ID, session.UserID,
	)
	if err != nil {
		return fmt.Errorf("failed to update study session: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("study session with ID %s not found or not owned by user", session.ID)
	}
	return nil
}

// DeleteStudySession deletes a study session from the database.
func (r *PGStudySessionRepository) DeleteStudySession(ctx context.Context, id string) error {
	query := `DELETE FROM study_sessions WHERE id = $1`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete study session: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("study session with ID %s not found", id)
	}
	return nil
}
