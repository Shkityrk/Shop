package services

import (
	"bytes"
	"encoding/json"
	"gateway/internal/domain/interfaces"
	"gateway/internal/domain/models"
	"net/http"
)

// AuthService handles authentication business logic
type AuthService struct {
	repo interfaces.AuthRepository
}

// NewAuthService creates a new auth service
func NewAuthService(repo interfaces.AuthRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

// Register handles user registration
func (s *AuthService) Register(userCreate models.UserCreate, req *http.Request) (*http.Response, error) {
	jsonData, err := json.Marshal(userCreate)
	if err != nil {
		return nil, err
	}
	return s.repo.Register(bytes.NewBuffer(jsonData), req)
}

// Login handles user login
func (s *AuthService) Login(userLogin models.UserLogin, req *http.Request) (*http.Response, error) {
	jsonData, err := json.Marshal(userLogin)
	if err != nil {
		return nil, err
	}
	return s.repo.Login(bytes.NewBuffer(jsonData), req)
}

// Logout handles user logout
func (s *AuthService) Logout(req *http.Request) (*http.Response, error) {
	return s.repo.Logout(req)
}

// Info gets current user info
func (s *AuthService) Info(req *http.Request) (*http.Response, error) {
	return s.repo.Info(req)
}

// GetStaff получает список сотрудников
func (s *AuthService) GetStaff() (*http.Response, error) {
	return s.repo.GetStaff()
}

