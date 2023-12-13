package repository

const (
	forecastSelect = `SELECT f.date, f.wind_speed, f.wind_direction, f.alert FROM forecast f WHERE f.date >= $1 ORDER BY f.date;`

	forecastInsert = `INSERT INTO forecast (date, wind_speed, wind_direction, alert) VALUES (:date, :wind_speed, :wind_direction, :alert) 
	ON CONFLICT (date) DO UPDATE SET wind_speed = excluded.wind_speed, wind_direction = excluded.wind_direction, alert = excluded.alert;`

	forecastAlertSelect = `SELECT f.date, f.wind_speed, f.wind_direction, f.alert FROM forecast f WHERE f.date >= $1 AND f.alert = true ORDER BY f.date;`

	configurationSelect = `SELECT c.last_update FROM configuration c WHERE c.id = 1;`
)
