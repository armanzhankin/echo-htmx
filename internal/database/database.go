package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func createPGPool(ctx context.Context, dbUrl string) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}
	if err = dbPool.Ping(ctx); err != nil {
		return nil, err
	}
	return dbPool, nil
}

type Repository struct {
	User UserRepoI
}

func New(ctx context.Context, dbUrl string) (*Repository, error) {
	pgPool, err := createPGPool(ctx, dbUrl)
	if err != nil {
		return nil, err
	}

	userRepo, err := NewUserRepo(pgPool)

	return &Repository{
		User: userRepo,
	}, nil
}
