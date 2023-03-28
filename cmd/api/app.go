package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/reantg/currency-api/internal/config"
	"github.com/reantg/currency-api/internal/currency"
	currencyPairHttp "github.com/reantg/currency-api/internal/currency/delivery/http"
	currencyPairPostgres "github.com/reantg/currency-api/internal/currency/repo/postgres"
	"github.com/reantg/currency-api/pkg/openexchange"
	"github.com/reantg/currency-api/pkg/worker"
	"log"
	"os"
	"os/signal"
	"time"
)

func runApp() error {
	ctx := context.Background()

	configData, err := config.Init()
	if err != nil {
		log.Fatal("config init error", err)
	}

	dbPool, err := pgxpool.Connect(ctx, configData.DbUri)
	if err != nil {
		return err
	}

	openexchangeratesClient := openexchange.New(configData.OpenexchangeratesUrl, configData.OpenexchangeratesApiKey)

	currencyPairRepo := currencyPairPostgres.New(dbPool)
	currencyEngine := currency.New(currencyPairRepo, openexchangeratesClient)
	currencyHandler := currencyPairHttp.New(currencyEngine)

	wrk := worker.New()
	err = wrk.Periodic(ctx, "currency-update", currencyEngine.UpdateAllRates, time.Second*5)
	if err != nil {
		log.Println("error add job", err)
	}

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
	apiRouter := router.Group("/api")
	currencyHandler.Register(apiRouter)

	serverShutdown := make(chan struct{})
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		_ = <-c
		log.Println("Gracefully shutting down...")

		timeoutFunc := time.AfterFunc(time.Second*time.Duration(configData.ForceShutdownTimeout), func() {
			log.Printf("timeout %d ms has been elapsed, force exit", configData.ForceShutdownTimeout)
			os.Exit(0)
		})
		defer timeoutFunc.Stop()

		_ = router.Shutdown()
		dbPool.Close()

		serverShutdown <- struct{}{}
	}()

	err = router.Listen(configData.Port)
	if err != nil {
		return err
	}

	<-serverShutdown
	log.Println("Running cleanup tasks...")

	return nil
}
