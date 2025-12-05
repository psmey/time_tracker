package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/psmey/time_tracker/internal/auth/keycloak"
	authmiddleware "github.com/psmey/time_tracker/internal/auth/middleware"
	"github.com/psmey/time_tracker/internal/config"
	"github.com/psmey/time_tracker/internal/http/api"
	"github.com/psmey/time_tracker/internal/logger"
	loggermiddleware "github.com/psmey/time_tracker/internal/logger/middleware"
)

func main() {
	config, err := config.Load("./config.yaml")
	if err != nil {
		logger.LogError("failed to load config", err)
		os.Exit(1)
	}

	logger.Init(config.Logger)
	logger.LogDebug("Initialized Logger with config", "config", config.Logger)

	logger.LogDebug("Trying to initialize Keycloak with config", "config", config.Keycloak)
	keycloakClient, err := keycloak.New(config.Keycloak)
	if err != nil {
		logger.LogError("failed to initialize Keycloak client", err)
		os.Exit(1)
	}

	logger.LogDebug("Setting up middleware")
	middleware := []api.MiddlewareFunc{
		loggermiddleware.New(),
		authmiddleware.New(keycloakClient),
	}

	logger.LogDebug("Registering routes")
	// calendarController, err := calendar.NewCalendarController(config.Database)
	// if err != nil {
	// 	logger.LogError("Failed to load controller", err, "controller", "calendar")
	// }

	// server := server.New(calendarController)
	server := api.Unimplemented{}

	router := chi.NewRouter()

	handler := api.HandlerWithOptions(server, api.ChiServerOptions{
		BaseURL:     "/api/v1",
		BaseRouter:  router,
		Middlewares: middleware,
	})

	logger.LogInfo("starting server and listening on port", "port", 8080)
	http.ListenAndServe(":8080", handler)
}
