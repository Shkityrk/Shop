package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler обработчик для проверки здоровья сервиса
type HealthHandler struct{}

// NewHealthHandler создает новый экземпляр HealthHandler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck проверяет здоровье сервиса
// @Summary Health check
// @Description Check if service is healthy
// @Tags health
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "data-service",
	})
}

