package storage

import (
	"sync"

	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/model"
)

type Memory struct {
	mu    sync.Mutex
	store map[string][]model.Message
}

func NewMemory() *Memory {
	return &Memory{
		store: make(map[string][]model.Message),
	}
}

func (m *Memory) Add(userID string, msg model.Message) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.store[userID] = append(m.store[userID], msg)

	if len(m.store[userID]) > 20 {
		m.store[userID] = m.store[userID][1:]
	}
}

func (m *Memory) Get(userID string) []model.Message {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.store[userID]
}

func (m *Memory) Clear(userID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.store, userID)
}
