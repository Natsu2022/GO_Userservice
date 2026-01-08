package model

import "github.com/google/uuid"

// Role table
type Role struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}

// Permission table
type Permission struct {
	ID          uuid.UUID `db:"id"`
	Permission  string    `db:"permission"`
	Description string    `db:"description"`
}

// RolePermission table
type RolePermission struct {
	RoleID       uuid.UUID `db:"role_id"`
	PermissionID uuid.UUID `db:"permission_id"`
}

// UserPermission table
type UserPermission struct {
	UserID       uuid.UUID `db:"user_id"`
	PermissionID uuid.UUID `db:"permission_id"`
}
