package controller

import (
	"net/http"

	"github.com/psmey/time_tracker/internal/features/calendar/domain"
	"github.com/psmey/time_tracker/internal/features/calendar/service"
	"github.com/psmey/time_tracker/internal/http/api"
)

type ControllerPort interface {
	DeleteCalendar(writer http.ResponseWriter, request *http.Request, id api.CalendarId)
}

type ControllerAdapter struct {
	service service.ServicePort
}

var _ ControllerPort = &ControllerAdapter{}

func New(service service.ServicePort) *ControllerAdapter {
	return &ControllerAdapter{service: service}
}

func (controller *ControllerAdapter) DeleteCalendar(writer http.ResponseWriter, request *http.Request, id api.CalendarId) {
	err := controller.service.DeleteCalendar(id)
	if err != nil {
		switch err {
		case domain.ErrCalendarNotFound:
			writer.WriteHeader(http.StatusNotFound)
		default:
			writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
