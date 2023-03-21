-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS currency_pairs (
  id BIGSERIAL PRIMARY KEY,
  currency_from varchar(3) NOT NULL,
  currency_to varchar(3) NOT NULL,
  well float,
  updated_at timestamptz default now() 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS currency_pairs;
-- +goose StatementEnd
