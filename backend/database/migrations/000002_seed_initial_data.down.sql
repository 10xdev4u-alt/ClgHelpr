-- Migration: 000002_seed_initial_data.down.sql

-- Delete timetable slots for Prince
DELETE FROM timetable_slots WHERE user_id = (SELECT id FROM users WHERE email = 'prince@svce.ac.in');

-- Delete the seeded Prince user
DELETE FROM users WHERE email = 'prince@svce.ac.in';

-- Delete sample venues
DELETE FROM venues WHERE name IN ('CSE Lab 1', 'Room 401', 'Central Library');

-- Delete staff members
DELETE FROM staff WHERE name IN (
'Dr. S. Vijayanand', 'Ms. Banupriya P', 'Dr. Janarthanan P', 'Ms. N. Suriya',
'Mr. E. Raghuvaran', 'Ms. Kapilvani R K', 'Dr. Poorani S', 'Ms. Rupa Kesavan',
'Mr. Selvamani P', 'Mr. Arunachalam Narayanan', 'Mr. S. Ashok Kumar', 'Ms. G.R. Khanagavalle'
);

-- Delete subjects
DELETE FROM subjects WHERE code IN (
'CS22601', 'CS22602', 'AD22501', 'CS22603', 'CS22604', 'CS22021',
'CS22611', 'CS22612', 'OE22705', 'HM001'
);
