package auth

import (
	"context"

	"github.com/reantg/currency-api/internal/auth/types"
)

type Repository interface {
	User
}

type User interface {
	CreateUser(ctx context.Context, login types.Login) error
}

type Engine struct {
	repo Repository
}

func (e Engine) CreateUser(ctx context.Context, login types.Login) error {
	return e.repo.CreateUser(ctx, login)
}
