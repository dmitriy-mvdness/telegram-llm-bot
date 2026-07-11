package storage

import (
	"sync"

	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/model"
)

type Memory struct {
	mu    sync.Mutex
	store map[int64][]model.Message
}

func NewMemory() *Memory {
	return &Memory{
		store: make(map[int64][]model.Message),
	}
}

func (m *Memory) Add(chatID int64, msg model.Message) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.store[chatID] = append(m.store[chatID], msg)

	if len(m.store[chatID]) > 20 {
		m.store[chatID] = m.store[chatID][1:]
	}

	return nil
}

func (m *Memory) Get(chatID int64) ([]model.Message, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	history := m.store[chatID]

	result := make([]model.Message, len(history))
	copy(result, history)

	return result, nil
}

func (m *Memory) Clear(chatID int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.store, chatID)
	return nil
}
