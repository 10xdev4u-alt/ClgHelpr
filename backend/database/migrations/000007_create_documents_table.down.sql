-- Migration: 000007_create_documents_table.down.sql

DROP TRIGGER IF EXISTS update_documents_updated_at ON documents;
DROP TABLE IF EXISTS documents;
