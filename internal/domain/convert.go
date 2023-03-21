package domain

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type ConvertResult struct {
	Well  float64
	Value float64
}

var (
	CurrencyPairNotFound = errors.New("currency pair not found")
)

func (m *Model) Convert(ctx context.Context, currencyFrom string, currencyTo string, value float64) (*ConvertResult, error) {
	pair, err := m.currencyPairRepo.Get(ctx, currencyFrom, currencyTo)
	if err != nil {
		if errors.Is(err, CurrencyPairNotFound) {
			return nil, fiber.NewError(404, err.Error())
		}
		return nil, err
	}

	var convertedValue float64
	if pair.CurrencyFrom == currencyFrom {
		convertedValue = pair.Well * value
	} else {
		convertedValue = value / pair.Well
	}

	resp := ConvertResult{
		Well:  pair.Well,
		Value: convertedValue,
	}

	return &resp, nil
}
