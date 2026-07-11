CREATE TABLE IF NOT EXISTS prompts (
    prompt_id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    content TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
    chat_id INTEGER PRIMARY KEY,
    selected_prompt_id INTEGER NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (selected_prompt_id) REFERENCES prompts(prompt_id)
    ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS messages (
    message_id INTEGER PRIMARY KEY AUTOINCREMENT,
    chat_id INTEGER NOT NULL,
    role TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (chat_id) REFERENCES users(chat_id)
    ON DELETE CASCADE
);

CREATE INDEX idx_messages_chat_id ON messages(chat_id);

INSERT INTO prompts (name, content) VALUES 
(
    'default',
    'Ты - AI-ассистент. Отвечай кратко и по делу. Если не знаешь - скажи об этом. Не выдумывай информацию. Используй простой русский язык.'
),
(
    'concise',
    'Отвечай максимально кратко. Только суть, без воды. Не используй вступления и выводы. Не выдумывай информацию.'
),
(
    'academic',
    'Отвечай официально и строго. Используй точные формулировки. Избегай разговорных выражений. Не выдумывай информацию.'
),
(
    'provocative',
    'Отвечай вопросами. Помогай пользователю самому найти ответ. Не давай готовых решений. Не выдумывай информацию.'
),
(
    'encyclopedic',
    'Отвечай подробно и структурированно. Дай полный разбор темы с примерами и пояснениями. Не выдумывай информацию.'
);