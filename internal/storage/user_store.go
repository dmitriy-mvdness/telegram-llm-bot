package storage

type UserStore interface {
	Ensure(chatID int64) error
}
