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

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	return s.repo.Create(ctx, req, string(hash))
}
