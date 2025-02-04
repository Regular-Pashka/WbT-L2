package handler

import (
	"github.com/Regular-Pashka/WbT-L2/develop/dev11/calendar/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}
