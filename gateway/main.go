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
	warehouseServiceURL := config.App.WarehouseServiceURL
	shippingServiceURL := config.App.ShippingServiceURL

	router := gin.Default()

	// Добавляем CORS middleware
	router.Use(middleware.CORSMiddleware())

	// Создаем репозитории
	authRepo := http.NewAuthRepositoryHTTP(authServiceURL)
	productRepo := http.NewProductRepositoryHTTP(productServiceURL)
	cartRepo := http.NewCartRepositoryHTTP(cartServiceURL)
	wmsRepo := http.NewWmsRepositoryHTTP(warehouseServiceURL)
	shippingRepo := http.NewShippingRepositoryHTTP(shippingServiceURL)

	// Создаем сервисы
	authService := services.NewAuthService(authRepo)
	productService := services.NewProductService(productRepo)
	cartService := services.NewCartService(cartRepo)
	wmsService := services.NewWmsService(wmsRepo)
	shippingService := services.NewShippingService(shippingRepo)

	// Создаем handlers
	authHandler := handlers.NewAuthHandler(authService)
	productHandler := handlers.NewProductHandler(productService)
	cartHandler := handlers.NewCartHandler(cartService)
	wmsHandler := handlers.NewWmsHandler(wmsService)
	shippingHandler := handlers.NewShippingHandler(shippingService)

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
		// /auth/staff - список сотрудников
		authGroup.GET("/staff", authHandler.GetStaff)
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

	// WMS routes (auth required)
	wmsGroup := router.Group("/wms")
	wmsGroup.Use(authMiddleware)
	{
		wmsGroup.POST("/check", wmsHandler.Check)
		wmsGroup.POST("/commit", wmsHandler.Commit)
	}

	// Warehouse routes (auth required)
	warehouseGroup := router.Group("/warehouse")
	warehouseGroup.Use(authMiddleware)
	{
		// WMS operations
		warehouseGroup.POST("/wms/check", wmsHandler.Check)
		warehouseGroup.POST("/wms/commit", wmsHandler.Commit)

		// Warehouse management
		warehouseGroup.POST("/warehouses", wmsHandler.CreateWarehouse)
		warehouseGroup.GET("/warehouses", wmsHandler.ListWarehouses)
		warehouseGroup.GET("/warehouses/:id", wmsHandler.GetWarehouse)
		warehouseGroup.DELETE("/warehouses/:id", wmsHandler.DeleteWarehouse)

		// Storage rules
		warehouseGroup.POST("/storage-rules", wmsHandler.CreateStorageRule)
		warehouseGroup.GET("/storage-rules", wmsHandler.ListStorageRules)
		warehouseGroup.DELETE("/storage-rules/:id", wmsHandler.DeleteStorageRule)

		// Bin locations
		warehouseGroup.POST("/locations/bins", wmsHandler.CreateBinLocation)
		warehouseGroup.GET("/locations/bins", wmsHandler.ListBinLocations)
		warehouseGroup.DELETE("/locations/bins/:id", wmsHandler.DeleteBinLocation)

		// Inventory
		warehouseGroup.POST("/inventory/add", wmsHandler.AddInventoryItem)
		warehouseGroup.POST("/inventory/move", wmsHandler.MoveInventory)
		warehouseGroup.GET("/inventory/items", wmsHandler.ListInventoryItems)
		warehouseGroup.GET("/inventory/product/:product_id/total", wmsHandler.GetProductTotalQuantity)
		warehouseGroup.GET("/inventory/totals", wmsHandler.GetAllProductsTotals)
	}

	// Shipping routes
	shippingPublic := router.Group("/shipping")
	{
		shippingPublic.GET("/:tracking_code", shippingHandler.Get)
	}
	shippingAuth := router.Group("/shipping")
	shippingAuth.Use(authMiddleware)
	{
		shippingAuth.GET("/list", shippingHandler.List)
		shippingAuth.POST("", shippingHandler.Create)
		shippingAuth.PATCH("/:tracking_code/status", shippingHandler.UpdateStatus)
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
