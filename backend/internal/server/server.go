package server

import (
	"github.com/psmey/time_tracker/internal/api"
	calendar "github.com/psmey/time_tracker/internal/calendar/controller"
)

type Server struct {
	calendar.CalendarControllerPort
}

var _ api.ServerInterface = &Server{}

func New(calendarController calendar.CalendarControllerPort) *Server {
	return &Server{
		CalendarControllerPort: calendarController,
	}
}
