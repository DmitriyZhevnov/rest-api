package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/DmitriyZhevnov/rest-api/internal/apperror"
	"github.com/DmitriyZhevnov/rest-api/internal/model"
	"github.com/DmitriyZhevnov/rest-api/pkg/client/postgresql"
	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
	"github.com/jackc/pgconn"
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

func (r *authorPostgres) Create(ctx context.Context, author model.Author) (string, error) {
	q := `
        INSERT INTO author
            (name, age)
        VALUES
               ($1, $2)
        RETURNING id
    `
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
	if err := r.client.QueryRow(ctx, q, author.Name, author.Age).Scan(&author.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := apperror.NewInternalServerError(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()), "3435456765")
			r.logger.Error(newErr)
			return "", newErr
		}
		return "", apperror.NewInternalServerError(fmt.Sprintf("error creating author: %v", err), "43656743346")
	}

	return author.ID, nil
}
