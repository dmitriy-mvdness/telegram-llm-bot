package busy

import "sync"

const noticeReserved = -1

type State struct {
	Busy            bool
	NoticeMessageID int
}

type Manager struct {
	mu    sync.Mutex
	users map[int64]*State
}

func NewManager() *Manager {
	return &Manager{
		users: make(map[int64]*State),
	}
}

func (m *Manager) TryLock(chatID int64) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	state, exists := m.users[chatID]

	if exists && state.Busy {
		return false
	}

	m.users[chatID] = &State{
		Busy: true,
	}
	return true
}

func (m *Manager) Unlock(chatID int64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.users, chatID)
}

func (m *Manager) TryReserveNotice(chatID int64) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	state, exists := m.users[chatID]
	if !exists {
		return false
	}

	if state.NoticeMessageID != 0 {
		return false
	}

	state.NoticeMessageID = noticeReserved
	return true
}

func (m *Manager) ReleaseNoticeReservation(chatID int64) {
	m.mu.Lock()
	defer m.mu.Unlock()

	state, exists := m.users[chatID]
	if !exists {
		return
	}

	if state.NoticeMessageID == noticeReserved {
		state.NoticeMessageID = 0
	}
}

func (m *Manager) SetNoticeMessage(chatID int64, messageID int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	state, exists := m.users[chatID]

	if !exists {
		return
	}

	state.NoticeMessageID = messageID
}

func (m *Manager) GetNoticeMessage(chatID int64) int {
	m.mu.Lock()
	defer m.mu.Unlock()

	state, exists := m.users[chatID]

	if !exists {
		return 0
	}

	if state.NoticeMessageID == noticeReserved {
		return 0
	}

	return state.NoticeMessageID
}