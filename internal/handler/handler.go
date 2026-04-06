package handler

type Handler struct {
	Health *HealthHandler
}

func NewHandler() *Handler {
	return &Handler{
		Health: NewHealthHandler(),
	}
}
