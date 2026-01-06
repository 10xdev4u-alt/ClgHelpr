package models

import (
	"time"
	// Ensure that 'github.com/google/uuid' is imported correctly
	// NewUUID() is called in user_repo.go, so this import is needed here.
	// We are also importing this for the UUID generation.
	_ "github.com/google/uuid"
)

// User represents a user in the system.
type User struct {
	ID                     string         `json:"id"`
	Email                  string         `json:"email"`
	PasswordHash           string         `json:"-"` // Omit password hash from JSON responses
	FullName               string         `json:"fullName"`
	AvatarURL              *string        `json:"avatarUrl"` // Use pointer for nullable fields
	Phone                  *string        `json:"phone"`

	// Academic Info
	RegisterNumber         *string        `json:"registerNumber"`
	Department             string         `json:"department"`
	Year                   *int           `json:"year"`
	Semester               *int           `json:"semester"`
	Section                *string        `json:"section"`
	Batch                  *string        `json:"batch"`
	IsHosteler             bool           `json:"isHosteler"`

	// Settings
	NotificationPreferences map[string]interface{} `json:"notificationPreferences"` // JSONB type
	Theme                  string         `json:"theme"`
	Timezone               string         `json:"timezone"`

	// OAuth
	GoogleID               *string        `json:"googleId"`
	GithubID               *string        `json:"githubId"`

	// Metadata
	IsActive               bool           `json:"isActive"`
	IsVerified             bool           `json:"isVerified"`
	LastLoginAt            *time.Time     `json:"lastLoginAt"`
	CreatedAt              time.Time      `json:"createdAt"`
	UpdatedAt              time.Time      `json:"updatedAt"`
}

// UserRegistrationInput defines the expected input for user registration.
type UserRegistrationInput struct {
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=8"`
	FullName     string `json:"fullName" validate:"required"`
	RegisterNumber string `json:"registerNumber" validate:"required"`
	Department   string `json:"department" validate:"required"`
	Year         int    `json:"year" validate:"required,min=1,max=4"`
	Semester     int    `json:"semester" validate:"required,min=1,max=8"`
}

