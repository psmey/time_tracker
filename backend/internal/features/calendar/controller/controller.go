package controller

import (
	"net/http"

	"github.com/psmey/time_tracker/internal/features/calendar/domain"
	"github.com/psmey/time_tracker/internal/features/calendar/service"
	"github.com/psmey/time_tracker/internal/http/api"
)

type ControllerInterface interface {
	DeleteCalendar(writer http.ResponseWriter, request *http.Request, id api.CalendarId)
}

type Controller struct {
	service service.ServicePort
}

var _ ControllerInterface = &Controller{}

func New(service service.ServicePort) *Controller {
	return &Controller{service: service}
}

func (controller *Controller) DeleteCalendar(writer http.ResponseWriter, request *http.Request, id api.CalendarId) {
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
