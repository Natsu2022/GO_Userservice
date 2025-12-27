-- ----------------------
-- Table: source_list (1:N)
-- ----------------------
CREATE TABLE source_list (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    source_name TEXT,
    ip_address TEXT NOT NULL
);