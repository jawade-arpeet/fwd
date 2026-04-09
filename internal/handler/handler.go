package handler

import "fwd/internal/service"

type Handler struct {
	Health  *HealthHandler
	Account *AccountHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Health:  NewHealthHandler(),
		Account: NewAccountHandler(service.Account),
	}
}
