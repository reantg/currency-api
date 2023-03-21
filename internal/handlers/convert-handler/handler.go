package convert_handler

import (
	"context"
	"errors"

	"github.com/reantg/currency-api/internal/domain"
)

type Handler struct {
	businessLogic *domain.Model
}

func New(businessLogic *domain.Model) *Handler {
	return &Handler{
		businessLogic: businessLogic,
	}
}

type Request struct {
	CurrencyFrom string  `json:"currencyFrom"`
	CurrencyTo   string  `json:"currencyTo"`
	Value        float64 `json:"value"`
}

var (
	CurrencyFromEmpty = errors.New("currencyFrom cannot be empty")
	CurrencyToEmpty   = errors.New("currencyTo cannot be empty")
	ValueEmpty        = errors.New("value cannot be empty")
)

func (r Request) Validate() error {
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

type Response struct {
	CurrencyFrom string  `json:"currencyFrom"`
	CurrencyTo   string  `json:"currencyTo"`
	Well         float64 `json:"well"`
	Value        float64 `json:"value"`
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	var res Response
	converted, err := h.businessLogic.Convert(ctx, req.CurrencyFrom, req.CurrencyTo, req.Value)
	if err != nil {
		return res, err
	}
	res = Response{
		CurrencyFrom: req.CurrencyFrom,
		CurrencyTo:   req.CurrencyTo,
		Well:         converted.Well,
		Value:        converted.Value,
	}

	return res, nil
}
