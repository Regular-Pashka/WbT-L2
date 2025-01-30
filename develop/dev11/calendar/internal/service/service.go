package service

type Event interface {
	CreateEvent()
	GetByDay()
	GetByWeek()
	GetByMonth()
	UpdateEvent()
}

type Service struct {
	// Authorization
	Event
}

func NewService() *Service {
	return &Service{
		Event: NewEventService(repos.Event),
	}
}