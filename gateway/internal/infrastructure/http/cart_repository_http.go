package http

import (
	"gateway/internal/domain/interfaces"
	"io"
	"net/http"
)

// CartRepositoryHTTP implements CartRepository using HTTP client
type CartRepositoryHTTP struct {
	serviceURL string
	client     *http.Client
}

// NewCartRepositoryHTTP creates a new HTTP cart repository
func NewCartRepositoryHTTP(serviceURL string) interfaces.CartRepository {
	return &CartRepositoryHTTP{
		serviceURL: serviceURL,
		client:     &http.Client{},
	}
}

// Get sends get request
func (r *CartRepositoryHTTP) Get(originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("GET", r.serviceURL+"/cart", nil)
	if err != nil {
		return nil, err
	}
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// Add sends add request
func (r *CartRepositoryHTTP) Add(body io.Reader, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("POST", r.serviceURL+"/cart/add", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// Update sends update request
func (r *CartRepositoryHTTP) Update(itemID string, body io.Reader, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("PUT", r.serviceURL+"/cart/update/"+itemID, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// Delete sends delete request
func (r *CartRepositoryHTTP) Delete(itemID string, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", r.serviceURL+"/cart/delete/"+itemID, nil)
	if err != nil {
		return nil, err
	}
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

