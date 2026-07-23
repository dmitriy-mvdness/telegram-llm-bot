package handler

import (
	"sync"
	"sync/atomic"

	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/busy"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/generation"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/service"
)

type Handler struct {
	svc                *service.Service
	busy               *busy.Manager
	generation         *generation.Manager
	processingMessages sync.Map
	isShuttingDown     atomic.Bool
}

func New(svc *service.Service, busy *busy.Manager, generation *generation.Manager) *Handler {
	return &Handler{
		svc:        svc,
		busy:       busy,
		generation: generation,
	}
}
