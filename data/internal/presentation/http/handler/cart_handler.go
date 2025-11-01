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

// CartHandler обработчик HTTP запросов для корзины
type CartHandler struct {
	service *service.CartService
	logger  *logrus.Logger
}

// NewCartHandler создает новый экземпляр CartHandler
func NewCartHandler(service *service.CartService, logger *logrus.Logger) *CartHandler {
	return &CartHandler{
		service: service,
		logger:  logger,
	}
}

// CreateCartItem создает новый элемент корзины
// @Summary Create cart item
// @Description Create a new cart item
// @Tags cart
// @Accept json
// @Produce json
// @Param item body dto.CreateCartItemRequest true "Cart item data"
// @Success 201 {object} dto.CartItemResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/cart [post]
func (h *CartHandler) CreateCartItem(c *gin.Context) {
	var req dto.CreateCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Failed to bind request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.service.CreateCartItem(c.Request.Context(), &req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to create cart item")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

// GetCartItemByID получает элемент корзины по ID
// @Summary Get cart item by ID
// @Description Get cart item by ID
// @Tags cart
// @Produce json
// @Param id path int true "Cart item ID"
// @Success 200 {object} dto.CartItemResponse
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/cart/{id} [get]
func (h *CartHandler) GetCartItemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart item ID"})
		return
	}

	item, err := h.service.GetCartItemByID(c.Request.Context(), id)
	if err != nil {
		h.logger.WithError(err).Error("Failed to get cart item by ID")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// GetCartItemsByUserID получает все элементы корзины пользователя
// @Summary Get cart items by user ID
// @Description Get all cart items for a user
// @Tags cart
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} dto.CartItemResponse
// @Failure 500 {object} map[string]interface{}
// @Router /api/cart/user/{user_id} [get]
func (h *CartHandler) GetCartItemsByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	items, err := h.service.GetCartItemsByUserID(c.Request.Context(), userID)
	if err != nil {
		h.logger.WithError(err).Error("Failed to get cart items by user ID")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

// GetCartItemByUserAndProduct получает элемент корзины по user_id и product_id
// @Summary Get cart item by user and product
// @Description Get cart item by user ID and product ID
// @Tags cart
// @Produce json
// @Param user_id path int true "User ID"
// @Param product_id path int true "Product ID"
// @Success 200 {object} dto.CartItemResponse
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/cart/user/{user_id}/product/{product_id} [get]
func (h *CartHandler) GetCartItemByUserAndProduct(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	item, err := h.service.GetCartItemByUserAndProduct(c.Request.Context(), userID, productID)
	if err != nil {
		h.logger.WithError(err).Error("Failed to get cart item by user and product")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// ListCartItems получает список элементов корзины с фильтрацией
// @Summary List cart items
// @Description Get list of cart items with optional filtering
// @Tags cart
// @Produce json
// @Param user_id query int false "Filter by user ID"
// @Param product_id query int false "Filter by product ID"
// @Success 200 {array} dto.CartItemResponse
// @Failure 500 {object} map[string]interface{}
// @Router /api/cart [get]
func (h *CartHandler) ListCartItems(c *gin.Context) {
	filter := &entity.CartFilter{}

	if userIDStr := c.Query("user_id"); userIDStr != "" {
		userID, err := strconv.Atoi(userIDStr)
		if err == nil {
			filter.UserID = &userID
		}
	}

	if productIDStr := c.Query("product_id"); productIDStr != "" {
		productID, err := strconv.Atoi(productIDStr)
		if err == nil {
			filter.ProductID = &productID
		}
	}

	items, err := h.service.ListCartItems(c.Request.Context(), filter)
	if err != nil {
		h.logger.WithError(err).Error("Failed to list cart items")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

// UpdateCartItem обновляет элемент корзины
// @Summary Update cart item
// @Description Update cart item data
// @Tags cart
// @Accept json
// @Produce json
// @Param id path int true "Cart item ID"
// @Param item body dto.UpdateCartItemRequest true "Cart item data"
// @Success 200 {object} dto.CartItemResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/cart/{id} [put]
func (h *CartHandler) UpdateCartItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart item ID"})
		return
	}

	var req dto.UpdateCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Failed to bind request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.service.UpdateCartItem(c.Request.Context(), id, &req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to update cart item")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// DeleteCartItem удаляет элемент корзины
// @Summary Delete cart item
// @Description Delete cart item by ID
// @Tags cart
// @Param id path int true "Cart item ID"
// @Success 204
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/cart/{id} [delete]
func (h *CartHandler) DeleteCartItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart item ID"})
		return
	}

	if err := h.service.DeleteCartItem(c.Request.Context(), id); err != nil {
		h.logger.WithError(err).Error("Failed to delete cart item")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteCartItemsByUserID удаляет все элементы корзины пользователя
// @Summary Delete cart items by user ID
// @Description Delete all cart items for a user
// @Tags cart
// @Param user_id path int true "User ID"
// @Success 204
// @Failure 500 {object} map[string]interface{}
// @Router /api/cart/user/{user_id} [delete]
func (h *CartHandler) DeleteCartItemsByUserID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.service.DeleteCartItemsByUserID(c.Request.Context(), userID); err != nil {
		h.logger.WithError(err).Error("Failed to delete cart items by user ID")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// DeleteCartItemByUserAndProduct удаляет элемент корзины по user_id и product_id
// @Summary Delete cart item by user and product
// @Description Delete cart item by user ID and product ID
// @Tags cart
// @Param user_id path int true "User ID"
// @Param product_id path int true "Product ID"
// @Success 204
// @Failure 500 {object} map[string]interface{}
// @Router /api/cart/user/{user_id}/product/{product_id} [delete]
func (h *CartHandler) DeleteCartItemByUserAndProduct(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := h.service.DeleteCartItemByUserAndProduct(c.Request.Context(), userID, productID); err != nil {
		h.logger.WithError(err).Error("Failed to delete cart item by user and product")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

