package postgres_repository

import (
	"context"
	"errors"
	"time"

	"github.com/georgysavva/scany/pgxscan"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/reantg/currency-api/internal/domain"
)

type CurrencyPairRepository struct {
	pool      *pgxpool.Pool
	pgBuilder sq.StatementBuilderType
}

func NewCurrencyRepository(pool *pgxpool.Pool) *CurrencyPairRepository {
	return &CurrencyPairRepository{
		pool:      pool,
		pgBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

const (
	currencyPairsTable = "currency_pairs"
)

func (r *CurrencyPairRepository) List(ctx context.Context) ([]*domain.CurrencyPair, error) {
	query := r.pgBuilder.Select("id", "currency_from", "currency_to", "well", "updated_at").
		From(currencyPairsTable)
	rawSql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var items []*CurrencyPair
	err = pgxscan.Select(ctx, r.pool, &items, rawSql, args...)
	if err != nil {
		return nil, err
	}

	return ToCurrencyPairList(items), nil
}

func (r *CurrencyPairRepository) Add(ctx context.Context, currencyFrom string, currencyTo string, well float64) error {
	query := r.pgBuilder.Insert(currencyPairsTable).
		Columns("currency_from", "currency_to", "well").
		Values(currencyFrom, currencyTo, well)
	rawSql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	rows, err := r.pool.Query(ctx, rawSql, args...)
	if err != nil {
		return err
	}
	rows.Close()
	return nil
}

func (r *CurrencyPairRepository) Update(ctx context.Context, currencyFrom string, currencyTo string, well float64) error {
	query := r.pgBuilder.Update(currencyPairsTable).
		Set("well", well).
		Set("updated_at", time.Now()).
		Where(sq.And{
			sq.Eq{"currency_from": currencyFrom},
			sq.Eq{"currency_to": currencyTo},
		})
	rawSql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	rows, err := r.pool.Query(ctx, rawSql, args...)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func (r *CurrencyPairRepository) Get(ctx context.Context, currencyFrom string, currencyTo string) (*domain.CurrencyPair, error) {
	query := r.pgBuilder.Select("id", "currency_from", "currency_to", "well", "updated_at").
		From(currencyPairsTable).
		Where(sq.Or{
			sq.And{
				sq.Eq{"currency_from": currencyFrom},
				sq.Eq{"currency_to": currencyTo},
			},
			sq.And{
				sq.Eq{"currency_to": currencyFrom},
				sq.Eq{"currency_from": currencyTo},
			},
		})
	rawSql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var item CurrencyPair
	err = pgxscan.Get(ctx, r.pool, &item, rawSql, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.CurrencyPairNotFound
		}
		return nil, err
	}

	return ToCurrencyPair(&item), nil
}
