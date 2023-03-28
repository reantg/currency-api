package http

import (
	"context"
	"github.com/pkg/errors"
)

type v1PostCurrencyRequest struct {
	CurrencyFrom string `json:"currencyFrom"`
	CurrencyTo   string `json:"currencyTo"`
}

var (
	CurrencyFromEmpty = errors.New("currencyFrom cannot be empty")
	CurrencyToEmpty   = errors.New("currencyTo cannot be empty")
)

func (r v1PostCurrencyRequest) Validate() error {
	if r.CurrencyFrom == "" {
		return CurrencyFromEmpty
	}

	if r.CurrencyTo == "" {
		return CurrencyToEmpty
	}

	return nil
}

type v1PostCurrencyResponse struct {
}

func (h *Handler) v1PostCurrencyHandler(ctx context.Context, req v1PostCurrencyRequest) (v1PostCurrencyResponse, error) {
	var res v1PostCurrencyResponse
	err := h.engine.Add(ctx, req.CurrencyFrom, req.CurrencyTo)
	if err != nil {
		return res, err
	}
	return res, nil
}
