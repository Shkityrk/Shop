package handler

import (
	"data-service/internal/application/dto"
	"data-service/internal/application/service"
	"data-service/internal/domain/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ProductHandler обработчик HTTP запросов для продуктов
type ProductHandler struct {
	service *service.ProductService
	logger  *logrus.Logger
}

// NewProductHandler создает новый экземпляр ProductHandler
func NewProductHandler(service *service.ProductService, logger *logrus.Logger) *ProductHandler {
	return &ProductHandler{
		service: service,
		logger:  logger,
	}
}

// CreateProduct создает новый продукт
// @Summary Create product
// @Description Create a new product
// @Tags products
// @Accept json
// @Produce json
// @Param product body dto.CreateProductRequest true "Product data"
// @Success 201 {object} dto.ProductResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Failed to bind request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.service.CreateProduct(c.Request.Context(), &req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to create product")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// GetProductByID получает продукт по ID
// @Summary Get product by ID
// @Description Get product by ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} dto.ProductResponse
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/products/{id} [get]
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.service.GetProductByID(c.Request.Context(), id)
	if err != nil {
		h.logger.WithError(err).Error("Failed to get product by ID")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetProductByName получает продукт по названию
// @Summary Get product by name
// @Description Get product by name
// @Tags products
// @Produce json
// @Param name path string true "Product name"
// @Success 200 {object} dto.ProductResponse
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/products/name/{name} [get]
func (h *ProductHandler) GetProductByName(c *gin.Context) {
	name := c.Param("name")

	product, err := h.service.GetProductByName(c.Request.Context(), name)
	if err != nil {
		h.logger.WithError(err).Error("Failed to get product by name")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// ListProducts получает список всех продуктов
// @Summary List products
// @Description Get list of all products
// @Tags products
// @Produce json
// @Param name query string false "Filter by name"
// @Success 200 {array} dto.ProductResponse
// @Failure 500 {object} map[string]interface{}
// @Router /api/products [get]
func (h *ProductHandler) ListProducts(c *gin.Context) {
	filter := &entity.ProductFilter{}

	if name := c.Query("name"); name != "" {
		filter.Name = &name
	}

	products, err := h.service.ListProducts(c.Request.Context(), filter)
	if err != nil {
		h.logger.WithError(err).Error("Failed to list products")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// UpdateProduct обновляет данные продукта
// @Summary Update product
// @Description Update product data
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body dto.UpdateProductRequest true "Product data"
// @Success 200 {object} dto.ProductResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var req dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Failed to bind request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.service.UpdateProduct(c.Request.Context(), id, &req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to update product")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct удаляет продукт
// @Summary Delete product
// @Description Delete product by ID
// @Tags products
// @Param id path int true "Product ID"
// @Success 204
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := h.service.DeleteProduct(c.Request.Context(), id); err != nil {
		h.logger.WithError(err).Error("Failed to delete product")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// CheckProductExists проверяет существование продукта
// @Summary Check product exists
// @Description Check if product exists by ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} dto.CheckProductExistsResponse
// @Failure 500 {object} map[string]interface{}
// @Router /api/products/{id}/exists [get]
func (h *ProductHandler) CheckProductExists(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	exists, err := h.service.CheckProductExists(c.Request.Context(), id)
	if err != nil {
		h.logger.WithError(err).Error("Failed to check product existence")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.CheckProductExistsResponse{Exists: exists})
}

