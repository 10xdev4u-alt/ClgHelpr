-- Migration: 000002_seed_initial_data.up.sql

-- UUIDs for consistent linking
-- Subject IDs
SELECT uuid_generate_v4() AS cns_id, uuid_generate_v4() AS spm_id, uuid_generate_v4() AS iot_id,
       uuid_generate_v4() AS cc_id, uuid_generate_v4() AS cd_id, uuid_generate_v4() AS eda_id,
       uuid_generate_v4() AS cnslab_id, uuid_generate_v4() AS cclab_id, uuid_generate_v4() AS es_id,
       uuid_generate_v4() AS hm_id
INTO TEMP subject_uuids;

-- Staff IDs
SELECT uuid_generate_v4() AS vijayanand_id, uuid_generate_v4() AS banupriya_id, uuid_generate_v4() AS janarthanan_id,
       uuid_generate_v4() AS suriya_id, uuid_generate_v4() AS raghuvaran_id, uuid_generate_v4() AS kapilvani_id,
       uuid_generate_v4() AS poorani_id, uuid_generate_v4() AS rupa_id, uuid_generate_v4() AS selvamani_id,
       uuid_generate_v4() AS arunachalam_id, uuid_generate_v4() AS ashok_id, uuid_generate_v4() AS khanagavalle_id
INTO TEMP staff_uuids;

-- Venue IDs
SELECT uuid_generate_v4() AS cse_lab1_id, uuid_generate_v4() AS room401_id, uuid_generate_v4() AS library_hall_id
INTO TEMP venue_uuids;


-- Insert subjects
INSERT INTO subjects (id, code, name, short_name, type, credits, department, semester, color) VALUES
( (SELECT cns_id FROM subject_uuids), 'CS22601', 'Cryptography and Network Security', 'CNS', 'core', 4, 'CSE', 6, '#EF4444'),
( (SELECT spm_id FROM subject_uuids), 'CS22602', 'Software Project Management', 'SPM', 'core', 3, 'CSE', 6, '#F59E0B'),
( (SELECT iot_id FROM subject_uuids), 'AD22501', 'Internet of Things and Applications', 'IoT', 'core', 4, 'CSE', 6, '#10B981'),
( (SELECT cc_id FROM subject_uuids), 'CS22603', 'Cloud Computing', 'CC', 'core', 4, 'CSE', 6, '#3B82F6'),
( (SELECT cd_id FROM subject_uuids), 'CS22604', 'Compiler Design', 'CD', 'core', 4, 'CSE', 6, '#8B5CF6'),
( (SELECT eda_id FROM subject_uuids), 'CS22021', 'Exploratory Data Analysis', 'EDA', 'elective', 3, 'CSE', 6, '#EC4899'),
( (SELECT cnslab_id FROM subject_uuids), 'CS22611', 'CNS Laboratory', 'CNS Lab', 'lab', 2, 'CSE', 6, '#EF4444'),
( (SELECT cclab_id FROM subject_uuids), 'CS22612', 'Cloud Computing Laboratory', 'CC Lab', 'lab', 2, 'CSE', 6, '#3B82F6'),
( (SELECT es_id FROM subject_uuids), 'OE22705', 'Embedded Systems and its Application', 'ES', 'open_elective', 3, 'ECE', 6, '#6366F1'),
( (SELECT hm_id FROM subject_uuids), 'HM001', 'Honor/Minor Subject', 'H/M', 'honor', 3, 'CSE', 6, '#14B8A6');

-- Insert staff
INSERT INTO staff (id, name, title, department, designation) VALUES
( (SELECT vijayanand_id FROM staff_uuids), 'Dr. S. Vijayanand', 'Dr.', 'ECE', 'Associate Professor'),
( (SELECT banupriya_id FROM staff_uuids), 'Ms. Banupriya P', 'Ms.', 'CSE', 'Assistant Professor'),
( (SELECT janarthanan_id FROM staff_uuids), 'Dr. Janarthanan P', 'Dr.', 'CSE', 'Associate Professor'),
( (SELECT suriya_id FROM staff_uuids), 'Ms. N. Suriya', 'Ms.', 'External', 'Trainer - VYVoxel'),
( (SELECT raghuvaran_id FROM staff_uuids), 'Mr. E. Raghuvaran', 'Mr.', 'CSE', 'Assistant Professor'),
( (SELECT kapilvani_id FROM staff_uuids), 'Ms. Kapilvani R K', 'Ms.', 'CSE', 'Assistant Professor'),
( (SELECT poorani_id FROM staff_uuids), 'Dr. Poorani S', 'Dr.', 'CSE', 'Assistant Professor'),
( (SELECT rupa_id FROM staff_uuids), 'Ms. Rupa Kesavan', 'Ms.', 'CSE', 'Assistant Professor'),
( (SELECT selvamani_id FROM staff_uuids), 'Mr. Selvamani P', 'Mr.', 'CSE', 'Assistant Professor'),
( (SELECT arunachalam_id FROM staff_uuids), 'Mr. Arunachalam Narayanan', 'Mr.', 'CSE', 'Assistant Professor'),
( (SELECT ashok_id FROM staff_uuids), 'Mr. S. Ashok Kumar', 'Mr.', 'Placement', 'Training Officer'),
( (SELECT khanagavalle_id FROM staff_uuids), 'Ms. G.R. Khanagavalle', 'Ms.', 'Library', 'Librarian');

