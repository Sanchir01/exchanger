-- +goose Up
-- +goose StatementBegin
TRUNCATE exchange_rates CASCADE;
TRUNCATE exchange_rates_history CASCADE;

INSERT INTO exchange_rates (from_currency, to_currency, rate)
VALUES
    ('USD', 'RUB', 75.000000),
    ('RUB', 'USD', 0.013333),
    ('USD', 'EUR', 0.900000),
    ('EUR', 'USD', 1.111111),
    ('EUR', 'RUB', 83.000000),
    ('RUB', 'EUR', 0.012048);

INSERT INTO exchange_rates_history (from_currency, to_currency, rate)
SELECT from_currency, to_currency, rate FROM exchange_rates;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
