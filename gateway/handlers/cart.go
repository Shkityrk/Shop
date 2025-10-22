package handlers

import (
	"bytes"
	"encoding/json"
	"gateway/models"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CartHandler handles cart related requests
type CartHandler struct {
	serviceURL string
	client     *http.Client
}

// NewCartHandler creates a new cart handler
func NewCartHandler(serviceURL string) *CartHandler {
	return &CartHandler{
		serviceURL: serviceURL,
		client:     &http.Client{},
	}
}

// Get godoc
// @Summary Get cart items
// @Description Get all items in the current user's cart
// @Tags Cart
// @Produce json
// @Success 200 {array} models.CartItem
// @Failure 401 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /cart [get]
func (h *CartHandler) Get(c *gin.Context) {
	req, err := http.NewRequest("GET", h.serviceURL+"/cart", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create request"})
		return
	}

	copyHeaders(c.Request, req)

	resp, err := h.client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with cart service"})
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
// @Summary Add item to cart
// @Description Add a new item to the current user's cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param item body models.CartItemCreate true "Cart item data"
// @Success 201 {object} models.CartItem
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /cart/add [post]
func (h *CartHandler) Add(c *gin.Context) {
	var cartItemCreate models.CartItemCreate
	if err := c.ShouldBindJSON(&cartItemCreate); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, err := json.Marshal(cartItemCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to marshal request"})
		return
	}

	req, err := http.NewRequest("POST", h.serviceURL+"/cart/add", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create request"})
		return
	}

	req.Header.Set("Content-Type", "application/json")
	copyHeaders(c.Request, req)

	resp, err := h.client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with cart service"})
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
// @Summary Update cart item
// @Description Update quantity of an item in the cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param item_id path int true "Cart item ID"
// @Param item body models.CartItemCreate true "Updated cart item data"
// @Success 200 {object} models.CartItem
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /cart/update/{item_id} [put]
func (h *CartHandler) Update(c *gin.Context) {
	itemID := c.Param("item_id")
	var cartItemCreate models.CartItemCreate
	if err := c.ShouldBindJSON(&cartItemCreate); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, err := json.Marshal(cartItemCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to marshal request"})
		return
	}

	req, err := http.NewRequest("PUT", h.serviceURL+"/cart/update/"+itemID, bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create request"})
		return
	}

	req.Header.Set("Content-Type", "application/json")
	copyHeaders(c.Request, req)

	resp, err := h.client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with cart service"})
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

// Delete godoc
// @Summary Delete cart item
// @Description Remove an item from the cart
// @Tags Cart
// @Param item_id path int true "Cart item ID"
// @Success 204 "No Content"
// @Failure 401 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /cart/delete/{item_id} [delete]
func (h *CartHandler) Delete(c *gin.Context) {
	itemID := c.Param("item_id")
	req, err := http.NewRequest("DELETE", h.serviceURL+"/cart/delete/"+itemID, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create request"})
		return
	}

	copyHeaders(c.Request, req)

	resp, err := h.client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with cart service"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		c.Status(http.StatusNoContent)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to read response"})
		return
	}

	c.Data(resp.StatusCode, "application/json", body)
}

