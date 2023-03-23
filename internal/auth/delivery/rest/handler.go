package rest

import (
	"context"

	"github.com/reantg/currency-api/internal/auth/types"
)

type User interface {
	CreateUser(ctx context.Context, login types.Login) error
}

type Handler struct {
	// init http framework

	user User
}

func New() *Handler {
	return &Handler{}
}
