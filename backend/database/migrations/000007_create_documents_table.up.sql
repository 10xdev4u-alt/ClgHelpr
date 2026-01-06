-- Migration: 000007_create_documents_table.up.sql

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
    storage_key VARCHAR(500), -- S3/MinIO key
    
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

-- Apply the auto-update trigger to the new documents table
CREATE TRIGGER update_documents_updated_at BEFORE UPDATE ON documents
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
