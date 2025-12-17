package http

import (
	"gateway/internal/domain/interfaces"
	"io"
	"net/http"
)

// WmsRepositoryHTTP implements WmsRepository using HTTP client
type WmsRepositoryHTTP struct {
	serviceURL string
	client     *http.Client
}

// NewWmsRepositoryHTTP creates a new HTTP wms repository
func NewWmsRepositoryHTTP(serviceURL string) interfaces.WmsRepository {
	return &WmsRepositoryHTTP{
		serviceURL: serviceURL,
		client:     &http.Client{},
	}
}

// Check sends availability check
func (r *WmsRepositoryHTTP) Check(body io.Reader, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("POST", r.serviceURL+"/warehouse/wms/check", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// Commit sends commit request
func (r *WmsRepositoryHTTP) Commit(body io.Reader, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("POST", r.serviceURL+"/warehouse/wms/commit", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// CreateWarehouse creates a new warehouse
func (r *WmsRepositoryHTTP) CreateWarehouse(body io.Reader, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("POST", r.serviceURL+"/warehouse/warehouses/", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// ListWarehouses lists all warehouses
func (r *WmsRepositoryHTTP) ListWarehouses(originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("GET", r.serviceURL+"/warehouse/warehouses/", nil)
	if err != nil {
		return nil, err
	}
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// GetWarehouse gets a warehouse by ID
func (r *WmsRepositoryHTTP) GetWarehouse(id string, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("GET", r.serviceURL+"/warehouse/warehouses/"+id, nil)
	if err != nil {
		return nil, err
	}
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// DeleteWarehouse deletes a warehouse by ID
func (r *WmsRepositoryHTTP) DeleteWarehouse(id string, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", r.serviceURL+"/warehouse/warehouses/"+id, nil)
	if err != nil {
		return nil, err
	}
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// CreateStorageRule creates a new storage rule
func (r *WmsRepositoryHTTP) CreateStorageRule(body io.Reader, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("POST", r.serviceURL+"/warehouse/storage-rules/", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// ListStorageRules lists all storage rules
func (r *WmsRepositoryHTTP) ListStorageRules(originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("GET", r.serviceURL+"/warehouse/storage-rules/", nil)
	if err != nil {
		return nil, err
	}
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// DeleteStorageRule deletes a storage rule by ID
func (r *WmsRepositoryHTTP) DeleteStorageRule(id string, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", r.serviceURL+"/warehouse/storage-rules/"+id, nil)
	if err != nil {
		return nil, err
	}
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// CreateBinLocation creates a new bin location
func (r *WmsRepositoryHTTP) CreateBinLocation(body io.Reader, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("POST", r.serviceURL+"/warehouse/locations/bins", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// ListBinLocations lists all bin locations
func (r *WmsRepositoryHTTP) ListBinLocations(queryParams string, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("GET", r.serviceURL+"/warehouse/locations/bins", nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = queryParams
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// DeleteBinLocation deletes a bin location by ID
func (r *WmsRepositoryHTTP) DeleteBinLocation(id string, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", r.serviceURL+"/warehouse/locations/bins/"+id, nil)
	if err != nil {
		return nil, err
	}
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// AddInventoryItem adds an inventory item
func (r *WmsRepositoryHTTP) AddInventoryItem(body io.Reader, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("POST", r.serviceURL+"/warehouse/inventory/add", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// MoveInventory moves inventory between bins
func (r *WmsRepositoryHTTP) MoveInventory(body io.Reader, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("POST", r.serviceURL+"/warehouse/inventory/move", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// ListInventoryItems lists inventory items with optional filters
func (r *WmsRepositoryHTTP) ListInventoryItems(queryParams string, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("GET", r.serviceURL+"/warehouse/inventory/items", nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = queryParams
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// GetProductTotalQuantity gets total quantity for a product
func (r *WmsRepositoryHTTP) GetProductTotalQuantity(productID string, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("GET", r.serviceURL+"/warehouse/inventory/product/"+productID+"/total", nil)
	if err != nil {
		return nil, err
	}
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// GetAllProductsTotals gets totals for all products
func (r *WmsRepositoryHTTP) GetAllProductsTotals(originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("GET", r.serviceURL+"/warehouse/inventory/totals", nil)
	if err != nil {
		return nil, err
	}
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

