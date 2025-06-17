-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE exchange_rates (
                                id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                                from_currency TEXT NOT NULL,
                                to_currency TEXT NOT NULL,
                                rate NUMERIC(12,6) NOT NULL,
                                created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                                updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
                                UNIQUE (from_currency, to_currency)
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_exchange_rates_updated_at
    BEFORE UPDATE ON exchange_rates
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();
CREATE TABLE exchange_rates_history (
                                        id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
                                        from_currency VARCHAR(3) NOT NULL,
                                        to_currency VARCHAR(3) NOT NULL,
                                        rate NUMERIC(12,6) NOT NULL,
                                        recorded_at TIMESTAMP NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_exchange_rates_history ON exchange_rates_history (from_currency, to_currency, recorded_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
