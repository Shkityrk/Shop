package services

import (
	"bytes"
	"encoding/json"
	"gateway/internal/domain/interfaces"
	"gateway/internal/domain/models"
	"net/http"
)

// ProductService handles product business logic
type ProductService struct {
	repo interfaces.ProductRepository
}

// NewProductService creates a new product service
func NewProductService(repo interfaces.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

// List gets list of products
func (s *ProductService) List(queryParams string, req *http.Request) (*http.Response, error) {
	return s.repo.List(queryParams, req)
}

// Add adds a new product
func (s *ProductService) Add(productCreate models.ProductCreate, req *http.Request) (*http.Response, error) {
	jsonData, err := json.Marshal(productCreate)
	if err != nil {
		return nil, err
	}
	return s.repo.Add(bytes.NewBuffer(jsonData), req)
}

// Update updates a product
func (s *ProductService) Update(id string, productCreate models.ProductCreate, req *http.Request) (*http.Response, error) {
	jsonData, err := json.Marshal(productCreate)
	if err != nil {
		return nil, err
	}
	return s.repo.Update(id, bytes.NewBuffer(jsonData), req)
}

// Verify verifies if product exists
func (s *ProductService) Verify(name string, req *http.Request) (*http.Response, error) {
	return s.repo.Verify(name, req)
}

// Info gets product info
func (s *ProductService) Info(id string, req *http.Request) (*http.Response, error) {
	return s.repo.Info(id, req)
}

