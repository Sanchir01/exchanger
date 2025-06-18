package exchanger

import (
	"context"
	"log/slog"
)

type Services struct {
	repository Repositories
	log        *slog.Logger
}

//go:generate go run github.com/vektra/mockery/v3@v3.4.0 --name=Repositories
type Repositories interface {
	GetCurrencyByRate(ctx context.Context, fromCurrency, toCurrency string) (*ExchangeDB, error)
	GetAllCurrency(ctx context.Context) ([]ExchangeDB, error)
}

func NewService(repo Repositories, l *slog.Logger) *Services {
	return &Services{repository: repo, log: l}
}

func (s *Services) AllCurrency(ctx context.Context) ([]ExchangeDB, error) {
	const op = "Exchanger.Handler.GetExchangeRate"
	log := s.log.With(
		slog.String("op", op),
	)
	data, err := s.repository.GetAllCurrency(ctx)
	if err != nil {
		return nil, err
	}
	log.Info("success get all currency data")
	return data, nil
}

func (s *Services) GetExchangeRate(ctx context.Context, fromCurrency, toCurrency string) (*ExchangeDB, error) {
	const op = "Exchanger.Handler.GetExchangeRate"
	log := s.log.With(
		slog.String("op", op),
	)
	data, err := s.repository.GetCurrencyByRate(ctx, fromCurrency, toCurrency)
	if err != nil {
		return nil, err
	}
	log.Info("success service get exchange rate data")
	return data, nil
}
