package domain

import "context"

func (m *Model) Add(ctx context.Context, currencyFrom string, currencyTo string) error {
	rate, err := m.ratesApiClient.Latest(ctx, currencyFrom, currencyTo)
	if err != nil {
		return err
	}

	err = m.currencyPairRepo.Add(ctx, currencyFrom, currencyTo, rate)

	if err != nil {
		return err
	}

	return nil
}
