package updater_service

import (
	"context"
	"log"
	"time"

	"github.com/reantg/currency-api/internal/domain"
)

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
				u.updatePairs(ctx)
			}
		}
	}()
}

func (u *updater) updatePairs(ctx context.Context) {
	pairs, err := u.businessLogic.List(ctx)
	if err != nil {
		log.Println("cannot get pairs list err", err)
		return
	}

	for _, pair := range pairs {
		err := u.businessLogic.Update(ctx, pair.CurrencyFrom, pair.CurrencyTo)
		if err != nil {
			log.Println("cannot update pair err", err)
		}
	}
}
