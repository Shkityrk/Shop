package interfaces

import (
	"io"
	"net/http"
)

// WmsRepository defines interface for WMS operations
type WmsRepository interface {
	Check(body io.Reader, req *http.Request) (*http.Response, error)
	Commit(body io.Reader, req *http.Request) (*http.Response, error)

	// Warehouse operations
	CreateWarehouse(body io.Reader, req *http.Request) (*http.Response, error)
	ListWarehouses(req *http.Request) (*http.Response, error)
	GetWarehouse(id string, req *http.Request) (*http.Response, error)
	DeleteWarehouse(id string, req *http.Request) (*http.Response, error)

	// Storage rules operations
	CreateStorageRule(body io.Reader, req *http.Request) (*http.Response, error)
	ListStorageRules(req *http.Request) (*http.Response, error)
	DeleteStorageRule(id string, req *http.Request) (*http.Response, error)

	// Bin location operations
	CreateBinLocation(body io.Reader, req *http.Request) (*http.Response, error)
	ListBinLocations(queryParams string, req *http.Request) (*http.Response, error)
	DeleteBinLocation(id string, req *http.Request) (*http.Response, error)

	// Inventory operations
	AddInventoryItem(body io.Reader, req *http.Request) (*http.Response, error)
	MoveInventory(body io.Reader, req *http.Request) (*http.Response, error)
	ListInventoryItems(queryParams string, req *http.Request) (*http.Response, error)
	GetProductTotalQuantity(productID string, req *http.Request) (*http.Response, error)
	GetAllProductsTotals(req *http.Request) (*http.Response, error)
}
