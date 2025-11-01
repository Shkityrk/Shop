package repository

import (
	"context"
	"data-service/internal/domain/entity"
)

// ProductRepository определяет интерфейс для работы с продуктами
type ProductRepository interface {
	// Create создает новый продукт
	Create(ctx context.Context, product *entity.Product) (*entity.Product, error)
	
	// GetByID получает продукт по ID
	GetByID(ctx context.Context, id int) (*entity.Product, error)
	
	// GetByName получает продукт по названию
	GetByName(ctx context.Context, name string) (*entity.Product, error)
	
	// List получает список всех продуктов
	List(ctx context.Context, filter *entity.ProductFilter) ([]*entity.Product, error)
	
	// Update обновляет данные продукта
	Update(ctx context.Context, product *entity.Product) (*entity.Product, error)
	
	// Delete удаляет продукт по ID
	Delete(ctx context.Context, id int) error
	
	// Exists проверяет существование продукта по ID
	Exists(ctx context.Context, id int) (bool, error)
}

