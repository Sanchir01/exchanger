package app

import (
	"github.com/Sanchir01/exchanger/internal/feature/exchanger"
	"github.com/Sanchir01/exchanger/pkg/db"
)

type Services struct {
	ExchangerService exchanger.Service
}

func NewServices(repos *Repository, db *db.Database) *Services {
	return &Services{
		ExchangerService: exchanger.NewService(repos.ExchangerRepo),
	}
}
