package model

import (
	"time"

	"github.com/google/uuid"
)

// User table
type User struct {
	ID                          uuid.UUID `json:"id"`
	FirstName                   string    `json:"first_name"`
	MiddleName                  string    `json:"middle_name"`
	LastName                    string    `json:"last_name"`
	DisplayName                 string    `json:"display_name"`
	RoleID                      uuid.UUID `json:"role_id"`
	Email                       string    `json:"email"`
	EmailVerify                 bool      `json:"email_verify"`
	Password                    string    `json:"password"` // hashed
	PhoneNumber                 string    `json:"phone_number"`
	AccountCreated              time.Time `json:"account_created"`
	AccountStatus               bool      `json:"account_status"`
	SignupSource                string    `json:"signup_source"`
	ReferralCode                string    `json:"referral_code"`
	SourceList                  string    `json:"source_list"` // JSONB
	PhysicalGender              string    `json:"physical_gender"`
	Birthday                    string    `json:"birthday"`
	ProfilePictureURL           string    `json:"profile_picture_url"`
	ProfilePictureBackgroundURL string    `json:"profile_picture_background_url"`
	IsDeleted                   bool      `json:"is_deleted"`
}

type CreateUserRequest struct {
	FirstName      string `json:"first_name"`
	MiddleName     string `json:"middle_name"`
	LastName       string `json:"last_name"`
	PhysicalGender string `json:"physical_gender"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	PhoneNumber    string `json:"phone_number"`
	SignupSource   string `json:"signup_source"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SourceList table
type SourceList struct {
	ID         int       `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	SourceName string    `json:"source_name"`
	IPAddress  string    `json:"ip_address"`
}

// Location table
type Location struct {
	UserID   uuid.UUID `json:"user_id"`
	Country  string    `json:"country"`
	Province string    `json:"province"`
	City     string    `json:"city"`
	Street   string    `json:"street"`
	ZipCode  string    `json:"zip_code"`
}

// Garbage table
type Garbage struct {
	UserID    uuid.UUID `json:"user_id"`
	DeletedAt time.Time `json:"deleted_at"`
	DeletedBy uuid.UUID `json:"deleted_by"` // ถ้า FK เป็น users.id
}

// Session table
type Session struct {
	ID         uuid.UUID `json:"id"` // ถ้าอยากเปลี่ยนเป็น UUID ก็ uuid.UUID
	UserID     uuid.UUID `json:"user_id"`
	IPAddress  string    `json:"ip_address"`
	SignInWith string    `json:"signin_with"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiresAt  time.Time `json:"expires_at"`
}

// Preference table
type Preference struct {
	UserID    uuid.UUID `json:"user_id"`
	Theme     string    `json:"theme"`
	Language  string    `json:"language"`
	EmailNoti bool      `json:"email_noti"`
	SMSNoti   bool      `json:"sms_noti"`
}
