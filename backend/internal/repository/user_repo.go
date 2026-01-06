package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
)

// UserRepository defines the interface for user data operations.
type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
}

// PGUserRepository implements UserRepository for PostgreSQL.
type PGUserRepository struct {
	db *pgxpool.Pool
}

// NewPGUserRepository creates a new PostgreSQL user repository.
func NewPGUserRepository(db *pgxpool.Pool) *PGUserRepository {
	return &PGUserRepository{db: db}
}

// CreateUser inserts a new user into the database.
func (r *PGUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (
			id, email, password_hash, full_name, avatar_url, phone,
			register_number, department, year, semester, section, batch, is_hosteler,
			notification_preferences, theme, timezone,
			google_id, github_id,
			is_active, is_verified, last_login_at, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6,
			$7, $8, $9, $10, $11, $12, $13,
			$14, $15, $16,
			$17, $18,
			$19, $20, $21, $22, $23
		) RETURNING id, created_at, updated_at
	`

	user.ID = models.NewUUID() // Assuming models.NewUUID() exists to generate a UUID
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.IsActive = true
	user.IsVerified = false
	user.NotificationPreferences = map[string]interface{}{
		"push":            true,
		"email":           true,
		"morning_briefing": true,
	}
	user.Theme = "system"
	user.Timezone = "Asia/Kolkata"

	err := r.db.QueryRow(ctx, query,
		user.ID, user.Email, user.PasswordHash, user.FullName, user.AvatarURL, user.Phone,
		user.RegisterNumber, user.Department, user.Year, user.Semester, user.Section, user.Batch, user.IsHosteler,
		user.NotificationPreferences, user.Theme, user.Timezone,
		user.GoogleID, user.GithubID,
		user.IsActive, user.IsVerified, user.LastLoginAt, user.CreatedAt, user.UpdatedAt,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	return err
}

// GetUserByEmail retrieves a user by their email address.
func (r *PGUserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT
			id, email, password_hash, full_name, avatar_url, phone,
			register_number, department, year, semester, section, batch, is_hosteler,
			notification_preferences, theme, timezone,
			google_id, github_id,
			is_active, is_verified, last_login_at, created_at, updated_at
		FROM users
		WHERE email = $1
	`
	user := &models.User{}
	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.FullName, &user.AvatarURL, &user.Phone,
		&user.RegisterNumber, &user.Department, &user.Year, &user.Semester, &user.Section, &user.Batch, &user.IsHosteler,
		&user.NotificationPreferences, &user.Theme, &user.Timezone,
		&user.GoogleID, &user.GithubID,
		&user.IsActive, &user.IsVerified, &user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByID retrieves a user by their ID.
func (r *PGUserRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	query := `
		SELECT
			id, email, password_hash, full_name, avatar_url, phone,
			register_number, department, year, semester, section, batch, is_hosteler,
			notification_preferences, theme, timezone,
			google_id, github_id,
			is_active, is_verified, last_login_at, created_at, updated_at
		FROM users
		WHERE id = $1
	`
	user := &models.User{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.FullName, &user.AvatarURL, &user.Phone,
		&user.RegisterNumber, &user.Department, &user.Year, &user.Semester, &user.Section, &user.Batch, &user.IsHosteler,
		&user.NotificationPreferences, &user.Theme, &user.Timezone,
		&user.GoogleID, &user.GithubID,
		&user.IsActive, &user.IsVerified, &user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
