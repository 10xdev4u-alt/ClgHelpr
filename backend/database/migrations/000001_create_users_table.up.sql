-- Migration: 000001_create_users_table.up.sql

-- Extension for UUID generation
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users Table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255), -- NULL for OAuth users
    full_name VARCHAR(255) NOT NULL,
    avatar_url TEXT,
    phone VARCHAR(20),
    
    -- Academic Info
    register_number VARCHAR(50),
    department VARCHAR(100) DEFAULT 'CSE',
    year INT CHECK (year BETWEEN 1 AND 4),
    semester INT CHECK (semester BETWEEN 1 AND 8),
    section VARCHAR(10),
    batch VARCHAR(10), -- e.g., 'B2' for labs
    is_hosteler BOOLEAN DEFAULT false,
    
    -- Settings
    notification_preferences JSONB DEFAULT '{"push": true, "email": true, "morning_briefing": true}',
    theme VARCHAR(20) DEFAULT 'system',
    timezone VARCHAR(50) DEFAULT 'Asia/Kolkata',
    
    -- OAuth
    google_id VARCHAR(255),
    github_id VARCHAR(255),
    
    -- Metadata
    is_active BOOLEAN DEFAULT true,
    is_verified BOOLEAN DEFAULT false,
    last_login_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_register_number ON users(register_number);

-- Auto-update updated_at timestamp trigger
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
