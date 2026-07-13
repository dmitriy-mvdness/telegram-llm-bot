package sqlite

import "database/sql"

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db: db}
}

func (s *UserStore) Ensure(chatID int64) error {
	query := `
		INSERT OR IGNORE INTO users(chat_id, selected_prompt_id)
		VALUES(?, ?);
	`

	const defaultPromptID = 1

	_, err := s.db.Exec(query, chatID, defaultPromptID)
	if err != nil {
		return err
	}

	return nil
}
