package service

import (
	"context"
	"errors"

	"github.com/armanzhankin/echo-htmx/internal/database"
	"github.com/armanzhankin/echo-htmx/internal/models"
)

type UserServiceI interface {
	GetUsers(ctx context.Context) ([]models.User, error)
}

type UserService struct {
	userRepo *database.Repository
}

func NewUserService(ctx context.Context, db *database.Repository) (*UserService, error) {
	if db == nil {
		return nil, errors.New("user db is nil")
	}

	return &UserService{
		userRepo: db,
	}, nil
}

func (u UserService) GetUsers(ctx context.Context) ([]models.User, error) {
	return nil, nil
}
