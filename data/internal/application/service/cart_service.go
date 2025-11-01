package service

import (
	"context"
	"data-service/internal/application/dto"
	"data-service/internal/domain/entity"
	"data-service/internal/domain/repository"
	"errors"
	"time"
)

var (
	ErrCartItemNotFound = errors.New("cart item not found")
)

// CartService сервис для работы с корзиной
type CartService struct {
	repo repository.CartRepository
}

// NewCartService создает новый экземпляр CartService
func NewCartService(repo repository.CartRepository) *CartService {
	return &CartService{repo: repo}
}

// CreateCartItem создает новый элемент корзины
func (s *CartService) CreateCartItem(ctx context.Context, req *dto.CreateCartItemRequest) (*dto.CartItemResponse, error) {
	// Проверяем, существует ли уже такой элемент
	existingItem, err := s.repo.GetByUserAndProduct(ctx, req.UserID, req.ProductID)
	if err == nil && existingItem != nil {
		// Если элемент существует, обновляем количество
		existingItem.Quantity += req.Quantity
		existingItem.UpdatedAt = time.Now()
		
		updatedItem, err := s.repo.Update(ctx, existingItem)
		if err != nil {
			return nil, err
		}
		return s.toCartItemResponse(updatedItem), nil
	}

	item := &entity.CartItem{
		UserID:    req.UserID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdItem, err := s.repo.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	return s.toCartItemResponse(createdItem), nil
}

// GetCartItemByID получает элемент корзины по ID
func (s *CartService) GetCartItemByID(ctx context.Context, id int) (*dto.CartItemResponse, error) {
	item, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, ErrCartItemNotFound
	}

	return s.toCartItemResponse(item), nil
}

// GetCartItemsByUserID получает все элементы корзины пользователя
func (s *CartService) GetCartItemsByUserID(ctx context.Context, userID int) ([]*dto.CartItemResponse, error) {
	items, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.CartItemResponse, 0, len(items))
	for _, item := range items {
		responses = append(responses, s.toCartItemResponse(item))
	}

	return responses, nil
}

// GetCartItemByUserAndProduct получает элемент корзины по user_id и product_id
func (s *CartService) GetCartItemByUserAndProduct(ctx context.Context, userID, productID int) (*dto.CartItemResponse, error) {
	item, err := s.repo.GetByUserAndProduct(ctx, userID, productID)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, ErrCartItemNotFound
	}

	return s.toCartItemResponse(item), nil
}

// ListCartItems получает список элементов корзины с фильтрацией
func (s *CartService) ListCartItems(ctx context.Context, filter *entity.CartFilter) ([]*dto.CartItemResponse, error) {
	items, err := s.repo.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.CartItemResponse, 0, len(items))
	for _, item := range items {
		responses = append(responses, s.toCartItemResponse(item))
	}

	return responses, nil
}

// UpdateCartItem обновляет элемент корзины
func (s *CartService) UpdateCartItem(ctx context.Context, id int, req *dto.UpdateCartItemRequest) (*dto.CartItemResponse, error) {
	item, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, ErrCartItemNotFound
	}

	// Обновляем только переданные поля
	if req.Quantity != nil {
		item.Quantity = *req.Quantity
	}
	item.UpdatedAt = time.Now()

	updatedItem, err := s.repo.Update(ctx, item)
	if err != nil {
		return nil, err
	}

	return s.toCartItemResponse(updatedItem), nil
}

// DeleteCartItem удаляет элемент корзины по ID
func (s *CartService) DeleteCartItem(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

// DeleteCartItemsByUserID удаляет все элементы корзины пользователя
func (s *CartService) DeleteCartItemsByUserID(ctx context.Context, userID int) error {
	return s.repo.DeleteByUserID(ctx, userID)
}

// DeleteCartItemByUserAndProduct удаляет элемент корзины по user_id и product_id
func (s *CartService) DeleteCartItemByUserAndProduct(ctx context.Context, userID, productID int) error {
	return s.repo.DeleteByUserAndProduct(ctx, userID, productID)
}

// toCartItemResponse преобразует entity.CartItem в dto.CartItemResponse
func (s *CartService) toCartItemResponse(item *entity.CartItem) *dto.CartItemResponse {
	return &dto.CartItemResponse{
		ID:        item.ID,
		UserID:    item.UserID,
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
		CreatedAt: item.CreatedAt.Format(time.RFC3339),
		UpdatedAt: item.UpdatedAt.Format(time.RFC3339),
	}
}

