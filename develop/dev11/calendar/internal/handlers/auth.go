package domain

import "time"
type Event struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Date time.Time `json:"date"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
}