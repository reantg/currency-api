package http

import (
	"context"
	"github.com/pkg/errors"
)

type v1PutCurrencyRequest struct {
	CurrencyFrom string  `json:"currencyFrom"`
	CurrencyTo   string  `json:"currencyTo"`
	Value        float64 `json:"value"`
}

var (
	ValueEmpty = errors.New("value cannot be empty")
)

func (r v1PutCurrencyRequest) Validate() error {
	if r.CurrencyFrom == "" {
		return CurrencyFromEmpty
	}

	if r.CurrencyTo == "" {
		return CurrencyToEmpty
	}

	if r.Value == 0 {
		return ValueEmpty
	}

	return nil
}

type v1PutCurrencyResponse struct {
	CurrencyFrom string  `json:"currencyFrom"`
	CurrencyTo   string  `json:"currencyTo"`
	Well         float64 `json:"well"`
	Value        float64 `json:"value"`
}

func (h *Handler) v1PutCurrencyHandler(ctx context.Context, req v1PutCurrencyRequest) (v1PutCurrencyResponse, error) {
	var res v1PutCurrencyResponse
	converted, err := h.engine.Convert(ctx, req.CurrencyFrom, req.CurrencyTo, req.Value)
	if err != nil {
		return res, err
	}
	res = v1PutCurrencyResponse{
		CurrencyFrom: req.CurrencyFrom,
		CurrencyTo:   req.CurrencyTo,
		Well:         converted.Well,
		Value:        converted.Value,
	}

	return res, nil
}
