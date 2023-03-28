package http

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/reantg/currency-api/internal/currency/types"
	serverWrapper "github.com/reantg/currency-api/pkg/wrappers/server"
)

type Engine interface {
	Currency
}

type Currency interface {
	List(ctx context.Context) ([]*types.CurrencyPair, error)
	Add(ctx context.Context, currencyFrom string, currencyTo string) error
	Convert(ctx context.Context, currencyFrom string, currencyTo string, value float64) (*types.ConvertResult, error)
}

type Handler struct {
	engine Engine
}

func New(engine Engine) *Handler {
	handler := &Handler{
		engine: engine,
	}

	return handler
}

func (h *Handler) Register(router fiber.Router) {
	router.Post("/currency", serverWrapper.Wrap(h.v1PostCurrencyHandler))
	router.Put("/currency", serverWrapper.Wrap(h.v1PutCurrencyHandler))
	router.Get("/currency", serverWrapper.Wrap(h.v1GetCurrencyHandler))
}
