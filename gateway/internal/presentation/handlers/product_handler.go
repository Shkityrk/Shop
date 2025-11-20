package handlers

import (
	"gateway/internal/application/services"
	"gateway/internal/domain/models"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProductHandler handles product related requests
type ProductHandler struct {
	service *services.ProductService
}

// NewProductHandler creates a new product handler
func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

// List godoc
// @Summary List all products
// @Description Get a list of all available products
// @Tags Product
// @Produce json
// @Param skip query int false "Number of products to skip" default(0)
// @Param limit query int false "Maximum number of products to return" default(100)
// @Success 200 {array} models.Product
// @Router /product/list [get]
func (h *ProductHandler) List(c *gin.Context) {
	resp, err := h.service.List(c.Request.URL.RawQuery, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with product service"})
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

// Add godoc
// @Summary Add a new product
// @Description Create a new product
// @Tags Product
// @Accept json
// @Produce json
// @Param product body models.ProductCreate true "Product data"
// @Success 200 {object} models.Product
// @Failure 400 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /product/add [post]
func (h *ProductHandler) Add(c *gin.Context) {
	var productCreate models.ProductCreate
	if err := c.ShouldBindJSON(&productCreate); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := h.service.Add(productCreate, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with product service"})
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

// Update godoc
// @Summary Update a product
// @Description Update an existing product by ID
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.ProductCreate true "Updated product data"
// @Success 200 {object} models.Product
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /product/update/{id} [put]
func (h *ProductHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var productCreate models.ProductCreate
	if err := c.ShouldBindJSON(&productCreate); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := h.service.Update(id, productCreate, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with product service"})
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

// Verify godoc
// @Summary Verify product exists
// @Description Check if a product with the given name exists
// @Tags Product
// @Produce json
// @Param name path string true "Product name"
// @Success 200 {object} models.VerifyResponse
// @Router /product/verify/{name} [get]
func (h *ProductHandler) Verify(c *gin.Context) {
	name := c.Param("name")
	resp, err := h.service.Verify(name, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with product service"})
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

// Info godoc
// @Summary Get product info
// @Description Get detailed information about a product by ID
// @Tags Product
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Failure 404 {object} models.ErrorResponse
// @Router /product/info/{id} [get]
func (h *ProductHandler) Info(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.service.Info(id, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with product service"})
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

