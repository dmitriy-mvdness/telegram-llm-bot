package memory

import (
	"sync"

	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/config"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/model"
)

type MemoryStore struct {
	mu    sync.Mutex
	store map[int64][]model.Message
}

func New() *MemoryStore {
	return &MemoryStore{
		store: make(map[int64][]model.Message),
	}
}

func (m *MemoryStore) Add(chatID int64, msg model.Message) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.store[chatID] = append(m.store[chatID], msg)

	if len(m.store[chatID]) > config.HistoryLimit {
		m.store[chatID] = m.store[chatID][1:]
	}

	return nil
}

func (m *MemoryStore) Get(chatID int64) ([]model.Message, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	history := m.store[chatID]

	result := make([]model.Message, len(history))
	copy(result, history)

	return result, nil
}

func (m *MemoryStore) Clear(chatID int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.store, chatID)
	return nil
}
