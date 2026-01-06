package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
)

// AuthService defines the interface for authentication-related business logic.
type AuthService interface {
	RegisterUser(ctx context.Context, input *models.UserRegistrationInput) (*models.User, error)
	Login(ctx context.Context, input *models.LoginUserInput) (string, error)
	GetUserProfile(ctx context.Context, userID string) (*models.User, error)
}

// authService implements AuthService.
type authService struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

// NewAuthService creates a new authentication service.
func NewAuthService(userRepo repository.UserRepository, jwtSecret string) AuthService {
	return &authService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

// RegisterUser handles the user registration process.
func (s *authService) RegisterUser(ctx context.Context, input *models.UserRegistrationInput) (*models.User, error) {
	// ... (existing RegisterUser code)
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
	}

	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	user.PasswordHash = ""
	return user, nil
}

// Login handles the user login process.
func (s *authService) Login(ctx context.Context, input *models.LoginUserInput) (string, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", errors.New("invalid credentials")
		}
		return "", fmt.Errorf("failed to get user: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Create JWT token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(), // Token expires in 7 days
		"iat": time.Now().Unix(),
	})

	token, err := claims.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return token, nil
}

// GetUserProfile retrieves a user's profile.
func (s *authService) GetUserProfile(ctx context.Context, userID string) (*models.User, error) {
	user, err := s.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user profile: %w", err)
	}

	// For security, clear the password hash before returning the user object
	user.PasswordHash = ""
	return user, nil
}
