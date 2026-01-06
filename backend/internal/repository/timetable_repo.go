package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/princetheprogrammer/campus-pilot/backend/internal/models"
)

// --- Subject Repository ---

// SubjectRepository defines the interface for subject data operations.
type SubjectRepository interface {
	CreateSubject(ctx context.Context, subject *models.Subject) error
	GetSubjectByID(ctx context.Context, id string) (*models.Subject, error)
	GetSubjectByCode(ctx context.Context, code string) (*models.Subject, error)
	GetAllSubjects(ctx context.Context) ([]models.Subject, error)
}

// PGSubjectRepository implements SubjectRepository for PostgreSQL.
type PGSubjectRepository struct {
	db *pgxpool.Pool
}

// NewPGSubjectRepository creates a new PostgreSQL subject repository.
func NewPGSubjectRepository(db *pgxpool.Pool) *PGSubjectRepository {
	return &PGSubjectRepository{db: db}
}

// CreateSubject inserts a new subject into the database.
func (r *PGSubjectRepository) CreateSubject(ctx context.Context, subject *models.Subject) error {
	query := `
		INSERT INTO subjects (id, code, name, short_name, type, credits, department, semester, color, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id, created_at
	`
	subject.ID = models.NewUUID()
	subject.CreatedAt = time.Now()

	return r.db.QueryRow(ctx, query,
		subject.ID, subject.Code, subject.Name, subject.ShortName, subject.Type, subject.Credits,
		subject.Department, subject.Semester, subject.Color, subject.CreatedAt,
	).Scan(&subject.ID, &subject.CreatedAt)
}

