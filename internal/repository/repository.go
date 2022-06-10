package repository

import (
	"context"

	"github.com/DmitriyZhevnov/rest-api/internal/model"
	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
	"go.mongodb.org/mongo-driver/mongo"
)

type User interface {
	Create(ctx context.Context, user model.User) (string, error)
	FindOne(ctx context.Context, id string) (model.User, error)
	FindAll(ctx context.Context) ([]model.User, error)
	Update(ctx context.Context, user model.User) error
	Delete(ctx context.Context, id string) error
}

type Repository struct {
	User
}

func NewRepository(db *mongo.Database, mongoCollection string, logger *logging.Logger) *Repository {
	return &Repository{
		User: NewUserMongo(db, mongoCollection, logger),
	}
}
