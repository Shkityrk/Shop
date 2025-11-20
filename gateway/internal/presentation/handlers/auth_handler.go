package handlers

import (
	"gateway/internal/application/services"
	"gateway/internal/domain/models"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthHandler handles authentication related requests
type AuthHandler struct {
	service *services.AuthService
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
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

	resp, err := h.service.Register(userCreate, c.Request)
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

	resp, err := h.service.Login(userLogin, c.Request)
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
	resp, err := h.service.Logout(c.Request)
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
	resp, err := h.service.Info(c.Request)
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
