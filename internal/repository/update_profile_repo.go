package repository

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"OSCIRA.com/service/user_service/internal/model"
	"github.com/google/uuid"
)

func (r *userRepository) UpdateProfile(ctx context.Context, userID uuid.UUID, req model.UpdateProfileRequest,
) error {

	set := []string{}
	args := []any{}
	i := 1

	if req.DisplayName != nil {
		set = append(set, fmt.Sprintf("display_name = $%d", i))
		args = append(args, *req.DisplayName)
		i++
	}
	if req.PhoneNumber != nil {
		set = append(set, fmt.Sprintf("phone_number = $%d", i))
		args = append(args, *req.PhoneNumber)
		i++
	}
	if req.PhysicalGender != nil {
		set = append(set, fmt.Sprintf("physical_gender = $%d", i))
		args = append(args, *req.PhysicalGender)
		i++
	}
	if req.Birthday != nil {
		set = append(set, fmt.Sprintf("birthday = $%d", i))
		args = append(args, *req.Birthday)
		i++
	}

	if len(set) == 0 {
		return nil
	}

	query := `
		UPDATE user_profiles
		SET ` + strings.Join(set, ", ") + `,
		    updated_at = now()
		WHERE user_id = $` + strconv.Itoa(i)

	args = append(args, userID)

	_, err := r.db.Exec(ctx, query, args...)
	return err
}

func (r *userRepository) GetProfileByUserID(ctx context.Context, userID uuid.UUID,
) (*model.UserProfile, error) {

	query := `
		SELECT user_id, display_name, phone_number, physical_gender, birthday
		FROM user_profiles
		WHERE user_id = $1
	`

	var p model.UserProfile
	err := r.db.QueryRow(ctx, query, userID).Scan(
		&p.UserID,
		&p.DisplayName,
		&p.PhoneNumber,
		&p.PhysicalGender,
		&p.Birthday,
	)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