-- Insert sample venues
INSERT INTO venues (id, name, building, floor, capacity, type) VALUES
( (SELECT cse_lab1_id FROM venue_uuids), 'CSE Lab 1', 'Main Block', 2, 60, 'lab'),
( (SELECT room401_id FROM venue_uuids), 'Room 401', 'Main Block', 4, 70, 'classroom'),
( (SELECT library_hall_id FROM venue_uuids), 'Central Library', 'Library Block', 1, 200, 'library');

-- Insert sample timetable slots for a dummy user (replace with actual user ID later or through API)
-- Assuming a user with ID 'c3e2e8e0-1b7e-4b7e-8c7e-1c7e2b8e2b8e' will be created by the registration system.
-- For initial seeding, we'll use a placeholder UUID for user_id.
-- Once a real user is registered, we can update these or provide UI for it.
-- Or better: create a default "Prince" user in the migration.

-- Dummy user ID
SELECT uuid_generate_v4() AS prince_id
INTO TEMP user_uuid;

INSERT INTO users (id, email, password_hash, full_name, register_number, department, year, semester, is_verified) VALUES
( (SELECT prince_id FROM user_uuid), 'prince@svce.ac.in', '$2a$10$wY.u9fH5R/Lp7x7C.a.n.e.s.e.c.r.e.t.P.a.s.s.w.o.r.d.F.o.r.P.r.i.n.c.e', 'PrinceTheProgrammer', '20UCS001', 'CSE', 3, 6, TRUE);

-- Timetable slots for Prince (based on provided schedule - Monday example)
-- Note: Start and End times are 'time without timezone' and only date part will be ignored by Go's time.Time
-- Monday (Day 1)
INSERT INTO timetable_slots (id, user_id, subject_id, staff_id, venue_id, day_of_week, start_time, end_time, period_number, slot_type, is_recurring, is_active) VALUES
(uuid_generate_v4(), (SELECT prince_id FROM user_uuid), (SELECT es_id FROM subject_uuids), (SELECT vijayanand_id FROM staff_uuids), (SELECT room401_id FROM venue_uuids), 1, '08:30:00', '09:20:00', 1, 'lecture', TRUE, TRUE),
(uuid_generate_v4(), (SELECT cd_id FROM subject_uuids), (SELECT banupriya_id FROM staff_uuids), (SELECT room401_id FROM venue_uuids), 1, '09:20:00', '10:10:00', 2, 'lecture', TRUE, TRUE),
(uuid_generate_v4(), (SELECT iot_id FROM subject_uuids), (SELECT janarthanan_id FROM staff_uuids), (SELECT room401_id FROM venue_uuids), 1, '10:25:00', '11:15:00', 3, 'lecture', TRUE, TRUE),
(uuid_generate_v4(), (SELECT eda_id FROM subject_uuids), (SELECT suriya_id FROM staff_uuids), (SELECT room401_id FROM venue_uuids), 1, '11:15:00', '12:05:00', 4, 'lecture', TRUE, TRUE),
(uuid_generate_v4(), (SELECT cclab_id FROM subject_uuids), (SELECT raghuvaran_id FROM staff_uuids), (SELECT cse_lab1_id FROM venue_uuids), 1, '12:45:00', '15:15:00', 5, 'lab', TRUE, TRUE); -- Spans 3 periods

-- Drop temporary tables
DROP TABLE subject_uuids;
DROP TABLE staff_uuids;
DROP TABLE venue_uuids;
DROP TABLE user_uuid;
