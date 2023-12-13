package service

import (
	"context"
	"time"

	"github.com/gbeletti/weatherforecast/api/model"
)

// ForecastGetter is the interface to get the forecast from storage
type ForecastGetter interface {
	GetForecast(ctx context.Context, fromDate time.Time) ([]model.ForecastModel, error)
	GetAlerts(ctx context.Context, fromDate time.Time) ([]model.ForecastModel, error)
	GetLastUpdate(ctx context.Context) (time.Time, error)
}
