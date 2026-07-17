package sqlite

import (
	"database/sql"

	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/model"
)

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

func (s *UserStore) GetUserPrompt(chatID int64) (model.Prompt, error) {
	query := `
		SELECT p.display_name, p.content
		FROM users u
		JOIN prompts p ON u.selected_prompt_id = p.prompt_id
		WHERE u.chat_id = ?;
	`

	row := s.db.QueryRow(query, chatID)

	var prompt model.Prompt

	err := row.Scan(&prompt.DisplayName, &prompt.Content)
	if err != nil {
		return model.Prompt{}, err
	}

	return prompt, nil
}

func (s *UserStore) UpdatePrompt(chatID int64, promptID int) error {
	query := `
		UPDATE users
		SET selected_prompt_id = ?
		WHERE chat_id = ?
	`

	_, err := s.db.Exec(query, promptID, chatID)
	if err != nil {
		return err
	}

	return nil
}
