package router

import (
	"data-service/internal/presentation/http/handler"
	"data-service/internal/presentation/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Router представляет HTTP роутер
type Router struct {
	engine         *gin.Engine
	userHandler    *handler.UserHandler
	productHandler *handler.ProductHandler
	cartHandler    *handler.CartHandler
	healthHandler  *handler.HealthHandler
	logger         *logrus.Logger
}

// NewRouter создает новый роутер
func NewRouter(
	userHandler *handler.UserHandler,
	productHandler *handler.ProductHandler,
	cartHandler *handler.CartHandler,
	healthHandler *handler.HealthHandler,
	logger *logrus.Logger,
) *Router {
	return &Router{
		engine:         gin.Default(),
		userHandler:    userHandler,
		productHandler: productHandler,
		cartHandler:    cartHandler,
		healthHandler:  healthHandler,
		logger:         logger,
	}
}

// Setup настраивает маршруты
func (r *Router) Setup() {
	// Middleware
	r.engine.Use(middleware.Logger(r.logger))
	r.engine.Use(middleware.Recovery(r.logger))
	r.engine.Use(middleware.CORS())

	// Health check
	r.engine.GET("/health", r.healthHandler.HealthCheck)

	// Swagger
	r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	api := r.engine.Group("/api")
	{
		// User routes
		users := api.Group("/users")
		{
			users.POST("", r.userHandler.CreateUser)
			users.GET("", r.userHandler.ListUsers)
			users.GET("/:id", r.userHandler.GetUserByID)
			users.GET("/username/:username", r.userHandler.GetUserByUsername)
			users.GET("/email/:email", r.userHandler.GetUserByEmail)
			users.PUT("/:id", r.userHandler.UpdateUser)
			users.DELETE("/:id", r.userHandler.DeleteUser)
			users.POST("/exists", r.userHandler.CheckUserExists)
		}

		// Product routes
		products := api.Group("/products")
		{
			products.POST("", r.productHandler.CreateProduct)
			products.GET("", r.productHandler.ListProducts)
			products.GET("/:id", r.productHandler.GetProductByID)
			products.GET("/name/:name", r.productHandler.GetProductByName)
			products.PUT("/:id", r.productHandler.UpdateProduct)
			products.DELETE("/:id", r.productHandler.DeleteProduct)
			products.GET("/:id/exists", r.productHandler.CheckProductExists)
		}

		// Cart routes
		cart := api.Group("/cart")
		{
			cart.POST("", r.cartHandler.CreateCartItem)
			cart.GET("", r.cartHandler.ListCartItems)
			cart.GET("/:id", r.cartHandler.GetCartItemByID)
			cart.GET("/user/:user_id", r.cartHandler.GetCartItemsByUserID)
			cart.GET("/user/:user_id/product/:product_id", r.cartHandler.GetCartItemByUserAndProduct)
			cart.PUT("/:id", r.cartHandler.UpdateCartItem)
			cart.DELETE("/:id", r.cartHandler.DeleteCartItem)
			cart.DELETE("/user/:user_id", r.cartHandler.DeleteCartItemsByUserID)
			cart.DELETE("/user/:user_id/product/:product_id", r.cartHandler.DeleteCartItemByUserAndProduct)
		}
	}
}

// GetEngine возвращает Gin engine
func (r *Router) GetEngine() *gin.Engine {
	return r.engine
}
