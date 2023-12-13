package worker

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/gbeletti/weatherforecast/api/model"
	"github.com/gbeletti/weatherforecast/clients"
	"github.com/gbeletti/weatherforecast/config"
	"github.com/gbeletti/weatherforecast/contracts"
	"github.com/jasonlvhit/gocron"
)

const (
	alertWindSpeed      = 20
	timeFormatOpenMeteo = "2006-01-02T15:04"
)

var openmeteoClient ForecastGetter

type ForecastGetter interface {
	GetForecast(ctx context.Context) (contracts.ForecastModel, error)
}

type ForecastSaver interface {
	SaveForecast(ctx context.Context, forecasts []model.ForecastModel) error
}

// Start starts the worker
func Start(ctx context.Context, saver ForecastSaver) {
	configApp := config.NewConfig()
	openmeteoClient = clients.NewOpenmeteoClient(configApp)
	go func() {
		err := gocron.Every(5).Minutes().From(gocron.NextTick()).Do(doForecast, saver)
		if err != nil {
			panic(err)
		}
		select {
		case <-ctx.Done():
			slog.Info("worker stopped")
		case <-gocron.Start():
		}
		gocron.Clear()
	}()
}

func doForecast(saver any) {
	// must convert it
	slog.Info("getting forecast")
	ctx := context.Background()
	forecast, err := filterForecast(ctx, openmeteoClient)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	if len(forecast) == 0 {
		slog.Info("no forecast to save")
		return
	}

	saverRepo, ok := saver.(ForecastSaver)
	if !ok {
		slog.Error("invalid saver")
		return
	}
	err = saverRepo.SaveForecast(ctx, forecast)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info("forecast saved")
}

func filterForecast(ctx context.Context, getter ForecastGetter) ([]model.ForecastModel, error) {
	forecast, err := getter.GetForecast(ctx)
	if err != nil {
		return nil, err
	}
	// filter forecast
	if len(forecast.Hourly.Time) != len(forecast.Hourly.WindDirection180M) || len(forecast.Hourly.Time) != len(forecast.Hourly.WindSpeed180M) {
		return nil, fmt.Errorf("invalid forecast data total hourly time [%d] total hourly wind direction [%d] total hourly wind speed [%d]", len(forecast.Hourly.Time), len(forecast.Hourly.WindDirection180M), len(forecast.Hourly.WindSpeed180M))
	}
	var result []model.ForecastModel
	for i := range forecast.Hourly.Time {
		if forecast.Hourly.WindSpeed180M[i] >= 15 && forecast.Hourly.WindDirection180M[i] >= 130 && forecast.Hourly.WindDirection180M[i] <= 230 {
			date, err := time.Parse(timeFormatOpenMeteo, forecast.Hourly.Time[i])
			if err != nil {
				return nil, fmt.Errorf("invalid date format [%s] error [%s]", forecast.Hourly.Time[i], err)
			}
			result = append(result, model.ForecastModel{
				Date:          date,
				WindSpeed:     forecast.Hourly.WindSpeed180M[i],
				WindDirection: forecast.Hourly.WindDirection180M[i],
				Alert:         forecast.Hourly.WindSpeed180M[i] > alertWindSpeed,
			})
		}
	}
	return result, nil
}
