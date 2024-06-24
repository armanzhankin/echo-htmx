package database

import (
	"context"
	"errors"

	"github.com/armanzhankin/echo-htmx/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepoI interface {
	GetUsers(ctx context.Context) ([]models.User, error)
}

type UserPgRepo struct {
	pgPool *pgxpool.Pool
}

func NewUserRepo(pgPool *pgxpool.Pool) (*UserPgRepo, error) {
	if pgPool == nil {
		return nil, errors.New("pg pool is nil")
	}

	return &UserPgRepo{pgPool: pgPool}, nil
}

func (r UserPgRepo) GetUsers(ctx context.Context) ([]models.User, error) {
	return nil, nil
}
