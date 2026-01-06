Core outcomes (what this app must do)
Daily schedule that never fails
Your timetable as recurring events
“Today view” + “Next class” + venue + staff
Auto reminders: before class, before labs, previous-day lab print/record reminders
Task + assignment system with deadlines
Tasks per subject / per class period / general
Deadlines, priority, status, attachments, checklist
Labs workflow
For each lab session:
“Do in lab” checklist
“After lab” checklist (record writing, print, submit, get signature)
Evidence upload (pics/pdf)
Study planner
Your Saturday 1.5 hr per subject plan
Weekly review: “What’s pending from this week?”
Storage + quick retrieval
Upload PDFs/images/links
Tag by subject + type (FAT paper / notes / important questions / record)
Exam tracker
CAT/FAT schedule, syllabus scope, revision tasks, reminders
AI assistant that uses your data
“What do I have tomorrow?”
“Summarize pending CNS tasks”
“Make Saturday plan using my last-week logs”
This must be RAG (retrieve your saved tasks/files/notes), not generic chatbot replies.
“Nice to have” (high value)
Attendance tracking (manual or timetable-based)
Pomodoro / focus sessions + streaks
“Placement mode”: DSA sheet tracker + mock test reminders
Calendar export/import (ICS)
Friend sharing: share a task/exam doc with batchmates

CAMPUS PILOT - Enterprise Academic Management System
Yo Prince! Let's build something LEGENDARY! I'm super hyped about this project, bro! Let me break this down properly - we're going enterprise-grade, no half measures!

📋 PROJECT OVERVIEW
text

┌─────────────────────────────────────────────────────────────────────────────┐
│                           CAMPUS PILOT                                       │
│         "Your AI-Powered Academic Co-Pilot for College Success"             │
├─────────────────────────────────────────────────────────────────────────────┤
│  Target Users: Engineering Students (Hostelers & Day Scholars)               │
│  Primary User: Prince (You!) - 3rd Year CSE, 6th Semester                   │
│  Secondary: Your friends, batchmates, potentially entire department          │
│  Scale Target: 100 → 1000 → 10,000+ users                                   │
└─────────────────────────────────────────────────────────────────────────────┘
🎯 FEATURE MATRIX (Comprehensive)
TIER 1: Core Features (MVP)
text

┌────────────────────────────────────────────────────────────────────────────┐
│ FEATURE                    │ DESCRIPTION                    │ PRIORITY    │
├────────────────────────────────────────────────────────────────────────────┤
│ 🔐 Authentication          │ Login/Signup, OAuth, Sessions  │ P0 - MUST   │
│ 📅 Smart Timetable         │ Weekly view, venue tracking    │ P0 - MUST   │
│ 📝 Period Logger           │ Notes, links, attachments      │ P0 - MUST   │
│ 📋 Assignment Tracker      │ Deadlines, status, reminders   │ P0 - MUST   │
│ 📚 Exam Manager            │ Schedule, syllabus, countdown  │ P0 - MUST   │
│ 🔔 Smart Reminders         │ Push, Email, Morning Briefing  │ P0 - MUST   │
│ 🤖 LLM Integration         │ Your SVCE AI for assistance    │ P0 - MUST   │
└────────────────────────────────────────────────────────────────────────────┘
TIER 2: Enhanced Features
text

┌────────────────────────────────────────────────────────────────────────────┐
│ FEATURE                    │ DESCRIPTION                    │ PRIORITY    │
├────────────────────────────────────────────────────────────────────────────┤
│ 🧪 Lab Record Manager      │ Track, upload, print status    │ P1 - HIGH   │
│ 📁 Document Vault          │ FAT papers, notes, resources   │ P1 - HIGH   │
│ 📊 Study Planner           │ Saturday sessions, goals       │ P1 - HIGH   │
│ 🌅 Morning Briefing        │ Daily summary, tasks ahead     │ P1 - HIGH   │
│ 📈 Analytics Dashboard     │ Progress, attendance, stats    │ P1 - HIGH   │
│ 🏷️ Smart Tags & Search     │ Find anything instantly        │ P1 - HIGH   │
└────────────────────────────────────────────────────────────────────────────┘
TIER 3: Power Features
text

┌────────────────────────────────────────────────────────────────────────────┐
│ FEATURE                    │ DESCRIPTION                    │ PRIORITY    │
├────────────────────────────────────────────────────────────────────────────┤
│ 👥 Friend Groups           │ Share schedules, collaborate   │ P2 - MEDIUM │
│ 🎮 Gamification            │ Streaks, XP, achievements      │ P2 - MEDIUM │
│ 🔄 Calendar Sync           │ Google/Apple Calendar          │ P2 - MEDIUM │
│ 💼 Placement Tracker       │ Companies, prep, interviews    │ P2 - MEDIUM │
│ 📱 Offline Mode            │ PWA with sync                  │ P2 - MEDIUM │
│ 🌐 Multi-tenant            │ Different departments          │ P2 - MEDIUM │
└────────────────────────────────────────────────────────────────────────────┘
TIER 4: Enterprise Features
text

┌────────────────────────────────────────────────────────────────────────────┐
│ FEATURE                    │ DESCRIPTION                    │ PRIORITY    │
├────────────────────────────────────────────────────────────────────────────┤
│ 👨‍🏫 Faculty Portal          │ Push announcements, materials  │ P3 - FUTURE │
│ 🏛️ Admin Dashboard         │ Department-wide management     │ P3 - FUTURE │
│ 📊 Batch Analytics         │ Class-wide insights            │ P3 - FUTURE │
│ 🔗 LMS Integration         │ Connect with college LMS       │ P3 - FUTURE │
│ 📧 Smart Email Parsing     │ Auto-extract from dept emails  │ P3 - FUTURE │
│ 🤖 AI Study Recommendations│ Personalized study paths       │ P3 - FUTURE │
└────────────────────────────────────────────────────────────────────────────┘
🏗️ HIGH-LEVEL DESIGN (HLD)
System Context Diagram
text

                                    ┌──────────────────┐
                                    │   EXTERNAL       │
                                    │   SERVICES       │
                                    ├──────────────────┤
                                    │ • Google OAuth   │
                                    │ • SVCE AI (LLM)  │
                                    │ • FCM (Push)     │
                                    │ • SendGrid       │
                                    │ • Cloudinary     │
                                    └────────┬─────────┘
                                             │
                                             ▼
┌─────────────┐    HTTPS     ┌──────────────────────────────────────────────┐
│             │◄────────────►│              CAMPUS PILOT                    │
│   STUDENTS  │              │         (Your Application)                   │
│             │              ├──────────────────────────────────────────────┤
│  📱 Mobile  │              │  • Schedule Management                       │
│  💻 Desktop │              │  • Assignment & Exam Tracking                │
│  🌐 PWA     │              │  • Document Storage                          │
│             │              │  • AI-Powered Assistance                     │
└─────────────┘              │  • Smart Notifications                       │
                             │  • Analytics & Insights                      │
                             └──────────────────────────────────────────────┘
                                             │
                                             ▼
                             ┌──────────────────────────────────────────────┐
                             │              DATA STORES                     │
                             ├──────────────────────────────────────────────┤
                             │  🐘 PostgreSQL  │  Primary Database          │
                             │  🔴 Redis       │  Cache & Sessions          │
                             │  📦 MinIO/S3    │  File Storage              │
                             └──────────────────────────────────────────────┘
