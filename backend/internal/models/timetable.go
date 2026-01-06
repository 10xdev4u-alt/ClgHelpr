package models

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// Subject represents a course subject.
type Subject struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	ShortName sql.NullString `json:"shortName"` // Use sql.NullString for nullable fields
	Type      string    `json:"type"`
	Credits   sql.NullInt32  `json:"credits"`
	Department sql.NullString `json:"department"`
	Semester  sql.NullInt32  `json:"semester"`
	Color     sql.NullString `json:"color"`
	CreatedAt time.Time `json:"createdAt"`
}

// Staff represents a faculty member.
type Staff struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Title       sql.NullString `json:"title"`
	Email       sql.NullString `json:"email"`
	Phone       sql.NullString `json:"phone"`
	Department  sql.NullString `json:"department"`
	Designation sql.NullString `json:"designation"`
	Cabin       sql.NullString `json:"cabin"`
	CreatedAt   time.Time `json:"createdAt"`
}

// Venue represents a physical location on campus.
type Venue struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Building  sql.NullString `json:"building"`
	Floor     sql.NullInt32  `json:"floor"`
	Capacity  sql.NullInt32  `json:"capacity"`
	Type      string         `json:"type"`
	Facilities pgtype.JSONB  `json:"facilities"` // JSONB type
	CreatedAt time.Time      `json:"createdAt"`
}

// TimetableSlot represents a single entry in a student's timetable.
type TimetableSlot struct {
	ID           string         `json:"id"`
	UserID       string         `json:"userId"`
	SubjectID    sql.NullString `json:"subjectId"`
	StaffID      sql.NullString `json:"staffId"`
	VenueID      sql.NullString `json:"venueId"`
	DayOfWeek    int32          `json:"dayOfWeek"`
	StartTime    time.Time      `json:"startTime"` // Only Time part is relevant
	EndTime      time.Time      `json:"endTime"`   // Only Time part is relevant
	PeriodNumber sql.NullInt32  `json:"periodNumber"`
	SlotType     string         `json:"slotType"`
	IsRecurring  bool           `json:"isRecurring"`
	SpecificDate sql.NullTime   `json:"specificDate"`
	Notes        sql.NullString `json:"notes"`
	BatchFilter  sql.NullString `json:"batchFilter"`
	IsActive     bool           `json:"isActive"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
}
