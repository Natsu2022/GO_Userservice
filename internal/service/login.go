package service

import (
	"context"
	"errors"
	"time"

	"OSCIRA.com/service/user_service/internal/model"
	"OSCIRA.com/service/user_service/internal/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, req model.LoginRequest, ip string) (*model.Session, error)
	Logout(ctx context.Context, sessionID uuid.UUID) error
}

type authService struct {
	userRepo    repository.UserRepository
	sessionRepo repository.SessionRepository
}

func NewAuthService(u repository.UserRepository, s repository.SessionRepository) AuthService {
	return &authService{userRepo: u, sessionRepo: s}
}

func (s *authService) Login(ctx context.Context, req model.LoginRequest, ip string) (*model.Session, error) {

	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email and password required")
	}

	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return nil, errors.New("invalid email or password")
	}

	if !user.AccountStatus {
		s.userRepo.UpdateAccountStatus(ctx, user.ID, true)
	}

	session := model.Session{
		ID:         uuid.New(),
		UserID:     user.ID,
		IPAddress:  ip,
		SignInWith: "password",
		CreatedAt:  time.Now(),
		ExpiresAt:  time.Now().Add(24 * time.Hour),
	}

	if err := s.sessionRepo.CreateSession(ctx, session); err != nil {
		return nil, errors.New("failed to create session")
	}

	return &session, nil
}
func (s *authService) Logout(ctx context.Context, sessionID uuid.UUID) error {
	return s.sessionRepo.DeleteSession(ctx, sessionID)
}
