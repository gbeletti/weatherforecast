package contracts

// ForecastModel is the forecast model
type ForecastModel struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	HourlyUnits          struct {
		Time              string `json:"time"`
		WindSpeed180M     string `json:"wind_speed_180m"`
		WindDirection180M string `json:"wind_direction_180m"`
	} `json:"hourly_units"`
	Hourly struct {
		Time              []string  `json:"time"`
		WindSpeed180M     []float64 `json:"wind_speed_180m"`
		WindDirection180M []int     `json:"wind_direction_180m"`
	} `json:"hourly"`
}
