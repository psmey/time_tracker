package calendar

import (
	"github.com/psmey/time_tracker/internal/calendar/controller"
	"github.com/psmey/time_tracker/internal/calendar/repository"
	"github.com/psmey/time_tracker/internal/calendar/service"
)

func NewCalendarController() *controller.ControllerAdapter {
	// TODO add DB connection
	calendarRepository := repository.NewPostgresRepository()
	calendarService := service.New(calendarRepository)
	calendarController := controller.New(calendarService)

	return calendarController
}
