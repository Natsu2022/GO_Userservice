package service

import (
	"context"
	"errors"
	"regexp"

	"OSCIRA.com/service/user_service/internal/model"
	"OSCIRA.com/service/user_service/internal/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(ctx context.Context, req model.CreateUserRequest) (uuid.UUID, error)
	UpdateMyProfile(ctx context.Context, userID uuid.UUID, req model.UpdateProfileRequest) error
	GetMyProfile(ctx context.Context, userID uuid.UUID) (*model.UserProfile, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(ctx context.Context, req model.CreateUserRequest) (uuid.UUID, error) {

	if req.Email == "" || req.Password == "" {
		return uuid.Nil, errors.New("email and password are required")
	}

	if !regexp.MustCompile(`.+@.+\..+`).MatchString(req.Email) {
		return uuid.Nil, errors.New("invalid email format")
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return uuid.Nil, err
	}

	// 1. create user
	userID, err := s.repo.Create(ctx, req, string(hash))
	if err != nil {
		return uuid.Nil, err
	}

	// 2. assign default role = guest
	if err := s.repo.AssignRole(ctx, userID, "guest"); err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}
