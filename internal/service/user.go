package service

import (
	"context"
	"fmt"

	"github.com/DmitriyZhevnov/rest-api/internal/model"
	"github.com/DmitriyZhevnov/rest-api/internal/repository"
	"github.com/DmitriyZhevnov/rest-api/pkg/hash"
	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
)

type userService struct {
	repository repository.User
	hasher     hash.PasswordHasher
	logger     *logging.Logger
}

func NewUserSevice(hasher hash.PasswordHasher, repos repository.User, logger *logging.Logger) *userService {
	return &userService{
		repository: repos,
		hasher:     hasher,
		logger:     logger,
	}
}

func (s *userService) Create(ctx context.Context, dto model.CreateUserDTO) (string, error) {
	passwordHash, err := s.hasher.Hash(dto.Password)
	if err != nil {
		return "", err
	}

	user := model.User{
		Email:        dto.Email,
		Username:     dto.Username,
		PasswordHash: passwordHash,
	}

	return s.repository.Create(ctx, user)
}

func (s *userService) FindUser(ctx context.Context, id string) (model.User, error) {
	return s.repository.FindOne(ctx, id)
}

func (s *userService) FindAll(ctx context.Context) ([]model.User, error) {
	return s.repository.FindAll(ctx)
}

func (s *userService) Update(ctx context.Context, id string, dto model.UpdateUserDTO) error {
	if dto.Email == "" && dto.Password == "" && dto.Username == "" {
		// TODO: add new error
		return fmt.Errorf("invalid body")
	}

	passwordHash, err := s.hasher.Hash(dto.Password)
	if err != nil {
		return err
	}

	user := model.User{
		ID:           id,
		Email:        dto.Email,
		Username:     dto.Username,
		PasswordHash: passwordHash,
	}

	return s.repository.Update(ctx, user)
}

func (s *userService) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
