package repository

import (
    
)

type Event interface {
    CreateEvent()
	GetByDay()
	GetByWeek()
	GetByMonth()
	UpdateEvent()
}

type Repository struct {
    Event
}

func NewRepository() *Repository {
    return &Repository{
        Event: NewEventPostgres(),
    }
}