package add_handler

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
	CurrencyFrom string `json:"currencyFrom"`
	CurrencyTo   string `json:"currencyTo"`
}

var (
	CurrencyFromEmpty = errors.New("currencyFrom cannot be empty")
	CurrencyToEmpty   = errors.New("currencyTo cannot be empty")
)

func (r Request) Validate() error {
	if r.CurrencyFrom == "" {
		return CurrencyFromEmpty
	}

	if r.CurrencyTo == "" {
		return CurrencyToEmpty
	}

	return nil
}

type Response struct {
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	var res Response
	err := h.businessLogic.Add(ctx, req.CurrencyFrom, req.CurrencyTo)
	if err != nil {
		return res, err
	}
	return res, nil
}
