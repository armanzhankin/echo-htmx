package service

import (
	"context"

	"github.com/armanzhankin/echo-htmx/internal/database"
)

type Service struct {
	UserService UserServiceI
}

func NewService(ctx context.Context, db *database.Repository) (*Service, error) {
	userService, err := NewUserService(ctx, db)
	if err != nil {
		return nil, err
	}

	return &Service{
		UserService: userService,
	}, nil
}
