package handler

import (
	"data-service/internal/application/dto"
	"data-service/internal/application/service"
	"data-service/internal/domain/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// UserHandler обработчик HTTP запросов для пользователей
type UserHandler struct {
	service *service.UserService
	logger  *logrus.Logger
}

// NewUserHandler создает новый экземпляр UserHandler
func NewUserHandler(service *service.UserService, logger *logrus.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

// CreateUser создает нового пользователя
// @Summary Create user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserRequest true "User data"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Failed to bind request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.CreateUser(c.Request.Context(), &req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to create user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUserByID получает пользователя по ID
// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.UserResponse
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.service.GetUserByID(c.Request.Context(), id)
	if err != nil {
		h.logger.WithError(err).Error("Failed to get user by ID")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserByUsername получает пользователя по username
// @Summary Get user by username
// @Description Get user by username
// @Tags users
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} dto.UserResponse
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/users/username/{username} [get]
func (h *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")

	user, err := h.service.GetUserByUsername(c.Request.Context(), username)
	if err != nil {
		h.logger.WithError(err).Error("Failed to get user by username")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserByEmail получает пользователя по email
// @Summary Get user by email
// @Description Get user by email
// @Tags users
// @Produce json
// @Param email path string true "Email"
// @Success 200 {object} dto.UserResponse
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/users/email/{email} [get]
func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")

	user, err := h.service.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		h.logger.WithError(err).Error("Failed to get user by email")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// ListUsers получает список всех пользователей
// @Summary List users
// @Description Get list of all users
// @Tags users
// @Produce json
// @Param username query string false "Filter by username"
// @Param email query string false "Filter by email"
// @Success 200 {array} dto.UserResponse
// @Failure 500 {object} map[string]interface{}
// @Router /api/users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
	filter := &entity.UserFilter{}

	if username := c.Query("username"); username != "" {
		filter.Username = &username
	}
	if email := c.Query("email"); email != "" {
		filter.Email = &email
	}

	users, err := h.service.ListUsers(c.Request.Context(), filter)
	if err != nil {
		h.logger.WithError(err).Error("Failed to list users")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUser обновляет данные пользователя
// @Summary Update user
// @Description Update user data
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body dto.UpdateUserRequest true "User data"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Failed to bind request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.UpdateUser(c.Request.Context(), id, &req)
	if err != nil {
		h.logger.WithError(err).Error("Failed to update user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser удаляет пользователя
// @Summary Delete user
// @Description Delete user by ID
// @Tags users
// @Param id path int true "User ID"
// @Success 204
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.service.DeleteUser(c.Request.Context(), id); err != nil {
		h.logger.WithError(err).Error("Failed to delete user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// CheckUserExists проверяет существование пользователя
// @Summary Check user exists
// @Description Check if user exists by username or email
// @Tags users
// @Accept json
// @Produce json
// @Param data body dto.CheckUserExistsRequest true "Check data"
// @Success 200 {object} dto.CheckUserExistsResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/users/exists [post]
func (h *UserHandler) CheckUserExists(c *gin.Context) {
	var req dto.CheckUserExistsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithError(err).Error("Failed to bind request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := ""
	if req.Username != nil {
		username = *req.Username
	}

	email := ""
	if req.Email != nil {
		email = *req.Email
	}

	exists, err := h.service.CheckUserExists(c.Request.Context(), username, email)
	if err != nil {
		h.logger.WithError(err).Error("Failed to check user existence")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.CheckUserExistsResponse{Exists: exists})
}

