#!/bin/bash

echo "Generating Swagger documentation..."

# Переход в корневую директорию проекта
cd "$(dirname "$0")/.."

# Генерация Swagger документации
swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal

echo "Swagger documentation generated successfully!"
echo "Available at: http://localhost:8004/swagger/index.html"
