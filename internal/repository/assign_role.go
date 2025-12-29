package repository

import (
	"context"

	"github.com/google/uuid"
)

func (r *userRepository) AssignRole(ctx context.Context,userID uuid.UUID,roleName string,) error {
	query := `
		INSERT INTO user_roles (user_id, role_id)
		SELECT $1, r.id
		FROM roles r
		WHERE r.name = $2
		ON CONFLICT DO NOTHING;
	`

	_, err := r.db.Exec(ctx, query, userID, roleName)
	return err
}
