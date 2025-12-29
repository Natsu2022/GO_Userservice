-- ----------------------
-- Table: user_permissions (many-to-many)
-- ----------------------
CREATE TABLE user_permissions (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    permission_id UUID NOT NULL REFERENCES permissions(id) ON DELETE CASCADE,
    PRIMARY KEY(user_id, permission_id)
);