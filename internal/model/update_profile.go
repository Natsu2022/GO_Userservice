package model

import (
	"github.com/google/uuid"
)

type UpdateProfileRequest struct {
	DisplayName    *string `json:"display_name"`
	PhoneNumber    *string `json:"phone_number"`
	PhysicalGender *string `json:"physical_gender"`
	Birthday       *string `json:"birthday"` // YYYY-MM-DD
}

type UserProfile struct {
	UserID         uuid.UUID `json:"user_id"`
	DisplayName    string    `json:"display_name"`
	PhoneNumber    *string   `json:"phone_number,omitempty"`
	PhysicalGender *string   `json:"physical_gender,omitempty"`
	Birthday       *string   `json:"birthday,omitempty"`
}
