package entity

import "time"

// User представляет доменную модель пользователя
type User struct {
	ID             int       `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// UserFilter представляет фильтры для поиска пользователей
type UserFilter struct {
	ID       *int
	Username *string
	Email    *string
}

