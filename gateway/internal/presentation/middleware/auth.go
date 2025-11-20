package middleware

import (
	"encoding/json"
	"gateway/internal/application/services"
	"gateway/internal/domain/models"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	UserIDHeader  = "X-User-Id"
	UserContextKey = "user_id"
)

// AuthMiddleware проверяет токен и добавляет user_id в контекст и заголовки
func AuthMiddleware(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Извлекаем токен из cookie или Authorization header
		token := extractToken(c)
		if token == "" {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error: "Токен не найден",
			})
			c.Abort()
			return
		}

		// Проверяем токен через auth service
		user, err := verifyToken(authService, token, c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error: "Не удалось проверить учетные данные",
			})
			c.Abort()
			return
		}

		// Добавляем user_id в контекст
		c.Set(UserContextKey, user.ID)

		// Добавляем user_id в заголовки для downstream сервисов
		c.Request.Header.Set(UserIDHeader, strconv.Itoa(user.ID))

		c.Next()
	}
}

// extractToken извлекает токен из cookie или Authorization header
func extractToken(c *gin.Context) string {
	// Сначала пробуем извлечь из cookie
	cookie, err := c.Cookie("access_token")
	if err == nil && cookie != "" {
		// Убираем префикс "Bearer " если есть
		token := strings.TrimPrefix(cookie, "Bearer ")
		return strings.TrimSpace(token)
	}

	// Если нет в cookie, пробуем из Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		// Убираем префикс "Bearer " если есть
		token := strings.TrimPrefix(authHeader, "Bearer ")
		return strings.TrimSpace(token)
	}

	return ""
}

// verifyToken проверяет токен через auth service
func verifyToken(authService *services.AuthService, token string, originalReq *http.Request) (*models.UserOut, error) {
	// Создаем запрос для проверки токена
	// AuthService.Info использует repo.Info, который создаст правильный запрос с URL
	// и скопирует cookies из переданного запроса
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		return nil, err
	}

	// Проверяем, есть ли токен в оригинальных cookies
	hasTokenCookie := false
	for _, cookie := range originalReq.Cookies() {
		if cookie.Name == "access_token" {
			req.AddCookie(cookie)
			hasTokenCookie = true
			break
		}
	}

	// Если токен не в cookie оригинального запроса, добавляем его
	if !hasTokenCookie {
		cookie := &http.Cookie{
			Name:  "access_token",
			Value: "Bearer " + token,
		}
		req.AddCookie(cookie)
	}

	// Вызываем Info через auth service
	// Repo создаст новый запрос с правильным URL и скопирует cookies из req
	resp, err := authService.Info(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Если статус не 200, токен невалиден
	if resp.StatusCode != http.StatusOK {
		return nil, &AuthError{Message: "Invalid token"}
	}

	// Декодируем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user models.UserOut
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// AuthError представляет ошибку аутентификации
type AuthError struct {
	Message string
}

func (e *AuthError) Error() string {
	return e.Message
}

