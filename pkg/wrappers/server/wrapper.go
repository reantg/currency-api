package server

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type Validator interface {
	Validate() error
}

func Wrap[Req Validator, Res any](fn func(ctx context.Context, req Req) (Res, error)) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		var request Req
		if c.Route().Method == fiber.MethodGet {
			if err := c.QueryParser(&request); err != nil {
				return fiber.NewError(400, errors.Wrap(err, "error parse request").Error())
			}
		} else {
			if err := c.BodyParser(&request); err != nil {
				return fiber.NewError(400, errors.Wrap(err, "error parse request").Error())
			}
		}

		err := request.Validate()
		if err != nil {
			return fiber.NewError(400, errors.Wrap(err, "validation error").Error())
		}

		response, err := fn(ctx, request)
		if err != nil {
			return err
		}

		rawJson, err := json.Marshal(response)
		if err != nil {
			return errors.Wrap(err, "marshal response error")
		}

		_, err = c.Response().BodyWriter().Write(rawJson)
		if err != nil {
			return errors.Wrap(err, "body write error")
		}
		c.Response().Header.Add("content-type", "application/json")
		return nil
	}
}
