package server

import (
	calendar "github.com/psmey/time_tracker/internal/features/calendar/controller"
	"github.com/psmey/time_tracker/internal/http/api"
)

type Server struct {
	api.Unimplemented

	calendarController calendar.ControllerInterface
}

var _ api.ServerInterface = &Server{}

func New(calendarController calendar.ControllerInterface) *Server {
	return &Server{
		calendarController: calendarController,
	}
}
