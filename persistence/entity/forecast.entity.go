package entity

import (
	"time"

	"github.com/gbeletti/weatherforecast/api/model"
)

// Forecast represents the forecast entity in database
type Forecast struct {
	Date          time.Time `db:"date"`
	WindSpeed     float64   `db:"wind_speed"`
	WindDirection int       `db:"wind_direction"`
	Alert         bool      `db:"alert"`
}

type Configuration struct {
	ID         int        `db:"id"`
	LastUpdate *time.Time `db:"last_update"`
}

// ToModel converts the entity to model
func (f Forecast) ToModel() model.ForecastModel {
	return model.ForecastModel{
		Date:          f.Date,
		WindSpeed:     f.WindSpeed,
		WindDirection: f.WindDirection,
		Alert:         f.Alert,
	}
}
