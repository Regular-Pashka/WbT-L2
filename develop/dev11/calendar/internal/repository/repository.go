package repository

import (
	"github.com/jmoiron/sqlx"
    "github.com/Regular-Pashka/WbT-L2/develop/dev11/calendar/internal/domain"
)

type Event interface {
	CreateEvent(event domain.Event) (int, error)
	// GetByDay()
	// GetByWeek()
	// GetByMonth()
	// UpdateEvent()
}

type Repository struct {
	Event
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Event: NewEventPostgres(db),
	}
}
