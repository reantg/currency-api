package main

import (
	"context"
	"log"

	"github.com/reantg/currency-api/internal/app"
)

func main() {
	ctx := context.Background()
	app, err := app.New(ctx)
	if err != nil {
		log.Fatal("config err", err)
	}

	if err := app.Run(ctx); err != nil {
		log.Fatal("app run err", err)
	}
}
