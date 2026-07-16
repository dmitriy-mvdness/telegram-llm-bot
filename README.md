# Telegram LLM Assistant

<div align="center">

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Ollama](https://img.shields.io/badge/Ollama-Local-green?style=for-the-badge&logo=ollama&logoColor=white)](https://ollama.com/)
[![Status](https://img.shields.io/badge/Status-Development-orange?style=for-the-badge)](https://github.com/dmitriy-mvdness/telegram-llm-bot)
[![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)](LICENSE)

**Telegram-бот с локальной LLM на Go + Ollama**

</div>

---

## О проекте

Telegram-бот для общения с локальной языковой моделью через [Ollama](https://ollama.com/).

Все AI-запросы обрабатываются локально — данные не передаются сторонним AI-провайдерам. Единственное внешнее взаимодействие происходит через Telegram Bot API.

**Быстрый старт:**

1. Клонировать репозиторий
2. Настроить переменные окружения
3. Запустить приложение

После запуска бот готов принимать сообщения и отвечать через выбранную LLM-модель.

> Проект находится в активной разработке. Создан для изучения возможностей Go, локальных LLM и Telegram Bot API.

---

## Особенности

- **Локальная LLM** — запуск моделей через Ollama без использования внешних AI-сервисов
- **Go** — высокая производительность, низкое потребление ресурсов и простой деплой
- **Контекст диалога** — сохранение истории сообщений для поддержания контекста общения
- **Гибкое хранилище** — история может храниться в оперативной памяти (RAM) или базе данных

---

## Технологии

- [Go](https://go.dev/) — основной язык разработки
- [Ollama](https://ollama.com/) — локальный запуск LLM-моделей
- [Telegram Bot API](https://core.telegram.org/bots/api) — взаимодействие с Telegram
- [go-telegram/bot](https://github.com/go-telegram/bot) — Go-библиотека для Telegram Bot API

---

## Команды

| Команда | Описание |
|---------|----------|
| `/start` | Запуск бота и приветственное сообщение |
| `/help` | Просмотр доступных команд |
| `/settings` | Настройки ассистента и управление историей диалога |

> Новые команды будут добавляться по мере развития проекта.

---

## Установка и запуск

### Требования
- Установленный [Go](https://go.dev/dl/) версии **1.21+**
- Установленный и запущенный [Ollama](https://ollama.com/download)
- Загруженная LLM модель (например, `llama3.1` или `qwen3`)
- Telegram Bot Token для подключения к API

---

### 1. Клонирование репозитория

```bash
git clone https://github.com/dmitriy-mvdness/telegram-llm-bot.git

cd telegram-llm-bot
````

---

### 2. Установка зависимостей

```bash
go mod download
```

---

### 3. Настройка переменных окружения

Создай файл `.env`:

```bash
cp .env.example .env
```

Отредактируй настройки:

```env
TELEGRAM_TOKEN=твой_токен_бота

OLLAMA_BASE_URL=http://localhost:11434

OLLAMA_MODEL=llama3.1
```

---

## Конфигурация

Основные параметры:

| Переменная | Описание | Пример |
|------------|----------|--------|
| `TELEGRAM_TOKEN` | Токен Telegram-бота | `123456:ABC...` |
| `OLLAMA_MODEL` | Используемая LLM-модель | `llama3.1:8b` |
| `OLLAMA_BASE_URL` | Адрес Ollama API | `http://localhost:11434` |
| `OLLAMA_NUM_CTX` | Размер контекстного окна модели | `8192` |
| `OLLAMA_NUM_PREDICT` | Максимальное количество токенов в ответе | `2048` |
| `OLLAMA_TEMPERATURE` | Настройка креативности ответов модели | `0.4` |
| `OLLAMA_HEALTH_TIMEOUT` | Таймаут проверки доступности Ollama | `30` |
| `OLLAMA_THINK` | Включение режима thinking модели | `false` |
| `DATABASE_URL` | Путь к базе данных | `file:storage.db` |

---

### 4. Запуск

Запуск через Go:

```bash
go run cmd/app/main.go
```

Или сборка бинарника:

```bash
go build -o telegram-bot cmd/app/main.go

./telegram-bot
```

---

### Запуск через Makefile

```bash
# Запуск приложения
make run

# Сборка бинарника
make build

# Удаление бинарника
make clean

# Обновление зависимостей
make tidy
```

---

## Архитектура хранения контекста

Бот поддерживает несколько вариантов хранения истории диалога:

* **RAM** — быстрый вариант без внешних зависимостей
* **Database** — постоянное хранение истории между перезапусками

Способ хранения можно заменить через изменение реализации хранилища без изменения основной логики приложения.

---

## Как помочь проекту

Буду рад любой помощи:

* Поставить ⭐ проекту на GitHub
* Сообщить об ошибке через [Issues](https://github.com/dmitriy-mvdness/telegram-llm-bot/issues)
* Предложить новые идеи
* Создать Pull Request с улучшениями

---
