INSERT INTO user_roles (user_id, role_id)
SELECT u.id, r.id
FROM users u
JOIN roles r ON r.name = 'superadmin'
WHERE u.email = 'superadmin@system.local'
ON CONFLICT DO NOTHING;
