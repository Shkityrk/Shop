package services

import (
	"bytes"
	"encoding/json"
	"gateway/internal/domain/interfaces"
	"gateway/internal/domain/models"
	"net/http"
)

// WmsService handles WMS logic
type WmsService struct {
	repo interfaces.WmsRepository
}

// NewWmsService creates WMS service
func NewWmsService(repo interfaces.WmsRepository) *WmsService {
	return &WmsService{repo: repo}
}

// Check checks availability
func (s *WmsService) Check(payload models.WmsRequest, req *http.Request) (*http.Response, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return s.repo.Check(bytes.NewBuffer(jsonData), req)
}

// Commit commits reservation/stock deduction
func (s *WmsService) Commit(payload models.WmsRequest, req *http.Request) (*http.Response, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return s.repo.Commit(bytes.NewBuffer(jsonData), req)
}

// CreateWarehouse creates a new warehouse
func (s *WmsService) CreateWarehouse(payload models.WarehouseCreate, req *http.Request) (*http.Response, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return s.repo.CreateWarehouse(bytes.NewBuffer(jsonData), req)
}

// ListWarehouses lists all warehouses
func (s *WmsService) ListWarehouses(req *http.Request) (*http.Response, error) {
	return s.repo.ListWarehouses(req)
}

// GetWarehouse gets a warehouse by ID
func (s *WmsService) GetWarehouse(id string, req *http.Request) (*http.Response, error) {
	return s.repo.GetWarehouse(id, req)
}

// DeleteWarehouse deletes a warehouse by ID
func (s *WmsService) DeleteWarehouse(id string, req *http.Request) (*http.Response, error) {
	return s.repo.DeleteWarehouse(id, req)
}

// CreateStorageRule creates a new storage rule
func (s *WmsService) CreateStorageRule(payload models.StorageRuleCreate, req *http.Request) (*http.Response, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return s.repo.CreateStorageRule(bytes.NewBuffer(jsonData), req)
}

// ListStorageRules lists all storage rules
func (s *WmsService) ListStorageRules(req *http.Request) (*http.Response, error) {
	return s.repo.ListStorageRules(req)
}

// DeleteStorageRule deletes a storage rule by ID
func (s *WmsService) DeleteStorageRule(id string, req *http.Request) (*http.Response, error) {
	return s.repo.DeleteStorageRule(id, req)
}

// CreateBinLocation creates a new bin location
func (s *WmsService) CreateBinLocation(payload models.BinLocationCreate, req *http.Request) (*http.Response, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return s.repo.CreateBinLocation(bytes.NewBuffer(jsonData), req)
}

// ListBinLocations lists all bin locations
func (s *WmsService) ListBinLocations(queryParams string, req *http.Request) (*http.Response, error) {
	return s.repo.ListBinLocations(queryParams, req)
}

// DeleteBinLocation deletes a bin location by ID
func (s *WmsService) DeleteBinLocation(id string, req *http.Request) (*http.Response, error) {
	return s.repo.DeleteBinLocation(id, req)
}

// AddInventoryItem adds an inventory item
func (s *WmsService) AddInventoryItem(payload models.InventoryItemCreate, req *http.Request) (*http.Response, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return s.repo.AddInventoryItem(bytes.NewBuffer(jsonData), req)
}

// MoveInventory moves inventory between bins
func (s *WmsService) MoveInventory(payload models.InventoryMoveRequest, req *http.Request) (*http.Response, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return s.repo.MoveInventory(bytes.NewBuffer(jsonData), req)
}

// ListInventoryItems lists inventory items with optional filters
func (s *WmsService) ListInventoryItems(queryParams string, req *http.Request) (*http.Response, error) {
	return s.repo.ListInventoryItems(queryParams, req)
}

// GetProductTotalQuantity gets total quantity for a product
func (s *WmsService) GetProductTotalQuantity(productID string, req *http.Request) (*http.Response, error) {
	return s.repo.GetProductTotalQuantity(productID, req)
}

// GetAllProductsTotals gets totals for all products
func (s *WmsService) GetAllProductsTotals(req *http.Request) (*http.Response, error) {
	return s.repo.GetAllProductsTotals(req)
}

