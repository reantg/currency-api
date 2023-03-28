package postgres

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/reantg/currency-api/internal/currency/types"
	"time"
)

var (
	CurrencyPairNotFound = errors.New("currency pair not found")
)

type Repository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) List(ctx context.Context) ([]*types.CurrencyPairRepo, error) {
	query := `
SELECT id, currency_from, currency_to, well, updated_at FROM currency_pairs
`

	var items []*types.CurrencyPairRepo
	err := pgxscan.Select(ctx, r.pool, &items, query)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (r *Repository) Add(ctx context.Context, from string, to string, well float64) error {
	query := `
INSERT INTO currency_pairs(currency_from, currency_to, well) VALUES ($1, $2, $3)
`
	rows, err := r.pool.Query(ctx, query, from, to, well)
	if err != nil {
		return err
	}
	rows.Close()
	return nil
}

func (r *Repository) Update(ctx context.Context, from string, to string, well float64) error {
	query := `
UPDATE currency_pairs SET well=$1, updated_at=$2 WHERE currency_from=$3 and currency_to=$4
`

	rows, err := r.pool.Query(ctx, query, well, time.Now(), from, to)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func (r *Repository) Get(ctx context.Context, from string, to string) (*types.CurrencyPairRepo, error) {
	query := `
SELECT id, currency_from, currency_to, well, updated_at FROM currency_pairs WHERE (currency_from=$1 AND currency_to=$2) OR (currency_from=$2 AND currency_to=$1)
`

	var item types.CurrencyPairRepo
	err := pgxscan.Get(ctx, r.pool, &item, query, from, to)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, CurrencyPairNotFound
		}
		return nil, err
	}

	return &item, nil
}
