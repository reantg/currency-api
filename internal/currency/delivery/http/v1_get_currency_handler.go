package http

import (
	"context"
	"github.com/reantg/currency-api/internal/currency/types"
)

type v1GetCurrencyRequest struct {
}

func (r v1GetCurrencyRequest) Validate() error {
	return nil
}

type v1GetCurrencyResponse = []types.CurrencyPair

func (h *Handler) v1GetCurrencyHandler(ctx context.Context, req v1GetCurrencyRequest) (v1GetCurrencyResponse, error) {
	var res v1GetCurrencyResponse
	currencyPairs, err := h.engine.List(ctx)
	if err != nil {
		return res, err
	}

	res = make([]types.CurrencyPair, 0, len(currencyPairs))

	for _, item := range currencyPairs {
		res = append(res, types.CurrencyPair{
			CurrencyFrom: item.CurrencyFrom,
			CurrencyTo:   item.CurrencyTo,
			Well:         item.Well,
			UpdatedAt:    item.UpdatedAt,
		})
	}

	return res, nil
}
