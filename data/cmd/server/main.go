package main

import (
	"context"
	"data-service/config"
	_ "data-service/docs"
	"data-service/internal/application/service"
	domrepo "data-service/internal/domain/repository"
	"data-service/internal/infrastructure/database"
	infrarepo "data-service/internal/infrastructure/repository"
	"data-service/internal/presentation/http/handler"
	"data-service/internal/presentation/http/router"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

// @title Data Service API
// @version 1.0
// @description Централизованный сервис для работы с базой данных
// @description Предоставляет REST API для управления пользователями, продуктами и корзинами

// @contact.name API Support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8004
// @BasePath /
// @schemes http https

func main() {
	// Инициализация логгера
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)

	logger.Info("Starting data-service...")

	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		logger.WithError(err).Fatal("Failed to load configuration")
	}

	logger.WithFields(logrus.Fields{
		"host":    cfg.Server.Host,
		"port":    cfg.Server.Port,
		"db_host": cfg.Database.Host,
	}).Info("Configuration loaded")

	// Подключение к базе данных и инициализация репозиториев
	var (
		userRepo    domrepo.UserRepository
		productRepo domrepo.ProductRepository
		cartRepo    domrepo.CartRepository
	)

	db, err := database.NewPostgresDB(cfg.Database.DSN())
	if err != nil {
		logger.WithError(err).Fatal("Failed to connect to database")
	}
	defer func() {
		if err := db.Close(); err != nil {
			logger.WithError(err).Error("Failed to close database connection")
		}
	}()

	userRepo = infrarepo.NewUserRepository(db.GetDB())
	productRepo = infrarepo.NewProductRepository(db.GetDB())
	cartRepo = infrarepo.NewCartRepository(db.GetDB())

	// Инициализация сервисов
	userService := service.NewUserService(userRepo)
	productService := service.NewProductService(productRepo)
	cartService := service.NewCartService(cartRepo)

	// Инициализация обработчиков
	userHandler := handler.NewUserHandler(userService, logger)
	productHandler := handler.NewProductHandler(productService, logger)
	cartHandler := handler.NewCartHandler(cartService, logger)
	healthHandler := handler.NewHealthHandler()

	// Инициализация роутера
	r := router.NewRouter(
		userHandler,
		productHandler,
		cartHandler,
		healthHandler,
		logger,
	)
	r.Setup()

	// Настройка HTTP сервера
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      r.GetEngine(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Запуск сервера в отдельной горутине
	go func() {
		logger.WithField("address", addr).Info("Starting HTTP server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.WithError(err).Fatal("Failed to start server")
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.WithError(err).Error("Server forced to shutdown")
	}

	logger.Info("Server exited")
}