High-Level Component Architecture
text

┌─────────────────────────────────────────────────────────────────────────────────────┐
│                                    CAMPUS PILOT                                      │
├─────────────────────────────────────────────────────────────────────────────────────┤
│                                                                                      │
│  ┌─────────────────────────────────────────────────────────────────────────────┐   │
│  │                           PRESENTATION LAYER                                 │   │
│  │  ┌───────────────────┐  ┌───────────────────┐  ┌───────────────────┐       │   │
│  │  │    Web App        │  │    PWA Module     │  │   Admin Panel     │       │   │
│  │  │   (Next.js 14)    │  │  (Service Worker) │  │   (Same Stack)    │       │   │
│  │  │                   │  │                   │  │                   │       │   │
│  │  │ • Dashboard       │  │ • Offline Cache   │  │ • User Mgmt       │       │   │
│  │  │ • Timetable       │  │ • Push Handler    │  │ • Analytics       │       │   │
│  │  │ • Assignments     │  │ • Sync Engine     │  │ • Content Mgmt    │       │   │
│  │  │ • Documents       │  │ • Install Prompt  │  │ • Reports         │       │   │
│  │  │ • AI Chat         │  │                   │  │                   │       │   │
│  │  └───────────────────┘  └───────────────────┘  └───────────────────┘       │   │
│  └─────────────────────────────────────────────────────────────────────────────┘   │
│                                         │                                           │
│                                         ▼                                           │
│  ┌─────────────────────────────────────────────────────────────────────────────┐   │
│  │                              API GATEWAY                                     │   │
│  │                          (Traefik / Kong)                                    │   │
│  │  • Rate Limiting  • Auth Verification  • Request Routing  • Load Balancing  │   │
│  └─────────────────────────────────────────────────────────────────────────────┘   │
│                                         │                                           │
│                    ┌────────────────────┼────────────────────┐                     │
│                    ▼                    ▼                    ▼                     │
│  ┌─────────────────────────────────────────────────────────────────────────────┐   │
│  │                           APPLICATION LAYER (Go Services)                    │   │
│  │                                                                              │   │
│  │  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐    │   │
│  │  │   AUTH       │  │  SCHEDULE    │  │  ACADEMIC    │  │  ASSISTANT   │    │   │
│  │  │   SERVICE    │  │  SERVICE     │  │  SERVICE     │  │  SERVICE     │    │   │
│  │  ├──────────────┤  ├──────────────┤  ├──────────────┤  ├──────────────┤    │   │
│  │  │ • Register   │  │ • Timetable  │  │ • Exams      │  │ • LLM Proxy  │    │   │
│  │  │ • Login      │  │ • Period Log │  │ • Labs       │  │ • Context    │    │   │
│  │  │ • OAuth      │  │ • Venues     │  │ • Documents  │  │ • History    │    │   │
│  │  │ • JWT Mgmt   │  │ • Reminders  │  │ • Resources  │  │ • Prompts    │    │   │
│  │  │ • Sessions   │  │ • Conflicts  │  │ • Placement  │  │ • RAG Engine │    │   │
│  │  └──────────────┘  └──────────────┘  └──────────────┘  └──────────────┘    │   │
│  │                                                                              │   │
│  │  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐    │   │
│  │  │ NOTIFICATION │  │  STORAGE     │  │  ANALYTICS   │  │  SCHEDULER   │    │   │
│  │  │   SERVICE    │  │  SERVICE     │  │  SERVICE     │  │  SERVICE     │    │   │
│  │  ├──────────────┤  ├──────────────┤  ├──────────────┤  ├──────────────┤    │   │
│  │  │ • Push       │  │ • Upload     │  │ • Progress   │  │ • Cron Jobs  │    │   │
│  │  │ • Email      │  │ • CDN        │  │ • Stats      │  │ • Reminders  │    │   │
│  │  │ • In-App     │  │ • Preview    │  │ • Reports    │  │ • Digests    │    │   │
│  │  │ • Digest     │  │ • Organize   │  │ • Export     │  │ • Cleanup    │    │   │
│  │  └──────────────┘  └──────────────┘  └──────────────┘  └──────────────┘    │   │
│  └─────────────────────────────────────────────────────────────────────────────┘   │
│                                         │                                           │
│                                         ▼                                           │
│  ┌─────────────────────────────────────────────────────────────────────────────┐   │
│  │                              DATA LAYER                                      │   │
│  │                                                                              │   │
│  │  ┌──────────────────┐  ┌──────────────────┐  ┌──────────────────┐          │   │
│  │  │   PostgreSQL     │  │      Redis       │  │    MinIO/S3      │          │   │
│  │  │                  │  │                  │  │                  │          │   │
│  │  │ • Users          │  │ • Sessions       │  │ • PDFs           │          │   │
│  │  │ • Schedules      │  │ • Cache          │  │ • Images         │          │   │
│  │  │ • Assignments    │  │ • Rate Limits    │  │ • Documents      │          │   │
│  │  │ • Documents      │  │ • Real-time      │  │ • Backups        │          │   │
│  │  │ • Analytics      │  │ • Pub/Sub        │  │                  │          │   │
│  │  └──────────────────┘  └──────────────────┘  └──────────────────┘          │   │
│  └─────────────────────────────────────────────────────────────────────────────┘   │
│                                                                                      │
└─────────────────────────────────────────────────────────────────────────────────────┘
🔧 LOW-LEVEL DESIGN (LLD)
Database Schema Design
SQL

-- ============================================================================
-- CAMPUS PILOT - DATABASE SCHEMA (PostgreSQL)
-- Version: 1.0.0
-- ============================================================================

-- Extension for UUID generation and full-text search
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";

-- ============================================================================
-- CORE ENTITIES
-- ============================================================================

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
    hostel_block VARCHAR(50),
    
    -- Settings
    notification_preferences JSONB DEFAULT '{"push": true, "email": true, "morning_briefing": true}',
    theme VARCHAR(20) DEFAULT 'system',
    timezone VARCHAR(50) DEFAULT 'Asia/Kolkata',
    
    -- OAuth
    google_id VARCHAR(255),
    github_id VARCHAR(255),
    
    -- Gamification
    xp_points INT DEFAULT 0,
    current_streak INT DEFAULT 0,
    longest_streak INT DEFAULT 0,
    level INT DEFAULT 1,
    
    -- Metadata
    is_active BOOLEAN DEFAULT true,
    is_verified BOOLEAN DEFAULT false,
    last_login_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_department ON users(department);
CREATE INDEX idx_users_register_number ON users(register_number);

-- ============================================================================
-- SCHEDULE & TIMETABLE
-- ============================================================================

