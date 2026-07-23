package generation

import (
	"context"
	"sync"
)

type Manager struct {
	mu     sync.Mutex
	cancel map[int64]context.CancelFunc
}

func NewManager() *Manager {
	return &Manager{
		cancel: make(map[int64]context.CancelFunc),
	}
}

func (m *Manager) Set(cancel context.CancelFunc, chatID int64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cancel[chatID] = cancel
}

func (m *Manager) Cancel(chatID int64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	cancel, ok := m.cancel[chatID]
	if !ok {
		return
	}

	cancel()
	delete(m.cancel, chatID)
}

func (m *Manager) Delete(chatID int64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.cancel, chatID)
}

func (m *Manager) Exists(chatID int64) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.cancel[chatID]
	return ok
}
