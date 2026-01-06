-- Migration: 000002_create_timetable_tables.down.sql

DROP TRIGGER IF EXISTS update_timetable_slots_updated_at ON timetable_slots;
DROP TABLE IF EXISTS timetable_slots;
DROP TABLE IF EXISTS venues;
DROP TABLE IF EXISTS staff;
DROP TABLE IF EXISTS subjects;
