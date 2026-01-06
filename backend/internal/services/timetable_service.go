package services

import (
	"context"
	"time"

	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/repository"
)

// TimetableService defines the interface for timetable-related business logic.
type TimetableService interface {
	CreateSubject(ctx context.Context, subject *models.Subject) error
	GetAllSubjects(ctx context.Context) ([]models.Subject, error)
	CreateStaff(ctx context.Context, staff *models.Staff) error
	GetAllStaff(ctx context.Context) ([]models.Staff, error)
	CreateVenue(ctx context.Context, venue *models.Venue) error
	GetAllVenues(ctx context.Context) ([]models.Venue, error)
	CreateTimetableSlot(ctx context.Context, slot *models.TimetableSlot) error
	GetUserTimetableByDay(ctx context.Context, userID string, dayOfWeek int32) ([]models.TimetableSlot, error)
	GetUserTimetableByDateRange(ctx context.Context, userID string, start, end time.Time) ([]models.TimetableSlot, error)
}

// timetableService implements TimetableService.
type timetableService struct {
	subjectRepo repository.SubjectRepository
	staffRepo   repository.StaffRepository
	venueRepo   repository.VenueRepository
	slotRepo    repository.TimetableSlotRepository
}

// NewTimetableService creates a new timetable service.
func NewTimetableService(
	subjectRepo repository.SubjectRepository,
	staffRepo repository.StaffRepository,
	venueRepo repository.VenueRepository,
	slotRepo repository.TimetableSlotRepository,
) TimetableService {
	return &timetableService{
		subjectRepo: subjectRepo,
		staffRepo:   staffRepo,
		venueRepo:   venueRepo,
		slotRepo:    slotRepo,
	}
}

// CreateSubject creates a new subject.
func (s *timetableService) CreateSubject(ctx context.Context, subject *models.Subject) error {
	return s.subjectRepo.CreateSubject(ctx, subject)
}

// GetAllSubjects retrieves all subjects.
func (s *timetableService) GetAllSubjects(ctx context.Context) ([]models.Subject, error) {
	return s.subjectRepo.GetAllSubjects(ctx)
}

// CreateStaff creates a new staff member.
func (s *timetableService) CreateStaff(ctx context.Context, staff *models.Staff) error {
	return s.staffRepo.CreateStaff(ctx, staff)
}

// GetAllStaff retrieves all staff members.
func (s *timetableService) GetAllStaff(ctx context.Context) ([]models.Staff, error) {
	return s.staffRepo.GetAllStaff(ctx)
}

// CreateVenue creates a new venue.
func (s *timetableService) CreateVenue(ctx context.Context, venue *models.Venue) error {
	return s.venueRepo.CreateVenue(ctx, venue)
}

// GetAllVenues retrieves all venues.
func (s *timetableService) GetAllVenues(ctx context.Context) ([]models.Venue, error) {
	return s.venueRepo.GetAllVenues(ctx)
}

// CreateTimetableSlot creates a new timetable slot.
func (s *timetableService) CreateTimetableSlot(ctx context.Context, slot *models.TimetableSlot) error {
	return s.slotRepo.CreateTimetableSlot(ctx, slot)
}

// GetUserTimetableByDay retrieves timetable slots for a specific user and day.
func (s *timetableService) GetUserTimetableByDay(ctx context.Context, userID string, dayOfWeek int32) ([]models.TimetableSlot, error) {
	return s.slotRepo.GetTimetableSlotsByUserIDAndDay(ctx, userID, dayOfWeek)
}

// GetUserTimetableByDateRange retrieves timetable slots for a specific user within a date range.
func (s *timetableService) GetUserTimetableByDateRange(ctx context.Context, userID string, start, end time.Time) ([]models.TimetableSlot, error) {
	return s.slotRepo.GetTimetableSlotsByUserIDAndDateRange(ctx, userID, start, end)
}
