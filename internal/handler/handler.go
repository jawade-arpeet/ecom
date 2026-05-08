package handler

import "ecom/internal/service"

type Handler struct {
	Health *HealthHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Health: newHealthHandler(),
	}
}
