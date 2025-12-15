package repository

import (
	"context"

	"OSCIRA.com/service/user_service/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, s model.Session) error
	DeleteSession(ctx context.Context, sessionID uuid.UUID) error
	GetSession(ctx context.Context, sessionID uuid.UUID) (*model.Session, error)
}

type sessionRepository struct {
	db *pgxpool.Pool
}

func NewSessionRepository(db *pgxpool.Pool) SessionRepository {
	return &sessionRepository{db}
}

func (r *sessionRepository) CreateSession(ctx context.Context, s model.Session) error {

	query := `
		INSERT INTO session (id, user_id, ip_address, signin_with, created_at, expires_at)
		VALUES ($1,$2,$3,$4,$5,$6)
	`

	_, err := r.db.Exec(ctx, query,
		s.ID, s.UserID, s.IPAddress, s.SignInWith, s.CreatedAt, s.ExpiresAt,
	)
	return err
}
func (r *sessionRepository) DeleteSession(ctx context.Context, sessionID uuid.UUID) error {
	query := `DELETE FROM session WHERE id = $1`
	_, err := r.db.Exec(ctx, query, sessionID)
	return err
}

func (r *sessionRepository) GetSession(ctx context.Context, sessionID uuid.UUID) (*model.Session, error) {

	query := `
		SELECT id, user_id, ip_address, signin_with, created_at, expires_at
		FROM session
		WHERE id = $1
	`

	var s model.Session

	err := r.db.QueryRow(ctx, query, sessionID).Scan(
		&s.ID,
		&s.UserID,
		&s.IPAddress,
		&s.SignInWith,
		&s.CreatedAt,
		&s.ExpiresAt,
	)

	if err != nil {
		return nil, err
	}

	return &s, nil
}
