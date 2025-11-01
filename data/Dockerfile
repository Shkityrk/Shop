# Используем multi-stage build для оптимизации размера образа
# Stage 1: Builder
FROM golang:1.23-alpine AS builder

# Устанавливаем необходимые пакеты
RUN apk add --no-cache git gcc musl-dev

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum (если есть)
COPY go.mod go.sum* ./

# Копируем весь исходный код для правильного разрешения зависимостей
COPY . .

# Добавляем все недостающие зависимости и скачиваем их
RUN go mod tidy && go mod download

# Устанавливаем swag CLI
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Генерируем Swagger документацию
RUN swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal

# Собираем приложение
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server/main.go

# Stage 2: Runtime
FROM alpine:latest

# Устанавливаем необходимые пакеты для runtime
RUN apk --no-cache add ca-certificates tzdata

# Создаем непривилегированного пользователя
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

WORKDIR /app

# Копируем бинарный файл из builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

# Меняем владельца файлов
RUN chown -R appuser:appuser /app

# Переключаемся на непривилегированного пользователя
USER appuser

# Открываем порт
EXPOSE 8004

# Запускаем приложение
CMD ["./main"]

