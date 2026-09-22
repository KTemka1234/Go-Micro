# Auth Service

Микросервис аутентификации пользователей (gRPC, JWT access/refresh токены).

## Быстрый старт

### Предварительные требования

- Go 1.27+
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

# Только запуск (если уже собран)
task run-only
```

#### Docker

```bash
# Сборка образа
docker build -t auth-service .

# Запуск контейнера
docker run -p 50052:50052 --env-file .env auth-service
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
task --list
```
