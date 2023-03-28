package convertor

import "github.com/reantg/currency-api/internal/currency/types"

func ToCurrencyPair(item *types.CurrencyPairRepo) *types.CurrencyPair {
	if item == nil {
		return nil
	}

	return &types.CurrencyPair{
		CurrencyFrom: item.CurrencyFrom,
		CurrencyTo:   item.CurrencyTo,
		Well:         item.Well,
		UpdatedAt:    item.UpdatedAt,
	}
}

func ToCurrencyPairList(items []*types.CurrencyPairRepo) []*types.CurrencyPair {
	resp := make([]*types.CurrencyPair, 0, len(items))
	for _, item := range items {
		resp = append(resp, ToCurrencyPair(item))
	}

	return resp
}
