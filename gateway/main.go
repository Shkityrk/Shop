package main

import (
	"gateway/config"
	"gateway/handlers"
	"gateway/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"

	_ "gateway/docs"
)

func main() {
	// Получаем порт из переменной окружения или используем значение по умолчанию
	port := config.App.Gateway_port

	// Получаем URL'ы сервисов из переменных окружения
	authServiceURL := config.App.Auth_service_url
	productServiceURL := config.App.Product_service_url
	cartServiceURL := config.App.Cart_service_url

	router := gin.Default()

	// Добавляем CORS middleware
	router.Use(middleware.CORSMiddleware())

	// Создаем handlers
	authHandler := handlers.NewAuthHandler(authServiceURL)
	productHandler := handlers.NewProductHandler(productServiceURL)
	cartHandler := handlers.NewCartHandler(cartServiceURL)

	// Auth routes
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/logout", authHandler.Logout)
		authGroup.GET("/info", authHandler.Info)
	}

	// Product routes
	productGroup := router.Group("/product")
	{
		productGroup.GET("/list", productHandler.List)
		productGroup.POST("/add", productHandler.Add)
		productGroup.PUT("/update/:id", productHandler.Update)
		productGroup.GET("/verify/:name", productHandler.Verify)
		productGroup.GET("/info/:id", productHandler.Info)
	}

	// Cart routes
	cartGroup := router.Group("/cart")
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
