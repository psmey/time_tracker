package server

import (
	"net/http"

	calendar "github.com/psmey/time_tracker/internal/features/calendar/controller"
	"github.com/psmey/time_tracker/internal/http/api"
)

type Server struct {
	api.Unimplemented

	calendarController calendar.ControllerPort
}

var _ api.ServerInterface = &Server{}

func New(calendarController calendar.ControllerPort) *Server {
	return &Server{
		calendarController: calendarController,
	}
}

func (server *Server) DeleteCalendar(writer http.ResponseWriter, request *http.Request, id api.CalendarId) {
	server.calendarController.DeleteCalendar(writer, request, id)
}
