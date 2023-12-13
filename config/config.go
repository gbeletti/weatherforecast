package config

import (
	"os"

	"github.com/gbeletti/weatherforecast/constants"
)

type Config struct {
	openMeteoURL string
}

func NewConfig() *Config {
	opURL := os.Getenv("OPENMETEO_URL")
	if len(opURL) == 0 {
		opURL = constants.OpenMeteoURL
	}
	return &Config{
		openMeteoURL: opURL,
	}
}

func (c *Config) GetOpenmeteoBaseURL() string {
	return c.openMeteoURL
}
