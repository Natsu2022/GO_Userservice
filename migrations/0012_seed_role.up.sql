INSERT INTO roles (name)
VALUES
    ('superadmin'),
    ('admin'),
    ('studen'),
    ('guest')
ON CONFLICT (name) DO NOTHING;
