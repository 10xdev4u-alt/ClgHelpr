package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/arran4/golang-ical" // Import the golang-ical library
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
	GenerateICSCalendar(ctx context.Context, userID string, start, end time.Time) (string, error)
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

// GenerateICSCalendar generates an ICS calendar string for a user's timetable within a date range.
func (s *timetableService) GenerateICSCalendar(ctx context.Context, userID string, start, end time.Time) (string, error) {
	slots, err := s.slotRepo.GetTimetableSlotsByUserIDAndDateRange(ctx, userID, start, end)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve timetable slots: %w", err)
	}

	cal := ics.NewCalendar()
	cal.SetProdid("-//Campus Pilot//NONSGML Timetable//EN")
	cal.SetName("Campus Pilot Timetable")
	cal.SetDescription("Your personalized Campus Pilot Timetable")

	for _, slot := range slots {
		event := cal.AddEvent(models.NewUUID()) // Use a unique ID for each event

		// Get Subject and Staff details for event description
		subject, err := s.subjectRepo.GetSubjectByID(ctx, slot.SubjectID.String)
		if err != nil {
			log.Printf("Warning: Could not find subject for slot %s: %v", slot.ID, err)
			subject = &models.Subject{Name: "Unknown Subject"}
		}

		staffName := "N/A"
		if slot.StaffID.Valid {
			staff, err := s.staffRepo.GetStaffByID(ctx, slot.StaffID.String)
			if err != nil {
				log.Printf("Warning: Could not find staff for slot %s: %v", slot.ID, err)
			} else {
				staffName = staff.Name
			}
		}

		venueName := "N/A"
		if slot.VenueID.Valid {
			venue, err := s.venueRepo.GetVenueByID(ctx, slot.VenueID.String)
			if err != nil {
				log.Printf("Warning: Could not find venue for slot %s: %v", slot.ID, err)
			} else {
				venueName = venue.Name
			}
		}

		// Calculate event start and end times
		// For recurring events, we need to create an occurrence for each week within the range
		// For simplicity, this basic implementation will create single events.
		// A more advanced implementation would use RRULE for recurring events.

		// Combine specific date (if not recurring) or an arbitrary date (for recurring) with start/end time
		eventDate := time.Now() // Default to today, will be overwritten
		if slot.SpecificDate.Valid {
			eventDate = slot.SpecificDate.Time // Use specific date if available
		} else {
			// For recurring events, we need to find the correct date for the given dayOfWeek
			// and then create an event for each week in the range.
			// This is a simplified approach for now - a full RRULE implementation is more complex.
			// For now, let's just create one event on the specific day if it falls within the range.
			// This will generate only one event, not a recurring series.
			currentDay := int32(start.Weekday()) // 0=Sunday, 1=Monday...
			daysToAdd := (slot.DayOfWeek - currentDay + 7) % 7
			eventDate = start.AddDate(0, 0, int(daysToAdd))
		}

		// Combine eventDate with slot.StartTime and slot.EndTime
		eventStart := time.Date(eventDate.Year(), eventDate.Month(), eventDate.Day(),
			slot.StartTime.Hour(), slot.StartTime.Minute(), slot.StartTime.Second(), 0, eventDate.Location())
		eventEnd := time.Date(eventDate.Year(), eventDate.Month(), eventDate.Day(),
			slot.EndTime.Hour(), slot.EndTime.Minute(), slot.EndTime.Second(), 0, eventDate.Location())

		event.SetSummary(fmt.Sprintf("%s - %s", subject.Name, slot.SlotType))
		event.SetDescription(fmt.Sprintf("Subject: %s (%s)\nStaff: %s\nVenue: %s\nType: %s",
			subject.Name, subject.Code, staffName, venueName, slot.SlotType))
		event.SetLocation(venueName)
		event.SetDtStart(eventStart)
		event.SetDtEnd(eventEnd)

		// For recurring events, a proper RRULE would be needed here
		// Example: event.AddRrule(fmt.Sprintf("FREQ=WEEKLY;BYDAY=%s", daysOfWeekShort[slot.DayOfWeek]))
	}

	return cal.Serialize(), nil
}
