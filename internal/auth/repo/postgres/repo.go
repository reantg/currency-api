package postgres

import (
	"context"

	"github.com/reantg/currency-api/internal/auth/types"
)

type Repo struct {
	// db
}

func (r Repo) CreateUser(ctx context.Context, login types.Login) error {
	query := `
insert into users(id, .../)
`
	_ = query

	// execute query
	return nil
}
