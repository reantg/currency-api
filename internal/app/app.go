package app

import (
	"context"
	"errors"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/reantg/currency-api/internal/clients/openexchangerates"
	"github.com/reantg/currency-api/internal/config"
	"github.com/reantg/currency-api/internal/domain"
	add "github.com/reantg/currency-api/internal/handlers/add-handler"
	convert "github.com/reantg/currency-api/internal/handlers/convert-handler"
	list "github.com/reantg/currency-api/internal/handlers/list-handler"
	postgresRepository "github.com/reantg/currency-api/internal/repositories/postgres-repository"
	updaterService "github.com/reantg/currency-api/internal/services/updater-service"
	serverWrapper "github.com/reantg/currency-api/internal/wrappers/server-wrapper"
)

type App struct {
	router  *fiber.App
	dbPool  *pgxpool.Pool
	updater updaterService.Updater
}

func New(ctx context.Context) (*App, error) {
	a := &App{}
	if err := a.init(ctx); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) init(ctx context.Context) error {
	err := config.Init()
	if err != nil {
		return err
	}
	dbPool, err := pgxpool.Connect(ctx, config.ConfigData.DbUri)
	if err != nil {
		return err
	}

	currencyPairRepo := postgresRepository.NewCurrencyRepository(dbPool)

	openexchangeratesClient := openexchangerates.New(config.ConfigData.OpenexchangeratesUrl, config.ConfigData.OpenexchangeratesApiKey)

	businessLogic := domain.New(currencyPairRepo, openexchangeratesClient)

	updater := updaterService.New(time.Hour*1, businessLogic)
	router := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			err = ctx.Status(code).SendString(err.Error())
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			return nil
		}})

	convertHandler := convert.New(businessLogic)
	listHandler := list.New(businessLogic)
	addHandler := add.New(businessLogic)

	router.Post("/api/currency", serverWrapper.New(addHandler.Handle).Wrap)
	router.Put("/api/currency", serverWrapper.New(convertHandler.Handle).Wrap)
	router.Get("/api/currency", serverWrapper.New(listHandler.Handle).Wrap)

	a.dbPool = dbPool
	a.router = router
	a.updater = updater

	return nil
}

func (a *App) Run(ctx context.Context) error {
	defer a.dbPool.Close()
	a.updater.Run(ctx)

	err := a.router.Listen(config.ConfigData.Port)
	if err != nil {
		return err
	}

	return nil
}
