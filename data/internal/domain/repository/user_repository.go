package repository

import (
	"context"
	"data-service/internal/domain/entity"
)

// UserRepository определяет интерфейс для работы с пользователями
type UserRepository interface {
	// Create создает нового пользователя
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	
	// GetByID получает пользователя по ID
	GetByID(ctx context.Context, id int) (*entity.User, error)
	
	// GetByUsername получает пользователя по username
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	
	// GetByEmail получает пользователя по email
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	
	// List получает список всех пользователей
	List(ctx context.Context, filter *entity.UserFilter) ([]*entity.User, error)
	
	// Update обновляет данные пользователя
	Update(ctx context.Context, user *entity.User) (*entity.User, error)
	
	// Delete удаляет пользователя по ID
	Delete(ctx context.Context, id int) error
	
	// Exists проверяет существование пользователя по username или email
	Exists(ctx context.Context, username, email string) (bool, error)
}

