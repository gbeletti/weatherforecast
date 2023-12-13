package router

import (
	"log/slog"
	"net/http"

	"github.com/gbeletti/weatherforecast/api/handler"
	"github.com/gbeletti/weatherforecast/api/service"
	"github.com/go-chi/chi/v5"
)

func NewRouter(forecast *service.Forecast) http.Handler {
	h := handler.NewHandler(forecast)
	r := chi.NewRouter()
	r.Get("/", rootHandler)
	r.Get("/forecast", h.GetForecast)
	r.Get("/previsao", h.GetForecast)
	r.Get("/alerts", h.GetAlerts)
	r.Get("/alerta", h.GetAlerts)
	return r
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("ok"))
	if err != nil {
		slog.Error(err.Error())
	}
}
