# Contracts

Общие gRPC/Protobuf-контракты для сервисов Go-Micro. Здесь лежат `.proto`-файлы
и сгенерированный из них Go-код, который подключают остальные сервисы.

Модуль: `github.com/KTemka1234/go-micro/contracts`

## Структура

```
account/         account_model.proto, account_service.proto — сервис Account
account/go/      сгенерированный Go-код
pagination/      pagination.proto — общий тип Pagination
Dockerfile       образ proto-builder (protoc + плагины + googleapis)
```

Пакеты `go/` генерируются автоматически — руками их не правят.

## Генерация

Требуется Docker.

```bash
# Сгенерировать Go-код из всех .proto
task gen

# Пересобрать образ proto-builder
task docker-build

# Удалить сгенерированный код
task clean
```

## Использование в сервисе

```bash
go get github.com/KTemka1234/go-micro/contracts
```

```go
import accountpb "github.com/KTemka1234/go-micro/contracts/account/go"
```
