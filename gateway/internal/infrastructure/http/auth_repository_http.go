package http

import (
	"gateway/internal/domain/interfaces"
	"io"
	"net/http"
)

// AuthRepositoryHTTP implements AuthRepository using HTTP client
type AuthRepositoryHTTP struct {
	serviceURL string
	client     *http.Client
}

// NewAuthRepositoryHTTP creates a new HTTP auth repository
func NewAuthRepositoryHTTP(serviceURL string) interfaces.AuthRepository {
	return &AuthRepositoryHTTP{
		serviceURL: serviceURL,
		client:     &http.Client{},
	}
}

// Register sends registration request
func (r *AuthRepositoryHTTP) Register(body io.Reader, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("POST", r.serviceURL+"/auth/register", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// Login sends login request
func (r *AuthRepositoryHTTP) Login(body io.Reader, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("POST", r.serviceURL+"/auth/login", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// Logout sends logout request
func (r *AuthRepositoryHTTP) Logout(originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("POST", r.serviceURL+"/auth/logout", nil)
	if err != nil {
		return nil, err
	}
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// Info sends info request
func (r *AuthRepositoryHTTP) Info(originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("GET", r.serviceURL+"/auth/info", nil)
	if err != nil {
		return nil, err
	}
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// GetStaff получает список сотрудников из data service
func (r *AuthRepositoryHTTP) GetStaff() (*http.Response, error) {
	// Обращаемся к data service напрямую (data:8004)
	req, err := http.NewRequest("GET", "http://data:8004/api/users/staff", nil)
	if err != nil {
		return nil, err
	}
	return r.client.Do(req)
}

