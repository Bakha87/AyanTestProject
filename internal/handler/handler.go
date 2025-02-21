package handler

import (
	"AyanTestProject/internal/service"
)

type Handler struct {
	srv service.Service
}

func New(srv service.Service) *Handler {
	return &Handler{srv: srv}
}
