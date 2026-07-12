package sqlite

import (
	"database/sql"
	"slices"

	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/config"
	"github.com/dmitriy-mvdness/telegram-llm-bot/internal/model"
)

type SQLiteStore struct {
	db *sql.DB
}

func New(db *sql.DB) *SQLiteStore {
	return &SQLiteStore{db: db}
}

func (s *SQLiteStore) Add(chatID int64, msg model.Message) error {
	query := `
		INSERT INTO messages(chat_id, role, content)
		VALUES (?, ?, ?);
	`

	_, err := s.db.Exec(query, chatID, msg.Role, msg.Content)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLiteStore) Get(chatID int64) ([]model.Message, error) {
	query := `
		SELECT role, content
		FROM messages
		WHERE chat_id = ?
		ORDER BY message_id DESC
		LIMIT ?;
	`

	rows, err := s.db.Query(query, chatID, config.HistoryLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := make([]model.Message, 0)
	for rows.Next() {
		var msg model.Message
		if err := rows.Scan(&msg.Role, &msg.Content); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	slices.Reverse(messages)

	return messages, nil
}

func (s *SQLiteStore) Clear(chatID int64) error {
	query := `
		DELETE FROM messages
		WHERE chat_id = ?
	`

	_, err := s.db.Exec(query, chatID)
	if err != nil {
		return err
	}

	return nil
}
