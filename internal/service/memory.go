package service

import "sync"

type Message struct {
	Role    string
	Content string
}

type Memory struct {
	mu    sync.Mutex
	store map[string][]Message
}

func NewMemory() *Memory {
	return &Memory{
		store: make(map[string][]Message),
	}
}

func (m *Memory) Add(userID string, msg Message) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.store[userID] = append(m.store[userID], msg)

	if len(m.store[userID]) > 20 {
		m.store[userID] = m.store[userID][1:]
	}
}

func (m *Memory) Get(userID string) []Message {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.store[userID]
}
