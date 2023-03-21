package postgres_repository

import "time"

type CurrencyPair struct {
	ID           int64     `db:"id"`
	CurrencyFrom string    `db:"currency_from"`
	CurrencyTo   string    `db:"currency_to"`
	Well         float64   `db:"well"`
	UpdatedAt    time.Time `db:"updated_at"`
}
