package handler

import "github.com/dmitriy-mvdness/telegram-llm-bot/internal/service"

type Handler struct {
	svc *service.Service
}

func New(svc *service.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) Handle(input string) string {
	return h.svc.Process(input)
}
