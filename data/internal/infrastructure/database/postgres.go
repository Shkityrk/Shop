package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const (
	maxOpenConns    = 25
	maxIdleConns    = 5
	connMaxLifetime = 5 * time.Minute
	connMaxIdleTime = 10 * time.Minute
)

// PostgresDB обертка над *sql.DB
type PostgresDB struct {
	DB *sql.DB
}

// NewPostgresDB создает новое подключение к PostgreSQL
func NewPostgresDB(dsn string) (*PostgresDB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Настройка пула соединений
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(connMaxLifetime)
	db.SetConnMaxIdleTime(connMaxIdleTime)

	// Проверка подключения
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logrus.Info("Successfully connected to PostgreSQL database")

	return &PostgresDB{DB: db}, nil
}

// Close закрывает соединение с базой данных
func (p *PostgresDB) Close() error {
	if p.DB != nil {
		logrus.Info("Closing database connection")
		return p.DB.Close()
	}
	return nil
}

// GetDB возвращает *sql.DB для прямого использования
func (p *PostgresDB) GetDB() *sql.DB {
	return p.DB
}

// InitCartItemsTable создает таблицу cart_items, если она не существует
func (p *PostgresDB) InitCartItemsTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS cart_items (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL,
			product_id INTEGER NOT NULL,
			quantity INTEGER NOT NULL DEFAULT 1,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
		
		CREATE INDEX IF NOT EXISTS idx_cart_items_user_id ON cart_items(user_id);
		CREATE INDEX IF NOT EXISTS idx_cart_items_product_id ON cart_items(product_id);
		CREATE INDEX IF NOT EXISTS idx_cart_items_user_product ON cart_items(user_id, product_id);
	`

	_, err := p.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create cart_items table: %w", err)
	}

	logrus.Info("Cart items table initialized successfully")
	return nil
}

