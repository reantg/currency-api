package domain

import (
	"context"
	"time"
)

type CurrencyPair struct {
	CurrencyFrom string
	CurrencyTo   string
	Well         float64
	UpdatedAt    time.Time
}

type RatesApiClient interface {
	Latest(ctx context.Context, currencyFrom string, currencyTo string) (float64, error)
}

type CurrencyPairRepo interface {
	List(ctx context.Context) ([]*CurrencyPair, error)
	Get(ctx context.Context, from, to string) (*CurrencyPair, error)
	Add(ctx context.Context, pair CurrencyPair) error
	Update(ctx context.Context, pair CurrencyPair) error
}

type Model struct {
	currencyPairRepo CurrencyPairRepo
	ratesApiClient   RatesApiClient
}

func New(currencyPairRepo CurrencyPairRepo, ratesApiClient RatesApiClient) *Model {
	return &Model{
		currencyPairRepo: currencyPairRepo,
		ratesApiClient:   ratesApiClient,
	}
}
