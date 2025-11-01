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

type userRepositoryImpl struct {
	db *sql.DB
}

// NewUserRepository создает новую реализацию UserRepository
func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepositoryImpl{db: db}
}

// Create создает нового пользователя
func (r *userRepositoryImpl) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	query := `
		INSERT INTO users (first_name, last_name, username, email, hashed_password)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, first_name, last_name, username, email, hashed_password
	`

	var created entity.User
	err := r.db.QueryRowContext(
		ctx, query,
		user.FirstName, user.LastName, user.Username, user.Email, user.HashedPassword,
	).Scan(
		&created.ID, &created.FirstName, &created.LastName,
		&created.Username, &created.Email, &created.HashedPassword,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &created, nil
}

// GetByID получает пользователя по ID
func (r *userRepositoryImpl) GetByID(ctx context.Context, id int) (*entity.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, hashed_password
		FROM users
		WHERE id = $1
	`

	var user entity.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.FirstName, &user.LastName,
		&user.Username, &user.Email, &user.HashedPassword,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}

	return &user, nil
}

// GetByUsername получает пользователя по username
func (r *userRepositoryImpl) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, hashed_password
		FROM users
		WHERE username = $1
	`

	var user entity.User
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID, &user.FirstName, &user.LastName,
		&user.Username, &user.Email, &user.HashedPassword,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}

	return &user, nil
}

// GetByEmail получает пользователя по email
func (r *userRepositoryImpl) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, hashed_password
		FROM users
		WHERE email = $1
	`

	var user entity.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.FirstName, &user.LastName,
		&user.Username, &user.Email, &user.HashedPassword,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return &user, nil
}

// List получает список всех пользователей
func (r *userRepositoryImpl) List(ctx context.Context, filter *entity.UserFilter) ([]*entity.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, hashed_password
		FROM users
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
		if filter.Username != nil {
			conditions = append(conditions, fmt.Sprintf("username = $%d", argPos))
			args = append(args, *filter.Username)
			argPos++
		}
		if filter.Email != nil {
			conditions = append(conditions, fmt.Sprintf("email = $%d", argPos))
			args = append(args, *filter.Email)
			argPos++
		}
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(
			&user.ID, &user.FirstName, &user.LastName,
			&user.Username, &user.Email, &user.HashedPassword,
		); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return users, nil
}

// Update обновляет данные пользователя
func (r *userRepositoryImpl) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	query := `
		UPDATE users
		SET first_name = $1, last_name = $2, username = $3, email = $4, hashed_password = $5
		WHERE id = $6
		RETURNING id, first_name, last_name, username, email, hashed_password
	`

	var updated entity.User
	err := r.db.QueryRowContext(
		ctx, query,
		user.FirstName, user.LastName, user.Username, user.Email, user.HashedPassword, user.ID,
	).Scan(
		&updated.ID, &updated.FirstName, &updated.LastName,
		&updated.Username, &updated.Email, &updated.HashedPassword,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return &updated, nil
}

// Delete удаляет пользователя по ID
func (r *userRepositoryImpl) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// Exists проверяет существование пользователя по username или email
func (r *userRepositoryImpl) Exists(ctx context.Context, username, email string) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1 FROM users WHERE username = $1 OR email = $2
		)
	`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, username, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check user existence: %w", err)
	}

	return exists, nil
}

