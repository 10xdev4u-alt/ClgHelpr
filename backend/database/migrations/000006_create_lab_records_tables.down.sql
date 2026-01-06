-- Migration: 000006_create_lab_records_tables.down.sql

DROP TRIGGER IF EXISTS update_lab_records_updated_at ON lab_records;
DROP TABLE IF EXISTS lab_record_attachments;
DROP TABLE IF EXISTS lab_records;
