package repository

import (
	"context"
	"data-service/internal/domain/entity"
)

// CartRepository определяет интерфейс для работы с корзиной
type CartRepository interface {
	// Create создает новый элемент корзины
	Create(ctx context.Context, item *entity.CartItem) (*entity.CartItem, error)
	
	// GetByID получает элемент корзины по ID
	GetByID(ctx context.Context, id int) (*entity.CartItem, error)
	
	// GetByUserID получает все элементы корзины пользователя
	GetByUserID(ctx context.Context, userID int) ([]*entity.CartItem, error)
	
	// GetByUserAndProduct получает элемент корзины по user_id и product_id
	GetByUserAndProduct(ctx context.Context, userID, productID int) (*entity.CartItem, error)
	
	// List получает список элементов корзины с фильтрацией
	List(ctx context.Context, filter *entity.CartFilter) ([]*entity.CartItem, error)
	
	// Update обновляет элемент корзины
	Update(ctx context.Context, item *entity.CartItem) (*entity.CartItem, error)
	
	// Delete удаляет элемент корзины по ID
	Delete(ctx context.Context, id int) error
	
	// DeleteByUserID удаляет все элементы корзины пользователя
	DeleteByUserID(ctx context.Context, userID int) error
	
	// DeleteByUserAndProduct удаляет элемент корзины по user_id и product_id
	DeleteByUserAndProduct(ctx context.Context, userID, productID int) error
}

