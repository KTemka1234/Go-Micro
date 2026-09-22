# Account Service

Микросервис для управления аккаунтами пользователей.

## Быстрый старт

### Предварительные требования

- Go 1.25+
- PostgreSQL

### Установка зависимостей

```bash
task deps
```

### Настройка окружения

```bash
cp .env.example .env
```

### Запуск

#### Локальная разработка

```bash
# Сборка и запуск
task run

# Только запуск
task run-only
```

#### Docker

```bash
# Сборка образа
docker build -t account-service .

# Запуск контейнера
docker run -p 50051:50051 --env-file .env account-service
```

### Полезные команды

```bash
# Форматирование кода
task fmt

# Линтинг
task lint

# Тесты
task test

# Тесты с покрытием
task test-coverage

# Очистка артефактов сборки
task clean

# Показать все доступные команды
task help
```
