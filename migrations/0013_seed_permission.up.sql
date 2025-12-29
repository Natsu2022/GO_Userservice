INSERT INTO permissions (name, description)
VALUES
    -- System level
    ('system.manage', 'Full system management'),

    -- User management
    ('user.read', 'Read user information'),
    ('user.create', 'Create new users'),
    ('user.update', 'Update user information'),
    ('user.delete', 'Delete users'),

    -- Role management
    ('role.read', 'Read roles'),
    ('role.assign', 'Assign roles to users'),

    -- Permission management
    ('permission.read', 'Read permissions'),
    ('permission.assign', 'Assign permissions to roles')

ON CONFLICT (name) DO NOTHING;