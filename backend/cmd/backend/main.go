package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/psmey/time_tracker/internal/configloader"
	"github.com/psmey/time_tracker/internal/features/calendar"
	"github.com/psmey/time_tracker/internal/http/api"
	"github.com/psmey/time_tracker/internal/http/server"
)

func main() {
	configLoader := configloader.New()
	config, err := configLoader.Load("./config.yaml")
	if err != nil {
		log.Print(err)
		return
	}

	calendarController, err := calendar.NewCalendarController(config.Database)
	if err != nil {
		log.Print(err)
		return
	}

	server := server.New(calendarController)

	router := chi.NewRouter()

	handler := api.HandlerFromMux(server, router)

	http.ListenAndServe(":8080", handler)
}
