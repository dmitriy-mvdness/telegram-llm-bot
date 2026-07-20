package busy

import "sync"

type Manager struct {
	mu    sync.Mutex
	users map[int64]struct{}
}

func NewManager() *Manager {
	return &Manager{
		users: make(map[int64]struct{}),
	}
}

func (m *Manager) TryLock(chatID int64) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.users[chatID]; exists {
		return false
	}

	m.users[chatID] = struct{}{}
	return true
}

func (m *Manager) Unlock(chatID int64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.users, chatID)
}
