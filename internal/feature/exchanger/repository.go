package exchanger

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/Sanchir01/exchanger/pkg/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	primaryDB *pgxpool.Pool
}

func NewRepository(primaryDB *pgxpool.Pool) *Repository {
	return &Repository{primaryDB: primaryDB}
}

func (r *Repository) GetAllCourse(ctx context.Context) ([]*ExchangeDB, error) {
	conn, err := r.primaryDB.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	query, args, err := sq.
		Select("id, from_currency, to_currency, rate,created_at, updated_at").
		From("exchange_rates").
		ToSql()
	if err != nil {
		return nil, utils.ErrorQueryString
	}
	var results []*ExchangeDB
	rows, err := conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var r ExchangeDB
		err := rows.Scan(
			&r.ID,
			&r.FromCurrency,
			&r.ToCurrency,
			&r.Rate,
			&r.CreatedAt,
			&r.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, &r)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return nil, nil
}
