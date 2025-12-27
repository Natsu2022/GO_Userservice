-- ----------------------
-- Table: location (1:1)
-- ----------------------
CREATE TABLE location (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    country TEXT NOT NULL,
    province TEXT NOT NULL,
    city TEXT NOT NULL,
    street TEXT NOT NULL,
    zip_code TEXT NOT NULL
);