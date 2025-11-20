package http

import (
	"gateway/internal/domain/interfaces"
	"io"
	"net/http"
)

// ProductRepositoryHTTP implements ProductRepository using HTTP client
type ProductRepositoryHTTP struct {
	serviceURL string
	client     *http.Client
}

// NewProductRepositoryHTTP creates a new HTTP product repository
func NewProductRepositoryHTTP(serviceURL string) interfaces.ProductRepository {
	return &ProductRepositoryHTTP{
		serviceURL: serviceURL,
		client:     &http.Client{},
	}
}

// List sends list request
func (r *ProductRepositoryHTTP) List(queryParams string, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("GET", r.serviceURL+"/product/list", nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = queryParams
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// Add sends add request
func (r *ProductRepositoryHTTP) Add(body io.Reader, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("POST", r.serviceURL+"/product/add", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// Update sends update request
func (r *ProductRepositoryHTTP) Update(id string, body io.Reader, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("PUT", r.serviceURL+"/product/update/"+id, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// Verify sends verify request
func (r *ProductRepositoryHTTP) Verify(name string, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("GET", r.serviceURL+"/product/verify/"+name, nil)
	if err != nil {
		return nil, err
	}
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

// Info sends info request
func (r *ProductRepositoryHTTP) Info(id string, originalReq *http.Request) (*http.Response, error) {
	req, err := http.NewRequest("GET", r.serviceURL+"/product/info/"+id, nil)
	if err != nil {
		return nil, err
	}
	copyHeaders(originalReq, req)
	return r.client.Do(req)
}

