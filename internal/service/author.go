package service

import (
	"context"

	"github.com/DmitriyZhevnov/rest-api/internal/model"
	"github.com/DmitriyZhevnov/rest-api/internal/repository"
	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
)

type authorService struct {
	repository repository.Author
	logger     *logging.Logger
}

func NewAuthorSevice(repos repository.Author, logger *logging.Logger) *authorService {
	return &authorService{
		repository: repos,
		logger:     logger,
	}
}

func (s *authorService) FindAll(ctx context.Context) ([]model.Author, error) {
	return s.repository.FindAll(ctx)
}
