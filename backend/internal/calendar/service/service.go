package service

import (
	"github.com/google/uuid"
	"github.com/psmey/time_tracker/internal/calendar/repository"
)

type ServicePort interface {
	DeleteCalendar(id uuid.UUID) error
}

type ServiceAdapter struct {
	repository repository.RepositoryPort
}

var _ ServicePort = &ServiceAdapter{}

func New(repository repository.RepositoryPort) *ServiceAdapter {
	return &ServiceAdapter{repository: repository}
}

func (service *ServiceAdapter) DeleteCalendar(id uuid.UUID) error {
	return service.repository.DeleteCalendar(id)
}
