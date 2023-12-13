package repository

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/gbeletti/weatherforecast/api/model"
	"github.com/gbeletti/weatherforecast/persistence/entity"
)

func (r *ForecastRepo) GetForecast(ctx context.Context, fromDate time.Time) ([]model.ForecastModel, error) {
	entities := []entity.Forecast{}
	err := r.db.SelectContext(ctx, &entities, forecastSelect, fromDate)
	if err != nil {
		slog.Error(fmt.Sprintf("error during get forecast %s", err))
	}
	result := make([]model.ForecastModel, 0, len(entities))
	for _, e := range entities {
		result = append(result, e.ToModel())
	}
	return result, err
}

func (r *ForecastRepo) SaveForecast(ctx context.Context, forecasts []model.ForecastModel) error {
	entities := make([]entity.Forecast, 0, len(forecasts))
	for _, f := range forecasts {
		entities = append(entities, entity.Forecast{
			Date:          f.Date,
			WindSpeed:     f.WindSpeed,
			WindDirection: f.WindDirection,
			Alert:         f.Alert,
		})
	}
	_, err := r.db.NamedExecContext(ctx, forecastInsert, entities)
	if err != nil {
		slog.Error(fmt.Sprintf("error during save forecast %s", err))
		return err
	}
	_, err = r.db.ExecContext(ctx, `UPDATE configuration SET last_update = $1 WHERE id = 1;`, time.Now())
	if err != nil {
		slog.Error(fmt.Sprintf("error during update configuration %s", err))
		return err
	}
	return nil
}

func (r *ForecastRepo) GetAlerts(ctx context.Context, fromDate time.Time) ([]model.ForecastModel, error) {
	entities := []entity.Forecast{}
	err := r.db.SelectContext(ctx, &entities, forecastAlertSelect, fromDate)
	if err != nil {
		slog.Error(fmt.Sprintf("error during get alerts %s", err))
	}
	result := make([]model.ForecastModel, 0, len(entities))
	for _, e := range entities {
		result = append(result, e.ToModel())
	}
	return result, err
}

func (r *ForecastRepo) GetLastUpdate(ctx context.Context) (time.Time, error) {
	var lastUpdate time.Time
	err := r.db.GetContext(ctx, &lastUpdate, configurationSelect)
	if err != nil {
		slog.Error(fmt.Sprintf("error during get last update %s", err))
	}
	return lastUpdate, err
}
