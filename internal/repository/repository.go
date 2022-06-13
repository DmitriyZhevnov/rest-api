package repository

import (
	"context"

	"github.com/DmitriyZhevnov/rest-api/internal/model"
	"github.com/DmitriyZhevnov/rest-api/pkg/client/postgresql"
	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	User
	Author
}

type User interface {
	Create(ctx context.Context, user model.User) (string, error)
	FindOne(ctx context.Context, id string) (model.User, error)
	FindAll(ctx context.Context) ([]model.User, error)
	Update(ctx context.Context, user model.User) error
	Delete(ctx context.Context, id string) error
}

type Author interface {
	FindAll(ctx context.Context) (u []model.Author, err error)
	FindOne(ctx context.Context, id string) (model.Author, error)
	Create(ctx context.Context, aurhor model.Author) (string, error)
	Update(ctx context.Context, user model.Author) error
	Delete(ctx context.Context, id string) error
}

func NewRepository(db *mongo.Database, mongoCollection string, client postgresql.Client, logger *logging.Logger) *Repository {
	return &Repository{
		User:   NewUserMongo(db, mongoCollection, logger),
		Author: NewAuthorPostgres(client, logger),
	}
}
