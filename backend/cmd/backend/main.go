package main

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/psmey/time_tracker/internal/config"
	"github.com/psmey/time_tracker/internal/features/calendar"
	"github.com/psmey/time_tracker/internal/http/api"
	"github.com/psmey/time_tracker/internal/http/server"
	"github.com/psmey/time_tracker/internal/logging"
)

var logger *slog.Logger

func main() {
	config, err := config.Load("./config.yaml")
	if err != nil {
		slog.Default().Error("Failed to load config", "error", err.Error())
	}

	logger = logging.Init(config.Logger)

	calendarController, err := calendar.NewCalendarController(config.Database)
	if err != nil {
		logger.Error("Failed to load controller", "controller", "calendar", "error", err)
	}

	server := server.New(calendarController)

	router := chi.NewRouter()

	handler := api.HandlerFromMux(server, router)

	http.ListenAndServe(":8080", handler)
}
