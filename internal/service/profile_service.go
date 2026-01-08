package service

import (
	"context"
	"errors"

	"OSCIRA.com/service/user_service/internal/model"
	"github.com/google/uuid"
)

func (s *userService) UpdateMyProfile(
	ctx context.Context,
	userID uuid.UUID,
	req model.UpdateProfileRequest,
) error {

	if req.PhysicalGender != nil {
		valid := map[string]bool{
			"male":   true,
			"female": true,
			"other":  true,
		}
		if !valid[*req.PhysicalGender] {
			return errors.New("invalid physical_gender")
		}
	}

	return s.repo.UpdateProfile(ctx, userID, req)
}

func (s *userService) GetMyProfile(
	ctx context.Context,
	userID uuid.UUID,
) (*model.UserProfile, error) {
	return s.repo.GetProfileByUserID(ctx, userID)
}
