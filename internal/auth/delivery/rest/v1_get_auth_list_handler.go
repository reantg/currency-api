package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/reantg/currency-api/internal/auth/types"
)

// GET api/auth/list

func (h *Handler) v1GetAuthListHandler(c *fiber.Ctx) error {
	return h.user.CreateUser(
		c.Context(),
		types.Login{},
	)
}
