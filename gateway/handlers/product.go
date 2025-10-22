package handlers

import (
	"bytes"
	"encoding/json"
	"gateway/models"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProductHandler handles product related requests
type ProductHandler struct {
	serviceURL string
	client     *http.Client
}

// NewProductHandler creates a new product handler
func NewProductHandler(serviceURL string) *ProductHandler {
	return &ProductHandler{
		serviceURL: serviceURL,
		client:     &http.Client{},
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
	req, err := http.NewRequest("GET", h.serviceURL+"/product/list", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create request"})
		return
	}

	// Copy query parameters
	req.URL.RawQuery = c.Request.URL.RawQuery
	copyHeaders(c.Request, req)

	resp, err := h.client.Do(req)
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

	jsonData, err := json.Marshal(productCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to marshal request"})
		return
	}

	req, err := http.NewRequest("POST", h.serviceURL+"/product/add", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create request"})
		return
	}

	req.Header.Set("Content-Type", "application/json")
	copyHeaders(c.Request, req)

	resp, err := h.client.Do(req)
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

	jsonData, err := json.Marshal(productCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to marshal request"})
		return
	}

	req, err := http.NewRequest("PUT", h.serviceURL+"/product/update/"+id, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create request"})
		return
	}

	req.Header.Set("Content-Type", "application/json")
	copyHeaders(c.Request, req)

	resp, err := h.client.Do(req)
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
	req, err := http.NewRequest("GET", h.serviceURL+"/product/verify/"+name, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create request"})
		return
	}

	copyHeaders(c.Request, req)

	resp, err := h.client.Do(req)
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
	req, err := http.NewRequest("GET", h.serviceURL+"/product/info/"+id, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create request"})
		return
	}

	copyHeaders(c.Request, req)

	resp, err := h.client.Do(req)
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

