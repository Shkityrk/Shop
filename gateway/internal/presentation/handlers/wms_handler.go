package handlers

import (
	"gateway/internal/application/services"
	"gateway/internal/domain/models"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// WmsHandler handles WMS related requests
type WmsHandler struct {
	service *services.WmsService
}

// NewWmsHandler creates WMS handler
func NewWmsHandler(service *services.WmsService) *WmsHandler {
	return &WmsHandler{service: service}
}

// Check availability
// @Summary Check product availability
// @Description Check if products are available in the warehouse
// @Tags WMS
// @Accept json
// @Produce json
// @Param request body models.WmsRequest true "WMS check request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /warehouse/wms/check [post]
func (h *WmsHandler) Check(c *gin.Context) {
	var payload models.WmsRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := h.service.Check(payload, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with warehouse service"})
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

// Commit order items
// @Summary Commit order items
// @Description Commit products from the warehouse
// @Tags WMS
// @Accept json
// @Produce json
// @Param request body models.WmsRequest true "WMS commit request"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /warehouse/wms/commit [post]
func (h *WmsHandler) Commit(c *gin.Context) {
	var payload models.WmsRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := h.service.Commit(payload, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with warehouse service"})
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

// CreateWarehouse creates a new warehouse
// @Summary Create a new warehouse
// @Description Create a new warehouse
// @Tags Warehouse
// @Accept json
// @Produce json
// @Param warehouse body models.WarehouseCreate true "Warehouse data"
// @Success 200 {object} models.Warehouse
// @Failure 400 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /warehouse/warehouses [post]
func (h *WmsHandler) CreateWarehouse(c *gin.Context) {
	var payload models.WarehouseCreate
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := h.service.CreateWarehouse(payload, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with warehouse service"})
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

// ListWarehouses lists all warehouses
// @Summary List all warehouses
// @Description Get a list of all warehouses
// @Tags Warehouse
// @Produce json
// @Success 200 {array} models.Warehouse
// @Security BearerAuth
// @Router /warehouse/warehouses [get]
func (h *WmsHandler) ListWarehouses(c *gin.Context) {
	resp, err := h.service.ListWarehouses(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with warehouse service"})
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

// GetWarehouse gets a warehouse by ID
// @Summary Get warehouse by ID
// @Description Get a warehouse by its ID
// @Tags Warehouse
// @Produce json
// @Param id path int true "Warehouse ID"
// @Success 200 {object} models.Warehouse
// @Failure 404 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /warehouse/warehouses/{id} [get]
func (h *WmsHandler) GetWarehouse(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.service.GetWarehouse(id, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with warehouse service"})
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

// DeleteWarehouse deletes a warehouse by ID
// @Summary Delete warehouse
// @Description Delete a warehouse by its ID
// @Tags Warehouse
// @Produce json
// @Param id path int true "Warehouse ID"
// @Success 200 {object} models.StatusResponse
// @Failure 404 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /warehouse/warehouses/{id} [delete]
func (h *WmsHandler) DeleteWarehouse(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.service.DeleteWarehouse(id, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with warehouse service"})
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

// CreateStorageRule creates a new storage rule
// @Summary Create a new storage rule
// @Description Create a new storage rule
// @Tags Storage Rules
// @Accept json
// @Produce json
// @Param storage_rule body models.StorageRuleCreate true "Storage rule data"
// @Success 200 {object} models.StorageRule
// @Failure 400 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /warehouse/storage-rules [post]
func (h *WmsHandler) CreateStorageRule(c *gin.Context) {
	var payload models.StorageRuleCreate
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := h.service.CreateStorageRule(payload, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with warehouse service"})
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

// ListStorageRules lists all storage rules
// @Summary List all storage rules
// @Description Get a list of all storage rules
// @Tags Storage Rules
// @Produce json
// @Success 200 {array} models.StorageRule
// @Security BearerAuth
// @Router /warehouse/storage-rules [get]
func (h *WmsHandler) ListStorageRules(c *gin.Context) {
	resp, err := h.service.ListStorageRules(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with warehouse service"})
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

// DeleteStorageRule deletes a storage rule by ID
// @Summary Delete storage rule
// @Description Delete a storage rule by its ID
// @Tags Storage Rules
// @Produce json
// @Param id path int true "Storage Rule ID"
// @Success 200 {object} models.StatusResponse
// @Failure 404 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /warehouse/storage-rules/{id} [delete]
func (h *WmsHandler) DeleteStorageRule(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.service.DeleteStorageRule(id, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with warehouse service"})
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

// CreateBinLocation creates a new bin location
// @Summary Create a new bin location
// @Description Create a new bin location in a warehouse
// @Tags Bin Locations
// @Accept json
// @Produce json
// @Param bin_location body models.BinLocationCreate true "Bin location data"
// @Success 200 {object} models.BinLocation
// @Failure 400 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /warehouse/locations/bins [post]
func (h *WmsHandler) CreateBinLocation(c *gin.Context) {
	var payload models.BinLocationCreate
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := h.service.CreateBinLocation(payload, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with warehouse service"})
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

// ListBinLocations lists all bin locations
// @Summary List all bin locations
// @Description Get a list of all bin locations, optionally filtered by warehouse
// @Tags Bin Locations
// @Produce json
// @Param warehouse_id query int false "Filter by warehouse ID"
// @Success 200 {array} models.BinLocation
// @Security BearerAuth
// @Router /warehouse/locations/bins [get]
func (h *WmsHandler) ListBinLocations(c *gin.Context) {
	resp, err := h.service.ListBinLocations(c.Request.URL.RawQuery, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with warehouse service"})
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

// DeleteBinLocation deletes a bin location by ID
// @Summary Delete bin location
// @Description Delete a bin location by its ID
// @Tags Bin Locations
// @Produce json
// @Param id path int true "Bin Location ID"
// @Success 200 {object} models.StatusResponse
// @Failure 404 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /warehouse/locations/bins/{id} [delete]
func (h *WmsHandler) DeleteBinLocation(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.service.DeleteBinLocation(id, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with warehouse service"})
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

// AddInventoryItem adds an inventory item
// @Summary Add inventory item
// @Description Add a new inventory item to a warehouse bin
// @Tags Inventory
// @Accept json
// @Produce json
// @Param inventory_item body models.InventoryItemCreate true "Inventory item data"
// @Success 200 {object} models.InventoryItem
// @Failure 400 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /warehouse/inventory/add [post]
func (h *WmsHandler) AddInventoryItem(c *gin.Context) {
	var payload models.InventoryItemCreate
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := h.service.AddInventoryItem(payload, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with warehouse service"})
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

// MoveInventory moves inventory between bins
// @Summary Move inventory
// @Description Move inventory from one bin to another
// @Tags Inventory
// @Accept json
// @Produce json
// @Param move_request body models.InventoryMoveRequest true "Inventory move request"
// @Success 200 {object} models.StatusResponse
// @Failure 400 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /warehouse/inventory/move [post]
func (h *WmsHandler) MoveInventory(c *gin.Context) {
	var payload models.InventoryMoveRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := h.service.MoveInventory(payload, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with warehouse service"})
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

// ListInventoryItems lists inventory items
// @Summary List inventory items
// @Description Get a list of inventory items, optionally filtered by product or warehouse
// @Tags Inventory
// @Produce json
// @Param product_id query int false "Filter by product ID"
// @Param warehouse_id query int false "Filter by warehouse ID"
// @Success 200 {array} models.InventoryItem
// @Security BearerAuth
// @Router /warehouse/inventory/items [get]
func (h *WmsHandler) ListInventoryItems(c *gin.Context) {
	resp, err := h.service.ListInventoryItems(c.Request.URL.RawQuery, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with warehouse service"})
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

// GetProductTotalQuantity gets total quantity for a product
// @Summary Get product total quantity
// @Description Get the total quantity of a product across all warehouses
// @Tags Inventory
// @Produce json
// @Param product_id path int true "Product ID"
// @Success 200 {object} models.ProductTotalQuantity
// @Security BearerAuth
// @Router /warehouse/inventory/product/{product_id}/total [get]
func (h *WmsHandler) GetProductTotalQuantity(c *gin.Context) {
	productID := c.Param("product_id")
	resp, err := h.service.GetProductTotalQuantity(productID, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with warehouse service"})
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

// GetAllProductsTotals gets totals for all products
// @Summary Get all products totals
// @Description Get the total quantities of all products across all warehouses
// @Tags Inventory
// @Produce json
// @Success 200 {array} models.ProductTotalQuantity
// @Security BearerAuth
// @Router /warehouse/inventory/totals [get]
func (h *WmsHandler) GetAllProductsTotals(c *gin.Context) {
	resp, err := h.service.GetAllProductsTotals(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with warehouse service"})
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

