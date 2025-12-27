-- ----------------------
-- Table: permissions
-- ----------------------
CREATE TABLE permissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    permission TEXT NOT NULL,
    description TEXT
);