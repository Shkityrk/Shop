package interfaces

import (
	"io"
	"net/http"
)

// ProductRepository defines the interface for product service operations
type ProductRepository interface {
	List(queryParams string, req *http.Request) (*http.Response, error)
	Add(body io.Reader, req *http.Request) (*http.Response, error)
	Update(id string, body io.Reader, req *http.Request) (*http.Response, error)
	Verify(name string, req *http.Request) (*http.Response, error)
	Info(id string, req *http.Request) (*http.Response, error)
}

