package repository

import (
	"fmt"

	"github.com/Regular-Pashka/WbT-L2/develop/dev11/calendar/internal/domain"
	"github.com/jmoiron/sqlx"
)

type EventPostgres struct {
	db *sqlx.DB
}

func NewCalendarEventPostgres(db *sqlx.DB) *EventPostgres {
	return &EventPostgres{
		db: db,
	}
}

func (r *EventPostgres) CreateEvent(event domain.Event) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createEventQuery := fmt.Sprint("INSERT INTO %s (title, description, date, start_time, end_time) VALUES ($1, $2, $3, $4, $5) RETURNING id", event)
	row := tx.QueryRow(createEventQuery, event.Title, event.Description, event.Date, event.StartTime, event.EndTime)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
