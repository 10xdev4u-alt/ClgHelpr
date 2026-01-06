-- Migration: 000008_create_study_planner_tables.up.sql

-- Study Plans Table (e.g., Saturday schedules)
CREATE TABLE study_plans (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    
    title VARCHAR(255) NOT NULL,
    plan_date DATE NOT NULL,
    
    plan_type VARCHAR(30) CHECK (plan_type IN (
        'daily', 'weekly', 'weekend', 'exam_prep', 'revision', 'custom'
    )),
    
    -- Status
    status VARCHAR(20) DEFAULT 'planned' CHECK (status IN (
        'planned', 'in_progress', 'completed', 'partial', 'skipped'
    )),
    
    notes TEXT,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_study_plans_user_date ON study_plans(user_id, plan_date);

-- Study Sessions Table
CREATE TABLE study_sessions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    study_plan_id UUID REFERENCES study_plans(id) ON DELETE SET NULL,
    subject_id UUID REFERENCES subjects(id),
    
    -- Schedule
    planned_start_time TIMESTAMP WITH TIME ZONE,
    planned_end_time TIMESTAMP WITH TIME ZONE,
    planned_duration_minutes INT,
    
    -- Actual
    actual_start_time TIMESTAMP WITH TIME ZONE,
    actual_end_time TIMESTAMP WITH TIME ZONE,
    actual_duration_minutes INT,
    
    -- Content
    session_type VARCHAR(30) CHECK (session_type IN (
        'study', 'revision', 'practice', 'assignment', 'lab_prep', 'exam_prep'
    )),
    topics_to_cover TEXT[],
    topics_covered TEXT[],
    
    -- Progress
    status VARCHAR(20) DEFAULT 'planned' CHECK (status IN (
        'planned', 'in_progress', 'completed', 'partial', 'skipped'
    )),
    completion_percentage INT DEFAULT 0 CHECK (completion_percentage BETWEEN 0 AND 100),
    
    -- Reflection
    productivity_rating INT CHECK (productivity_rating BETWEEN 1 AND 5),
    notes TEXT,
    blockers TEXT,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_study_sessions_user ON study_sessions(user_id);
CREATE INDEX idx_study_sessions_plan ON study_sessions(study_plan_id);
CREATE INDEX idx_study_sessions_date ON study_sessions(planned_start_time);

-- Apply the auto-update trigger to the new tables
CREATE TRIGGER update_study_plans_updated_at BEFORE UPDATE ON study_plans
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_study_sessions_updated_at BEFORE UPDATE ON study_sessions
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
