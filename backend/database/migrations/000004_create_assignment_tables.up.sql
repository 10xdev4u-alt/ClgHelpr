-- Migration: 000004_create_assignment_tables.up.sql

-- Assignments Table
CREATE TABLE assignments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    subject_id UUID REFERENCES subjects(id),
    staff_id UUID REFERENCES staff(id),
    
    title VARCHAR(500) NOT NULL,
    description TEXT,
    instructions TEXT,
    
    assignment_type VARCHAR(30) CHECK (assignment_type IN (
        'assignment', 'lab_record', 'project', 'presentation',
        'viva', 'quiz', 'report', 'other'
    )),
    
    assigned_date DATE,
    due_date TIMESTAMP WITH TIME ZONE NOT NULL,
    submitted_at TIMESTAMP WITH TIME ZONE,
    
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN (
        'pending', 'in_progress', 'completed', 'submitted', 'graded', 'overdue'
    )),
    
    max_marks DECIMAL(5,2),
    obtained_marks DECIMAL(5,2),
    feedback TEXT,
    
    priority VARCHAR(10) DEFAULT 'medium' CHECK (priority IN ('low', 'medium', 'high', 'urgent')),
    estimated_hours DECIMAL(4,2),
    actual_hours DECIMAL(4,2),
    
    reminder_enabled BOOLEAN DEFAULT true,
    reminder_before_hours INT DEFAULT 24,
    last_reminded_at TIMESTAMP WITH TIME ZONE,
    
    tags TEXT[],
    is_recurring BOOLEAN DEFAULT false,
    recurrence_pattern VARCHAR(50),
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_assignments_user_status ON assignments(user_id, status);
CREATE INDEX idx_assignments_due_date ON assignments(due_date);
CREATE INDEX idx_assignments_subject ON assignments(subject_id);
CREATE INDEX idx_assignments_tags ON assignments USING GIN(tags);

-- Assignment Attachments Table
CREATE TABLE assignment_attachments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    assignment_id UUID REFERENCES assignments(id) ON DELETE CASCADE,
    
    file_name VARCHAR(255) NOT NULL,
    file_type VARCHAR(50),
    file_size BIGINT,
    file_url TEXT NOT NULL,
    storage_key VARCHAR(500),
    
    attachment_type VARCHAR(20) CHECK (attachment_type IN ('reference', 'submission', 'feedback')),
    
    uploaded_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_attachments_assignment ON assignment_attachments(assignment_id);

-- Apply the auto-update trigger to the new tables
CREATE TRIGGER update_assignments_updated_at BEFORE UPDATE ON assignments
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
