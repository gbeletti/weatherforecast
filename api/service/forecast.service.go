package service

import (
	"context"
	"time"

	"github.com/gbeletti/weatherforecast/api/model"
)

// Forecast is the forecast service
type Forecast struct {
	repo ForecastGetter
}

func NewForecast(repo ForecastGetter) *Forecast {
	return &Forecast{
		repo: repo,
	}
}

// GetForecast returns the forecast for the next 7 days
func (f *Forecast) GetForecast(ctx context.Context) ([]model.ForecastModel, *time.Time, error) {
	forecast, err := f.repo.GetForecast(ctx, time.Now())
	if err != nil {
		return nil, nil, err
	}
	lastUpdate, err := f.repo.GetLastUpdate(ctx)
	if err != nil {
		return nil, nil, err
	}
	return forecast, &lastUpdate, nil
}

// GetAlerts returns the alerts for the next 7 days
func (f *Forecast) GetAlerts(ctx context.Context) ([]model.ForecastModel, *time.Time, error) {
	alerts, err := f.repo.GetAlerts(ctx, time.Now())
	if err != nil {
		return nil, nil, err
	}
	lastUpdate, err := f.repo.GetLastUpdate(ctx)
	if err != nil {
		return nil, nil, err
	}
	return alerts, &lastUpdate, nil
}
