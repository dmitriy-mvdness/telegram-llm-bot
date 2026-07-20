package handler

import (
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/busy"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/service"
)

type Handler struct {
	svc  *service.Service
	busy *busy.Manager
}

func New(svc *service.Service, busy *busy.Manager) *Handler {
	return &Handler{
		svc:  svc,
		busy: busy,
	}
}
