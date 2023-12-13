package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gbeletti/weatherforecast/contracts"
)

const (
	defaultTimeout = time.Second * 30
	defaultBaseURL = "https://api.open-meteo.com/v1/forecast"
)

//https://api.open-meteo.com/v1/forecast?latitude=-27.6289&longitude=-48.4478&hourly=wind_speed_180m,wind_direction_180m&timezone=America%2FSao_Paulo

type OpenmeteoConfigger interface {
	GetOpenmeteoBaseURL() string
}

type OpenmeteoClient struct {
	httpClient *http.Client
	baseURL    string
}

func NewOpenmeteoClient(config OpenmeteoConfigger) *OpenmeteoClient {
	cli := &http.Client{
		Timeout: defaultTimeout,
	}
	return &OpenmeteoClient{
		httpClient: cli,
		baseURL:    config.GetOpenmeteoBaseURL(),
	}
}

// GetForecast returns the forecast from openmeteo api for the next 7 days
func (o *OpenmeteoClient) GetForecast(ctx context.Context) (contracts.ForecastModel, error) {
	var forecast contracts.ForecastModel
	url := fmt.Sprintf("%s?latitude=-27.6289&longitude=-48.4478&hourly=wind_speed_180m,wind_direction_180m&timezone=America%%2FSao_Paulo", o.baseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return forecast, err
	}
	resp, err := o.httpClient.Do(req)
	if err != nil {
		return forecast, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&forecast)
	return forecast, err
}
