-- ----------------------
-- Table: preference (1:1)
-- ----------------------
CREATE TABLE preference (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    theme TEXT NOT NULL DEFAULT 'System',
    language TEXT NOT NULL DEFAULT 'ENG',
    email_noti BOOLEAN NOT NULL DEFAULT TRUE,
    sms_noti BOOLEAN NOT NULL DEFAULT FALSE
);