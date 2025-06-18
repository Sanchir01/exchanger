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

func (r *Repository) GetAllCurrency(ctx context.Context) ([]ExchangeDB, error) {
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

	var results []ExchangeDB
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
		results = append(results, r)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *Repository) GetCurrencyByRate(ctx context.Context, fromCurrency, toCurrency string) (*ExchangeDB, error) {
	conn, err := r.primaryDB.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	query, args, err := sq.
		Select("id, from_currency, to_currency, rate,created_at, updated_at").
		From("exchange_rates").Where(sq.Eq{"from_currency": fromCurrency, "to_currency": toCurrency}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, utils.ErrorQueryString
	}
	res := &ExchangeDB{}
	if err := conn.QueryRow(ctx, query, args...).Scan(&res.ID, &res.FromCurrency, &res.ToCurrency, &res.Rate, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return nil, err
	}

	return res, nil
}
