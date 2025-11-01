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

type productRepositoryImpl struct {
	db *sql.DB
}

// NewProductRepository создает новую реализацию ProductRepository
func NewProductRepository(db *sql.DB) repository.ProductRepository {
	return &productRepositoryImpl{db: db}
}

// Create создает новый продукт
func (r *productRepositoryImpl) Create(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	query := `
		INSERT INTO product (name, short_description, full_description, composition, weight, price, photo)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, name, short_description, full_description, composition, weight, price, photo
	`

	var created entity.Product
	err := r.db.QueryRowContext(
		ctx, query,
		product.Name, product.ShortDescription, product.FullDescription,
		product.Composition, product.Weight, product.Price, product.Photo,
	).Scan(
		&created.ID, &created.Name, &created.ShortDescription, &created.FullDescription,
		&created.Composition, &created.Weight, &created.Price, &created.Photo,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return &created, nil
}

// GetByID получает продукт по ID
func (r *productRepositoryImpl) GetByID(ctx context.Context, id int) (*entity.Product, error) {
	query := `
		SELECT id, name, short_description, full_description, composition, weight, price, photo
		FROM product
		WHERE id = $1
	`

	var product entity.Product
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&product.ID, &product.Name, &product.ShortDescription, &product.FullDescription,
		&product.Composition, &product.Weight, &product.Price, &product.Photo,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get product by id: %w", err)
	}

	return &product, nil
}

// GetByName получает продукт по названию
func (r *productRepositoryImpl) GetByName(ctx context.Context, name string) (*entity.Product, error) {
	query := `
		SELECT id, name, short_description, full_description, composition, weight, price, photo
		FROM product
		WHERE name = $1
	`

	var product entity.Product
	err := r.db.QueryRowContext(ctx, query, name).Scan(
		&product.ID, &product.Name, &product.ShortDescription, &product.FullDescription,
		&product.Composition, &product.Weight, &product.Price, &product.Photo,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get product by name: %w", err)
	}

	return &product, nil
}

// List получает список всех продуктов
func (r *productRepositoryImpl) List(ctx context.Context, filter *entity.ProductFilter) ([]*entity.Product, error) {
	query := `
		SELECT id, name, short_description, full_description, composition, weight, price, photo
		FROM product
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
		if filter.Name != nil {
			conditions = append(conditions, fmt.Sprintf("name = $%d", argPos))
			args = append(args, *filter.Name)
			argPos++
		}
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list products: %w", err)
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(
			&product.ID, &product.Name, &product.ShortDescription, &product.FullDescription,
			&product.Composition, &product.Weight, &product.Price, &product.Photo,
		); err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return products, nil
}

// Update обновляет данные продукта
func (r *productRepositoryImpl) Update(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	query := `
		UPDATE product
		SET name = $1, short_description = $2, full_description = $3, 
		    composition = $4, weight = $5, price = $6, photo = $7
		WHERE id = $8
		RETURNING id, name, short_description, full_description, composition, weight, price, photo
	`

	var updated entity.Product
	err := r.db.QueryRowContext(
		ctx, query,
		product.Name, product.ShortDescription, product.FullDescription,
		product.Composition, product.Weight, product.Price, product.Photo, product.ID,
	).Scan(
		&updated.ID, &updated.Name, &updated.ShortDescription, &updated.FullDescription,
		&updated.Composition, &updated.Weight, &updated.Price, &updated.Photo,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("product not found")
		}
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	return &updated, nil
}

// Delete удаляет продукт по ID
func (r *productRepositoryImpl) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM product WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product not found")
	}

	return nil
}

// Exists проверяет существование продукта по ID
func (r *productRepositoryImpl) Exists(ctx context.Context, id int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM product WHERE id = $1)`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check product existence: %w", err)
	}

	return exists, nil
}

