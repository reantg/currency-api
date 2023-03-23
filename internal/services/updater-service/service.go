package updater_service

import (
	"context"
	"log"
	"time"

	"github.com/reantg/currency-api/internal/domain"
)

//  Сделать по следующей схеме
//wg := worker.New(Config{})
//wg.Async(ctx, "rates", func(ctx context.Context) error {})
//wg.Async(ctx, "auth-checker", func(ctx context.Context) error {})

type Updater interface {
	Run(ctx context.Context)
}

type updater struct {
	ticker        *time.Ticker
	businessLogic *domain.Model
}

func New(interval time.Duration, businessLogic *domain.Model) Updater {
	return &updater{
		ticker:        time.NewTicker(interval),
		businessLogic: businessLogic,
	}
}

func (u *updater) Run(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-u.ticker.C:
				if err := u.updatePairs(ctx); err != nil {
					log.Println(err)
				}
			}
		}
	}()
}

func (u *updater) updatePairs(ctx context.Context) error {
	pairs, err := u.businessLogic.List(ctx)
	if err != nil {
		return err
	}

	for _, pair := range pairs {
		err := u.businessLogic.Update(ctx, pair.CurrencyFrom, pair.CurrencyTo)
		if err != nil {
			return err
		}
	}
	return nil
}
