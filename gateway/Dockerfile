FROM golang:1.24.5-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# копируем весь код
COPY . .

# если main.go в корне проекта
RUN go build -o app ./cmd

# --- финальный образ ---
FROM alpine:3.18
WORKDIR /app

COPY --from=build /app/app .

CMD ["./app"]