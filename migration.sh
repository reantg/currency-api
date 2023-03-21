goose -dir ./migrations postgres "postgres://currency:password@localhost:5433/currency?sslmode=disable" status

goose -dir ./migrations postgres "postgres://currency:password@localhost:5433/currency?sslmode=disable" up
