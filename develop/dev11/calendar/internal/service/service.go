package service

import (
	"github.com/Regular-Pashka/WbT-L2/develop/dev11/calendar/internal/domain"
	"github.com/Regular-Pashka/WbT-L2/develop/dev11/calendar/internal/repository"
)

type Event interface {
	CreateEvent(event domain.Event) (int, error)
	// GetByDay()
	// GetByWeek()
	// GetByMonth()
	// UpdateEvent()
}

type Service struct {
	// Authorization
	Event
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Event: NewEventService(repos.Event),
	}
}

