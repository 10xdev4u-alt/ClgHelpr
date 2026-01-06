-- Migration: 000002_create_timetable_tables.up.sql

-- Subjects Table
CREATE TABLE subjects (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    code VARCHAR(20) NOT NULL,
    name VARCHAR(255) NOT NULL,
    short_name VARCHAR(50),
    type VARCHAR(20) CHECK (type IN ('core', 'lab', 'elective', 'open_elective', 'honor', 'minor')),
    credits INT,
    department VARCHAR(100),
    semester INT,
    color VARCHAR(7),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
CREATE INDEX idx_subjects_code ON subjects(code);

-- Staff/Faculty Table
CREATE TABLE staff (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    title VARCHAR(50),
    email VARCHAR(255),
    phone VARCHAR(20),
    department VARCHAR(100),
    designation VARCHAR(100),
    cabin VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Venues Table
CREATE TABLE venues (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    building VARCHAR(100),
    floor INT,
    capacity INT,
    type VARCHAR(50) CHECK (type IN ('classroom', 'lab', 'library', 'seminar_hall', 'auditorium')),
    facilities JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Timetable Slots Table
CREATE TABLE timetable_slots (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    subject_id UUID REFERENCES subjects(id),
    staff_id UUID REFERENCES staff(id),
    venue_id UUID REFERENCES venues(id),
    
    day_of_week INT CHECK (day_of_week BETWEEN 0 AND 6),
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    period_number INT,
    
    slot_type VARCHAR(30) CHECK (slot_type IN (
        'lecture', 'lab', 'tutorial', 'library', 
        'placement_training', 'honor_minor', 'free'
    )),
    
    is_recurring BOOLEAN DEFAULT true,
    specific_date DATE,
    
    notes TEXT,
    batch_filter VARCHAR(10),
    
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
CREATE INDEX idx_timetable_user_day ON timetable_slots(user_id, day_of_week);

-- Apply the auto-update trigger to the new timetable_slots table
CREATE TRIGGER update_timetable_slots_updated_at BEFORE UPDATE ON timetable_slots
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
