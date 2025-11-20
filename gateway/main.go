package main

import (
	"gateway/internal/application/services"
	"gateway/internal/infrastructure/config"
	"gateway/internal/infrastructure/http"
	"gateway/internal/presentation/handlers"
	"gateway/internal/presentation/middleware"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "gateway/docs"
)

func main() {
	// Получаем порт из переменной окружения или используем значение по умолчанию
	port := config.App.GatewayPort

	// Получаем URL'ы сервисов из переменных окружения
	authServiceURL := config.App.AuthServiceURL
	productServiceURL := config.App.ProductServiceURL
	cartServiceURL := config.App.CartServiceURL

	router := gin.Default()

	// Добавляем CORS middleware
	router.Use(middleware.CORSMiddleware())

	// Создаем репозитории
	authRepo := http.NewAuthRepositoryHTTP(authServiceURL)
	productRepo := http.NewProductRepositoryHTTP(productServiceURL)
	cartRepo := http.NewCartRepositoryHTTP(cartServiceURL)

	// Создаем сервисы
	authService := services.NewAuthService(authRepo)
	productService := services.NewProductService(productRepo)
	cartService := services.NewCartService(cartRepo)

	// Создаем handlers
	authHandler := handlers.NewAuthHandler(authService)
	productHandler := handlers.NewProductHandler(productService)
	cartHandler := handlers.NewCartHandler(cartService)

	// Создаем middleware для аутентификации
	authMiddleware := middleware.AuthMiddleware(authService)

	// Auth routes (некоторые требуют аутентификации)
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/logout", authHandler.Logout)
		// /auth/info требует аутентификации
		authGroup.GET("/info", authMiddleware, authHandler.Info)
	}

	// Product routes
	productGroup := router.Group("/product")
	{
		// Публичные роуты
		productGroup.GET("/list", productHandler.List)
		productGroup.GET("/verify/:name", productHandler.Verify)
		productGroup.GET("/info/:id", productHandler.Info)

		// Защищенные роуты (требуют аутентификации)
		productGroup.POST("/add", authMiddleware, productHandler.Add)
		productGroup.PUT("/update/:id", authMiddleware, productHandler.Update)
	}

	// Cart routes (все требуют аутентификации)
	cartGroup := router.Group("/cart")
	cartGroup.Use(authMiddleware)
	{
		cartGroup.GET("", cartHandler.Get)
		cartGroup.POST("/add", cartHandler.Add)
		cartGroup.PUT("/update/:item_id", cartHandler.Update)
		cartGroup.DELETE("/delete/:item_id", cartHandler.Delete)
	}

	// Healthcheck
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запускаем сервер
	log.Printf("Starting Gateway on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
