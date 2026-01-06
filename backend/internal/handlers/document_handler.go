package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/services"
)

// DocumentHandler handles HTTP requests related to documents.
type DocumentHandler struct {
	documentService services.DocumentService
	validator       *validator.Validate
}

// NewDocumentHandler creates a new DocumentHandler.
func NewDocumentHandler(documentService services.DocumentService) *DocumentHandler {
	return &DocumentHandler{
		documentService: documentService,
		validator:       validator.New(),
	}
}

// CreateDocument handles creating a new document.
// @Summary Create a new document
// @Description Create a new document for the authenticated user.
// @Tags Documents
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param document body models.DocumentCreationInput true "Document details"
// @Success 201 {object} models.Document
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /documents [post]
func (h *DocumentHandler) CreateDocument(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var input models.DocumentCreationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	document, err := h.documentService.CreateDocument(context.Background(), userID, &input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create document: " + err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(document)
}

// GetDocuments handles retrieving all documents for the authenticated user.
// @Summary Get all documents
// @Description Retrieve a list of all documents for the authenticated user.
// @Tags Documents
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Document
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /documents [get]
func (h *DocumentHandler) GetDocuments(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	documents, err := h.documentService.GetDocumentsByUserID(context.Background(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve documents: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(documents)
}

// GetDocumentByID handles retrieving a single document by ID.
// @Summary Get document by ID
// @Description Retrieve a single document by its ID for the authenticated user.
// @Tags Documents
// @Produce json
// @Security BearerAuth
// @Param id path string true "Document ID"
// @Success 200 {object} models.Document
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /documents/{id} [get]
func (h *DocumentHandler) GetDocumentByID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")
	document, err := h.documentService.GetDocumentByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Document not found or not owned by user"})
	}
	if document.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Document does not belong to user"})
	}
	return c.Status(fiber.StatusOK).JSON(document)
}

// GetDocumentsBySubjectID handles retrieving documents by subject ID.
// @Summary Get documents by subject ID
// @Description Retrieve a list of documents for a specific subject for the authenticated user.
// @Tags Documents
// @Produce json
// @Security BearerAuth
// @Param subjectId path string true "Subject ID"
// @Success 200 {array} models.Document
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /documents/subject/{subjectId} [get]
func (h *DocumentHandler) GetDocumentsBySubjectID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	subjectID := c.Params("subjectId")
	documents, err := h.documentService.GetDocumentsByUserIDAndSubjectID(context.Background(), userID, subjectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve documents: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(documents)
}

// UpdateDocument handles updating an existing document.
// @Summary Update a document
// @Description Update an existing document for the authenticated user.
// @Tags Documents
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Document ID"
// @Param document body models.DocumentCreationInput true "Updated document details"
// @Success 200 {object} models.Document
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /documents/{id} [put]
func (h *DocumentHandler) UpdateDocument(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")
	var input models.DocumentCreationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if err := h.validator.Struct(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	document, err := h.documentService.UpdateDocument(context.Background(), userID, id, &input)
	if err != nil {
		if err.Error() == "document does not belong to user" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
		}
		if err.Error() == "document not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update document: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(document)
}

// DeleteDocument handles deleting a document.
// @Summary Delete a document
// @Description Delete a document for the authenticated user.
// @Tags Documents
// @Security BearerAuth
// @Param id path string true "Document ID"
// @Success 204 "Document deleted"
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /documents/{id} [delete]
func (h *DocumentHandler) DeleteDocument(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	id := c.Params("id")

	// Check ownership before deleting
	document, err := h.documentService.GetDocumentByID(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Document not found"})
	}
	if document.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Document does not belong to user"})
	}

	if err := h.documentService.DeleteDocument(context.Background(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete document: " + err.Error()})
	}
	return c.Status(fiber.StatusNoContent).Send(nil)
}

// IncrementViewCount handles incrementing the view count for a document.
// @Summary Increment document view count
// @Description Increment the view count for a specific document.
// @Tags Documents
// @Produce json
// @Security BearerAuth
// @Param id path string true "Document ID"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /documents/{id}/view [post]
func (h *DocumentHandler) IncrementViewCount(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.documentService.IncrementViewCount(context.Background(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to increment view count: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "View count incremented"})
}

// IncrementDownloadCount handles incrementing the download count for a document.
// @Summary Increment document download count
// @Description Increment the download count for a specific document.
// @Tags Documents
// @Produce json
// @Security BearerAuth
// @Param id path string true "Document ID"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /documents/{id}/download-count [post]
func (h *DocumentHandler) IncrementDownloadCount(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.documentService.IncrementDownloadCount(context.Background(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to increment download count: " + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Download count incremented"})
}