// GetSubjectByID retrieves a subject by its ID.
func (r *PGSubjectRepository) GetSubjectByID(ctx context.Context, id string) (*models.Subject, error) {
	subject := &models.Subject{}
	query := `SELECT id, code, name, short_name, type, credits, department, semester, color, created_at FROM subjects WHERE id = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&subject.ID, &subject.Code, &subject.Name, &subject.ShortName, &subject.Type, &subject.Credits,
		&subject.Department, &subject.Semester, &subject.Color, &subject.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return subject, nil
}

// GetSubjectByCode retrieves a subject by its code.
func (r *PGSubjectRepository) GetSubjectByCode(ctx context.Context, code string) (*models.Subject, error) {
	subject := &models.Subject{}
	query := `SELECT id, code, name, short_name, type, credits, department, semester, color, created_at FROM subjects WHERE code = $1`
	err := r.db.QueryRow(ctx, query, code).Scan(
		&subject.ID, &subject.Code, &subject.Name, &subject.ShortName, &subject.Type, &subject.Credits,
		&subject.Department, &subject.Semester, &subject.Color, &subject.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return subject, nil
}

// GetAllSubjects retrieves all subjects.
func (r *PGSubjectRepository) GetAllSubjects(ctx context.Context) ([]models.Subject, error) {
	var subjects []models.Subject
	query := `SELECT id, code, name, short_name, type, credits, department, semester, color, created_at FROM subjects`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var subject models.Subject
		err := rows.Scan(
			&subject.ID, &subject.Code, &subject.Name, &subject.ShortName, &subject.Type, &subject.Credits,
			&subject.Department, &subject.Semester, &subject.Color, &subject.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, subject)
	}
	return subjects, nil
}

// --- Staff Repository ---

// StaffRepository defines the interface for staff data operations.
type StaffRepository interface {
	CreateStaff(ctx context.Context, staff *models.Staff) error
	GetStaffByID(ctx context.Context, id string) (*models.Staff, error)
	GetAllStaff(ctx context.Context) ([]models.Staff, error)
}

// PGStaffRepository implements StaffRepository for PostgreSQL.
type PGStaffRepository struct {
	db *pgxpool.Pool
}

// NewPGStaffRepository creates a new PostgreSQL staff repository.
func NewPGStaffRepository(db *pgxpool.Pool) *PGStaffRepository {
	return &PGStaffRepository{db: db}
}

// CreateStaff inserts a new staff member into the database.
func (r *PGStaffRepository) CreateStaff(ctx context.Context, staff *models.Staff) error {
	query := `
		INSERT INTO staff (id, name, title, email, phone, department, designation, cabin, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at
	`
	staff.ID = models.NewUUID()
	staff.CreatedAt = time.Now()

	return r.db.QueryRow(ctx, query,
		staff.ID, staff.Name, staff.Title, staff.Email, staff.Phone, staff.Department, staff.Designation, staff.Cabin, staff.CreatedAt,
	).Scan(&staff.ID, &staff.CreatedAt)
}

// GetStaffByID retrieves a staff member by their ID.
func (r *PGStaffRepository) GetStaffByID(ctx context.Context, id string) (*models.Staff, error) {
	staff := &models.Staff{}
	query := `SELECT id, name, title, email, phone, department, designation, cabin, created_at FROM staff WHERE id = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&staff.ID, &staff.Name, &staff.Title, &staff.Email, &staff.Phone, &staff.Department, &staff.Designation, &staff.Cabin, &staff.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return staff, nil
}

// GetAllStaff retrieves all staff members.
func (r *PGStaffRepository) GetAllStaff(ctx context.Context) ([]models.Staff, error) {
	var staffMembers []models.Staff
	query := `SELECT id, name, title, email, phone, department, designation, cabin, created_at FROM staff`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var staff models.Staff
		err := rows.Scan(
			&staff.ID, &staff.Name, &staff.Title, &staff.Email, &staff.Phone, &staff.Department, &staff.Designation, &staff.Cabin, &staff.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		staffMembers = append(staffMembers, staff)
	}
	return staffMembers, nil
}

// --- Venue Repository ---

// VenueRepository defines the interface for venue data operations.
type VenueRepository interface {
	CreateVenue(ctx context.Context, venue *models.Venue) error
	GetVenueByID(ctx context.Context, id string) (*models.Venue, error)
	GetAllVenues(ctx context.Context) ([]models.Venue, error)
}

// PGVenueRepository implements VenueRepository for PostgreSQL.
type PGVenueRepository struct {
	db *pgxpool.Pool
}

// NewPGVenueRepository creates a new PostgreSQL venue repository.
func NewPGVenueRepository(db *pgxpool.Pool) *PGVenueRepository {
	return &PGVenueRepository{db: db}
}

// CreateVenue inserts a new venue into the database.
func (r *PGVenueRepository) CreateVenue(ctx context.Context, venue *models.Venue) error {
	query := `
		INSERT INTO venues (id, name, building, floor, capacity, type, facilities, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at
	`
	venue.ID = models.NewUUID()
	venue.CreatedAt = time.Now()

	return r.db.QueryRow(ctx, query,
		venue.ID, venue.Name, venue.Building, venue.Floor, venue.Capacity, venue.Type, venue.Facilities, venue.CreatedAt,
	).Scan(&venue.ID, &venue.CreatedAt)
}

// GetVenueByID retrieves a venue by its ID.
func (r *PGVenueRepository) GetVenueByID(ctx context.Context, id string) (*models.Venue, error) {
	venue := &models.Venue{}
	query := `SELECT id, name, building, floor, capacity, type, facilities, created_at FROM venues WHERE id = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&venue.ID, &venue.Name, &venue.Building, &venue.Floor, &venue.Capacity, &venue.Type, &venue.Facilities, &venue.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return venue, nil
}

// GetAllVenues retrieves all venues.
func (r *PGVenueRepository) GetAllVenues(ctx context.Context) ([]models.Venue, error) {
	var venues []models.Venue
	query := `SELECT id, name, building, floor, capacity, type, facilities, created_at FROM venues`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var venue models.Venue
		err := rows.Scan(
			&venue.ID, &venue.Name, &venue.Building, &venue.Floor, &venue.Capacity, &venue.Type, &venue.Facilities, &venue.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		venues = append(venues, venue)
	}
	return venues, nil
}

// --- TimetableSlot Repository ---

// TimetableSlotRepository defines the interface for timetable slot data operations.
type TimetableSlotRepository interface {
	CreateTimetableSlot(ctx context.Context, slot *models.TimetableSlot) error
	GetTimetableSlotByID(ctx context.Context, id string) (*models.TimetableSlot, error)
	GetTimetableSlotsByUserIDAndDay(ctx context.Context, userID string, dayOfWeek int32) ([]models.TimetableSlot, error)
	GetTimetableSlotsByUserIDAndDateRange(ctx context.Context, userID string, start, end time.Time) ([]models.TimetableSlot, error)
}

// PGTimetableSlotRepository implements TimetableSlotRepository for PostgreSQL.
type PGTimetableSlotRepository struct {
	db *pgxpool.Pool
}

// NewPGTimetableSlotRepository creates a new PostgreSQL timetable slot repository.
func NewPGTimetableSlotRepository(db *pgxpool.Pool) *PGTimetableSlotRepository {
	return &PGTimetableSlotRepository{db: db}
}

// CreateTimetableSlot inserts a new timetable slot into the database.
func (r *PGTimetableSlotRepository) CreateTimetableSlot(ctx context.Context, slot *models.TimetableSlot) error {
	query := `
		INSERT INTO timetable_slots (
			id, user_id, subject_id, staff_id, venue_id, day_of_week,
			start_time, end_time, period_number, slot_type, is_recurring,
			specific_date, notes, batch_filter, is_active, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17
		) RETURNING id, created_at, updated_at
	`
	slot.ID = models.NewUUID()
	slot.CreatedAt = time.Now()
	slot.UpdatedAt = time.Now()
	slot.IsActive = true

	return r.db.QueryRow(ctx, query,
		slot.ID, slot.UserID, slot.SubjectID, slot.StaffID, slot.VenueID, slot.DayOfWeek,
		slot.StartTime, slot.EndTime, slot.PeriodNumber, slot.SlotType, slot.IsRecurring,
		slot.SpecificDate, slot.Notes, slot.BatchFilter, slot.IsActive, slot.CreatedAt, slot.UpdatedAt,
	).Scan(&slot.ID, &slot.CreatedAt, &slot.UpdatedAt)
}

// GetTimetableSlotByID retrieves a timetable slot by its ID.
func (r *PGTimetableSlotRepository) GetTimetableSlotByID(ctx context.Context, id string) (*models.TimetableSlot, error) {
	slot := &models.TimetableSlot{}
	query := `SELECT id, user_id, subject_id, staff_id, venue_id, day_of_week, start_time, end_time,
	          period_number, slot_type, is_recurring, specific_date, notes, batch_filter,
	          is_active, created_at, updated_at FROM timetable_slots WHERE id = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&slot.ID, &slot.UserID, &slot.SubjectID, &slot.StaffID, &slot.VenueID, &slot.DayOfWeek,
		&slot.StartTime, &slot.EndTime, &slot.PeriodNumber, &slot.SlotType, &slot.IsRecurring,
		&slot.SpecificDate, &slot.Notes, &slot.BatchFilter, &slot.IsActive, &slot.CreatedAt, &slot.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return slot, nil
}

// GetTimetableSlotsByUserIDAndDay retrieves timetable slots for a specific user and day of the week.
func (r *PGTimetableSlotRepository) GetTimetableSlotsByUserIDAndDay(ctx context.Context, userID string, dayOfWeek int32) ([]models.TimetableSlot, error) {
	var slots []models.TimetableSlot
	query := `
		SELECT id, user_id, subject_id, staff_id, venue_id, day_of_week, start_time, end_time,
	          period_number, slot_type, is_recurring, specific_date, notes, batch_filter,
	          is_active, created_at, updated_at
		FROM timetable_slots
		WHERE user_id = $1 AND day_of_week = $2 AND is_active = TRUE AND is_recurring = TRUE
		ORDER BY start_time ASC
	`
	rows, err := r.db.Query(ctx, query, userID, dayOfWeek)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var slot models.TimetableSlot
		err := rows.Scan(
			&slot.ID, &slot.UserID, &slot.SubjectID, &slot.StaffID, &slot.VenueID, &slot.DayOfWeek,
			&slot.StartTime, &slot.EndTime, &slot.PeriodNumber, &slot.SlotType, &slot.IsRecurring,
			&slot.SpecificDate, &slot.Notes, &slot.BatchFilter, &slot.IsActive, &slot.CreatedAt, &slot.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		slots = append(slots, slot)
	}
	return slots, nil
}

// GetTimetableSlotsByUserIDAndDateRange retrieves timetable slots for a specific user within a date range.
func (r *PGTimetableSlotRepository) GetTimetableSlotsByUserIDAndDateRange(ctx context.Context, userID string, start, end time.Time) ([]models.TimetableSlot, error) {
	var slots []models.TimetableSlot
	query := `
		SELECT id, user_id, subject_id, staff_id, venue_id, day_of_week, start_time, end_time,
	          period_number, slot_type, is_recurring, specific_date, notes, batch_filter,
	          is_active, created_at, updated_at
		FROM timetable_slots
		WHERE user_id = $1 AND is_active = TRUE AND
		((is_recurring = TRUE AND day_of_week BETWEEN EXTRACT(DOW FROM $2) AND EXTRACT(DOW FROM $3)) OR
		 (is_recurring = FALSE AND specific_date BETWEEN $2::date AND $3::date))
		ORDER BY start_time ASC
	`
	rows, err := r.db.Query(ctx, query, userID, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var slot models.TimetableSlot
		err := rows.Scan(
			&slot.ID, &slot.UserID, &slot.SubjectID, &slot.StaffID, &slot.VenueID, &slot.DayOfWeek,
			&slot.StartTime, &slot.EndTime, &slot.PeriodNumber, &slot.SlotType, &slot.IsRecurring,
			&slot.SpecificDate, &slot.Notes, &slot.BatchFilter, &slot.IsActive, &slot.CreatedAt, &slot.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		slots = append(slots, slot)
	}
	return slots, nil
}
