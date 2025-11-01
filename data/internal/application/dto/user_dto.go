package dto

// CreateUserRequest DTO для создания пользователя
type CreateUserRequest struct {
	FirstName      string `json:"first_name" binding:"required"`
	LastName       string `json:"last_name" binding:"required"`
	Username       string `json:"username" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	HashedPassword string `json:"hashed_password" binding:"required"`
}

// UpdateUserRequest DTO для обновления пользователя
type UpdateUserRequest struct {
	FirstName      *string `json:"first_name,omitempty"`
	LastName       *string `json:"last_name,omitempty"`
	Username       *string `json:"username,omitempty"`
	Email          *string `json:"email,omitempty"`
	HashedPassword *string `json:"hashed_password,omitempty"`
}

// UserResponse DTO для ответа с данными пользователя
type UserResponse struct {
	ID             int    `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

// CheckUserExistsRequest DTO для проверки существования пользователя
type CheckUserExistsRequest struct {
	Username *string `json:"username,omitempty"`
	Email    *string `json:"email,omitempty"`
}

// CheckUserExistsResponse DTO для ответа проверки существования
type CheckUserExistsResponse struct {
	Exists bool `json:"exists"`
}

