package services

import (
	"bytes"
	"encoding/json"
	"gateway/internal/domain/interfaces"
	"gateway/internal/domain/models"
	"net/http"
)

// ShippingService handles shipping logic
type ShippingService struct {
	repo interfaces.ShippingRepository
}

// NewShippingService creates Shipping service
func NewShippingService(repo interfaces.ShippingRepository) *ShippingService {
	return &ShippingService{repo: repo}
}

// Create creates shipment
func (s *ShippingService) Create(payload models.ShipmentCreate, req *http.Request) (*http.Response, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return s.repo.Create(bytes.NewBuffer(jsonData), req)
}

// Get returns shipment by tracking
func (s *ShippingService) Get(tracking string, req *http.Request) (*http.Response, error) {
	return s.repo.Get(tracking, req)
}

// UpdateStatus updates shipment status
func (s *ShippingService) UpdateStatus(tracking string, payload models.ShipmentStatusUpdate, req *http.Request) (*http.Response, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return s.repo.UpdateStatus(tracking, bytes.NewBuffer(jsonData), req)
}

// List returns all shipments
func (s *ShippingService) List(queryParams string, req *http.Request) (*http.Response, error) {
	return s.repo.List(queryParams, req)
}

