-- Migration: 000005_create_exam_tables.down.sql

DROP TRIGGER IF EXISTS update_exams_updated_at ON exams;
DROP TABLE IF EXISTS important_questions;
DROP TABLE IF EXISTS exams;
