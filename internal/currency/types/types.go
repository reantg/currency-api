package types

import "time"

type CurrencyPair struct {
	CurrencyFrom string    `json:"currencyFrom"`
	CurrencyTo   string    `json:"currencyTo"`
	Well         float64   `json:"well"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ConvertResult struct {
	Well  float64
	Value float64
}

type CurrencyPairRepo struct {
	ID           int64     `db:"id"`
	CurrencyFrom string    `db:"currency_from"`
	CurrencyTo   string    `db:"currency_to"`
	Well         float64   `db:"well"`
	UpdatedAt    time.Time `db:"updated_at"`
}
