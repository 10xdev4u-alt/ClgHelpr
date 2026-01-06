package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"github.com/jackc/pgx/v5" // Correct import for pgx.ErrNoRows
)

// AuthService defines the interface for authentication-related business logic.
type AuthService interface {
	RegisterUser(ctx context.Context, input *models.UserRegistrationInput) (*models.User, error)
}

// authService implements AuthService.
type authService struct {
	userRepo repository.UserRepository
}

// NewAuthService creates a new authentication service.
func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

// RegisterUser handles the user registration process.
func (s *authService) RegisterUser(ctx context.Context, input *models.UserRegistrationInput) (*models.User, error) {
	// Check if user with this email already exists
	existingUser, err := s.userRepo.GetUserByEmail(ctx, input.Email)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) { // Corrected check for no rows
		return nil, fmt.Errorf("failed to check for existing user: %w", err)
	}
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user model
	user := &models.User{
		Email:           input.Email,
		PasswordHash:    string(hashedPassword),
		FullName:        input.FullName,
		RegisterNumber:  &input.RegisterNumber, // Pointer to string
		Department:      input.Department,
		Year:            &input.Year, // Pointer to int
		Semester:        &input.Semester, // Pointer to int
		IsActive:        true,
		IsVerified:      false,
		// Default values for other fields will be set in the repository CreateUser method
	}

	// Save user to database
	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// For security, clear the password hash before returning the user object
	user.PasswordHash = ""

	return user, nil
}
