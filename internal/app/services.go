package app

import (
	"github.com/Sanchir01/exchanger/internal/feature/exchanger"
	"log/slog"
)

type Services struct {
	ExchangerService exchanger.Service
}

func NewServices(repos *Repository, l *slog.Logger) *Services {
	return &Services{
		ExchangerService: exchanger.NewService(repos.ExchangerRepo, l),
	}
}
