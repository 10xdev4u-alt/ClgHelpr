-- Migration: 000009_create_analytics_tables.up.sql

-- Activity Logs Table
CREATE TABLE activity_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    
    activity_type VARCHAR(50) NOT NULL,
    description TEXT,
    
    entity_type VARCHAR(30), -- 'assignment', 'exam', 'study_session', etc.
    entity_id UUID,
    
    metadata JSONB,
    
    ip_address INET,
    user_agent TEXT,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_activity_logs_user ON activity_logs(user_id);
CREATE INDEX idx_activity_logs_type ON activity_logs(activity_type);
CREATE INDEX idx_activity_logs_created ON activity_logs(created_at);

-- Daily Stats Table (Pre-aggregated for dashboard)
CREATE TABLE daily_stats (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    stat_date DATE NOT NULL,
    
    -- Study metrics
    study_minutes INT DEFAULT 0,
    sessions_completed INT DEFAULT 0,
    topics_covered INT DEFAULT 0,
    
    -- Task metrics
    assignments_completed INT DEFAULT 0,
    assignments_added INT DEFAULT 0,
    
    -- Attendance
    classes_attended INT DEFAULT 0,
    total_classes INT DEFAULT 0,
    
    -- XP
    xp_earned INT DEFAULT 0,
    
    UNIQUE(user_id, stat_date)
);

CREATE INDEX idx_daily_stats_user_date ON daily_stats(user_id, stat_date);
