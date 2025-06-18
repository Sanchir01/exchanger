package exchanger

import (
	"github.com/google/uuid"
	"time"
)

type GetExchangeRate struct {
	FromCurrency string `json:"from_currency" validate:"required"`
	ToCurrency   string `json:"to_currency" validate:"required"`
}

type ExchangeDB struct {
	ID           uuid.UUID `db:"id"`
	FromCurrency string    `db:"from_currency"`
	ToCurrency   string    `db:"to_currency"`
	Rate         float64   `db:"rate"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
type ExchangeRateHistoryDB struct {
	ID           uuid.UUID `db:"id"`
	FromCurrency string    `db:"from_currency"`
	ToCurrency   string    `db:"to_currency"`
	Rate         float64   `db:"rate"`
	RecordedAt   time.Time `db:"recorded_at"`
}
