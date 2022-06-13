package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/DmitriyZhevnov/rest-api/internal/apperror"
	"github.com/DmitriyZhevnov/rest-api/internal/model"
	"github.com/DmitriyZhevnov/rest-api/pkg/client/postgresql"
	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
)

type authorPostgres struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewAuthorPostgres(client postgresql.Client, logger *logging.Logger) *authorPostgres {
	return &authorPostgres{
		client: client,
		logger: logger,
	}
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (r *authorPostgres) FindAll(ctx context.Context) (u []model.Author, err error) {
	q := `
		SELECT id, name, age FROM public.author;
	`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, apperror.NewInternalServerError(fmt.Sprintf("failed to find all authors due to error: %v", err), "34547654")
	}

	authors := make([]model.Author, 0)

	for rows.Next() {
		var ath model.Author

		err = rows.Scan(&ath.ID, &ath.Name, &ath.Age)
		if err != nil {
			return nil, apperror.NewInternalServerError(fmt.Sprintf("failed to scan author (id = %s) due to error: %v", ath.ID, err), "94847574")
		}

		authors = append(authors, ath)
	}

	if err = rows.Err(); err != nil {
		return nil, apperror.NewInternalServerError(fmt.Sprintf("failed to find all authors due to error: %v", err), "345464534")
	}

	return authors, nil
}

func (r *authorPostgres) FindOne(ctx context.Context, id string) (model.Author, error) {
	q := `
		SELECT id, name, age FROM public.author WHERE id = $1
	`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	var ath model.Author
	err := r.client.QueryRow(ctx, q, id).Scan(&ath.ID, &ath.Name, &ath.Age)
	if err != nil {
		return model.Author{}, apperror.NewErrNotFound(fmt.Sprintf("user with id = %s not found", id), "34346234")
	}

	return ath, nil
}
