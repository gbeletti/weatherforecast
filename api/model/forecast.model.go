package model

import "time"

// ForecastModel is the forecast model
type ForecastModel struct {
	Date          time.Time `json:"date"`
	WindSpeed     float64   `json:"windSpeed"`
	WindDirection int       `json:"windDirection"`
	Alert         bool      `json:"alert"`
}
