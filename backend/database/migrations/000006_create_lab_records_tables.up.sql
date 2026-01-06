-- Migration: 000006_create_lab_records_tables.up.sql

-- Lab Records Table
CREATE TABLE lab_records (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    subject_id UUID REFERENCES subjects(id),
    
    experiment_number INT NOT NULL,
    title VARCHAR(500) NOT NULL,
    
    -- Dates
    lab_date DATE,
    record_written_date DATE,
    submitted_date DATE,
    
    -- Status Tracking
    status VARCHAR(30) DEFAULT 'pending' CHECK (status IN (
        'pending', 'practiced', 'written', 'printed', 'submitted', 'signed', 'returned'
    )),
    
    -- Content
    aim TEXT,
    algorithm TEXT,
    code TEXT,
    output TEXT,
    observations TEXT,
    result TEXT,
    viva_questions TEXT[],
    
    -- Print tracking
    print_required BOOLEAN DEFAULT true,
    pages_to_print INT,
    printed_at TIMESTAMP WITH TIME ZONE,
    
    -- Grading
    marks DECIMAL(5,2),
    staff_remarks TEXT,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    UNIQUE(user_id, subject_id, experiment_number)
);

CREATE INDEX idx_lab_records_user ON lab_records(user_id);
CREATE INDEX idx_lab_records_subject ON lab_records(subject_id);
CREATE INDEX idx_lab_records_status ON lab_records(status);

-- Lab Record Attachments
CREATE TABLE lab_record_attachments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    lab_record_id UUID REFERENCES lab_records(id) ON DELETE CASCADE,
    
    file_name VARCHAR(255) NOT NULL,
    file_type VARCHAR(50),
    file_url TEXT NOT NULL,
    storage_key VARCHAR(500),
    
    attachment_type VARCHAR(30) CHECK (attachment_type IN (
        'code_file', 'output_screenshot', 'record_pdf', 'signed_record', 'other'
    )),
    
    uploaded_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Apply the auto-update trigger to the new lab_records table
CREATE TRIGGER update_lab_records_updated_at BEFORE UPDATE ON lab_records
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
