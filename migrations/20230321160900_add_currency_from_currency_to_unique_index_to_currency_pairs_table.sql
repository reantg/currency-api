-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX currency_pairs_currency_from_currency_to_uindex ON currency_pairs (currency_from, currency_to);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX currency_pairs_currency_from_currency_to_uindex;
-- +goose StatementEnd
