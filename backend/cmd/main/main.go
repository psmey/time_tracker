package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/psmey/time_tracker/internal/api"
	"github.com/psmey/time_tracker/internal/calendar"
	"github.com/psmey/time_tracker/internal/server"
)

func main() {
	router := chi.NewRouter()

	calendarController := calendar.NewCalendarController()

	server := server.New(calendarController)

	handler := api.HandlerFromMux(server, router)

	http.ListenAndServe(":8080", handler)
}
