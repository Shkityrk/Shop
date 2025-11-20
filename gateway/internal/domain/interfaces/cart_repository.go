package interfaces

import (
	"io"
	"net/http"
)

// CartRepository defines the interface for cart service operations
type CartRepository interface {
	Get(req *http.Request) (*http.Response, error)
	Add(body io.Reader, req *http.Request) (*http.Response, error)
	Update(itemID string, body io.Reader, req *http.Request) (*http.Response, error)
	Delete(itemID string, req *http.Request) (*http.Response, error)
}

