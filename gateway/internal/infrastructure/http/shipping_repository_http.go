package http

import (
	"gateway/internal/domain/interfaces"
	"io"
	"net/http"
)

// ShippingRepositoryHTTP implements ShippingRepository using HTTP client
type ShippingRepositoryHTTP struct {
	serviceURL string
	client     *http.Client
}

// NewShippingRepositoryHTTP creates a new HTTP shipping repository
func NewShippingRepositoryHTTP(serviceURL string) interfaces.ShippingRepository {
	return &ShippingRepositoryHTTP{
		serviceURL: serviceURL,
		client:     &http.Client{},
	}
}

// Create sends create shipment request
func (r *ShippingRepositoryHTTP) Create(body io.Reader, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("POST", r.serviceURL+"/shipping", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// Get fetches shipment by tracking code
func (r *ShippingRepositoryHTTP) Get(tracking string, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("GET", r.serviceURL+"/shipping/"+tracking, nil)
	if err != nil {
		return nil, err
	}
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// UpdateStatus updates shipment status
func (r *ShippingRepositoryHTTP) UpdateStatus(tracking string, body io.Reader, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("PATCH", r.serviceURL+"/shipping/"+tracking+"/status", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// List fetches all shipments
func (r *ShippingRepositoryHTTP) List(queryParams string, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("GET", r.serviceURL+"/shipping/list", nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = queryParams
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