-- Subjects Table
CREATE TABLE subjects (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    code VARCHAR(20) NOT NULL, -- e.g., 'CS22601'
    name VARCHAR(255) NOT NULL, -- e.g., 'Cryptography and Network Security'
    short_name VARCHAR(50), -- e.g., 'CNS'
    type VARCHAR(20) CHECK (type IN ('core', 'lab', 'elective', 'open_elective', 'honor', 'minor')),
    credits INT,
    department VARCHAR(100),
    semester INT,
    color VARCHAR(7), -- Hex color for UI
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_subjects_code ON subjects(code);

-- Staff/Faculty Table
CREATE TABLE staff (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    title VARCHAR(50), -- e.g., 'Dr.', 'Mr.', 'Ms.'
    email VARCHAR(255),
    phone VARCHAR(20),
    department VARCHAR(100),
    designation VARCHAR(100), -- e.g., 'Associate Professor'
    cabin VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Venues Table
CREATE TABLE venues (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL, -- e.g., 'CSE Lab 1', 'Room 401'
    building VARCHAR(100),
    floor INT,
    capacity INT,
    type VARCHAR(50) CHECK (type IN ('classroom', 'lab', 'library', 'seminar_hall', 'auditorium')),
    facilities JSONB, -- e.g., {"projector": true, "ac": true, "computers": 60}
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Timetable Slots Table
CREATE TABLE timetable_slots (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    subject_id UUID REFERENCES subjects(id),
    staff_id UUID REFERENCES staff(id),
    venue_id UUID REFERENCES venues(id),
    
    day_of_week INT CHECK (day_of_week BETWEEN 0 AND 6), -- 0=Sunday, 1=Monday...
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    period_number INT, -- 1, 2, 3, etc.
    
    slot_type VARCHAR(30) CHECK (slot_type IN (
        'lecture', 'lab', 'tutorial', 'library', 
        'placement_training', 'honor_minor', 'free'
    )),
    
    -- For recurring vs specific dates
    is_recurring BOOLEAN DEFAULT true,
    specific_date DATE, -- For non-recurring slots
    
    -- Additional info
    notes TEXT,
    batch_filter VARCHAR(10), -- NULL means all, 'B1' or 'B2' for specific
    
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_timetable_user_day ON timetable_slots(user_id, day_of_week);
CREATE INDEX idx_timetable_specific_date ON timetable_slots(specific_date) WHERE specific_date IS NOT NULL;

-- ============================================================================
-- PERIOD LOGS & NOTES
-- ============================================================================

-- Period Logs Table (Daily logs for each period)
CREATE TABLE period_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    timetable_slot_id UUID REFERENCES timetable_slots(id),
    subject_id UUID REFERENCES subjects(id),
    
    log_date DATE NOT NULL,
    
    -- Attendance
    attendance_status VARCHAR(20) CHECK (attendance_status IN ('present', 'absent', 'late', 'od', 'medical')),
    
    -- Content
    topics_covered TEXT,
    personal_notes TEXT,
    important_points TEXT[],
    
    -- References
    page_references VARCHAR(100), -- e.g., "Chapter 5, Pages 120-145"
    external_links TEXT[],
    
    -- Mood/Engagement
    understanding_level INT CHECK (understanding_level BETWEEN 1 AND 5),
    engagement_level INT CHECK (engagement_level BETWEEN 1 AND 5),
    
    -- Metadata
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    UNIQUE(user_id, timetable_slot_id, log_date)
);

CREATE INDEX idx_period_logs_user_date ON period_logs(user_id, log_date);
CREATE INDEX idx_period_logs_subject ON period_logs(subject_id);

-- ============================================================================
-- ASSIGNMENTS & TASKS
-- ============================================================================

-- Assignments Table
CREATE TABLE assignments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    subject_id UUID REFERENCES subjects(id),
    staff_id UUID REFERENCES staff(id),
    
    title VARCHAR(500) NOT NULL,
    description TEXT,
    instructions TEXT,
    
    -- Type
    assignment_type VARCHAR(30) CHECK (assignment_type IN (
        'assignment', 'lab_record', 'project', 'presentation',
        'viva', 'quiz', 'report', 'other'
    )),
    
    -- Dates
    assigned_date DATE,
    due_date TIMESTAMP WITH TIME ZONE NOT NULL,
    submitted_at TIMESTAMP WITH TIME ZONE,
    
    -- Status
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN (
        'pending', 'in_progress', 'completed', 'submitted', 'graded', 'overdue'
    )),
    
    -- Grading
    max_marks DECIMAL(5,2),
    obtained_marks DECIMAL(5,2),
    feedback TEXT,
    
    -- Priority & Effort
    priority VARCHAR(10) DEFAULT 'medium' CHECK (priority IN ('low', 'medium', 'high', 'urgent')),
    estimated_hours DECIMAL(4,2),
    actual_hours DECIMAL(4,2),
    
    -- Reminder settings
    reminder_enabled BOOLEAN DEFAULT true,
    reminder_before_hours INT DEFAULT 24,
    last_reminded_at TIMESTAMP WITH TIME ZONE,
    
    -- Metadata
    tags TEXT[],
    is_recurring BOOLEAN DEFAULT false,
    recurrence_pattern VARCHAR(50), -- e.g., 'weekly', 'biweekly'
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_assignments_user_status ON assignments(user_id, status);
CREATE INDEX idx_assignments_due_date ON assignments(due_date);
CREATE INDEX idx_assignments_subject ON assignments(subject_id);

-- Assignment Attachments Table
CREATE TABLE assignment_attachments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    assignment_id UUID REFERENCES assignments(id) ON DELETE CASCADE,
    
    file_name VARCHAR(255) NOT NULL,
    file_type VARCHAR(50),
    file_size BIGINT,
    file_url TEXT NOT NULL,
    storage_key VARCHAR(500), -- S3/MinIO key
    
    attachment_type VARCHAR(20) CHECK (attachment_type IN ('reference', 'submission', 'feedback')),
    
    uploaded_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_attachments_assignment ON assignment_attachments(assignment_id);

-- ============================================================================
-- EXAMS
-- ============================================================================

-- Exams Table
CREATE TABLE exams (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    subject_id UUID REFERENCES subjects(id),
    
    title VARCHAR(255) NOT NULL, -- e.g., 'CAT-1', 'FAT', 'Model Exam'
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
    syllabus_units TEXT[], -- e.g., ['Unit 1', 'Unit 2']
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

-- Important Questions Table
CREATE TABLE important_questions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    subject_id UUID REFERENCES subjects(id),
    exam_id UUID REFERENCES exams(id),
    
    question_text TEXT NOT NULL,
    answer_text TEXT,
    source VARCHAR(100), -- e.g., 'Previous Year', 'Staff Provided', 'Self'
    
    unit VARCHAR(50),
    topic VARCHAR(255),
    marks INT,
    frequency_count INT DEFAULT 1, -- How often it appeared
    
    is_practiced BOOLEAN DEFAULT false,
    last_practiced_at TIMESTAMP WITH TIME ZONE,
    confidence_level INT CHECK (confidence_level BETWEEN 1 AND 5),
    
    tags TEXT[],
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_imp_questions_subject ON important_questions(subject_id);
CREATE INDEX idx_imp_questions_exam ON important_questions(exam_id);

-- ============================================================================
-- LAB RECORDS
-- ============================================================================

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

-- ============================================================================
-- DOCUMENT VAULT
-- ============================================================================

-- Documents Table
CREATE TABLE documents (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    subject_id UUID REFERENCES subjects(id),
    
    title VARCHAR(500) NOT NULL,
    description TEXT,
    
    document_type VARCHAR(30) CHECK (document_type IN (
        'notes', 'textbook', 'slides', 'previous_paper', 'fat_paper',
        'question_bank', 'formula_sheet', 'cheat_sheet', 'other'
    )),
    
    -- File info
    file_name VARCHAR(255),
    file_type VARCHAR(50),
    file_size BIGINT,
    file_url TEXT NOT NULL,
    storage_key VARCHAR(500),
    
    -- Organization
    folder VARCHAR(255), -- Virtual folder path
    tags TEXT[],
    
    -- Access
    is_public BOOLEAN DEFAULT false, -- Can be shared with friends
    shared_with UUID[], -- Specific user IDs
    
    -- Usage tracking
    view_count INT DEFAULT 0,
    download_count INT DEFAULT 0,
    last_accessed_at TIMESTAMP WITH TIME ZONE,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_documents_user ON documents(user_id);
CREATE INDEX idx_documents_subject ON documents(subject_id);
CREATE INDEX idx_documents_type ON documents(document_type);
CREATE INDEX idx_documents_tags ON documents USING GIN(tags);

-- ============================================================================
-- STUDY SESSIONS & PLANNING
-- ============================================================================

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

-- ============================================================================
-- PLACEMENT TRACKING
-- ============================================================================

-- Placement Drives Table
CREATE TABLE placement_drives (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    
    company_name VARCHAR(255) NOT NULL,
    company_logo_url TEXT,
    job_role VARCHAR(255),
    job_description TEXT,
    
    -- Eligibility
    min_cgpa DECIMAL(3,2),
    eligible_branches TEXT[],
    backlog_allowed BOOLEAN,
    
    -- Package
    ctc_lpa DECIMAL(5,2),
    stipend_if_intern DECIMAL(10,2),
    location TEXT[],
    
    -- Process
    rounds JSONB, -- [{name: "Online Test", date: "2024-02-15", status: "cleared"}]
    current_round VARCHAR(100),
    
    -- Dates
    registration_deadline TIMESTAMP WITH TIME ZONE,
    drive_date DATE,
    
    -- Status
    status VARCHAR(30) DEFAULT 'interested' CHECK (status IN (
        'interested', 'registered', 'in_process', 'selected', 'rejected', 'not_eligible', 'skipped'
    )),
    
    -- Notes
    prep_notes TEXT,
    interview_experience TEXT,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_placement_user ON placement_drives(user_id);
CREATE INDEX idx_placement_status ON placement_drives(status);

-- ============================================================================
-- AI ASSISTANT
-- ============================================================================

-- Chat Conversations Table
CREATE TABLE chat_conversations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    
    title VARCHAR(255),
    context_type VARCHAR(30) CHECK (context_type IN (
        'general', 'subject_specific', 'assignment_help', 'exam_prep',
        'study_planning', 'placement_prep', 'doubt_solving'
    )),
    context_subject_id UUID REFERENCES subjects(id),
    
    -- Metadata
    message_count INT DEFAULT 0,
    last_message_at TIMESTAMP WITH TIME ZONE,
    
    is_archived BOOLEAN DEFAULT false,
    is_pinned BOOLEAN DEFAULT false,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_conversations_user ON chat_conversations(user_id);

-- Chat Messages Table
CREATE TABLE chat_messages (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id UUID REFERENCES chat_conversations(id) ON DELETE CASCADE,
    
    role VARCHAR(20) CHECK (role IN ('user', 'assistant', 'system')),
    content TEXT NOT NULL,
    
    -- For rich responses
    code_blocks JSONB, -- [{language: "python", code: "..."}]
    attachments JSONB, -- [{type: "file", url: "..."}]
    
    -- Token tracking
    tokens_used INT,
    model_used VARCHAR(100),
    
    -- Feedback
    feedback_rating INT CHECK (feedback_rating BETWEEN 1 AND 5),
    feedback_text TEXT,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_messages_conversation ON chat_messages(conversation_id);
CREATE INDEX idx_messages_created ON chat_messages(created_at);

-- ============================================================================
-- NOTIFICATIONS & REMINDERS
-- ============================================================================

-- Notifications Table
CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    
    title VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
    
    notification_type VARCHAR(30) CHECK (notification_type IN (
        'assignment_reminder', 'exam_reminder', 'lab_reminder',
        'morning_briefing', 'streak_reminder', 'achievement',
        'system', 'custom'
    )),
    
    -- Reference
    reference_type VARCHAR(30), -- 'assignment', 'exam', 'lab_record', etc.
    reference_id UUID,
    
    -- Delivery
    channels TEXT[] DEFAULT ARRAY['in_app'], -- 'in_app', 'push', 'email'
    scheduled_for TIMESTAMP WITH TIME ZONE,
    sent_at TIMESTAMP WITH TIME ZONE,
    
    -- Status
    is_read BOOLEAN DEFAULT false,
    read_at TIMESTAMP WITH TIME ZONE,
    
    -- Action
    action_url TEXT, -- Deep link to relevant page
    action_text VARCHAR(100),
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_notifications_user ON notifications(user_id);
CREATE INDEX idx_notifications_unread ON notifications(user_id, is_read) WHERE is_read = false;
CREATE INDEX idx_notifications_scheduled ON notifications(scheduled_for) WHERE sent_at IS NULL;

-- Push Subscriptions Table
CREATE TABLE push_subscriptions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    
    endpoint TEXT NOT NULL,
    p256dh_key TEXT NOT NULL,
    auth_key TEXT NOT NULL,
    
    device_type VARCHAR(30),
    device_name VARCHAR(100),
    browser VARCHAR(50),
    
    is_active BOOLEAN DEFAULT true,
    last_used_at TIMESTAMP WITH TIME ZONE,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    UNIQUE(user_id, endpoint)
);

-- ============================================================================
-- GAMIFICATION
-- ============================================================================

-- Achievements Table
CREATE TABLE achievements (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    name VARCHAR(100) NOT NULL,
    description TEXT,
    icon VARCHAR(50), -- Emoji or icon name
    
    category VARCHAR(30) CHECK (category IN (
        'streak', 'study', 'assignment', 'exam', 'social', 'explorer', 'master'
    )),
    
    xp_reward INT DEFAULT 0,
    
    -- Requirements
    requirement_type VARCHAR(50),
    requirement_value INT,
    
    is_hidden BOOLEAN DEFAULT false, -- Secret achievements
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- User Achievements Table
CREATE TABLE user_achievements (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    achievement_id UUID REFERENCES achievements(id) ON DELETE CASCADE,
    
    unlocked_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    progress INT DEFAULT 0, -- For progressive achievements
    
    UNIQUE(user_id, achievement_id)
);

CREATE INDEX idx_user_achievements ON user_achievements(user_id);

-- ============================================================================
-- SOCIAL FEATURES
-- ============================================================================

-- Friend Connections Table
CREATE TABLE friendships (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    friend_id UUID REFERENCES users(id) ON DELETE CASCADE,
    
    status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending', 'accepted', 'blocked')),
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    accepted_at TIMESTAMP WITH TIME ZONE,
    
    UNIQUE(user_id, friend_id)
);

CREATE INDEX idx_friendships_user ON friendships(user_id);
CREATE INDEX idx_friendships_friend ON friendships(friend_id);

-- Study Groups Table
CREATE TABLE study_groups (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    
    name VARCHAR(100) NOT NULL,
    description TEXT,
    avatar_url TEXT,
    
    owner_id UUID REFERENCES users(id),
    
    is_public BOOLEAN DEFAULT false,
    invite_code VARCHAR(20) UNIQUE,
    
    member_count INT DEFAULT 1,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Study Group Members Table
CREATE TABLE study_group_members (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    group_id UUID REFERENCES study_groups(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    
    role VARCHAR(20) DEFAULT 'member' CHECK (role IN ('owner', 'admin', 'member')),
    
    joined_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    UNIQUE(group_id, user_id)
);

-- ============================================================================
-- ANALYTICS & LOGS
-- ============================================================================

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

-- ============================================================================
-- TRIGGERS
-- ============================================================================

-- Auto-update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Apply trigger to all relevant tables
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_timetable_slots_updated_at BEFORE UPDATE ON timetable_slots
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_period_logs_updated_at BEFORE UPDATE ON period_logs
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_assignments_updated_at BEFORE UPDATE ON assignments
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_exams_updated_at BEFORE UPDATE ON exams
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_lab_records_updated_at BEFORE UPDATE ON lab_records
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_documents_updated_at BEFORE UPDATE ON documents
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- ============================================================================
-- SEED DATA FOR PRINCE'S SCHEDULE
-- ============================================================================

-- Insert subjects
INSERT INTO subjects (code, name, short_name, type, credits, department, semester, color) VALUES
('CS22601', 'Cryptography and Network Security', 'CNS', 'core', 4, 'CSE', 6, '#EF4444'),
('CS22602', 'Software Project Management', 'SPM', 'core', 3, 'CSE', 6, '#F59E0B'),
('AD22501', 'Internet of Things and Applications', 'IoT', 'core', 4, 'CSE', 6, '#10B981'),
('CS22603', 'Cloud Computing', 'CC', 'core', 4, 'CSE', 6, '#3B82F6'),
('CS22604', 'Compiler Design', 'CD', 'core', 4, 'CSE', 6, '#8B5CF6'),
('CS22021', 'Exploratory Data Analysis', 'EDA', 'elective', 3, 'CSE', 6, '#EC4899'),
('CS22611', 'CNS Laboratory', 'CNS Lab', 'lab', 2, 'CSE', 6, '#EF4444'),
('CS22612', 'Cloud Computing Laboratory', 'CC Lab', 'lab', 2, 'CSE', 6, '#3B82F6'),
('OE22705', 'Embedded Systems and its Application', 'ES', 'open_elective', 3, 'ECE', 6, '#6366F1'),
('HM001', 'Honor/Minor Subject', 'H/M', 'honor', 3, 'CSE', 6, '#14B8A6');

-- Insert staff
INSERT INTO staff (name, title, department, designation) VALUES
('Dr. S. Vijayanand', 'Dr.', 'ECE', 'Associate Professor'),
('Ms. Banupriya P', 'Ms.', 'CSE', 'Assistant Professor'),
('Dr. Janarthanan P', 'Dr.', 'CSE', 'Associate Professor'),
('Ms. N. Suriya', 'Ms.', 'External', 'Trainer - VYVoxel'),
('Mr. E. Raghuvaran', 'Mr.', 'CSE', 'Assistant Professor'),
('Ms. Kapilvani R K', 'Ms.', 'CSE', 'Assistant Professor'),
('Dr. Poorani S', 'Dr.', 'CSE', 'Assistant Professor'),
('Ms. Rupa Kesavan', 'Ms.', 'CSE', 'Assistant Professor'),
('Mr. Selvamani P', 'Mr.', 'CSE', 'Assistant Professor'),
('Mr. Arunachalam Narayanan', 'Mr.', 'CSE', 'Assistant Professor'),
('Mr. S. Ashok Kumar', 'Mr.', 'Placement', 'Training Officer'),
('Ms. G.R. Khanagavalle', 'Ms.', 'Library', 'Librarian');
🔌 API DESIGN
RESTful API Structure
text

┌─────────────────────────────────────────────────────────────────────────────────┐
│                           CAMPUS PILOT API v1                                   │
│                         Base URL: /api/v1                                       │
├─────────────────────────────────────────────────────────────────────────────────┤

🔐 AUTHENTICATION
├── POST   /auth/register              → Create new account
├── POST   /auth/login                 → Login with email/password
├── POST   /auth/logout                → Logout & invalidate session
├── POST   /auth/refresh               → Refresh access token
├── POST   /auth/forgot-password       → Request password reset
├── POST   /auth/reset-password        → Reset password with token
├── GET    /auth/google                → Google OAuth redirect
├── GET    /auth/google/callback       → Google OAuth callback
├── GET    /auth/me                    → Get current user profile
├── PUT    /auth/me                    → Update current user profile
├── PUT    /auth/me/password           → Change password
├── PUT    /auth/me/preferences        → Update notification preferences

📅 SCHEDULE & TIMETABLE
├── GET    /schedule                   → Get current week schedule
├── GET    /schedule/today             → Get today's schedule
├── GET    /schedule/week/:weekOffset  → Get specific week schedule
├── GET    /schedule/date/:date        → Get schedule for specific date
├── POST   /schedule/slots             → Add timetable slot
├── PUT    /schedule/slots/:id         → Update timetable slot
├── DELETE /schedule/slots/:id         → Delete timetable slot
├── POST   /schedule/import            → Import schedule from file/template
├── GET    /schedule/export            → Export schedule as iCal/JSON

📝 PERIOD LOGS
├── GET    /logs                       → List all period logs (paginated)
├── GET    /logs/date/:date            → Get logs for specific date
├── GET    /logs/:slotId/:date         → Get specific period log
├── POST   /logs                       → Create period log
├── PUT    /logs/:id                   → Update period log
├── DELETE /logs/:id                   → Delete period log
├── GET    /logs/subject/:subjectId    → Get logs by subject

📋 ASSIGNMENTS
├── GET    /assignments                → List all assignments
├── GET    /assignments/pending        → List pending assignments
├── GET    /assignments/overdue        → List overdue assignments
├── GET    /assignments/:id            → Get assignment details
├── POST   /assignments                → Create assignment
├── PUT    /assignments/:id            → Update assignment
├── DELETE /assignments/:id            → Delete assignment
├── PATCH  /assignments/:id/status     → Update status only
├── POST   /assignments/:id/attachments → Upload attachment
├── DELETE /assignments/:id/attachments/:attachmentId → Delete attachment

📚 EXAMS
├── GET    /exams                      → List all exams
├── GET    /exams/upcoming             → List upcoming exams
├── GET    /exams/:id                  → Get exam details
├── POST   /exams                      → Create exam
├── PUT    /exams/:id                  → Update exam
├── DELETE /exams/:id                  → Delete exam
├── PATCH  /exams/:id/prep-status      → Update prep status
├── GET    /exams/:id/questions        → Get important questions
├── POST   /exams/:id/questions        → Add important question

🧪 LAB RECORDS
├── GET    /labs                       → List all lab records
├── GET    /labs/subject/:subjectId    → List by subject
├── GET    /labs/:id                   → Get lab record details
├── POST   /labs                       → Create lab record
├── PUT    /labs/:id                   → Update lab record
├── DELETE /labs/:id                   → Delete lab record
├── PATCH  /labs/:id/status            → Update status (printed, submitted, etc.)
├── POST   /labs/:id/attachments       → Upload attachment

📁 DOCUMENTS (Vault)
├── GET    /documents                  → List all documents
├── GET    /documents/folders          → List virtual folders
├── GET    /documents/folder/:path     → List documents in folder
├── GET    /documents/:id              → Get document details
├── GET    /documents/:id/download     → Download document
├── POST   /documents                  → Upload document
├── PUT    /documents/:id              → Update document metadata
├── DELETE /documents/:id              → Delete document
├── POST   /documents/:id/share        → Share document

📖 STUDY PLANNING
├── GET    /study-plans                → List all study plans
├── GET    /study-plans/date/:date     → Get plan for specific date
├── GET    /study-plans/:id            → Get study plan details
├── POST   /study-plans                → Create study plan
├── PUT    /study-plans/:id            → Update study plan
├── DELETE /study-plans/:id            → Delete study plan
├── GET    /study-plans/:id/sessions   → Get sessions in plan
├── POST   /study-sessions             → Create study session
├── PUT    /study-sessions/:id         → Update study session
├── PATCH  /study-sessions/:id/start   → Start a session
├── PATCH  /study-sessions/:id/end     → End a session
├── POST   /study-plans/generate       → AI-generate study plan

💼 PLACEMENT
├── GET    /placements                 → List all placement drives
├── GET    /placements/upcoming        → List upcoming drives
├── GET    /placements/:id             → Get drive details
├── POST   /placements                 → Add placement drive
├── PUT    /placements/:id             → Update drive details
├── DELETE /placements/:id             → Delete drive
├── PATCH  /placements/:id/status      → Update application status

🤖 AI ASSISTANT
├── GET    /assistant/conversations    → List all conversations
├── GET    /assistant/conversations/:id → Get conversation with messages
├── POST   /assistant/conversations    → Create new conversation
├── DELETE /assistant/conversations/:id → Delete conversation
├── POST   /assistant/chat             → Send message & get response
├── POST   /assistant/quick-actions    → Quick AI actions (summarize, explain, etc.)
├── GET    /assistant/context          → Get current context for AI
├── POST   /assistant/generate-plan    → AI-generate study/revision plan

🔔 NOTIFICATIONS
├── GET    /notifications              → List all notifications
├── GET    /notifications/unread       → List unread notifications
├── PATCH  /notifications/:id/read     → Mark as read
├── PATCH  /notifications/read-all     → Mark all as read
├── DELETE /notifications/:id          → Delete notification
├── POST   /notifications/subscribe    → Register push subscription
├── DELETE /notifications/subscribe    → Remove push subscription
├── GET    /notifications/preferences  → Get notification preferences
├── PUT    /notifications/preferences  → Update preferences

📊 ANALYTICS & DASHBOARD
├── GET    /dashboard                  → Get dashboard summary
├── GET    /dashboard/today            → Today's briefing
├── GET    /analytics/study            → Study time analytics
├── GET    /analytics/assignments      → Assignment completion stats
├── GET    /analytics/attendance       → Attendance stats
├── GET    /analytics/streak           → Streak & gamification stats
├── GET    /analytics/weekly-report    → Weekly summary report
├── GET    /analytics/export           → Export analytics data

🏆 GAMIFICATION
├── GET    /achievements               → List all achievements
├── GET    /achievements/unlocked      → List unlocked achievements
├── GET    /leaderboard                → Get leaderboard
├── GET    /stats                      → Get user stats (XP, level, streak)

👥 SOCIAL
├── GET    /friends                    → List friends
├── POST   /friends/request/:userId    → Send friend request
├── PATCH  /friends/accept/:userId     → Accept friend request
├── DELETE /friends/:userId            → Remove friend
├── GET    /groups                     → List study groups
├── POST   /groups                     → Create study group
├── POST   /groups/:id/join            → Join group (with invite code)
├── DELETE /groups/:id/leave           → Leave group

📚 SUBJECTS & STAFF
├── GET    /subjects                   → List all subjects
├── GET    /subjects/:id               → Get subject details
├── GET    /staff                      → List all staff
├── GET    /staff/:id                  → Get staff details

└─────────────────────────────────────────────────────────────────────────────────┘
API Response Format
Go

// Standard API Response Structure
type APIResponse struct {
    Success   bool        `json:"success"`
    Data      interface{} `json:"data,omitempty"`
    Error     *APIError   `json:"error,omitempty"`
    Meta      *APIMeta    `json:"meta,omitempty"`
    Timestamp int64       `json:"timestamp"`
}

type APIError struct {
    Code    string            `json:"code"`
    Message string            `json:"message"`
    Details map[string]string `json:"details,omitempty"`
}

type APIMeta struct {
    Page       int `json:"page,omitempty"`
    PerPage    int `json:"per_page,omitempty"`
    TotalPages int `json:"total_pages,omitempty"`
    TotalItems int `json:"total_items,omitempty"`
}
🛠️ TECHNOLOGY STACK
Backend (Go)
text

┌─────────────────────────────────────────────────────────────────────────────────┐
│                              GO BACKEND STACK                                   │
├─────────────────────────────────────────────────────────────────────────────────┤
│                                                                                  │
│  CORE FRAMEWORK                                                                  │
│  ├── Go 1.22+                    → Latest Go version                            │
│  ├── Fiber v2                    → Fast HTTP framework (Express-like)           │
│  └── Air                         → Hot reload for development                   │
│                                                                                  │
│  DATABASE & ORM                                                                  │
│  ├── PostgreSQL 16               → Primary database                             │
│  ├── GORM v2                     → ORM (or sqlc for type-safe SQL)             │
│  ├── golang-migrate              → Database migrations                          │
│  └── pgx                         → PostgreSQL driver                            │
│                                                                                  │
│  CACHING & SESSIONS                                                              │
│  ├── Redis 7                     → Caching, sessions, rate limiting            │
│  ├── go-redis/redis              → Redis client                                 │
│  └── scs                         → Session management                           │
│                                                                                  │
│  AUTHENTICATION                                                                  │
│  ├── golang-jwt/jwt              → JWT handling                                 │
│  ├── bcrypt                      → Password hashing                             │
│  └── goth                        → OAuth providers (Google, GitHub)            │
│                                                                                  │
│  FILE STORAGE                                                                    │
│  ├── MinIO                       → S3-compatible object storage                 │
│  ├── minio-go                    → MinIO client                                 │
│  └── Cloudinary (optional)       → Image optimization & CDN                     │
│                                                                                  │
│  BACKGROUND JOBS                                                                 │
│  ├── Asynq                       → Redis-based task queue                       │
│  └── robfig/cron                 → Cron job scheduler                           │
│                                                                                  │
│  NOTIFICATIONS                                                                   │
│  ├── webpush-go                  → Web push notifications                       │
│  ├── gomail                      → Email sending                                │
│  └── Firebase Admin SDK          → FCM for mobile push                          │
│                                                                                  │
│  VALIDATION & UTILS                                                              │
│  ├── go-playground/validator     → Struct validation                            │
│  ├── spf13/viper                 → Configuration management                     │
│  ├── uber-go/zap                 → Structured logging                           │
│  └── google/uuid                 → UUID generation                              │
│                                                                                  │
│  API DOCUMENTATION                                                               │
│  ├── swaggo/swag                 → Swagger/OpenAPI generation                   │
│  └── scalar                      → Beautiful API docs UI                        │
│                                                                                  │
│  TESTING                                                                         │
│  ├── testify                     → Assertions & mocking                         │
│  ├── testcontainers-go           → Integration testing with containers          │
│  └── go-sqlmock                  → Database mocking                             │
│                                                                                  │
└─────────────────────────────────────────────────────────────────────────────────┘
Frontend (Next.js 14)
text

┌─────────────────────────────────────────────────────────────────────────────────┐
│                            FRONTEND STACK                                        │
├─────────────────────────────────────────────────────────────────────────────────┤
│                                                                                  │
│  CORE FRAMEWORK                                                                  │
│  ├── Next.js 14                  → React framework with App Router              │
│  ├── React 18                    → UI library                                   │
│  ├── TypeScript 5                → Type safety                                  │
│  └── Bun                         → Fast package manager & runtime               │
│                                                                                  │
│  STYLING                                                                         │
│  ├── Tailwind CSS 3.4            → Utility-first CSS                            │
│  ├── shadcn/ui                   → Beautiful component library                  │
│  ├── Radix UI                    → Accessible primitives                        │
│  ├── Framer Motion               → Animations                                   │
│  └── lucide-react                → Icons                                        │
│                                                                                  │
│  STATE MANAGEMENT                                                                │
│  ├── TanStack Query v5           → Server state management                      │
│  ├── Zustand                     → Client state management                      │
│  └── React Hook Form             → Form handling                                │
│                                                                                  │
│  DATA FETCHING                                                                   │
│  ├── Axios                       → HTTP client                                  │
│  └── SWR (alternative)           → Data fetching hooks                          │
│                                                                                  │
│  UI COMPONENTS                                                                   │
│  ├── @fullcalendar/react         → Calendar views                               │
│  ├── react-big-calendar          → Alternative calendar                         │
│  ├── recharts                    → Charts & analytics                           │
│  ├── react-dropzone              → File uploads                                 │
│  ├── react-pdf                   → PDF viewing                                  │
│  ├── @tiptap/react               → Rich text editor                             │
│  ├── date-fns                    → Date manipulation                            │
│  └── react-hot-toast             → Toast notifications                          │
│                                                                                  │
│  PWA                                                                             │
│  ├── next-pwa                    → PWA configuration                            │
│  ├── workbox                     → Service worker                               │
│  └── idb-keyval                  → IndexedDB for offline                        │
│                                                                                  │
│  REAL-TIME                                                                       │
│  └── socket.io-client            → WebSocket for real-time updates             │
│                                                                                  │
│  TESTING                                                                         │
│  ├── Vitest                      → Unit testing                                 │
│  ├── Playwright                  → E2E testing                                  │
│  └── MSW                         → API mocking                                  │
│                                                                                  │
└─────────────────────────────────────────────────────────────────────────────────┘
DevOps & Infrastructure
text

┌─────────────────────────────────────────────────────────────────────────────────┐
│                          INFRASTRUCTURE STACK                                    │
├─────────────────────────────────────────────────────────────────────────────────┤
│                                                                                  │
│  CONTAINERIZATION                                                                │
│  ├── Docker                      → Container runtime                            │
│  ├── Docker Compose              → Local development                            │
│  └── Podman (alternative)        → Rootless containers                          │
│                                                                                  │
│  ORCHESTRATION (Future)                                                          │
│  ├── Kubernetes                  → Container orchestration                      │
│  └── K3s                         → Lightweight Kubernetes                       │
│                                                                                  │
│  REVERSE PROXY                                                                   │
│  ├── Traefik                     → Dynamic reverse proxy                        │
│  └── Caddy                       → Alternative with auto-SSL                    │
│                                                                                  │
│  CI/CD                                                                           │
│  ├── GitHub Actions              → CI/CD pipelines                              │
│  ├── GoReleaser                  → Go binary releases                           │
│  └── Docker Hub / GHCR           → Container registry                           │
│                                                                                  │
│  HOSTING OPTIONS                                                                 │
│  ├── Railway                     → Easy deployment                              │
│  ├── Render                      → Alternative PaaS                             │
│  ├── DigitalOcean                → VPS option                                   │
│  ├── AWS (EC2/ECS)               → Enterprise option                            │
│  └── Self-hosted (VPS)           → Full control                                 │
│                                                                                  │
│  MONITORING                                                                      │
│  ├── Prometheus                  → Metrics collection                           │
│  ├── Grafana                     → Dashboards                                   │
│  ├── Loki                        → Log aggregation                              │
│  └── Sentry                      → Error tracking                               │
│                                                                                  │
│  SECURITY                                                                        │
│  ├── Let's Encrypt               → Free SSL                                     │
│  ├── Vault (optional)            → Secrets management                           │
│  └── Fail2Ban                    → Intrusion prevention                         │
│                                                                                  │
└─────────────────────────────────────────────────────────────────────────────────┘
📁 PROJECT STRUCTURE
Go Backend Structure
text

campus-pilot/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
│
├── internal/
│   ├── config/
│   │   ├── config.go               # Configuration struct
│   │   └── loader.go               # Viper config loader
│   │
│   ├── database/
│   │   ├── postgres.go             # PostgreSQL connection
│   │   ├── redis.go                # Redis connection
│   │   └── migrations/             # SQL migrations
│   │       ├── 000001_init.up.sql
│   │       └── 000001_init.down.sql
│   │
│   ├── models/
│   │   ├── user.go
│   │   ├── subject.go
│   │   ├── staff.go
│   │   ├── timetable.go
│   │   ├── assignment.go
│   │   ├── exam.go
│   │   ├── lab_record.go
│   │   ├── document.go
│   │   ├── study_session.go
│   │   ├── notification.go
│   │   └── chat.go
│   │
│   ├── repository/
│   │   ├── interfaces.go           # Repository interfaces
│   │   ├── user_repo.go
│   │   ├── schedule_repo.go
│   │   ├── assignment_repo.go
│   │   ├── exam_repo.go
│   │   ├── lab_repo.go
│   │   ├── document_repo.go
│   │   └── chat_repo.go
│   │
│   ├── services/
│   │   ├── auth_service.go
│   │   ├── schedule_service.go
│   │   ├── assignment_service.go
│   │   ├── exam_service.go
│   │   ├── lab_service.go
│   │   ├── document_service.go
│   │   ├── notification_service.go
│   │   ├── ai_service.go           # LLM integration
│   │   ├── storage_service.go      # File upload/download
│   │   └── analytics_service.go
│   │
│   ├── handlers/
│   │   ├── auth_handler.go
│   │   ├── schedule_handler.go
│   │   ├── assignment_handler.go
│   │   ├── exam_handler.go
│   │   ├── lab_handler.go
│   │   ├── document_handler.go
│   │   ├── notification_handler.go
│   │   ├── ai_handler.go
│   │   ├── dashboard_handler.go
│   │   └── analytics_handler.go
│   │
│   ├── middleware/
│   │   ├── auth.go                 # JWT validation
│   │   ├── cors.go
│   │   ├── ratelimit.go
│   │   ├── logger.go
│   │   └── recovery.go
│   │
│   ├── routes/
│   │   └── routes.go               # All route definitions
│   │
│   ├── utils/
│   │   ├── jwt.go
│   │   ├── password.go
│   │   ├── validator.go
│   │   ├── response.go
│   │   └── helpers.go
│   │
│   └── workers/
│       ├── scheduler.go            # Cron jobs
│       ├── reminder_worker.go      # Send reminders
│       ├── digest_worker.go        # Daily digests
│       └── cleanup_worker.go       # Data cleanup
│
├── pkg/
│   ├── llm/
│   │   └── client.go               # Your SVCE AI client
│   └── push/
│       └── webpush.go              # Web push notifications
│
├── api/
│   └── openapi.yaml                # OpenAPI specification
│
├── scripts/
│   ├── migrate.sh
│   └── seed.sh
│
├── docker/
│   ├── Dockerfile
│   ├── Dockerfile.dev
│   └── docker-compose.yml
│
├── .env.example
├── .air.toml                       # Air hot reload config
├── go.mod
├── go.sum
├── Makefile
└── README.md
Next.js Frontend Structure
text

campus-pilot-web/
├── src/
│   ├── app/
│   │   ├── (auth)/
│   │   │   ├── login/
│   │   │   │   └── page.tsx
│   │   │   ├── register/
│   │   │   │   └── page.tsx
│   │   │   └── layout.tsx
│   │   │
│   │   ├── (dashboard)/
│   │   │   ├── layout.tsx          # Dashboard layout with sidebar
│   │   │   ├── page.tsx            # Main dashboard
│   │   │   │
│   │   │   ├── schedule/
│   │   │   │   ├── page.tsx        # Timetable view
│   │   │   │   └── [date]/
│   │   │   │       └── page.tsx    # Specific date view
│   │   │   │
│   │   │   ├── assignments/
│   │   │   │   ├── page.tsx        # List all assignments
│   │   │   │   ├── new/
│   │   │   │   │   └── page.tsx    # Create assignment
│   │   │   │   └── [id]/
│   │   │   │       └── page.tsx    # Assignment detail
│   │   │   │
│   │   │   ├── exams/
│   │   │   │   ├── page.tsx
│   │   │   │   └── [id]/
│   │   │   │       └── page.tsx
│   │   │   │
│   │   │   ├── labs/
│   │   │   │   ├── page.tsx
│   │   │   │   └── [id]/
│   │   │   │       └── page.tsx
│   │   │   │
│   │   │   ├── documents/
│   │   │   │   ├── page.tsx
│   │   │   │   └── upload/
│   │   │   │       └── page.tsx
│   │   │   │
│   │   │   ├── study/
│   │   │   │   ├── page.tsx        # Study planner
│   │   │   │   └── saturday/
│   │   │   │       └── page.tsx    # Saturday session planner
│   │   │   │
│   │   │   ├── assistant/
│   │   │   │   └── page.tsx        # AI chat interface
│   │   │   │
│   │   │   ├── placements/
│   │   │   │   └── page.tsx
│   │   │   │
│   │   │   ├── analytics/
│   │   │   │   └── page.tsx
│   │   │   │
│   │   │   └── settings/
│   │   │       ├── page.tsx
│   │   │       ├── profile/
│   │   │       │   └── page.tsx
│   │   │       └── notifications/
│   │   │           └── page.tsx
│   │   │
│   │   ├── api/                    # API routes (if needed)
│   │   │   └── push/
│   │   │       └── route.ts
│   │   │
│   │   ├── layout.tsx              # Root layout
│   │   ├── globals.css
│   │   └── manifest.json           # PWA manifest
│   │
│   ├── components/
│   │   ├── ui/                     # shadcn/ui components
│   │   │   ├── button.tsx
│   │   │   ├── card.tsx
│   │   │   ├── dialog.tsx
│   │   │   ├── dropdown-menu.tsx
│   │   │   ├── input.tsx
│   │   │   ├── select.tsx
│   │   │   ├── tabs.tsx
│   │   │   ├── toast.tsx
│   │   │   └── ...
│   │   │
│   │   ├── layout/
│   │   │   ├── sidebar.tsx
│   │   │   ├── header.tsx
│   │   │   ├── footer.tsx
│   │   │   └── mobile-nav.tsx
│   │   │
│   │   ├── schedule/
│   │   │   ├── timetable-grid.tsx
│   │   │   ├── period-card.tsx
│   │   │   ├── day-view.tsx
│   │   │   └── week-view.tsx
│   │   │
│   │   ├── assignments/
│   │   │   ├── assignment-card.tsx
│   │   │   ├── assignment-form.tsx
│   │   │   ├── assignment-list.tsx
│   │   │   └── deadline-badge.tsx
│   │   │
│   │   ├── exams/
│   │   │   ├── exam-card.tsx
│   │   │   ├── exam-countdown.tsx
│   │   │   └── syllabus-tracker.tsx
│   │   │
│   │   ├── labs/
│   │   │   ├── lab-record-card.tsx
│   │   │   ├── status-pipeline.tsx
│   │   │   └── lab-form.tsx
│   │   │
│   │   ├── documents/
│   │   │   ├── document-card.tsx
│   │   │   ├── upload-zone.tsx
│   │   │   ├── folder-tree.tsx
│   │   │   └── pdf-viewer.tsx
│   │   │
│   │   ├── study/
│   │   │   ├── session-timer.tsx
│   │   │   ├── plan-builder.tsx
│   │   │   └── progress-ring.tsx
│   │   │
│   │   ├── assistant/
│   │   │   ├── chat-interface.tsx
│   │   │   ├── message-bubble.tsx
│   │   │   ├── code-block.tsx
│   │   │   └── quick-actions.tsx
│   │   │
│   │   ├── dashboard/
│   │   │   ├── today-briefing.tsx
│   │   │   ├── upcoming-deadlines.tsx
│   │   │   ├── streak-counter.tsx
│   │   │   ├── quick-stats.tsx
│   │   │   └── recent-activity.tsx
│   │   │
│   │   ├── analytics/
│   │   │   ├── study-chart.tsx
│   │   │   ├── progress-chart.tsx
│   │   │   └── heatmap.tsx
│   │   │
│   │   └── shared/
│   │       ├── loading-spinner.tsx
│   │       ├── empty-state.tsx
│   │       ├── error-boundary.tsx
│   │       ├── confirm-dialog.tsx
│   │       └── file-upload.tsx
│   │
│   ├── hooks/
│   │   ├── use-auth.ts
│   │   ├── use-schedule.ts
│   │   ├── use-assignments.ts
│   │   ├── use-exams.ts
│   │   ├── use-labs.ts
│   │   ├── use-documents.ts
│   │   ├── use-notifications.ts
│   │   ├── use-ai-chat.ts
│   │   └── use-analytics.ts
│   │
│   ├── lib/
│   │   ├── api.ts                  # Axios instance
│   │   ├── auth.ts                 # Auth utilities
│   │   ├── utils.ts                # cn() and helpers
│   │   ├── constants.ts
│   │   └── validations.ts          # Zod schemas
│   │
│   ├── stores/
│   │   ├── auth-store.ts
│   │   ├── schedule-store.ts
│   │   └── ui-store.ts
│   │
│   ├── types/
│   │   ├── index.ts
│   │   ├── user.ts
│   │   ├── schedule.ts
│   │   ├── assignment.ts
│   │   ├── exam.ts
│   │   ├── lab.ts
│   │   └── api.ts
│   │
│   └── providers/
│       ├── query-provider.tsx      # TanStack Query
│       ├── auth-provider.tsx
│       └── theme-provider.tsx
│
├── public/
│   ├── icons/
│   ├── images/
│   └── sw.js                       # Service worker
│
├── .env.local.example
├── next.config.js
├── tailwind.config.ts
├── tsconfig.json
├── package.json
└── README.md

