package handlers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/services"
)

// AuthHandler handles HTTP requests related to authentication.
type AuthHandler struct {
	authService services.AuthService
	validator   *validator.Validate
}

// NewAuthHandler creates a new AuthHandler.
func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validator:   validator.New(),
	}
}

// RegisterUser handles user registration.
// @Summary Register a new user
// @Description Register a new user with email, password, full name, and academic info.
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.UserRegistrationInput true "User registration details"
// @Success 201 {object} models.User "User registered successfully"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 409 {object} map[string]string "User with email already exists"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /auth/register [post]
func (h *AuthHandler) RegisterUser(c *fiber.Ctx) error {
	var input models.UserRegistrationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.authService.RegisterUser(context.Background(), &input)
	if err != nil {
		// Differentiate between known errors (e.g., duplicate email) and internal errors
		if err.Error() == "user with this email already exists" {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
