package postgres_repository

import "github.com/reantg/currency-api/internal/domain"

func ToCurrencyPair(item *CurrencyPair) *domain.CurrencyPair {
	if item == nil {
		return nil
	}

	return &domain.CurrencyPair{
		CurrencyFrom: item.CurrencyFrom,
		CurrencyTo:   item.CurrencyTo,
		Well:         item.Well,
		UpdatedAt:    item.UpdatedAt,
	}
}

func ToCurrencyPairList(items []*CurrencyPair) []*domain.CurrencyPair {
	resp := make([]*domain.CurrencyPair, 0, len(items))
	for _, item := range items {
		resp = append(resp, ToCurrencyPair(item))
	}

	return resp
}
