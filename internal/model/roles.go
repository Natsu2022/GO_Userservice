package model

import "github.com/google/uuid"

// Role table
type Role struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

// Permission table
type Permission struct {
	ID          int    `db:"id"`
	Permission  string `db:"permission"`
	Description string `db:"description"`
}

// RolePermission table
type RolePermission struct {
	RoleID       int `db:"role_id"`
	PermissionID int `db:"permission_id"`
}

// UserPermission table
type UserPermission struct {
	UserID       uuid.UUID `db:"user_id"`
	PermissionID int       `db:"permission_id"`
}
