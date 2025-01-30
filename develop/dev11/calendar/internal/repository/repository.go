package repository

import (
	"github.com/jmoiron/sqlx"
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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Event: NewEventPostgres(db),
	}
}
