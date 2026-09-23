# Gateway Service

API Gateway для взаимодействия пользователя с внутренними сервисами.

## Быстрый старт

### Предварительные требования

- Go 1.27+

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
docker build -t gateway-service .

# Запуск контейнера
docker run -p 50053:50053 --env-file .env gateway-service
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
