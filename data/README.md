# Data Service

Data Service - централизованный сервис для работы с базой данных в архитектуре микросервисов Shop.

## Описание

Data Service предоставляет единую точку доступа к базе данных PostgreSQL для всех других микросервисов (auth, product, cart). Это повышает безопасность, так как только один сервис имеет прямой доступ к БД.

## Архитектура

Проект построен на принципах чистой архитектуры (Clean Architecture):

```
data/
├── cmd/server/          # Точка входа приложения
├── config/              # Конфигурация
├── internal/
│   ├── domain/          # Бизнес-логика, сущности, интерфейсы
│   │   ├── entity/      # Доменные модели
│   │   └── repository/  # Интерфейсы репозиториев
│   ├── application/     # Слой приложения
│   │   ├── dto/         # Data Transfer Objects
│   │   └── service/     # Бизнес-сервисы
│   ├── infrastructure/  # Инфраструктура
│   │   ├── database/    # Подключение к БД
│   │   └── repository/  # Реализация репозиториев
│   └── presentation/    # Слой представления
│       ├── http/        # HTTP handlers и роутинг
│       └── middleware/  # HTTP middleware
├── Dockerfile
├── go.mod
└── README.md
```

## API Endpoints

### Users

- `POST /api/users` - Создать пользователя
- `GET /api/users` - Получить список пользователей
- `GET /api/users/:id` - Получить пользователя по ID
- `GET /api/users/username/:username` - Получить пользователя по username
- `GET /api/users/email/:email` - Получить пользователя по email
- `PUT /api/users/:id` - Обновить пользователя
- `DELETE /api/users/:id` - Удалить пользователя
- `POST /api/users/exists` - Проверить существование пользователя

### Products

- `POST /api/products` - Создать продукт
- `GET /api/products` - Получить список продуктов
- `GET /api/products/:id` - Получить продукт по ID
- `GET /api/products/name/:name` - Получить продукт по названию
- `PUT /api/products/:id` - Обновить продукт
- `DELETE /api/products/:id` - Удалить продукт
- `GET /api/products/:id/exists` - Проверить существование продукта

### Cart

- `POST /api/cart` - Создать элемент корзины
- `GET /api/cart` - Получить список элементов корзины
- `GET /api/cart/:id` - Получить элемент корзины по ID
- `GET /api/cart/user/:user_id` - Получить корзину пользователя
- `GET /api/cart/user/:user_id/product/:product_id` - Получить конкретный элемент корзины
- `PUT /api/cart/:id` - Обновить элемент корзины
- `DELETE /api/cart/:id` - Удалить элемент корзины
- `DELETE /api/cart/user/:user_id` - Очистить корзину пользователя
- `DELETE /api/cart/user/:user_id/product/:product_id` - Удалить продукт из корзины

### Health

- `GET /health` - Проверка состояния сервиса

## Переменные окружения

```env
# Server
DATA_SERVICE_PORT=8004
DATA_SERVICE_HOST=0.0.0.0

# Database
DB_HOST=db
DB_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=shop
DB_SSLMODE=disable
```

## Запуск

### С помощью Docker

```bash
docker build -t data-service .
docker run -p 8004:8004 --env-file .env data-service
```

### С помощью Docker Compose

```bash
docker-compose up data
```

### Локально

```bash
# Установка зависимостей
go mod download

# Запуск
go run cmd/server/main.go
```

## Технологии

- **Go 1.21** - язык программирования
- **Gin** - HTTP фреймворк
- **PostgreSQL** - база данных
- **database/sql** - работа с БД
- **Logrus** - логирование

## Принципы

- **Clean Architecture** - разделение на слои
- **SOLID** - принципы ООП
- **Dependency Injection** - инверсия зависимостей
- **Interface Segregation** - использование интерфейсов
- **Single Responsibility** - каждый компонент отвечает за одну задачу

## Безопасность

- Только data-service имеет прямой доступ к базе данных
- Все остальные сервисы обращаются к данным через HTTP API
- Использование непривилегированного пользователя в Docker
- Graceful shutdown для корректного завершения работы

## Мониторинг

- Health check endpoint для проверки состояния сервиса
- Структурированное логирование (JSON format)
- Логирование всех HTTP запросов с метриками (время выполнения, статус код, и т.д.)

