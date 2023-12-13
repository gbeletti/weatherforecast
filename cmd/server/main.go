package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gbeletti/weatherforecast/api/router"
	"github.com/gbeletti/weatherforecast/api/service"
	"github.com/gbeletti/weatherforecast/persistence/repository"
	"github.com/gbeletti/weatherforecast/worker"
)

const defaultPort = "8080"

func main() {
	ctxStart, cancelStart := context.WithCancel(context.Background())
	repo, err := repository.GetPostgresRepo()
	if err != nil {
		slog.Error(err.Error())
		return
	}
	worker.Start(ctxStart, repo)
	srv := startWebServer(repo)
	waitShutdown()
	cancelStart()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error(err.Error())
	}
	if err := repo.Close(); err != nil {
		slog.Error(err.Error())
	}
}

func startWebServer(repo *repository.ForecastRepo) *http.Server {
	forecast := service.NewForecast(repo)
	router := router.NewRouter(forecast)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	srv := &http.Server{
		Addr:              fmt.Sprintf("0.0.0.0:%s", port),
		Handler:           router,
		ReadHeaderTimeout: time.Second * 7,
		ReadTimeout:       time.Second * 10,
		WriteTimeout:      time.Second * 20,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
	return srv
}

// waitShutdown waits for a shutdown signal
func waitShutdown() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-sigc
	slog.Info("got signal %s, shutting down", s)
}
