package calendar

import (
	"github.com/psmey/time_tracker/internal/database"
	"github.com/psmey/time_tracker/internal/features/calendar/controller"
	"github.com/psmey/time_tracker/internal/features/calendar/repository"
	"github.com/psmey/time_tracker/internal/features/calendar/service"
)

func NewCalendarController(config *database.Config) (*controller.Controller, error) {
	connector := database.NewConnector(config)
	db, err := connector.NewPostgresDB("calendars")
	if err != nil {
		return nil, err
	}
	calendarRepository := repository.NewPostgresRepository(db)
	calendarService := service.New(calendarRepository)
	calendarController := controller.New(calendarService)

	return calendarController, nil
}
