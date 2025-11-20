package services

import (
	"bytes"
	"encoding/json"
	"gateway/internal/domain/interfaces"
	"gateway/internal/domain/models"
	"net/http"
)

// CartService handles cart business logic
type CartService struct {
	repo interfaces.CartRepository
}

// NewCartService creates a new cart service
func NewCartService(repo interfaces.CartRepository) *CartService {
	return &CartService{
		repo: repo,
	}
}

// Get gets cart items
func (s *CartService) Get(req *http.Request) (*http.Response, error) {
	return s.repo.Get(req)
}

// Add adds item to cart
func (s *CartService) Add(cartItemCreate models.CartItemCreate, req *http.Request) (*http.Response, error) {
	jsonData, err := json.Marshal(cartItemCreate)
	if err != nil {
		return nil, err
	}
	return s.repo.Add(bytes.NewBuffer(jsonData), req)
}

// Update updates cart item
func (s *CartService) Update(itemID string, cartItemCreate models.CartItemCreate, req *http.Request) (*http.Response, error) {
	jsonData, err := json.Marshal(cartItemCreate)
	if err != nil {
		return nil, err
	}
	return s.repo.Update(itemID, bytes.NewBuffer(jsonData), req)
}

// Delete deletes cart item
func (s *CartService) Delete(itemID string, req *http.Request) (*http.Response, error) {
	return s.repo.Delete(itemID, req)
}

