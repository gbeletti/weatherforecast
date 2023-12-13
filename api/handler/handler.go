package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/gbeletti/weatherforecast/api/model"
	"github.com/gbeletti/weatherforecast/api/service"
)

type Handler struct {
	forecastService *service.Forecast
}

func NewHandler(forecast *service.Forecast) *Handler {
	return &Handler{
		forecastService: forecast,
	}
}

// GetForecast godoc
//
// @Summary Get forecast returns all the data for the next 7 days
func (h *Handler) GetForecast(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	forecast, lastUpdate, err := h.forecastService.GetForecast(ctx)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeForecastResponse(w, forecast, lastUpdate)
}

// GetAlerts godoc
//
// @Summary Get alerts returns all the alerts for the next 7 days
func (h *Handler) GetAlerts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	alerts, lastUpdate, err := h.forecastService.GetAlerts(ctx)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeForecastResponse(w, alerts, lastUpdate)
}

func writeForecastResponse(w http.ResponseWriter, forecast []model.ForecastModel, lastUpdate *time.Time) {
	resp := response{
		Forecast:   forecast,
		LastUpdate: lastUpdate,
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		slog.Error(err.Error())
	}
}

type response struct {
	Forecast   []model.ForecastModel `json:"forecast"`
	LastUpdate *time.Time            `json:"last_update"`
}
