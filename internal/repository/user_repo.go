package repository

import (
	"context"

	"OSCIRA.com/service/user_service/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Create(ctx context.Context, req model.CreateUserRequest, hash string) (uuid.UUID, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateAccountStatus(ctx context.Context, userID uuid.UUID, status bool) error
	AssignRole(ctx context.Context, userID uuid.UUID, roleName string) error
	UpdateProfile(ctx context.Context, userID uuid.UUID, req model.UpdateProfileRequest) error
	GetProfileByUserID(ctx context.Context, userID uuid.UUID) (*model.UserProfile, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, req model.CreateUserRequest, hash string) (uuid.UUID, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback(ctx)

	insertUserQuery := `
		INSERT INTO users (
			first_name, middle_name, last_name, display_name,
			physical_gender, email, password, phone_number, signup_source
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING id;
	`

	displayName := req.FirstName + " " + req.LastName
	var id uuid.UUID

	err = tx.QueryRow(ctx, insertUserQuery,
		req.FirstName,
		req.MiddleName,
		req.LastName,
		displayName,
		req.PhysicalGender,
		req.Email,
		hash,
		req.PhoneNumber,
		req.SignupSource,
	).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	insertProfileQuery := `
		INSERT INTO user_profiles (user_id, display_name)
		VALUES ($1, $2);
	`

	if _, err := tx.Exec(ctx, insertProfileQuery, id, displayName); err != nil {
		return uuid.Nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {

	query := `
		SELECT id, password, account_status
		FROM users
		WHERE email = $1
	`

	var u model.User

	err := r.db.QueryRow(ctx, query, email).Scan(
		&u.ID,
		&u.Password,
		&u.AccountStatus,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *userRepository) UpdateAccountStatus(ctx context.Context, userID uuid.UUID, status bool) error {
	query := `UPDATE users SET account_status=$1 WHERE id=$2`
	_, err := r.db.Exec(ctx, query, status, userID)
	return err
}
