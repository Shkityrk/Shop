package interfaces

import (
	"io"
	"net/http"
)

// ShippingRepository defines interface for shipping operations
type ShippingRepository interface {
	Create(body io.Reader, req *http.Request) (*http.Response, error)
	Get(tracking string, req *http.Request) (*http.Response, error)
	UpdateStatus(tracking string, body io.Reader, req *http.Request) (*http.Response, error)
	List(queryParams string, req *http.Request) (*http.Response, error)
}
