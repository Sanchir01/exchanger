package app

import (
	"github.com/Sanchir01/exchanger/internal/feature/exchanger"
	"github.com/Sanchir01/exchanger/pkg/db"
)

type Repository struct {
	ExchangerRepo *exchanger.Repository
}

func NewRepository(databases *db.Database) *Repository {
	return &Repository{
		ExchangerRepo: exchanger.NewRepository(databases.PrimaryDB),
	}
}
