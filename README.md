# Telegram LLM Bot

<div align="center">

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Ollama](https://img.shields.io/badge/Ollama-0.1.0+-green?style=for-the-badge&logo=ollama&logoColor=white)](https://ollama.ai/)
[![Status](https://img.shields.io/badge/Status-Development-orange?style=for-the-badge)](https://github.com/dmitriy-mvdness/telegram-llm-bot)
[![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)](LICENSE)

**Telegram бот с локальной LLM на Go + Ollama**

</div>

---

## О проекте

Этот Telegram бот позволяет общаться с локальной языковой моделью через Ollama. Всё работает полностью офлайн — никаких облачных провайдеров и платежей.

**Как начать:**

1. Клонируешь репозиторий
2. Настраиваешь переменные окружения (токен бота, параметры модели)
3. Запускаешь приложение — и бот готов к работе

> Проект на стадии активной разработки. Используется для личных целей и изучения возможностей Go + LLM + Telegram API.

### Особенности

- **Локальная LLM** — все запросы обрабатываются через Ollama, ничего не уходит в облако
- **Написан на Go** — эффективное использование ресурсов и простой деплой
- **Контекст диалога** — бот помнит историю сообщений

---

## Технологии

- [Go](https://go.dev/) — основной язык
- [Ollama](https://ollama.ai/) — запуск LLM локально
- [Telegram Bot API](https://github.com/go-telegram/bot) — взаимодействие с Telegram

---

## Команды

| Команда | Описание |
|---------|----------|
| `/start` | Приветственное сообщение и начало работы |
| `/help`  | Список доступных команд |
| `/clear` | Очистка истории диалога |

> Больше команд появится по мере развития проекта.

---

## Установка и запуск

### Требования
- Установленный [Go](https://go.dev/dl/) (версия 1.21+)
- Установленный и запущенный [Ollama](https://ollama.com/download)
- Загруженная LLM модель (например, `ollama run llama3.1` или `openchat:7b`)

### Шаги

1. **Клонируй репозиторий**
```bash
git clone https://github.com/dmitriy-mvdness/telegram-llm-bot.git
cd telegram-llm-bot
```

2. **Установи зависимости**
```bash
go mod download
```

3. **Настрой переменные окружения**
```bash
# Скопируй пример конфига
cp .env.example .env

# Отредактируй .env, указав свои данные и настроки для llm
# TELEGRAM_TOKEN=твой_токен_бота
# OLLAMA_BASE_URL=http://localhost:11434
# OLLAMA_MODEL=llama3.1
# ...
```

4. **Запусти бота**
```bash
go run cmd/app/main.go
```

Или собери бинарник:
```bash
go build -o telegram-bot cmd/app/main.go
./telegram-bot
```

Или же воспользуйся Makefile:
```bash
# Запуск
make run

# Сборка
make build
```

---

## Как помочь проекту

- Поставь звезду на GitHub
- Сообщай об ошибках в [Issues](https://github.com/dmitriy-mvdness/telegram-llm-bot/issues)
- Предлагай новые идеи
- Отправляй Pull Requests с улучшениями
