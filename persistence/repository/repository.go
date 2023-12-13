package repository

import (
	"github.com/jmoiron/sqlx"
)

// ForecastRepo is the forecast repository
type ForecastRepo struct {
	db *sqlx.DB
}

func (f *ForecastRepo) Close() error {
	return f.db.Close()
}

// NewForecastRepo creates a new forecast repository
func NewForecastRepo(db *sqlx.DB) *ForecastRepo {
	return &ForecastRepo{
		db: db,
	}
}
