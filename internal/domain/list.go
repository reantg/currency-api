package domain

import "context"

func (m *Model) List(ctx context.Context) ([]*CurrencyPair, error) {
	pairs, err := m.currencyPairRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	return pairs, nil
}
