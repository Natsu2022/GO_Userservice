-- =========================
-- Create superadmin user
-- =========================
INSERT INTO users (
  first_name,
  middle_name,
  last_name,
  display_name,
  email,
  email_verify,
  phone_number,
  password,
  signup_source,
  physical_gender
)
VALUES (
  'Phumin',
  '',
  'Tonglar',
  'The_First_SUPERADMIN',
  'superadmin@system.local',
  true,
  '0635797500',
  '$2a$10$HASHED_PASSWORD',
  'migrate_seed',
  'male'
)
ON CONFLICT (email) DO NOTHING;