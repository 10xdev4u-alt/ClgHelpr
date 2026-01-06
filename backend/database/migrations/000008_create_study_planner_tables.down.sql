-- Migration: 000008_create_study_planner_tables.down.sql

DROP TRIGGER IF EXISTS update_study_sessions_updated_at ON study_sessions;
DROP TRIGGER IF EXISTS update_study_plans_updated_at ON study_plans;
DROP TABLE IF EXISTS study_sessions;
DROP TABLE IF EXISTS study_plans;
