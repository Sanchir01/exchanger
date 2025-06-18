package models

import (
	"github.com/google/uuid"
	"time"
)

type ExchangeRate struct {
	ID           uuid.UUID `json:"id"`
	FromCurrency string    `json:"from_currency"`
	ToCurrency   string    `json:"to_currency"`
	Rate         float64   `json:"rate"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ExchangeRateHistory struct {
	ID           uuid.UUID `json:"id"`
	FromCurrency string    `json:"from_currency"`
	ToCurrency   string    `json:"to_currency"`
	Rate         float64   `json:"rate"`
	RecordedAt   time.Time `json:"recorded_at"`
}
