-- Migration: 000004_create_assignment_tables.down.sql

DROP TRIGGER IF EXISTS update_assignments_updated_at ON assignments;
DROP TABLE IF EXISTS assignment_attachments;
DROP TABLE IF EXISTS assignments;
