package list_handler

import (
	"context"
	"time"

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
}

func (r Request) Validate() error {
	return nil
}

type CurrencyPair struct {
	CurrencyFrom string    `json:"currencyFrom"`
	CurrencyTo   string    `json:"currencyTo"`
	Well         float64   `json:"well"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Response = []CurrencyPair

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	var res Response
	currencyPairs, err := h.businessLogic.List(ctx)
	if err != nil {
		return res, err
	}

	res = make([]CurrencyPair, 0, len(currencyPairs))

	for _, item := range currencyPairs {
		res = append(res, CurrencyPair{
			CurrencyFrom: item.CurrencyFrom,
			CurrencyTo:   item.CurrencyTo,
			Well:         item.Well,
			UpdatedAt:    item.UpdatedAt,
		})
	}

	return res, nil
}
