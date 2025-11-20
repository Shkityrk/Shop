package interfaces

import (
	"io"
	"net/http"
)

// AuthRepository defines the interface for authentication service operations
type AuthRepository interface {
	Register(body io.Reader, req *http.Request) (*http.Response, error)
	Login(body io.Reader, req *http.Request) (*http.Response, error)
	Logout(req *http.Request) (*http.Response, error)
	Info(req *http.Request) (*http.Response, error)
}

