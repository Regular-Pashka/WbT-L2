package service

import (
	"github.com/Regular-Pashka/WbT-L2/develop/dev11/calendar/internal/repository"
	"github.com/Regular-Pashka/WbT-L2/develop/dev11/calendar/internal/domain"
)

type EventService struct {
	repo repository.Event
}

func NewEventService(repo repository.Event) *EventService {
	return &EventService{
		repo: repo,
	}
}

func (s *EventService) CreateEvent(event domain.Event) (int, error) {
	return s.repo.CreateEvent(event)
}