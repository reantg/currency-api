package currency

import (
	"context"
	"github.com/reantg/currency-api/internal/currency/types"
	"github.com/reantg/currency-api/internal/currency/types/convertor"
)

type RatesApiClient interface {
	Latest(ctx context.Context, currencyFrom string, currencyTo string) (float64, error)
}

type Repository interface {
	Currency
}

type Currency interface {
	List(ctx context.Context) ([]*types.CurrencyPairRepo, error)
	Add(ctx context.Context, from string, to string, well float64) error
	Get(ctx context.Context, from string, to string) (*types.CurrencyPairRepo, error)
	Update(ctx context.Context, from string, to string, well float64) error
}

type Engine struct {
	repo           Repository
	ratesApiClient RatesApiClient
}

func New(repo Repository, ratesApiClient RatesApiClient) *Engine {
	return &Engine{
		repo:           repo,
		ratesApiClient: ratesApiClient,
	}
}

func (e *Engine) List(ctx context.Context) ([]*types.CurrencyPair, error) {
	pairs, err := e.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	return convertor.ToCurrencyPairList(pairs), nil
}

func (e *Engine) Add(ctx context.Context, currencyFrom string, currencyTo string) error {
	rate, err := e.ratesApiClient.Latest(ctx, currencyFrom, currencyTo)
	if err != nil {
		return err
	}

	err = e.repo.Add(ctx, currencyFrom, currencyTo, rate)

	if err != nil {
		return err
	}

	return nil
}

func (e *Engine) Convert(ctx context.Context, currencyFrom string, currencyTo string, value float64) (*types.ConvertResult, error) {
	pair, err := e.repo.Get(ctx, currencyFrom, currencyTo)
	if err != nil {
		return nil, err
	}

	var convertedValue float64
	if pair.CurrencyFrom == currencyFrom {
		convertedValue = pair.Well * value
	} else {
		convertedValue = value / pair.Well
	}

	resp := types.ConvertResult{
		Well:  pair.Well,
		Value: convertedValue,
	}

	return &resp, nil
}

func (e *Engine) Update(ctx context.Context, pair types.CurrencyPair) error {
	rate, err := e.ratesApiClient.Latest(ctx, pair.CurrencyFrom, pair.CurrencyTo)
	if err != nil {
		return err
	}

	err = e.repo.Update(ctx, pair.CurrencyFrom, pair.CurrencyTo, rate)

	if err != nil {
		return err
	}

	return nil
}

func (e *Engine) UpdateAllRates(ctx context.Context) error {
	pairs, err := e.List(ctx)
	if err != nil {
		return err
	}

	for _, pair := range pairs {
		err := e.Update(ctx, *pair)
		if err != nil {
			return err
		}
	}

	return nil
}
