package handlers

import (
	"bytes"
	"encoding/json"
	"gateway/models"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthHandler handles authentication related requests
type AuthHandler struct {
	serviceURL string
	client     *http.Client
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(serviceURL string) *AuthHandler {
	return &AuthHandler{
		serviceURL: serviceURL,
		client:     &http.Client{},
	}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username, email, and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.UserCreate true "User registration data"
// @Success 200 {object} models.UserOut
// @Failure 400 {object} models.ErrorResponse
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var userCreate models.UserCreate
	if err := c.ShouldBindJSON(&userCreate); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, err := json.Marshal(userCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to marshal request"})
		return
	}

	req, err := http.NewRequest("POST", h.serviceURL+"/auth/register", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create request"})
		return
	}

	req.Header.Set("Content-Type", "application/json")
	copyHeaders(c.Request, req)

	resp, err := h.client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with auth service"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to read response"})
		return
	}

	// Copy cookies from the response
	copyCookies(resp, c)

	c.Data(resp.StatusCode, "application/json", body)
}

// Login godoc
// @Summary Login user
// @Description Authenticate user with username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body models.UserLogin true "Login credentials"
// @Success 200 {object} models.MessageResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var userLogin models.UserLogin
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	jsonData, err := json.Marshal(userLogin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to marshal request"})
		return
	}

	req, err := http.NewRequest("POST", h.serviceURL+"/auth/login", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create request"})
		return
	}

	req.Header.Set("Content-Type", "application/json")
	copyHeaders(c.Request, req)

	resp, err := h.client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with auth service"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to read response"})
		return
	}

	// Copy cookies from the response
	copyCookies(resp, c)

	c.Data(resp.StatusCode, "application/json", body)
}

// Logout godoc
// @Summary Logout user
// @Description Logout the current user
// @Tags Auth
// @Produce json
// @Success 200 {object} models.MessageResponse
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	req, err := http.NewRequest("POST", h.serviceURL+"/auth/logout", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create request"})
		return
	}

	copyHeaders(c.Request, req)

	resp, err := h.client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with auth service"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to read response"})
		return
	}

	// Copy cookies from the response
	copyCookies(resp, c)

	c.Data(resp.StatusCode, "application/json", body)
}

// Info godoc
// @Summary Get current user info
// @Description Get information about the currently authenticated user
// @Tags Auth
// @Produce json
// @Success 200 {object} models.UserOut
// @Failure 401 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /auth/info [get]
func (h *AuthHandler) Info(c *gin.Context) {
	req, err := http.NewRequest("GET", h.serviceURL+"/auth/info", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create request"})
		return
	}

	copyHeaders(c.Request, req)

	resp, err := h.client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to communicate with auth service"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to read response"})
		return
	}

	c.Data(resp.StatusCode, "application/json", body)
}

