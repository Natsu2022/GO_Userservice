-- ----------------------
-- Table: users
-- ----------------------
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name TEXT NOT NULL,
    middle_name TEXT,
    last_name TEXT NOT NULL,
    display_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    email_verify BOOLEAN DEFAULT FALSE,
    password TEXT NOT NULL,
    phone_number TEXT,
    account_created TIMESTAMP NOT NULL DEFAULT now(),
    account_status BOOLEAN DEFAULT FALSE,
    signup_source TEXT NOT NULL,
    referral_code TEXT,
    source_list JSONB DEFAULT '{}',
    physical_gender TEXT,
    birthday DATE,
    profile_picture_url TEXT DEFAULT 'DEFAULT_picture',
    profile_picture_background_url TEXT DEFAULT 'DEFAULT_picture_background',
    is_deleted BOOLEAN DEFAULT FALSE
);