package handlers

import (
	"gateway/internal/application/services"
	"gateway/internal/domain/models"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShippingHandler handles shipping related requests
type ShippingHandler struct {
	service *services.ShippingService
}

// NewShippingHandler creates shipping handler
func NewShippingHandler(service *services.ShippingService) *ShippingHandler {
	return &ShippingHandler{service: service}
}

// Create shipment
func (h *ShippingHandler) Create(c *gin.Context) {
	var payload models.ShipmentCreate
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := h.service.Create(payload, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with shipping service"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to read response"})
		return
	}

	c.Data(resp.StatusCode, "application/json", body)
}

// Get shipment by tracking (public)
func (h *ShippingHandler) Get(c *gin.Context) {
	tracking := c.Param("tracking_code")
	resp, err := h.service.Get(tracking, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with shipping service"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to read response"})
		return
	}

	c.Data(resp.StatusCode, "application/json", body)
}

// UpdateStatus updates shipment status (requires auth)
func (h *ShippingHandler) UpdateStatus(c *gin.Context) {
	tracking := c.Param("tracking_code")
	var payload models.ShipmentStatusUpdate
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := h.service.UpdateStatus(tracking, payload, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with shipping service"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to read response"})
		return
	}

	c.Data(resp.StatusCode, "application/json", body)
}

// List godoc
// @Summary List all shipments
// @Description Get a list of all shipments, optionally filtered by status
// @Tags Shipping
// @Produce json
// @Param skip query int false "Number of shipments to skip" default(0)
// @Param limit query int false "Maximum number of shipments to return" default(100)
// @Param status query string false "Filter by status"
// @Success 200 {array} models.Shipment
// @Security BearerAuth
// @Router /shipping/list [get]
func (h *ShippingHandler) List(c *gin.Context) {
	resp, err := h.service.List(c.Request.URL.RawQuery, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with shipping service"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to read response"})
		return
	}

	c.Data(resp.StatusCode, "application/json", body)
}

