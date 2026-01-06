-- Migration: 000005_create_exam_tables.up.sql

-- Exams Table
CREATE TABLE exams (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    subject_id UUID REFERENCES subjects(id),
    
    title VARCHAR(255) NOT NULL,
    exam_type VARCHAR(30) CHECK (exam_type IN (
        'cat1', 'cat2', 'cat3', 'fat', 'model', 'retest', 'quiz', 'viva', 'practical'
    )),
    
    -- Schedule
    exam_date DATE NOT NULL,
    start_time TIME,
    end_time TIME,
    duration_minutes INT,
    venue_id UUID REFERENCES venues(id),
    
    -- Syllabus
    syllabus_units TEXT[],
    syllabus_topics TEXT[],
    syllabus_notes TEXT,
    
    -- Results
    max_marks DECIMAL(5,2),
    obtained_marks DECIMAL(5,2),
    grade VARCHAR(5),
    
    -- Prep Tracking
    prep_status VARCHAR(20) DEFAULT 'not_started' CHECK (prep_status IN (
        'not_started', 'in_progress', 'revision', 'ready'
    )),
    prep_notes TEXT,
    study_hours_logged DECIMAL(5,2) DEFAULT 0,
    
    -- Reminder
    reminder_enabled BOOLEAN DEFAULT true,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_exams_user_date ON exams(user_id, exam_date);
CREATE INDEX idx_exams_subject ON exams(subject_id);
CREATE INDEX idx_exams_syllabus_units ON exams USING GIN(syllabus_units);
CREATE INDEX idx_exams_syllabus_topics ON exams USING GIN(syllabus_topics);


-- Important Questions Table
CREATE TABLE important_questions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    subject_id UUID REFERENCES subjects(id),
    exam_id UUID REFERENCES exams(id),
    
    question_text TEXT NOT NULL,
    answer_text TEXT,
    source VARCHAR(100),
    
    unit VARCHAR(50),
    topic VARCHAR(255),
    marks INT,
    frequency_count INT DEFAULT 1,
    
    is_practiced BOOLEAN DEFAULT false,
    last_practiced_at TIMESTAMP WITH TIME ZONE,
    confidence_level INT CHECK (confidence_level BETWEEN 1 AND 5),
    
    tags TEXT[],
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_imp_questions_subject ON important_questions(subject_id);
CREATE INDEX idx_imp_questions_exam ON important_questions(exam_id);
CREATE INDEX idx_imp_questions_tags ON important_questions USING GIN(tags);

-- Apply the auto-update trigger to the new exams table
CREATE TRIGGER update_exams_updated_at BEFORE UPDATE ON exams
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
