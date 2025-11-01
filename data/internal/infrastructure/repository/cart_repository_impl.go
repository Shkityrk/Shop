package repository

import (
	"context"
	"database/sql"
	"data-service/internal/domain/entity"
	"data-service/internal/domain/repository"
	"errors"
	"fmt"
	"strings"
)

type cartRepositoryImpl struct {
	db *sql.DB
}

// NewCartRepository создает новую реализацию CartRepository
func NewCartRepository(db *sql.DB) repository.CartRepository {
	return &cartRepositoryImpl{db: db}
}

// Create создает новый элемент корзины
func (r *cartRepositoryImpl) Create(ctx context.Context, item *entity.CartItem) (*entity.CartItem, error) {
	query := `
		INSERT INTO cart_items (user_id, product_id, quantity)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, product_id, quantity
	`

	var created entity.CartItem
	err := r.db.QueryRowContext(
		ctx, query,
		item.UserID, item.ProductID, item.Quantity,
	).Scan(
		&created.ID, &created.UserID, &created.ProductID, &created.Quantity,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create cart item: %w", err)
	}

	return &created, nil
}

// GetByID получает элемент корзины по ID
func (r *cartRepositoryImpl) GetByID(ctx context.Context, id int) (*entity.CartItem, error) {
	query := `
		SELECT id, user_id, product_id, quantity
		FROM cart_items
		WHERE id = $1
	`

	var item entity.CartItem
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&item.ID, &item.UserID, &item.ProductID, &item.Quantity,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get cart item by id: %w", err)
	}

	return &item, nil
}

// GetByUserID получает все элементы корзины пользователя
func (r *cartRepositoryImpl) GetByUserID(ctx context.Context, userID int) ([]*entity.CartItem, error) {
	query := `
		SELECT id, user_id, product_id, quantity
		FROM cart_items
		WHERE user_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cart items by user id: %w", err)
	}
	defer rows.Close()

	var items []*entity.CartItem
	for rows.Next() {
		var item entity.CartItem
		if err := rows.Scan(&item.ID, &item.UserID, &item.ProductID, &item.Quantity); err != nil {
			return nil, fmt.Errorf("failed to scan cart item: %w", err)
		}
		items = append(items, &item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return items, nil
}

// GetByUserAndProduct получает элемент корзины по user_id и product_id
func (r *cartRepositoryImpl) GetByUserAndProduct(ctx context.Context, userID, productID int) (*entity.CartItem, error) {
	query := `
		SELECT id, user_id, product_id, quantity
		FROM cart_items
		WHERE user_id = $1 AND product_id = $2
	`

	var item entity.CartItem
	err := r.db.QueryRowContext(ctx, query, userID, productID).Scan(
		&item.ID, &item.UserID, &item.ProductID, &item.Quantity,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get cart item by user and product: %w", err)
	}

	return &item, nil
}

// List получает список элементов корзины с фильтрацией
func (r *cartRepositoryImpl) List(ctx context.Context, filter *entity.CartFilter) ([]*entity.CartItem, error) {
	query := `
		SELECT id, user_id, product_id, quantity
		FROM cart_items
	`

	var conditions []string
	var args []interface{}
	argPos := 1

	if filter != nil {
		if filter.ID != nil {
			conditions = append(conditions, fmt.Sprintf("id = $%d", argPos))
			args = append(args, *filter.ID)
			argPos++
		}
		if filter.UserID != nil {
			conditions = append(conditions, fmt.Sprintf("user_id = $%d", argPos))
			args = append(args, *filter.UserID)
			argPos++
		}
		if filter.ProductID != nil {
			conditions = append(conditions, fmt.Sprintf("product_id = $%d", argPos))
			args = append(args, *filter.ProductID)
			argPos++
		}
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list cart items: %w", err)
	}
	defer rows.Close()

	var items []*entity.CartItem
	for rows.Next() {
		var item entity.CartItem
		if err := rows.Scan(&item.ID, &item.UserID, &item.ProductID, &item.Quantity); err != nil {
			return nil, fmt.Errorf("failed to scan cart item: %w", err)
		}
		items = append(items, &item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return items, nil
}

// Update обновляет элемент корзины
func (r *cartRepositoryImpl) Update(ctx context.Context, item *entity.CartItem) (*entity.CartItem, error) {
	query := `
		UPDATE cart_items
		SET quantity = $1
		WHERE id = $2
		RETURNING id, user_id, product_id, quantity
	`

	var updated entity.CartItem
	err := r.db.QueryRowContext(
		ctx, query,
		item.Quantity, item.ID,
	).Scan(
		&updated.ID, &updated.UserID, &updated.ProductID, &updated.Quantity,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("cart item not found")
		}
		return nil, fmt.Errorf("failed to update cart item: %w", err)
	}

	return &updated, nil
}

// Delete удаляет элемент корзины по ID
func (r *cartRepositoryImpl) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM cart_items WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete cart item: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("cart item not found")
	}

	return nil
}

// DeleteByUserID удаляет все элементы корзины пользователя
func (r *cartRepositoryImpl) DeleteByUserID(ctx context.Context, userID int) error {
	query := `DELETE FROM cart_items WHERE user_id = $1`

	_, err := r.db.ExecContext(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("failed to delete cart items by user id: %w", err)
	}

	return nil
}

// DeleteByUserAndProduct удаляет элемент корзины по user_id и product_id
func (r *cartRepositoryImpl) DeleteByUserAndProduct(ctx context.Context, userID, productID int) error {
	query := `DELETE FROM cart_items WHERE user_id = $1 AND product_id = $2`

	result, err := r.db.ExecContext(ctx, query, userID, productID)
	if err != nil {
		return fmt.Errorf("failed to delete cart item by user and product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("cart item not found")
	}

	return nil
}

